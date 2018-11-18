package handler

import (
	info "github.com/dy-platform/user-srv-info/idl/platform/user/srv-info"
	"github.com/sirupsen/logrus"
	"context"
)

type Handle struct {

}

func (h *Handle) Hello(ctx context.Context, req *info.HelloReq, rsp *info.HelloRsp) error {
	logrus.Debugf("hello %s", req.Name)
	rsp.Msg = "my name is Jerry"
	return nil
}
