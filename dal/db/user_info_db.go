package db

import (
	"errors"
	"github.com/dy-platform/user-srv-info/util/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	defaultMgo *mgo.Session
)

func Mgo() *mgo.Session {
	return defaultMgo
}

func Init() {
	dialInfo := &mgo.DialInfo{
		Addrs:     uconfig.DefaultMgoConf.Addr,
		Direct:    false,
		Timeout:   time.Second * 3,
		PoolLimit: uconfig.DefaultMgoConf.PoolLimit,
		Username:  uconfig.DefaultMgoConf.Username,
		Password:  uconfig.DefaultMgoConf.Password,
	}

	ses, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		logrus.Fatalf("dail mgo server error:%v", err)
	}

	ses.SetMode(mgo.Monotonic, true)
	defaultMgo = ses
}

type UserInfo struct {
	UID       int64  `bson:"_id"`
	Nickname  string  `bson:"nickname"`
	Name      string `bson:"name"`
	AvatarURL string `bson:"avatar_url"`
	IDNumber  string `bson:"id_number"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}

var (
	DBUser = "dy_user"
	CUserInfo  = "user_info"
)

func InsertUserInfo(uid int64, nickname, name, avatar, idNumber string) error {
	ses := defaultMgo.Copy()
	if ses == nil {
		logrus.Warnf("mgo session is nil")
		return errors.New("mgo session is nil")
	}
	defer ses.Close()

	now := time.Now().Unix()
	info := &UserInfo{
		UID:       uid,
		Nickname:  nickname,
		Name:      name,
		AvatarURL: avatar,
		IDNumber:  idNumber,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return ses.DB(DBUser).C(CUserInfo).Insert(info)
}

func GetOneUserInfo(uid int64, fields []string) (ret *UserInfo, err error) {
	ses := defaultMgo.Copy()
	if ses == nil {
		logrus.Warnf("mgo session is nil")
		return nil, errors.New("mgo session is nil")
	}
	defer ses.Close()

	query := bson.M{"_id":uid}
	ret = &UserInfo{}
	if len(fields) == 0 {
		err = ses.DB(DBUser).C(CUserInfo).Find(query).One(ret)
	} else {

	}
	if err != nil {
		return nil, err
	}

	return ret, nil
}