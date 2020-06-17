package modulegate

import (
	"server/msg"
	"server/modulelogin"
)

func init() {
	// 这里指定消息 LoginACK 路由到 login 模块
	msg.Processor.SetRouter(&msg.PingREQ{}, modulelogin.ChanRPC)
	msg.Processor.SetRouter(&msg.LoginREQ{}, modulelogin.ChanRPC)
}
