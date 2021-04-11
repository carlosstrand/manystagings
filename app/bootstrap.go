package app

import (
	"github.com/carlosstrand/manystagings/models"
	"golang.org/x/crypto/bcrypt"
)

func createPasswordHash(password string) string {
	res, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(res)
}

// Initialize the root User
func (a *App) setupFirstUser() error {
	var adminCount int64
	if err := a.DB.Model(&models.User{}).Count(&adminCount).Error; err != nil {
		return err
	}
	if adminCount == 0 {
		user := models.User{
			Username:     "root",
			PasswordHash: createPasswordHash("root"),
		}
		a.DB.Create(&user)
	}

	return nil
}

// First-Time initialization for the server
func (a *App) bootstrap() error {
	if err := a.setupFirstUser(); err != nil {
		return err
	}
	return nil
}
