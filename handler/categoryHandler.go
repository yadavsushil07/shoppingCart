package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/yadavsushil07/shoppingCart/database"
	"github.com/yadavsushil07/shoppingCart/repository"
)

type CategoryHandler interface {
	AddCategory(w http.ResponseWriter, r *http.Request)
	GetCategories(w http.ResponseWriter, r *http.Request)
	GetCategory(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
}

type CategoryHandlerImpl struct {
	repo repository.CategoryRepository
	log  *zerolog.Logger
}

func NewCategoryHandler(logger *zerolog.Logger) (*CategoryHandlerImpl, error) {
	repo, err := repository.NewCategoryRepository(logger)
	if err != nil {
		return nil, err
	}
	return &CategoryHandlerImpl{
		repo: repo,
		log:  logger,
	}, nil
}

func (h *CategoryHandlerImpl) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.repo.GetCategories()
	if err != nil {
		responseError(w, http.StatusBadRequest, "url does not exsist")
		log.Error().Msg("Error in GetProducts Handler func ")
		return
	}
	responseJson(w, http.StatusOK, categories)
}

func (h *CategoryHandlerImpl) GetCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseError(w, http.StatusBadRequest, "url not exsist")
		log.Error().Msg("url not exsist ")
		return
	}
	category, err := h.repo.GetCategory(uint(id))
	if err != nil {
		responseError(w, http.StatusInternalServerError, "record not found")
	}
	responseJson(w, http.StatusOK, category)

}

func (h *CategoryHandlerImpl) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseError(w, http.StatusBadRequest, "url not exsist")
		log.Error().Msg("url not exsist ")
		return
	}
	category, err := h.repo.UpdateCategory(uint(id))
	if err != nil {
		responseError(w, http.StatusInternalServerError, "record not found")
	}
	responseJson(w, http.StatusOK, category)
}

func (h *CategoryHandlerImpl) AddCategory(w http.ResponseWriter, r *http.Request) {
	var category database.Category
	data := json.NewDecoder(r.Body)
	if err := data.Decode(&category); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	categoryName := category.CategoryName
	categoryresult, err := h.repo.AddCategory(category, categoryName)
	if err != nil {
		h.log.Error().Msg("error in Add category handler level")
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseJson(w, http.StatusCreated, categoryresult)
}

func (h *CategoryHandlerImpl) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseError(w, http.StatusBadRequest, "url not exsist")
		log.Error().Msg("url not exsist ")
		return
	}
	err = h.repo.DeleteCategory(uint(id))
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseJson(w, http.StatusOK, "Category deleted!")
}
