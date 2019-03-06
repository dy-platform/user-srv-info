package main

import (
	"github.com/dy-gopkg/kit"
	"github.com/dy-platform/user-srv-info/dal/db"
	info "github.com/dy-platform/user-srv-info/idl/platform/user/srv-info"
	"github.com/dy-platform/user-srv-info/handler"
	"github.com/sirupsen/logrus"
	"github.com/dy-platform/user-srv-info/util/config"
)

func main() {
	kit.Init()

	uconfig.Init()

	// 初始化数据库
	db.Init()

	err := info.RegisterUserInfoHandler(kit.Server(), &handler.Handle{})
	if err != nil {
		logrus.Fatalf("RegisterPassportHandler error:%v", err)
	}
	kit.Run()
}