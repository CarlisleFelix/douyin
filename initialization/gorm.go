package initialization

import (
	"douyin/global"
	"douyin/initialization/database"
	"douyin/model"
)

func InitializeGorm() {
	switch global.SERVER_CONFIG.Server.DbType {
	case "mysql":
		global.SERVER_DB = database.InitializeMysql()
	case "pgsql":
		global.SERVER_DB = database.InitializegormPgSql()
	default:
		global.SERVER_DB = database.InitializeMysql()
	}
	if err := global.SERVER_DB.AutoMigrate(&model.User{}, &model.Video{}, &model.Comment{}, &model.Relation{}, &model.Favorite{}, &model.Chat{}); err != nil {
		global.SERVER_LOG.Info("automigration failed")
	}
}
