package internal

import (
	"fmt"
	"reflect"
	"server/msg"
	"strconv"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	log.Debug("login init")
	handleMsg(&msg.LoginRequest{}, handleAuth)
	handleMsg(&msg.RegisteRequest{}, handleRegiste)
}

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleAuth(args []interface{}) {
	m := args[0].(*msg.LoginRequest)
	a := args[1].(gate.Agent)

	fmt.Println(a.RemoteAddr())
	fmt.Println(strconv.FormatInt(int64(playerIDQuene), 10) + "  " + m.GetAccount() + "  " + m.GetPassword())

	newPlayerBaseInfo := new(PlayerBaseInfo)
	newPlayerBaseInfo.PlayerID = strconv.FormatInt(int64(playerIDQuene), 10)
	newPlayerBaseInfo.Name = m.GetAccount()

	newPlayer := new(Player)
	newPlayer.Agent = a
	newPlayer.playerBaseInfo = newPlayerBaseInfo
	playerID2Player[playerIDQuene] = newPlayer

	playerIDQuene = playerIDQuene + 1

	for _, v := range playerID2Player {
		v.Agent.WriteMsg(&msg.LoginResponse{
			Code: msg.LoginResponse_SUCCESS,
			Uid:  v.playerBaseInfo.PlayerID,
		})
	}
}

func handleRegiste(args []interface{}) {

}
