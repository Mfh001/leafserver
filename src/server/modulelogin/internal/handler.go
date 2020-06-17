package internal

import (
	//"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	"server/msg"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	// 向当前模块（gate 模块）注册 LoginREQ 消息的消息处理函数 handleLoginREQ
	handleMsg(&msg.PingREQ{}, handlePingREQ)
	handleMsg(&msg.LoginREQ{}, handleLoginREQ)
}

func handlePingREQ(args []interface{}) {
	// 收到的 PingREQ 消息
	//m := args[0].(*msg.PingREQ)
	// 消息的发送者
	a := args[1].(gate.Agent)
	// 输出收到的消息的内容
	//log.Debug("PingREQ %v", m.GetUID())
	// 给发送者回应一个 PingACK 消息
	a.WriteMsg(&msg.PingACK{
	})
}

func handleLoginREQ(args []interface{}) {
	// 收到的 LoginREQ 消息
	m := args[0].(*msg.LoginREQ)
	// 消息的发送者
	a := args[1].(gate.Agent)
	// 输出收到的消息的内容
	log.Debug("LoginREQ %v", m.GetUID())
	// 给发送者回应一个 LoginACK 消息
	a.WriteMsg(&msg.LoginACK{
		GameSvcID: "ssss",
	})
}