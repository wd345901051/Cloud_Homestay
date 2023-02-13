package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	// 用户的唯一标识
	Identity   string `gorm:"column:identity;type:varchar(36);" json:"identity"`
	Name       string `gorm:"column:name;type:varchar(100);" json:"name"`
	Icon       string `gorm:"column:icon;type:varchar(50);" json:"icon"`
	Email      string `gorm:"column:email;type:varchar(100);" json:"email"`
	Password   string `gorm:"column:password;type:varchar(32);" json:"password"`
	Phone      string `gorm:"column:phone;type:varchar(20);" json:"phone"`
	Birthday   string `gorm:"column:birthday;type:varchar(50);" json:"birthday"`
	Attention  int    `gorm:"column:attention;type:tinyint(1);" json:"attention"`
	Isadmin    int    `gorm:"column:isadmin;type:tinyint(1);" json:"isadmin"`
	UserStatus int    `gorm:"column:user_status;type:tinyint(1);" json:"user_status"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
