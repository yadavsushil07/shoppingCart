package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yadavsushil07/shoppingCart/database"
)

type ProductRepository interface {
	GetProducts() ([]database.Product, error)
	GetProduct(uint) (database.Product, error)
	UpdateProduct(uint) (database.Product, error)
	AddProduct(database.Product) (database.Product, error)
	DeleteProduct(uint) error
}

type productRepository struct {
	connection *gorm.DB
}

func NewProductRepository() ProductRepository {
	return &productRepository{
		connection: database.DbConnection(),
	}
}

func (db *productRepository) GetProducts() (products []database.Product, err error) {
	return products, db.connection.Find(&products).Error
}

func (db *productRepository) GetProduct(id uint) (product database.Product, err error) {
	return product, db.connection.First(&product, id).Error
}

func (db *productRepository) AddProduct(product database.Product) (database.Product, error) {
	return product, db.connection.Create(&product).Error
}

func (db *productRepository) UpdateProduct(id uint) (database.Product, error) {
	var product database.Product
	if err := db.connection.First(&product, id).Error; err != nil {
		return product, err
	}
	return product, db.connection.Model(&product).Updates(&product).Error
}

func (db *productRepository) DeleteProduct(id uint) error {
	var product database.Product
	if err := db.connection.First(&product, id).Error; err != nil {
		return err
	}
	return db.connection.Delete(&product).Error
}

// func GetAllProduct() []database.Product {
// 	db := database.DbConnection()
// 	defer db.Close()
// 	rows, err := db.Query("select * from Product")
// 	if err != nil {
// 		log.Fatal().Msg("Query error")
// 	}
// 	var products []database.Product
// 	for rows.Next() {
// 		var Id int
// 		var Name, category, price string
// 		err := rows.Scan(&Id, &Name, &category, &price)
// 		if err != nil {
// 			log.Error().Msg("Row sacnning error")
// 		}
// 		products = append(products, database.Product{ProductId: Id, ProductName: Name, Category: category, Price: price})
// 	}

// 	return products
// }

// func GetProcductById(id int) database.Product {
// 	db := database.DbConnection()
// 	defer db.Close()
// 	var name, category, price string
// 	row, err := db.Query("select productName,category,price from Product where productId = ?", id)
// 	if err != nil {
// 		log.Error().Msg("DB Query Error")
// 	}
// 	err = row.Scan(&name, &category, &price)
// 	if err != nil {
// 		log.Error().Msg("DB scaning Error")
// 	}
// 	product := database.Product{
// 		ProductId:   id,
// 		ProductName: name,
// 		Category:    category,
// 		Price:       price,
// 	}
// 	return product
// }
