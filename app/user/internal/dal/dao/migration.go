package dao

import (
	"os"

	"douyin/app/user/internal/dal/model"
)

func migration() {
	// 自动迁移模式
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.User{},
		)
	if err != nil {
		os.Exit(0)
	}
}
