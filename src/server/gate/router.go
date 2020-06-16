package gate

import (
	"server/login"
	"server/msg"
)

func init() {
	// 这里指定消息 LoginACK 路由到 login 模块
	msg.Processor.SetRouter(&msg.PingREQ{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.LoginREQ{}, login.ChanRPC)
}
