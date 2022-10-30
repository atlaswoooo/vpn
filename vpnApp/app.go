package app

import (
	config "myvpn/common/vpnConfig"
	"myvpn/tuntap"

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
}
