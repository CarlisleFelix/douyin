package database

import (
	"douyin/global"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GormPgSql 初始化 Postgresql 数据库
func InitializegormPgSql() *gorm.DB {
	p := global.SERVER_CONFIG.PgSQL

	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}

	db, err := gorm.Open(postgres.New(pgsqlConfig), Gorm.Config(p.Singular))

	if err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)

		fmt.Println("====3-gorm====: gorm link PostgreSQL success")

		return db
	}
}
