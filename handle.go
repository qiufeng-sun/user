package main

import (
	"github.com/golang/protobuf/proto"

	"util/logs"

	"share/handler"
	"share/msg"
)

var _ = logs.Debug

//
func handleMsgs(msg []byte) []byte {
	return handler.HandleBytes(msg)
}

//
type handleMid func(int32, []byte) proto.Message

func (this handleMid) Handle(msgId int32, raw []byte) []byte {
	res := this(msgId, raw)
	if nil == res {
		return nil
	}

	r, e := handler.PackMsg(msgId, res)
	if e != nil {
		return nil
	}

	return r
}

//
func regFunc(msgId msg.EUserMsg, h func(int32, []byte) proto.Message) {
	handler.RegHandleBytes(int32(msgId), handleMid(h).Handle)
}

//
func init() {
	regFunc(msg.EUserMsg_ID_LoadUser, handleLoadUser)
}

//
func handleLoadUser(id int32, raw []byte) proto.Message {
	// parse
	var m msg.LoadUserReq
	e := handler.ParseMsgData(raw, &m)
	if e != nil {
		logs.Error("invalid msg! msgId:%v, error:%v", id, e)
		return nil
	}
	logs.Info("load user:%s", m.String())

	// process// to do

	// feedback// to do
	fb := &msg.LoadUserResp{UserId: m.UserId, UserName: proto.String("test1")}

	return fb
}
