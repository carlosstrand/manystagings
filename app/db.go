package app

import (
	"github.com/go-zepto/zepto/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func AutoMigrateDB(db *gorm.DB) {
	db.AutoMigrate()
}

func CreateDB() (*gorm.DB, error) {
	dbURI := utils.GetEnv("DB_URI", "root:root@/manystagings?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	AutoMigrateDB(db)

	if err != nil {
		return nil, err
	}
	return db, nil
}
