package routes

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/yadavsushil07/shoppingCart/handler"
)

func InitializeRoutes() {
	r := mux.NewRouter().StrictSlash(true)
	log := zerolog.New(os.Stdout)

	productHandler, err := handler.NewProductHandler(&log)
	if err != nil {
		log.Panic().Err(err)
	}

	r.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	r.HandleFunc("/products/{pageNo}", productHandler.GetPageAndFilter).Methods("GET")
	r.HandleFunc("/product/{id}", productHandler.GetProduct).Methods("GET")
	r.HandleFunc("/addProduct", productHandler.AddProduct).Methods("POST")
	r.HandleFunc("/updateProduct/{id}", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/deleteProduct/{id}", productHandler.DeleteProduct).Methods("DELETE")

	categoryHandler, err := handler.NewCategoryHandler(&log)
	if err != nil {
		log.Panic().Err(err)
	}

	r.HandleFunc("/categories", categoryHandler.GetCategories).Methods("GET")
	r.HandleFunc("/addCategory", categoryHandler.AddCategory).Methods("POST")
	r.HandleFunc("/category/{id}", categoryHandler.GetCategory).Methods("GET")
	r.HandleFunc("/updateCategory/{id}", categoryHandler.UpdateCategory).Methods("PUT")
	r.HandleFunc("/deleteCategory/{id}", categoryHandler.DeleteCategory).Methods("DELETE")

	cartHandler, err := handler.NewCartHandler(&log)
	if err != nil {
		log.Panic().Err(err)
	}

	r.HandleFunc("/veiwCart", cartHandler.VeiwCart).Methods("GET")
	r.HandleFunc("/cart/buy/{productId}", cartHandler.AddToCart)
	r.HandleFunc("/cart/remove/{productId}", cartHandler.RemoveFromCart)

	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal().Msg("ListenAndServer Error")
	}
}
