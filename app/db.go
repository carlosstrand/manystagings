package app

import (
	"github.com/carlosstrand/manystagings/models"
	"github.com/go-zepto/zepto/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func createDbDialector() gorm.Dialector {
	dbType := utils.GetEnv("DB_TYPE", "sqlite")
	dbURI := utils.GetEnv("DB_URI", "file::memory:?cache=shared")
	switch dbType {
	case "sqlite":
		return sqlite.Open(dbURI)
	case "postgres":
		return postgres.Open(dbURI)
	case "mysql":
		return mysql.Open(dbURI)
	}
	panic("unknown database type: " + dbType)
}

func AutoMigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Environment{},
		&models.Application{},
		&models.ApplicationEnvVar{},
		&models.User{},
	)
}

func CreateDB() (*gorm.DB, error) {
	db, err := gorm.Open(createDbDialector(), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	AutoMigrateDB(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}
