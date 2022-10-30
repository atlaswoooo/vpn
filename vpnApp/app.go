package app

import (
	"log"
	"myvpn/common/netutil"
	config "myvpn/common/vpnConfig"
	"myvpn/tuntap"
	"myvpn/udp"

	"github.com/net-byte/water"
)

// vtun app struct
type AppVpn struct {
	Config       *config.Config
	Version      string
	TunInterface *water.Interface
}

//创建app实例
//创建app实例
func NewApp(config *config.Config, version string) *AppVpn {

	return &AppVpn{
		Config:  config,
		Version: version,
	}
}

//初始化配置文件
//初始化配置文件
func (app *AppVpn) InitConfig() {
	app.Config.BufferSize = 64 * 1024
	app.TunInterface = tuntap.CreateTun(*app.Config)
	netutil.PrintStats(app.Config.Verbose)
}

//开始后台服务
//开始后台服务
func (app *AppVpn) StartApp() {

	switch app.Config.Protocol {
	case "udp":
		udp.StartServer(app.TunInterface, *app.Config)
	default:
		udp.StartServer(app.TunInterface, *app.Config)
	}
}

// StopApp stops the app
func (app *AppVpn) StopApp() {
	tuntap.ResetTun(*app.Config)
	app.TunInterface.Close()
	log.Println("vtun stopped")
}
