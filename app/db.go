package app

import (
	"errors"

	"github.com/carlosstrand/manystagings/models"
	"github.com/go-zepto/zepto/utils"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func AutoMigrateDB(db *gorm.DB) {
	db.AutoMigrate(
		&models.Environment{},
		&models.Application{},
		&models.ApplicationEnvVar{},
		&models.Config{},
	)
}

// Initialize the config dataset
func initConfigDataset(db *gorm.DB) error {
	var c models.Config
	if err := db.Where("key = ?", "TOKEN").Find(&c).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		tokenConfig := models.Config{
			Key:   "TOKEN",
			Value: uuid.NewString(),
		}
		return db.Create(&tokenConfig).Error
	}
	return nil
}

func CreateDB() (*gorm.DB, error) {
	dbURI := utils.GetEnv("DB_URI", "root:root@/manystagings?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	AutoMigrateDB(db)

	err = initConfigDataset(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
