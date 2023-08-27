package dao

import (
	"os"

	"douyin/app/favorite/internal/dal/model"
)

func migration() {
	// 自动迁移模式
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.Favorite{},
		)
	if err != nil {
		os.Exit(0)
	}
}
