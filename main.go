package main

import (
	"github.com/dy-gopkg/kit"
	"github.com/dy-platform/user-srv-info/dal/db"
	pb "github.com/dy-platform/user-srv-info/idl/platform/user/srv-info"
	h "github.com/dy-platform/user-srv-info/handler"
)

func main() {
	kit.Init()
	pb.RegisterUserInfoHandler(kit.Server(),&h.Handle{})
	db.DBInit()
	kit.Run()
}