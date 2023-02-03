package repository

import (
	"database/sql"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/jmoiron/sqlx"
	"github.com/yadavsushil07/shoppingCart/database"
)

type ProductRepository interface {
	GetProducts() ([]database.ResponseProduct, error)
	GetPageAndFilter(pageSize int, pagenumber int) (products []database.ResponseProduct, err error)
	GetProduct(uint) (database.ResponseProduct, error)
	UpdateProduct(database.RequestProduct, uint) error
	AddProduct(product database.RequestProduct) (database.RequestProduct, error)
	DeleteProduct(uint) error
}

type ProductRepositoryImpl struct {
	connection *sqlx.DB
	log        *zerolog.Logger
}

func NewProductRepository(logger *zerolog.Logger) (*ProductRepositoryImpl, error) {
	conn, err := database.DbConnection()
	if err != nil {
		return nil, err
	}
	return &ProductRepositoryImpl{
		connection: conn,
		log:        logger,
	}, err
}

func (db *ProductRepositoryImpl) GetProducts() (products []database.ResponseProduct, err error) {
	var product database.ResponseProduct
	row, err := db.connection.Query("SELECT product.productId, product.productName, product.categoryId,product.price,inventory.quantity FROM product INNER JOIN inventory ON inventory.productID=product.productId")
	if err != nil {
		db.log.Err(err)
		return nil, err
	} else {
		for row.Next() {
			row.Scan(&product.ProductID, &product.ProductName, &product.CategoryID, &product.Price, &product.Quantity)
			products = append(products, product)
		}
	}
	db.log.Info().Msg("Fetched all product record repo layer")
	return products, nil
}

func (db *ProductRepositoryImpl) GetProduct(id uint) (product database.ResponseProduct, err error) {
	row, err := db.connection.Query("SELECT product.productId, product.productName, product.categoryId,product.price,inventory.quantity FROM product INNER JOIN inventory ON inventory.productID=product.productId where product.productId = ?", id)
	if err != nil {
		db.log.Err(err)
		return product, err
	} else {
		for row.Next() {
			row.Scan(&product.ProductID, &product.ProductName, &product.CategoryID, &product.Price, &product.Quantity)
		}
	}
	if product.ProductID == 0 {
		db.log.Error().Msg("product does not exist")
		return
	}
	db.log.Info().Msg("fetch product by Id repo layer")
	return product, nil
}

func (db *ProductRepositoryImpl) AddProduct(product database.RequestProduct) (database.RequestProduct, error) {
	tx := db.connection.MustBegin()
	res := tx.MustExec("insert into category(categoryName) select * from (select ?) as tmp where not exists (select categoryName from category where categoryName = ?) limit 1", product.Category.CategoryName, product.Category.CategoryName)
	id, err := res.LastInsertId()
	if err != nil {
		db.log.Err(err)
	}
	var resproduct sql.Result
	if id == 0 {
		var categoryId int64
		tx.QueryRow("select categoryId from category where categoryName = ?", product.Category.CategoryName).Scan(&categoryId)
		resproduct = tx.MustExec("insert into product(productName,categoryId,price) values (?,?,?)", product.ProductName, categoryId, product.Price)
	} else {
		resproduct = tx.MustExec("insert into product(productName,categoryId,price) values (?,?,?)", product.ProductName, id, product.Price)
	}
	productId, err := resproduct.LastInsertId()
	if err != nil {
		db.log.Err(err)
	}
	tx.MustExec("insert into inventory(productId,quantity) values (?,?)", productId, product.Inventory.Quantity)
	tx.Commit()
	db.log.Info().Msg("add product successfully from repo layer")
	return product, nil
}

func (db *ProductRepositoryImpl) UpdateProduct(product database.RequestProduct, id uint) error {
	tx := db.connection.MustBegin()
	if product.ProductName != "" {
		tx.MustExec("update product set productName = ? where productId = ?", product.ProductName, id)
	}
	if product.Inventory.Quantity != 0 {
		tx.MustExec("update inventory set quantity = ? where productId = ?", product.Inventory.Quantity, id)
	}
	err := tx.Commit()
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	return nil
}

func (db *ProductRepositoryImpl) DeleteProduct(id uint) error {
	return nil
}

func (db *ProductRepositoryImpl) GetPageAndFilter(pageSize int, pagenumber int) (products []database.ResponseProduct, err error) {
	var product database.ResponseProduct
	offset := (pagenumber - 1) * pageSize
	row, err := db.connection.Query("SELECT product.productId, product.productName, product.categoryId,product.price,inventory.quantity FROM product INNER JOIN inventory ON inventory.productID=product.productId limit ? offset ?", pageSize, offset)
	if err != nil {
		db.log.Err(err)
		return nil, err
	} else {
		for row.Next() {
			row.Scan(&product.ProductID, &product.ProductName, &product.CategoryID, &product.Price, &product.Quantity)
			products = append(products, product)
		}
	}
	db.log.Info().Msg(fmt.Sprintf("Fetched product from pageNo %d record repo layer", pagenumber))
	return products, nil
}
