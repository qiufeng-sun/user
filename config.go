package main

import (
	"path/filepath"

	"github.com/astaxie/beego/config"

	"util/etcd"
	"util/logs"
	"util/logs/scribe"

	"core/net/lan"
)

var _ = logs.Debug

//
type Config struct {
	// scribe
	Scribed    bool
	ScribeAddr string

	// server
	LanCfg  *lan.LanCfg
	EtcdCfg *etcd.SrvCfg
	GoNum   int
}

func (this *Config) init(fileName string) bool {
	confd, e := config.NewConfig("ini", fileName)
	if e != nil {
		logs.Panicln("load config file failed! file:", fileName, "error:", e)
	}

	//[scribe]
	scribe.Init("user", confd)

	//[server]
	srvName := confd.String("server::name")
	srvAddr := confd.String("server::addr")
	this.LanCfg = lan.NewLanCfg(srvName, srvAddr)

	this.GoNum = confd.DefaultInt("server::gonum", 4)

	//[etcd]
	this.EtcdCfg = &etcd.SrvCfg{}
	this.EtcdCfg.EtcdAddrs = confd.Strings("etcd::addrs")
	this.EtcdCfg.SrvAddr = srvAddr
	this.EtcdCfg.SrvRegPath = confd.String("etcd::reg_path")
	this.EtcdCfg.SrvRegUpTick = confd.DefaultInt64("etcd::reg_uptick", 2000)

	// echo
	logs.Info("user config:%+v", *this)

	return true
}

//
var Cfg = &Config{}

//
func LoadConfig(confPath string) bool {
	// config
	confFile := filepath.Clean(confPath + "/self.ini")

	return Cfg.init(confFile)
}

//
func SrvId() string {
	return Cfg.LanCfg.ServerId()
}

//
func SrvName() string {
	return Cfg.LanCfg.Name
}

// to do add check func
