package fakers

import (
	"TokoOnline/apps/models"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.User {
	return &models.User{
		ID: uuid.NewString(),
		// Addresses:     "",
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		Password:  faker.Password(),
		// RememberToken: faker.RandomInt(),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
}
