package models

import "gorm.io/gorm"

type EquipBasic struct {
	gorm.Model
	EquipName   string `gorm:"column:equip_name;type:varchar(10);" json:"equip_name"`
	EquipTime   string `gorm:"column:equip_time;type:varchar(10);" json:"equip_time"`
	EquipSolve  string `gorm:"column:equip_solve;type:varchar(10);" json:"equip_solve"`
	EquipStatus int    `gorm:"column:equip_status;type:tinyint(1);" json:"equip_status"`
}

func (table *EquipBasic) TableName() string {
	return "equip_basic"
}
