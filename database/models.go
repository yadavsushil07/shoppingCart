package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	ProductName string   `json:"productName"`
	Category    Category `gorm:"foreignKey:CategoryID"`
	CategoryID  uint
	Price       string `json:"price"`
	Status      bool   `json:"status"`
}

type Category struct {
	gorm.Model
	CategoryName string `json:"categoryName"`
}

type Cart struct {
	gorm.Model
	Product         []Product `gorm:"foreignKey:ProductID"` // cart can have many product
	ProductID       uint
	ProductQuantity uint    `json:"productQuantity"`
	Subtotal        float64 `json:"subtotal"`
}

type Inventory struct {
	gorm.Model
	Product         []Product `gorm:"foreignKey:ProductID"` // there can be many products in inventory
	ProductID       uint
	QuantityInStock uint `json:"quantityInStock"`
}
