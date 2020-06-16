package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)
// 使用 Protobuf 消息处理器
var Processor = protobuf.NewProcessor()

func init() {
	// 注册消息
	Processor.Register(&PingREQ{})
	Processor.Register(&PingACK{})
	Processor.Register(&LoginREQ{})
	Processor.Register(&LoginACK{})
}