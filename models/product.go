package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(60); not null"`
	Price float32 `gorm:"type:decimal(8,10); not null"`
	Stock bool    `gorm:"type:tiny; not null"`
}

func (product *Product) Validate() error {
	if strings.TrimSpace(product.Name) == "" {
		return errors.New("name is required and must not be blank")
	}

	if product.Price == 0 {
		return errors.New("price is required and must not be blank")
	}

	return nil
}
