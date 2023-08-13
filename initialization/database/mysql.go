package database

import (
	"douyin/global"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化mysql数据库
func InitializeMysql() *gorm.DB {
	m := global.SERVER_CONFIG.MySQL
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}

	// 打开数据库连接
	db, err := gorm.Open(mysql.New(mysqlConfig), Gorm.Config(m.Singular))

	// 将引擎设置配置的引擎，并设置每个连接的最大空闲数和最大连接数。
	if err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)

		fmt.Println("====3-gorm====: gorm link mysql success")
		return db
	}
}
