package miner

import (
	"context"
)

var CoalTransferPoint chan int = make(chan int)

var NotificationPoint chan string = make(chan string, 10)

type Miner interface {
	Run(ctx context.Context, id int) //<- Coal
	//	Info() MinerInfo
	Info() *MinerInfo
}

type MinerInfo struct {
	Class  string `json:"class"`
	Cost   int    `json:"cost"`
	Energy int    `json:"energy"`
	Income int    `json:"income"`
	Rest   int    `json:"rest"`
}
