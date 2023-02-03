package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/yadavsushil07/shoppingCart/database"
)

type CartRepository interface {
	AddToCart(productID uint) (database.CartProduct, error)
	RemoveFromCart(productID uint)
	// viewCart(CartId uint)
}

type CartRepositoryImpl struct {
	connection *sqlx.DB
	log        *zerolog.Logger
}

func NewCartRepository(logger *zerolog.Logger) (*CartRepositoryImpl, error) {
	conn, err := database.DbConnection()
	if err != nil {
		return nil, err
	}

	return &CartRepositoryImpl{
		connection: conn,
		log:        logger,
	}, nil
}

func (db *CartRepositoryImpl) AddToCart(productID uint) (database.CartProduct, error) {
	var cartProduct database.CartProduct
	prod, _ := NewProductRepository(db.log)
	item, err := prod.GetProduct(productID)
	if err != nil {
		return database.CartProduct{}, err
	}
	cartProduct.ProductID = item.ProductID
	cartProduct.Price = item.Price
	cartProduct.ProductName = item.ProductName
	return cartProduct, nil
}

func (db *CartRepositoryImpl) RemoveFromCart(productID uint) {
	fmt.Println("Add to cart repository")
}

// func (db *CartRepositoryImpl) viewCart(cartID uint) {
// 	fmt.Println("Add to cart repository")
// }
