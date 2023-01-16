package database

import (
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	_ "github.com/lib/pq"              // PostgreSQL driver
)

func DbConnection() *gorm.DB {

	// Connect to a MySQL database
	db, err := gorm.Open("mysql", "sushil:suchil$123@tcp(127.0.0.1:3306)/shoppingCart")
	if err != nil {
		log.Fatal().Msg("database connection Error")
	}
	log.Info().Msg("connection established")

	// Connect to a PostgreSQL database
	// db, err = sql.Open("postgres", "user=username password=password dbname=dbname sslmode=disable")
	// if err != nil {
	//     panic(err)
	// }
	// defer db.Close()

	// Do something with the PostgreSQL database

	db.AutoMigrate(&Product{}, &Category{}, &Cart{}, &Inventory{})
	return db

}
