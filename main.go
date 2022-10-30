package main

import (
	"flag"
	config "myvpn/common/vpnConfig"
	app "myvpn/vpnApp"
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
	flag.Parse()

	//新建app实例
	app := app.NewApp(&config, _version)

	//初始化app
	app.InitConfig()

	//开始干活
	go app.StartApp()
}
