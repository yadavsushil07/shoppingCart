package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/rs/zerolog/log"
	"github.com/yadavsushil07/shoppingCart/database"
	"github.com/yadavsushil07/shoppingCart/repository"
)

type ProductHandler interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
	GetProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	AddProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
}

type productHandler struct {
	repo repository.ProductRepository
}

func NewProductHandler() ProductHandler {
	return &productHandler{
		repo: repository.NewProductRepository,
	}
}

func (h *productHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.repo.GetProducts()
	if err != nil {
		responseError(w, http.StatusBadRequest, "url does not exsist")
		log.Error().Msg("Error in GetProducts Handler func ")
		return
	}
	responseJson(w, http.StatusOK, products)
}

func (h *productHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseError(w, http.StatusBadRequest, "url not exsist")
		log.Error().Msg("url not exsist ")
		return
	}
	product, err := h.repo.GetProduct(uint(id))
	if err != nil {
		responseError(w, http.StatusInternalServerError, "record not found")
	}
	responseJson(w, http.StatusOK, product)

}

func (h *productHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseError(w, http.StatusBadRequest, "url not exsist")
		log.Error().Msg("url not exsist ")
		return
	}
	product, err := h.repo.UpdateProduct(uint(id))
	if err != nil {
		responseError(w, http.StatusInternalServerError, "record not found")
	}
	responseJson(w, http.StatusOK, product)
}

func (h *productHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var product database.Product
	data := json.NewDecoder(r.Body)
	if err := data.Decode(&product); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	prod, err := h.repo.AddProduct(product)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
	}
	responseJson(w, http.StatusCreated, prod)
}

func (h *productHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseError(w, http.StatusBadRequest, "url not exsist")
		log.Error().Msg("url not exsist ")
		return
	}
	err = &h.repo.DeleteProduct(uint(id))
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseJson(w, http.StatusOK, "product deleted!")
}
