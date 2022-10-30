package udp

import (
	"log"
	"myvpn/common/cipher"
	"myvpn/common/counter"
	"myvpn/common/netutil"
	config "myvpn/common/vpnConfig"
	"net"
	"time"

	"github.com/golang/snappy"
	"github.com/net-byte/water"
	"github.com/patrickmn/go-cache"
)

// the server struct
type Server struct {
	config    config.Config
	iface     *water.Interface
	localConn *net.UDPConn
	connCache *cache.Cache
}

// tunToUdp sends packets from tun to udp
// tunToUdp sends packets from tun to udp
func (svr *Server) tunToUdp() {
	packet := make([]byte, svr.config.BufferSize)
	for {
		n, err := svr.iface.Read(packet)
		if err != nil {
			netutil.PrintErr(err, svr.config.Verbose)
			break
		}
		//这种写法相当于读数组到n+1下标
		b := packet[:n]
		if key := netutil.GetDstKey(b); key != "" {
			if v, ok := svr.connCache.Get(key); ok {
				if svr.config.Obfs {
					b = cipher.XOR(b)
				}
				if svr.config.Compress {
					b = snappy.Encode(nil, b)
				}
				svr.localConn.WriteToUDP(b, v.(*net.UDPAddr))
				counter.IncrWrittenBytes(n)
			}
		}
	}
}

// udpToTun sends packets from udp to tun
// udpToTun sends packets from udp to tun
func (svr *Server) udpToTun() {
	packet := make([]byte, svr.config.BufferSize)
	for {
		n, cliAddr, err := svr.localConn.ReadFromUDP(packet)
		if err != nil || n == 0 {
			netutil.PrintErr(err, svr.config.Verbose)
			continue
		}
		b := packet[:n]
		if svr.config.Compress {
			b, err = snappy.Decode(nil, b)
			if err != nil {
				netutil.PrintErr(err, svr.config.Verbose)
				continue
			}
		}
		if svr.config.Obfs {
			b = cipher.XOR(b)
		}
		if key := netutil.GetSrcKey(b); key != "" {
			svr.iface.Write(b)
			svr.connCache.Set(key, cliAddr, cache.DefaultExpiration)
			counter.IncrReadBytes(n)
		}
	}
}

// StartServer starts the udp server
// StartServer starts the udp server
func StartServer(iface *water.Interface, config config.Config) {
	log.Printf("vtun udp server started on %v", config.LocalAddr)
	localAddr, err := net.ResolveUDPAddr("udp", config.LocalAddr)
	if err != nil {
		log.Fatalln("failed to get udp socket:", err)
	}
	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		log.Fatalln("failed to listen on udp socket:", err)
	}
	defer conn.Close()

	svr := &Server{config: config, iface: iface, localConn: conn, connCache: cache.New(30*time.Minute, 10*time.Minute)}
	go svr.tunToUdp()
	svr.udpToTun()
}
