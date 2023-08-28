package dao

import (
	"douyin/app/video/internal/dal/model"
	"os"
)

func migration() {
	// 自动迁移模式
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.Video{},
		)
	if err != nil {
		os.Exit(0)
	}
}
