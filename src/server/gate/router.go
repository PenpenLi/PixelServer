package gate

import (
	"server/login"
	"server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.LoginRequest{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.RegisteRequest{}, login.ChanRPC)
}
