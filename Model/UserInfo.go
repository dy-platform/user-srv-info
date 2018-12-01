package Model


type UserInfo struct{
	Uid int64 `gorm:"unique_index" json:"uid"`
	Nick string `gorm:"size:255" json:"nick"`
	Gender string `gorm:"size:32" json:"gender"`
	AvatarUrl string `gorm:"size:1024" json:"avatarUrl"`
	UserType int `gorm:"type:smallint" json:"userType"`
}



