package routes

import (
	"github.com/gorilla/mux"
	"github.com/yadavsushil07/shoppingCart/handler"
)

func Routes() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/products", handler.NewProductHandler().GetProducts).Methods("GET")
	r.HandleFunc("/product/{id}", handler.NewProductHandler().GetProduct).Methods("GET")
	r.HandleFunc("/addProduct", handler.NewProductHandler().AddProduct).Methods("POST")
	r.HandleFunc("/updateProduct/{id}", handler.NewProductHandler().UpdateProduct).Methods("PUT")
	r.HandleFunc("/deleteProduct/{id}", handler.NewProductHandler().DeleteProduct).Methods("DELETE")
	// Get All Product
	// 	r.HandleFunc("/allProducts", func(w http.ResponseWriter, r *http.Request) {
	// 		db := database.DbConnection()
	// 		rows, err := db.Query("select * from Product")
	// 		if err != nil {
	// 			log.Fatal().Msg("Query error")
	// 		}
	// 		var products []database.Product
	// 		for rows.Next() {
	// 			var Id int
	// 			var Name, category, price string
	// 			err := rows.Scan(&Id, &Name, &category, &price)
	// 			if err != nil {
	// 				log.Error().Msg("Row sacnning error")
	// 			}
	// 			products = append(products, database.Product{ProductId: Id, ProductName: Name, Category: category, Price: price})
	// 		}
	// 		json.NewEncoder(w).Encode(products)
	// 		log.Info().Msg("Inside Handlerfunc ")
	// 		db.Close()
	// 	}).Methods("GET")

	// 	// Get certain Product
	// 	r.HandleFunc("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 		var product database.Product
	// 		param := mux.Vars(r)
	// 		id, err := strconv.Atoi(param["id"])

	// 		if err != nil {
	// 			log.Error().Msg("parsing string to int Error")
	// 		}
	// 		db := database.DbConnection()
	// 		defer db.Close()
	// 		row, err := db.Query("select * from Product where productId = ?", id)
	// 		if err != nil {
	// 			log.Error().Msg("DB Query Error")
	// 		}
	// 		for row.Next() {
	// 			err = row.Scan(&product.ProductId, &product.ProductName, &product.Category, &product.Price)
	// 			if err != nil {
	// 				log.Error().Msg("DB scaning Error")
	// 			}
	// 		}
	// 		json.NewEncoder(w).Encode(product)
	// 	}).Methods("GET")

	// 	// Create Product
	// 	r.HandleFunc("/createProduct", func(w http.ResponseWriter, r *http.Request) {
	// 		db := database.DbConnection()
	// 		row, err := db.Prepare("insert into Product(productId,productName,category,price) values(?,?,?,?)")
	// 		body, err := ioutil.ReadAll(r.Body)
	// 		data := make(map[string]interface{})
	// 		json.Unmarshal(body, &data)
	// 		Id := data["productId"]
	// 		name := data["productName"]
	// 		category := data["category"]
	// 		price := data["price"]
	// 		_, err = row.Exec(Id, name, category, price)
	// 		if err != nil {
	// 			log.Error().Msg("Error in executing Query")
	// 		}
	// 	}).Methods("POST")
	// 	err := http.ListenAndServe(":8000", r)
	// 	if err != nil {
	// 		log.Fatal().Msg("ListenAndServer Error")
	// 	}
}
