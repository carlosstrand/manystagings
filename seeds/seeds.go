package seeds

import (
	"github.com/carlosstrand/manystagings/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreatePasswordHash(password string) string {
	res, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(res)
}

type SeedsData struct {
	Users       []models.User
	Environment []models.Environment
}

func DropAll(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&models.User{},
		&models.Environment{},
	)
}

func RunSeeds(db *gorm.DB, data SeedsData) error {
	tx := db.Session(&gorm.Session{SkipDefaultTransaction: true})
	tx.Create(data.Users)
	tx.Create(data.Environment)
	return tx.Error
}
