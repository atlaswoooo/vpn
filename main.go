package main

import (
	"flag"
	"fmt"
	config "myvpn/common/vpnConfig"
	app "myvpn/vpnApp"
	"os"
	"os/signal"
	"syscall"
)

const (
	_version = "1.1.0"
)

//main
//main
func main() {
	//解析参数,参数类型为"-dn devName"
	config := config.Config{}
	flag.StringVar(&config.DeviceName, "dn", "", "device name")
	flag.StringVar(&config.CIDR, "cidr", "", "tun interface cidr")
	flag.BoolVar(&config.Verbose, "verb", false, "enable verbose output")
	flag.StringVar(&config.Protocol, "p", "udp", "protocol udp/tls/grpc/ws/wss")
	flag.StringVar(&config.LocalAddr, "l", ":3000", "local address")
	flag.IntVar(&config.MTU, "mtu", 1500, "tun mtu")
	flag.Parse()

	//debug
	fmt.Printf("config:%+v\r\n", config)

	//新建app实例
	app := app.NewApp(&config, _version)

	//初始化app
	app.InitConfig()

	//开始干活
	go app.StartApp()

	//注册停止条件
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-quit
	app.StopApp()
}

/*
启动方式:
go run main.go -dn tun10 -cidr 172.31.0.1/20 -verb true
*/
