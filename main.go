package main

import (
	"github.com/dy-gopkg/kit/micro"
	"github.com/dy-platform/user-srv-info/dal/db"
	"github.com/dy-platform/user-srv-info/handler"
	info "github.com/dy-platform/user-srv-info/idl/platform/user/srv-info"
	"github.com/dy-platform/user-srv-info/util/config"
	"github.com/sirupsen/logrus"
)

func main() {
	micro.Init()

	uconfig.Init()

	// 初始化数据库
	db.Init()

	err := info.RegisterUserInfoHandler(micro.Server(), &handler.Handle{})
	if err != nil {
		logrus.Fatalf("RegisterPassportHandler error:%v", err)
	}
	micro.Run()
}