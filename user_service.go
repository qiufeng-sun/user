package main

import (
	"util/logs"

	"core/server"

	"share/rpc"
)

var _ = logs.Debug

//
type UserSrv struct {
	server.Server
}

//
func NewUserSrv() *UserSrv {
	return &UserSrv{}
}

//
func (this *UserSrv) Init() bool {
	// config
	if !LoadConfig("conf/") {
		return false
	}

	// start rpc service
	rpc.InitServer(Cfg.LanCfg, Cfg.EtcdCfg, Cfg.GoNum, handleMsgs)

	return true
}

//
func (this UserSrv) String() string {
	return "userSrv"
}
