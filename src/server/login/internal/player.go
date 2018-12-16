package internal

import (
	"github.com/name5566/leaf/gate"
)

var (
	playerIDQuene   = uint32(0)
	playerID2Player = make(map[uint32]*Player)
)

type Player struct {
	gate.Agent
	playerBaseInfo *PlayerBaseInfo
}
