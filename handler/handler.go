package handler

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/dy-platform/user-srv-info/dal/db"
	"github.com/dy-platform/user-srv-info/idl"
	info "github.com/dy-platform/user-srv-info/idl/platform/user/srv-info"
)

type Handle struct {

}


func (h *Handle)CreateUser(ctx context.Context,req *info.CreateUserReq, resp *info.CreateUserResp) error {
	resp.BaseResp = &base.Resp{
		Code:int32(base.CODE_OK),
	}

	if req.UserId == 0 {
		resp.BaseResp.Code = int32(base.CODE_INVALID_PARAMETER)
		return nil
	}

	err := db.InsertUserInfo(req.UserId, "", req.NickName, req.AvatarUrl, "")
	if err != nil {
		logrus.Warnf("db.InsertUserInfo error:%v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}
	return nil
}

func (h *Handle) GetUserInfo(ctx context.Context, req *info.GetUserInfoReq, resp *info.GetUserInfoResp) error {
	resp.BaseResp = &base.Resp{
		Code:int32(base.CODE_OK),
	}

	if req.UserId == 0 {
		resp.BaseResp.Code = int32(base.CODE_INVALID_PARAMETER)
		return nil
	}

	ret, err := db.GetOneUserInfo(req.UserId, req.Fields)
	if err != nil {
		logrus.Warnf("db.GetOneUserInfo error:%v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}

	data, err := json.Marshal(ret)
	if err != nil {
		logrus.Warnf("json marshal error:%v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
	}

	resp.Data = string(data)
	return nil
}
