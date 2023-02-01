package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/rs/zerolog/log"
	"github.com/yadavsushil07/shoppingCart/database"
	"github.com/yadavsushil07/shoppingCart/repository"
)

type cartHandler interface {
	// GetCarts(w http.ResponseWriter, r *http.Request)
	ViewCart(w http.ResponseWriter, r *http.Request)
	AddToCart(w http.ResponseWriter, r *http.Request)
	RemoveFromCart(w http.ResponseWriter, r *http.Request)
}

type CartHandlerImpl struct {
	repo repository.CartRepository
}

var store = sessions.NewCookieStore([]byte("sushilyadav"))

func NewCartHandler() (*CartHandlerImpl, error) {
	repo, err := repository.NewCartRepository()
	if err != nil {
		return nil, err
	}
	return &CartHandlerImpl{
		repo: repo,
	}, err
}

func (ch *CartHandlerImpl) ViewCart(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "sushilyadav")
	strCart := session.Values["cart"].(string)
	var cart []database.Item
	json.Unmarshal([]byte(strCart), &cart)

	data := map[string]interface{}{
		"cart":  cart,
		"total": total(cart),
	}

	responseJson(w, http.StatusOK, data)
}

func (ch *CartHandlerImpl) AddToCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["productId"])
	if err != nil {
		responseError(w, http.StatusBadRequest, "url not exsist")
		log.Error().Msg("url not exsist ")
		return
	}
	prod, _ := repository.NewProductRepository()
	item, err := prod.GetProduct(uint(id))
	if err != nil {
		responseError(w, http.StatusBadRequest, "product not found")
		log.Error().Msg("product not found")
		return
	}
	session, _ := store.Get(r, "sushilyadav")
	cart := session.Values["cart"]

	if cart == nil {
		var cart []database.Item
		cart = append(cart, database.Item{
			Product:  item,
			Quantity: 1,
		})
		bytesCart, _ := json.Marshal(cart)
		session.Values["cart"] = string(bytesCart)
	} else {
		strCart := session.Values["cart"].(string)
		var cart []database.Item
		json.Unmarshal([]byte(strCart), &cart)

		index := exists(uint(id), cart)
		if index == -1 {
			cart = append(cart, database.Item{
				Product:  item,
				Quantity: 1,
			})
		} else {
			cart[index].Quantity++
		}
		bytesCart, _ := json.Marshal(cart)
		session.Values["cart"] = string(bytesCart)
	}
	session.Save(r, w)
	responseJson(w, http.StatusOK, "product add to the cart successfully")
}

func exists(id uint, cart []database.Item) int {
	for i := 0; i < len(cart); i++ {
		if cart[i].Product.ProductID == id {
			return i
		}
	}
	return -1
}

func total(cart []database.Item) float64 {
	var totalPrice float64
	for _, i := range cart {
		price, _ := strconv.ParseFloat(i.Product.Price, 64)
		totalPrice += price * float64(i.Quantity)
	}
	return totalPrice
}

func remove(cart []database.Item, index int) []database.Item {
	copy(cart[index:], cart[index+1:])
	return cart[:len(cart)-1]
}

func (ch *CartHandlerImpl) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["productId"])
	if err != nil {
		responseError(w, http.StatusBadRequest, "url not exsist")
		log.Error().Msg("url not exsist ")
		return
	}
	session, _ := store.Get(r, "sushilyadav")
	strCart := session.Values["cart"].(string)
	var cart []database.Item
	json.Unmarshal([]byte(strCart), &cart)

	index := exists(uint(id), cart)
	cart = remove(cart, index)

	bytesCart, _ := json.Marshal(cart)
	session.Values["cart"] = string(bytesCart)
	session.Save(r, w)
	responseJson(w, http.StatusOK, "product removed from the cart successfully")
}
