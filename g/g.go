package g

import (
	"github.com/fatedier/frp/models/config"
)

var (
	DefaultClientConfigIni = `
[common]
server_addr = 118.24.115.224
server_port = 7000
token = QWER6183z2lJKL

[socks5]
type = tcp
remote_port = 45187
plugin = socks5
`
)

var (
	GlbClientCfg *ClientCfg
	GlbServerCfg *ServerCfg
)

func init() {
	GlbClientCfg = &ClientCfg{
		ClientCommonConf: *config.GetDefaultClientConf(),
	}
	GlbServerCfg = &ServerCfg{
		ServerCommonConf: *config.GetDefaultServerConf(),
	}
}

type ClientCfg struct {
	config.ClientCommonConf

	CfgFile       string
	ServerUdpPort int // this is configured by login response from frps
}

type ServerCfg struct {
	config.ServerCommonConf

	CfgFile string
}
