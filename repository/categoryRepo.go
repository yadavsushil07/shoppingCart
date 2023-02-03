package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/yadavsushil07/shoppingCart/database"
)

type CategoryRepository interface {
	GetCategories() ([]database.Category, error)
	GetCategory(uint) (database.Category, error)
	UpdateCategory(uint) (database.Category, error)
	AddCategory(category database.Category, categoryName string) (bool, error)
	DeleteCategory(uint) error
}

type CategoryRepositoryImpl struct {
	connection *sqlx.DB
	log        *zerolog.Logger
}

func NewCategoryRepository(logger *zerolog.Logger) (*CategoryRepositoryImpl, error) {
	conn, err := database.DbConnection()
	if err != nil {
		return nil, err
	}

	return &CategoryRepositoryImpl{
		connection: conn,
		log:        logger,
	}, nil
}

func (db *CategoryRepositoryImpl) GetCategories() (categorys []database.Category, err error) {
	var category database.Category
	rows, err := db.connection.Query("select * from category")
	if err != nil {
		db.log.Err(err)
	} else {
		for rows.Next() {
			rows.Scan(&category.CategoryName)
			categorys = append(categorys, category)
		}
	}
	return categorys, nil
}

func (db *CategoryRepositoryImpl) GetCategory(id uint) (category database.Category, err error) {
	row, err := db.connection.Query("select * from category where categoryId = ?", id)
	if err != nil {
		db.log.Err(err)
	} else {
		for row.Next() {
			row.Scan(&category.CategoryName)
		}

	}
	return category, nil
}

func (db *CategoryRepositoryImpl) AddCategory(category database.Category, categoryName string) (bool, error) {
	_, err := db.connection.Query("insert ignore into category(categoryName) values(?)", category.CategoryName)
	if err != nil {
		db.log.Err(err)
		return false, err
	}
	return true, nil
}

func (db *CategoryRepositoryImpl) UpdateCategory(id uint) (database.Category, error) {
	var category database.Category
	// if err := db.connection.First(&category, id).Error; err != nil {
	// 	return category, err
	// }
	return category, nil
}

func (db *CategoryRepositoryImpl) DeleteCategory(id uint) error {
	// var category database.Category
	// if err := db.connection.First(&category, id).Error; err != nil {
	// 	return err
	// }
	return nil
}
