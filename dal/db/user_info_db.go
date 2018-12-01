package db

import (
	"fmt"
	"github.com/dy-platform/user-srv-info/util/config"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

type UserPassportDB struct {
	rClient *gorm.DB
	wClient *gorm.DB

	sync.RWMutex
}

var (
	db = &UserPassportDB{}
	dbArgsFormat = "%s:%s@tcp(%s)/%s?timeout=%s&readTimeout=%s&writeTimeout=%s&maxAllowedPacket=536870912"
)

func DBInit() {
	//logrus.Infof()
	var err error
	rArgs := fmt.Sprintf(dbArgsFormat,
		uconfig.DefaultConfig.UserInfoDB.User,
		uconfig.DefaultConfig.UserInfoDB.Password,
		uconfig.DefaultConfig.UserInfoDB.ReadAddress,
		uconfig.DefaultConfig.UserInfoDB.DBName,
		uconfig.DefaultConfig.UserInfoDB.ReadTimeout,
		uconfig.DefaultConfig.UserInfoDB.WriteTimeout)

	db.rClient, err = gorm.Open("mysql", rArgs)
	if err != nil {
		logrus.Warnf("open read-mysqldb failed. args:%s", rArgs)
		os.Exit(1)
	}

	wArgs := fmt.Sprintf(dbArgsFormat,
		uconfig.DefaultConfig.UserInfoDB.User,
		uconfig.DefaultConfig.UserInfoDB.Password,
		uconfig.DefaultConfig.UserInfoDB.ReadAddress,
		uconfig.DefaultConfig.UserInfoDB.DBName,
		uconfig.DefaultConfig.UserInfoDB.ReadTimeout,
		uconfig.DefaultConfig.UserInfoDB.WriteTimeout)

	db.wClient, err = gorm.Open("mysql", wArgs)
	if err != nil {
		logrus.Warnf("open write-mysqldb failed. args:%s", wArgs)
		os.Exit(1)
	}

	// TODO PING

	db.wClient.DB().SetMaxIdleConns(10)
	db.wClient.DB().SetMaxOpenConns(20)


}

func WriteDB()*gorm.DB{
	return db.wClient
}


func ReadDB()*gorm.DB{
	return db.rClient
}