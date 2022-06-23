package models

import (
	"gorm.io/gorm"
)

type OrderBasic struct {
	gorm.Model
	// 订单编号
	OrderNo string `gorm:"column:order_no;type:varchar(36)" json:"order_no"`
	// 订单状态，0-已付款，1交易结束
	OrderStatus int        `gorm:"column:order_status;type:tinyint(1)" json:"order_status"`
	UserId      int        `gorm:"column:user_id;type:int(11)" json:"user_id"`
	UserName    string     `gorm:"column:user_name;type:varchar(10)" json:"user_name"`
	UserPhone   string     `gorm:"column:user_phone;type:varchar(15)" json:"user_phone"`
	UserNum     string     `gorm:"column:user_num;type:varchar(20)" json:"user_num"`
	PeopleCnt   int        `gorm:"column:people_cnt;type:tinyint(1)" json:"people_cnt"`
	RoomId      int        `gorm:"column:room_id;type:int(11)" json:"room_id"`
	RoomIn      string     `gorm:"column:room_in;type:varchar(20)" json:"room_in"`
	RoomOut     string     `gorm:"column:room_out;type:varchar(20)" json:"room_out"`
	Content     string     `gorm:"column:content;type:varchar(50)" json:"content"`
	RoomExit    int        `gorm:"column:room_exit;type:tinyint(1)" json:"room_exit"`
	RoomInfo    *RoomBasic `gorm:"foreignKey:id;references:room_id"`
	UserInfo    *UserBasic `gorm:"foreignKey:id;references:user_id"`
}

func (table *OrderBasic) TableName() string {
	return "order_basic"
}
