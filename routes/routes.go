package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/yadavsushil07/shoppingCart/handler"
)

func InitializeRoutes() {
	r := mux.NewRouter().StrictSlash(true)

	paroductHandler, err := handler.NewProductHandler()
	if err != nil {
		log.Panic().Err(err)
	}

	r.HandleFunc("/products", paroductHandler.GetProducts).Methods("GET")
	r.HandleFunc("/product/{id}", paroductHandler.GetProduct).Methods("GET")
	r.HandleFunc("/addProduct", paroductHandler.AddProduct).Methods("POST")
	r.HandleFunc("/updateProduct/{id}", paroductHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/deleteProduct/{id}", paroductHandler.DeleteProduct).Methods("DELETE")

	categoryHandler, err := handler.NewCategoryHandler()
	if err != nil {
		log.Panic().Err(err)
	}

	r.HandleFunc("/categories", categoryHandler.GetCategories).Methods("GET")
	r.HandleFunc("/addCategory", categoryHandler.GetCategories).Methods("POST")
	r.HandleFunc("/category/{id}", categoryHandler.GetCategories).Methods("GET")
	r.HandleFunc("/updateCategory/{id}", categoryHandler.GetCategories).Methods("PUT")
	r.HandleFunc("/deleteCategory/{id}", categoryHandler.GetCategories).Methods("DELETE")

	cartHandler, err := handler.NewCartHandler()
	if err != nil {
		log.Panic().Err(err)
	}

	r.HandleFunc("/veiwCart", cartHandler.ViewCart).Methods("GET")
	r.HandleFunc("/cart/buy/{productId}", cartHandler.AddToCart)
	r.HandleFunc("/cart/remove/{productId}", cartHandler.RemoveFromCart)

	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal().Msg("ListenAndServer Error")
	}
}
