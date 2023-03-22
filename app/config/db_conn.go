package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitDb(cfg *Config) *gorm.DB {

	fmt.Println("Trying to connect database :" + cfg.Db.DbName)
	fmt.Println("Trying to connect MYSQL_HOST :" + cfg.Db.Address)

	cfgDb := cfg.Db

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfgDb.User,
		cfgDb.Password,
		cfgDb.Address,
		cfgDb.Port,
		cfgDb.DbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}

	dbSQL.SetConnMaxIdleTime(5 * time.Minute)
	dbSQL.SetMaxOpenConns(20)
	dbSQL.SetConnMaxLifetime(6 * time.Minute)
	dbSQL.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
