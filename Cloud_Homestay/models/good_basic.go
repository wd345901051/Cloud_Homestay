package models

import "gorm.io/gorm"

type GoodBasic struct {
	gorm.Model
	GoodName    string `gorm:"column:good_name;type:varchar(10);" json:"good_name"`
	GoodPrice   int    `gorm:"column:good_price;type:int(10);" json:"good_price"`
	GoodNum     int    `gorm:"column:good_num;type:int(10);" json:"good_num"`
	GoodContent string `gorm:default:"无";"column:good_content;type:varchar(20);" json:"good_content"`
}

func (table *GoodBasic) TableName() string {
	return "good_basic"
}
