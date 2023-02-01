package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/yadavsushil07/shoppingCart/config"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	_ "github.com/lib/pq"              // PostgreSQL driver
)

const (
	categoryTable = `
CREATE TABLE if not exists category (
	CategoryID int NOT NULL AUTO_INCREMENT,
	CategoryName varchar(255) not null,
	PRIMARY KEY (CategoryID),
	UNIQUE (CategoryName)
);
`
	productTable = `

CREATE TABLE if not exists product (
    ProductID int NOT NULL AUTO_INCREMENT,
    ProductName varchar(255) NOT NULL,
	CategoryID int NOT NULL,
	Price varchar(255),
	PRIMARY KEY (ProductID),
	FOREIGN KEY (CategoryID) REFERENCES category(CategoryID)
);
`

	inventoryTable = `CREATE TABLE if not exists inventory (
    InventoryID int NOT NULL AUTO_INCREMENT,
    ProductID int NOT NULL,
    Quantity int not null,
    Primary KEY (InventoryID),
	FOREIGN KEY (ProductID) REFERENCES product(ProductID)
);
`
)

func DbConnection() (*sqlx.DB, error) {
	env := config.NewEnv()
	var dbURI string
	switch env.DB_DRIVER {
	case "mysql":
		dbURI = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True", env.DB_USER, env.DB_PASS, env.DB_HOST, env.DB_PORT, env.DB_NAME, env.DB_CHARSET)
	case "postgres":
		dbURI = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", env.DB_HOST, env.DB_PORT, env.DB_USER, env.DB_NAME, env.DB_PASS)
	default:
		return nil, fmt.Errorf("invalid database driver")
	}

	db, err := sqlx.Connect(env.DB_DRIVER, dbURI)
	if err != nil {
		log.Panic().Err(err).Msg("error connecting to database")
	}

	db.MustExec(categoryTable)
	db.MustExec(productTable)
	db.MustExec(inventoryTable)

	return db, nil
}
