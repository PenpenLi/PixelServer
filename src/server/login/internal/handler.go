package internal

import (
	"database/sql"
	"reflect"
	"server/msg"
	"strings"

	_ "github.com/Go-SQL-Driver/MySQL"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	log.Debug("[login init]")
	handleMsg(&msg.LoginRequest{}, handleAuth)
	handleMsg(&msg.RegisteRequest{}, handleRegiste)
}

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleAuth(args []interface{}) {
	m := args[0].(*msg.LoginRequest)
	a := args[1].(gate.Agent)

	db, err := sql.Open("mysql", "root:A845240287a@tcp(rm-wz9sw694mi8020vigo.mysql.rds.aliyuncs.com)/pixel_farm?charset=utf8")
	checkErr(err)

	var uid string
	err = db.QueryRow("SELECT uid FROM pixel_user WHERE account=? and password = ?", m.GetAccount(), m.GetPassword()).Scan(&uid)
	checkErr(err)
	log.Debug("[login handleAuth] uid = " + uid)
	if strings.Count(uid, "")-1 > 0 {
		a.WriteMsg(&msg.LoginResponse{
			Code: msg.LoginResponse_SUCCESS,
			Uid:  uid,
		})
	} else {
		a.WriteMsg(&msg.LoginResponse{
			Code: msg.LoginResponse_FAIL,
			Err: &msg.Error{
				Code: 100,
				Msg:  "用户不存在",
			},
		})
		tx, err := db.Begin()
		err = db.QueryRow("SELECT REPLACE(UUID(),\"-\",\"\") FROM dual").Scan(&uid)
		checkErr(err)
		stmt, err1 := tx.Prepare("INSERT INTO pixel_user (uid, account, password) VALUES (?, ?, ?)")
		checkErr(err1)
		_, err2 := stmt.Exec(uid, m.GetAccount(), m.GetPassword())
		checkErr(err2)
		err3 := tx.Commit()
		// err3 := tx.Rollback()
		checkErr(err3)
	}

	newPlayerBaseInfo := new(PlayerBaseInfo)
	newPlayerBaseInfo.PlayerID = uid
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func handleRegiste(args []interface{}) {
	m := args[0].(*msg.RegisteRequest)
	a := args[1].(gate.Agent)

	db, err := sql.Open("mysql", "root:A845240287a@tcp(rm-wz9sw694mi8020vigo.mysql.rds.aliyuncs.com)/pixel_farm?charset=utf8")
	checkErr(err)

	var uid string
	err = db.QueryRow("SELECT uid FROM pixel_user WHERE account=? and password = ?", m.GetAccount(), m.GetPassword()).Scan(&uid)
	checkErr(err)
	log.Debug("[login handleRegiste] uid = " + uid)

	if strings.Count(uid, "")-1 > 0 {
		a.WriteMsg(&msg.RegisteResponse{
			Code: msg.RegisteResponse_FAIL,
			Err: &msg.Error{
				Code: 101,
				Msg:  "用户已存在",
			},
		})
	} else {
		tx, err := db.Begin()
		err = db.QueryRow("SELECT REPLACE(UUID(),\"-\",\"\") FROM dual").Scan(&uid)
		checkErr(err)
		stmt, err1 := tx.Prepare("INSERT INTO pixel_user (uid, account, password) VALUES (?, ?, ?)")
		checkErr(err1)
		_, err2 := stmt.Exec(uid, m.GetAccount(), m.GetPassword())
		checkErr(err2)
		err3 := tx.Commit()
		// err3 := tx.Rollback()
		checkErr(err3)

		a.WriteMsg(&msg.RegisteResponse{
			Code: msg.RegisteResponse_SUCCESS,
			Uid:  uid,
		})
	}
}
