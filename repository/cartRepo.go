package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/yadavsushil07/shoppingCart/database"
)

type CartRepository interface {
	AddToCart(productID uint)
	RemoveFromCart(productID uint)
	// viewCart(CartId uint)
}

type CartRepositoryImpl struct {
	connection *sqlx.DB
}

func NewCartRepository() (*CartRepositoryImpl, error) {
	conn, err := database.DbConnection()
	if err != nil {
		return nil, err
	}

	return &CartRepositoryImpl{
		connection: conn,
	}, nil
}

func (db *CartRepositoryImpl) AddToCart(productID uint) {
	fmt.Println("Add to cart repository")
}

func (db *CartRepositoryImpl) RemoveFromCart(productID uint) {
	fmt.Println("Add to cart repository")
}

// func (db *CartRepositoryImpl) viewCart(cartID uint) {
// 	fmt.Println("Add to cart repository")
// }
