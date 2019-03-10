package util

import (
	"github.com/micro/go-config"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

type (
	MongoDBConfig struct {
		Addr      []string `json:"addr"`
		Username  string   `json:"username"`
		Password  string   `json:"password"`
		PoolLimit int      `json:"poolLimit"`
	}
)

var (
	DefaultMgoConf MongoDBConfig
)

func Init() {
	// 加载mongo配置
	err := config.Get("mongodb").Scan(&DefaultMgoConf)
	if err != nil {
		logrus.Fatalf("get mgo config error: %s", err)
	}

	if len(DefaultMgoConf.Addr) == 0 {
		logrus.Fatalf("invalid mgo addr")
	}

	for k, _ := range DefaultMgoConf.Addr {
		if strings.Index(DefaultMgoConf.Addr[k], "$") == 0 {
			DefaultMgoConf.Addr[k] = os.Getenv(strings.TrimPrefix(DefaultMgoConf.Addr[k], "$"))
		}
	}

}
