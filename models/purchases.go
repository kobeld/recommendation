package models

import (
	"github.com/jinzhu/gorm"
)

type Purchase struct {
	gorm.Model
	CustomerID string
	ProductID  string
	Quantity   uint64 `gorm:"default:1"`
}
