package internal

import (
	"fmt"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("Login_Login", rpcLogin)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcLogin(args []interface{}) {
	fmt.Println("login request")
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	fmt.Println("close agent")
	a := args[0].(gate.Agent)
	_ = a
}
