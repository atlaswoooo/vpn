package tuntap

import (
	"log"
	config "myvpn/common/vpnConfig"
	"net"
	"runtime"
	"strconv"
	"vtun/common/netutil"

	"github.com/net-byte/water"
)

//配置tun口
//配置tun口
func configTun(conf config.Config, tunInferface *water.Interface) {
	ip, _, err := net.ParseCIDR(conf.CIDR)
	if err != nil {
		log.Panicf("error cidr %v", conf.CIDR)
	}
	ipv6, _, err := net.ParseCIDR(conf.CIDRv6)
	if err != nil {
		log.Panicf("error ipv6 cidr %v", conf.CIDRv6)
	}

	//f服务端分平台处理
	os := runtime.GOOS
	if os == "linux" {
		netutil.ExecCmd("/sbin/ip", "link", "set", "dev", tunInferface.Name(), "mtu", strconv.Itoa(config.MTU))
		netutil.ExecCmd("/sbin/ip", "addr", "add", config.CIDR, "dev", iface.Name())
		netutil.ExecCmd("/sbin/ip", "-6", "addr", "add", config.CIDRv6, "dev", iface.Name())
		netutil.ExecCmd("/sbin/ip", "link", "set", "dev", iface.Name(), "up")
		if !config.ServerMode && config.GlobalMode {
			physicalIface := netutil.GetInterface()
			serverAddrIP := netutil.LookupServerAddrIP(config.ServerAddr)
			if physicalIface != "" && serverAddrIP != nil {
				if config.LocalGateway != "" {
					netutil.ExecCmd("/sbin/ip", "route", "add", "0.0.0.0/1", "dev", iface.Name())
					netutil.ExecCmd("/sbin/ip", "route", "add", "128.0.0.0/1", "dev", iface.Name())
					if serverAddrIP.To4() != nil {
						netutil.ExecCmd("/sbin/ip", "route", "add", serverAddrIP.To4().String()+"/32", "via", config.LocalGateway, "dev", physicalIface)
					}
				}
				if config.LocalGatewayV6 != "" {
					netutil.ExecCmd("/sbin/ip", "-6", "route", "add", "::/1", "dev", iface.Name())
					if serverAddrIP.To16() != nil {
						netutil.ExecCmd("/sbin/ip", "-6", "route", "add", serverAddrIP.To16().String()+"/128", "via", config.LocalGatewayV6, "dev", physicalIface)
					}
				}
				if net.ParseIP(config.DNSIP) != nil && net.ParseIP(config.DNSIP).To4() == nil {
					netutil.ExecCmd("/sbin/ip", "route", "add", config.DNSIP+"/128", "via", config.LocalGatewayV6, "dev", physicalIface)
				} else {
					netutil.ExecCmd("/sbin/ip", "route", "add", config.DNSIP+"/32", "via", config.LocalGateway, "dev", physicalIface)
				}
			}
		}

	}

	return
}

//创建tun口
//创建tun口
func CreateTun(conf config.Config) *water.Interface {
	c := water.Config{
		DeviceType: water.TUN,
		PlatformSpecificParams: water.PlatformSpecificParams{
			Name: conf.DeviceName,
		},
	}

	tunInferface, err := water.New(c)
	if err != nil {
		log.Fatalln("failed to create tun interface:", err)
	}
	log.Printf("interface created %v", tunInferface.Name())

	//配置tun口
	configTun(conf, tunInferface)

	return tunInferface
}
