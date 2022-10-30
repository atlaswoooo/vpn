package main

import (
	"flag"
	"fmt"
	config "myvpn/common/vpnConfig"
	"runtime"
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
	//app := app.NewApp(&config, _version)

	//初始化app
	//app.InitConfig()

	os := runtime.GOOS

	fmt.Printf("os:%+v\n", os)

}
