package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

var Processor = protobuf.NewProcessor()

func init() {
	Processor.SetByteOrder(true)
	Processor.Register(&LoginRequest{})
	Processor.Register(&LoginResponse{})
}
