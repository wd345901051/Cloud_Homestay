package models

import "fmt"

// 迁移函数
func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&RoomBasic{},
		)
	err = DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&UserBasic{},
		)
	err = DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&CommentRoom{},
		)
	err = DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&EquipBasic{},
		)
	err = DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&GoodBasic{},
		)
	err = DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&OrderBasic{},
		)
	if err != nil {
		fmt.Println("migration err", err)
	}
}
