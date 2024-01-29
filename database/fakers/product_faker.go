package fakers

import (
	"TokoOnline/apps/models"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func ProductFaker(db *gorm.DB) *models.Product {
	name := faker.Name()
	id := UserFaker(db)
	err := db.Create(&id).Error
	if err != nil {
		fmt.Println(err)
	}

	return &models.Product{
		ID:               uuid.NewString(),
		UserID:           "",
		Sku:              "",
		Name:             name,
		Slug:             slug.Make(name),
		Price:            decimal.NewFromFloat(FakePrice()),
		Stock:            rand.Intn(100),
		Weight:           decimal.NewFromFloat(rand.Float64()),
		ShortDescription: faker.Paragraph(),
		Description:      faker.Paragraph(),
		Status:           0,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        gorm.DeletedAt{},
	}
}

func FakePrice() float64 {
	return precision(rand.Float64()*math.Pow10(rand.Intn(8)), rand.Intn(2)+1)
}

func precision(val float64, pre int) float64 {
	div := math.Pow10(pre)
	return float64(int(val*div)) / div
}
