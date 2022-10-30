package tuntap

import (
	"log"
	netutil "myvpn/common/netutil"
	config "myvpn/common/vpnConfig"
	"runtime"
	"strconv"

	"github.com/net-byte/water"
)

//配置tun口
//配置tun口
func configTun(conf config.Config, tunInferface *water.Interface) {
	/*
		ip, _, err := net.ParseCIDR(conf.CIDR)
		if err != nil {
			log.Panicf("error cidr %v", conf.CIDR)
		}
	*/
	/*
		ipv6后续再考虑
		ipv6, _, err := net.ParseCIDR(conf.CIDRv6)
		if err != nil {
			log.Panicf("error ipv6 cidr %v", conf.CIDRv6)
		}
	*/

	//服务端分平台处理
	os := runtime.GOOS
	if os == "linux" {
		netutil.ExecCmd("/sbin/ip", "link", "set", "dev", tunInferface.Name(), "mtu", strconv.Itoa(conf.MTU))
		netutil.ExecCmd("/sbin/ip", "addr", "add", conf.CIDR, "dev", tunInferface.Name())
		//netutil.ExecCmd("/sbin/ip", "-6", "addr", "add", config.CIDRv6, "dev", iface.Name())
		//netutil.ExecCmd("/sbin/ip", "link", "set", "dev", iface.Name(), "up")
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

// ResetTun resets the tun interface
func ResetTun(config config.Config) {
	// reset gateway
	return
}
