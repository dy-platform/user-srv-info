package handler

import (
	"context"
	pb "github.com/dy-platform/user-srv-info/idl/platform/user/srv-info"
)

type Handle struct {

}


func (h *Handle)CreateUser(ctx context.Context,req *pb.CreateUserReq, resp *pb.CommonResp)error{
	return nil
}

func (h *Handle) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq, resp *pb.CommonResp) error {
	return nil
}
