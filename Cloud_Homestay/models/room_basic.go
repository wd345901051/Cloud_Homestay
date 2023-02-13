package models

import "gorm.io/gorm"

type RoomBasic struct {
	gorm.Model
	Title   string `gorm:"column:title;type:varchar(100);" json:"title"`
	Info    string `gorm:"column:info;type:varchar(100);" json:"info"`
	Content string `gorm:"column:content;type:varchar(100);" json:"content"`
	// 销量
	Sales int `gorm:"column:sales;type:int(10);" json:"sales"`
	Price int `gorm:"column:price;type:int(10);" json:"price"`
	// 房间折扣
	Img string `gorm:"default:"../static/images/room-1.png";column:img;type:varchar(50);" json:"img"`
	// 0-单人间,1-双人间,2-套房
	Type int `gorm:"column:type;type:tinyint(1);" json:"type"`
	// 房间号
	RoomNum string `gorm:"column:room_num;type:varchar(10);" json:"room_num"`
	// 0-未入住,1-已入住
	RoomStatus int `gorm:"column:room_status;type:tinyint(1);" json:"room_status"`
}

func (table *RoomBasic) TableName() string {
	return "room_basic"
}
