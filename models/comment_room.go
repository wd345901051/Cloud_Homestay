package models

import "gorm.io/gorm"

type CommentRoom struct {
	gorm.Model
	// 发表评论的房间
	RoomId string `gorm:"column:room_id;type:int(11)" json:"room_id"`
	// 评论的内容
	Content  string     `gorm:"column:content;type:varchar(200);" json:"content"`
	UserId   string     `gorm:"column:user_id;type:int(11)" json:"user_id"`
	UserIcon string     `gorm:"column:user_icon;type:varchar(50)" json:"user_icon"`
	RoomInfo *RoomBasic `gorm:"foreignKey:id;references:room_id"`
	UserInfo *UserBasic `gorm:"foreignKey:id;references:user_id"`
}

func (table *CommentRoom) TableName() string {
	return "comment_room"
}
