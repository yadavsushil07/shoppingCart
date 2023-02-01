package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/yadavsushil07/shoppingCart/database"
)

type CategoryRepository interface {
	GetCategories() ([]database.Category, error)
	GetCategory(uint) (database.Category, error)
	UpdateCategory(uint) (database.Category, error)
	AddCategory(category database.Category, categoryName string) (database.Category, error)
	DeleteCategory(uint) error
}

type CategoryRepositoryImpl struct {
	connection *sqlx.DB
}

func NewCategoryRepository() (*CategoryRepositoryImpl, error) {
	conn, err := database.DbConnection()
	if err != nil {
		return nil, err
	}

	return &CategoryRepositoryImpl{
		connection: conn,
	}, nil
}

func (db *CategoryRepositoryImpl) GetCategories() (categorys []database.Category, err error) {
	return categorys, nil
}

func (db *CategoryRepositoryImpl) GetCategory(id uint) (category database.Category, err error) {
	return category, nil
}

func (db *CategoryRepositoryImpl) AddCategory(category database.Category, categoryName string) (database.Category, error) {
	return category, nil
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
