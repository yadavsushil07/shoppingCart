package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/yadavsushil07/shoppingCart/database"
	"github.com/yadavsushil07/shoppingCart/repository"
)

type ProductHandler interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
	GetPageAndFilter(w http.ResponseWriter, r *http.Request)
	GetProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	AddProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
}

type ProductHandlerImpl struct {
	repo repository.ProductRepository
	log  *zerolog.Logger
}

func NewProductHandler(logger *zerolog.Logger) (*ProductHandlerImpl, error) {
	repo, err := repository.NewProductRepository(logger)
	if err != nil {
		return nil, err
	}
	return &ProductHandlerImpl{
		repo: repo,
		log:  logger,
	}, err
}

func (h *ProductHandlerImpl) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.repo.GetProducts()
	if err != nil {
		responseError(w, http.StatusBadRequest, "url does not exsist")
		h.log.Error().Msg("Error in GetProducts Handler func ")
		return
	}
	responseJson(w, http.StatusOK, products)
}

func (h *ProductHandlerImpl) GetPageAndFilter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageNo, err := strconv.Atoi(vars["pageNo"])
	fmt.Println(pageNo)
	pageSize := 10
	products, err := h.repo.GetPageAndFilter(pageSize, pageNo)
	if err != nil {
		responseError(w, http.StatusBadRequest, "url does not exsist")
		h.log.Error().Msg("Error in GetProducts Handler func in pagition ")
		return
	}
	responseJson(w, http.StatusOK, products)
}

func (h *ProductHandlerImpl) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseError(w, http.StatusBadRequest, "url not exsist")
		h.log.Error().Msg("url not exsist ")
		return
	}
	product, err := h.repo.GetProduct(uint(id))
	if err != nil {
		responseError(w, http.StatusInternalServerError, "record not found")
	}
	responseJson(w, http.StatusOK, product)

}

func (h *ProductHandlerImpl) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseError(w, http.StatusBadRequest, "url not exsist")
		h.log.Error().Msg("url not exsist ")
		return
	}
	var product database.RequestProduct
	data := json.NewDecoder(r.Body)
	if err := data.Decode(&product); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	err = h.repo.UpdateProduct(product, uint(id))
	if err != nil {
		responseError(w, http.StatusInternalServerError, "record not found")
	}
	responseJson(w, http.StatusOK, nil)
}

func (h *ProductHandlerImpl) AddProduct(w http.ResponseWriter, r *http.Request) {
	var product database.RequestProduct
	data := json.NewDecoder(r.Body)
	if err := data.Decode(&product); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	prod, err := h.repo.AddProduct(product)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
	}
	responseJson(w, http.StatusCreated, prod)
}

func (h *ProductHandlerImpl) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseError(w, http.StatusBadRequest, "url not exsist")
		h.log.Error().Msg("url not exsist ")
		return
	}
	err = h.repo.DeleteProduct(uint(id))
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseJson(w, http.StatusOK, "product deleted!")
}
