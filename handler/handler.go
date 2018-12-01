package handler

import (
	"context"
	"encoding/json"
	"github.com/dy-platform/user-srv-info/Model"
	"github.com/dy-platform/user-srv-info/dal/db"
	pb "github.com/dy-platform/user-srv-info/idl/platform/user/srv-info"
	basePB "github.com/dy-platform/user-srv-info/idl"
)

type Handle struct {

}


func (h *Handle)CreateUser(ctx context.Context,req *pb.CreateUserReq, resp *pb.CommonResp)error{
	userModel := Model.UserInfo{
		Uid:       req.UserId,
		Nick:      req.NickName,
		Gender:    req.Gender,
		AvatarUrl: req.AvatarUrl,
		UserType:  int(req.UserType),
	}

	wdb := db.WriteDB()
	wdb.NewRecord(userModel)
	return nil
}

func (h *Handle) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq, resp *pb.CommonResp) error {
	user := Model.UserInfo{}
	rdb := db.ReadDB()
	var fields  []string
	for _, attr := range req.AttrNameList {
		fields =append(fields, attr)
	}

	rdb.Select(fields).Where(&Model.UserInfo{Uid:req.UserId}).First(&user)

	data, err := json.Marshal(user)

	if err != nil{
		resp.BaseResp.Code = uint32(basePB.CODE_FAILED)
	}else{
		resp.BaseResp.Code = uint32(basePB.CODE_SUCESS)
	}
	resp.Data = string(data)
	return nil
}
