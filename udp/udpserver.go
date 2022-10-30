package udp

import (
	"log"
	config "myvpn/common/vpnConfig"
	"net"
	"time"

	"github.com/net-byte/water"
	"github.com/patrickmn/go-cache"
)

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
	s := &Server{config: config, iface: iface, localConn: conn, connCache: cache.New(30*time.Minute, 10*time.Minute)}
	go s.tunToUdp()
	s.udpToTun()
}
