package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/rs/zerolog"
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
	log  *zerolog.Logger
}

var store = sessions.NewCookieStore([]byte("sushil"))

func NewCartHandler(logger *zerolog.Logger) (*CartHandlerImpl, error) {
	repo, err := repository.NewCartRepository(logger)
	if err != nil {
		return nil, err
	}
	return &CartHandlerImpl{
		repo: repo,
		log:  logger,
	}, err
}

func (ch *CartHandlerImpl) VeiwCart(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "sushil")
	if session.Values["cart"] == nil {
		ch.log.Error().Msg("cart is empty")
		responseError(w, http.StatusBadRequest, "cart is empty")
		return
	}
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
		ch.log.Error().Msg("url not exsist ")
		return
	}
	item, err := ch.repo.AddToCart(uint(id))
	if err != nil {
		responseError(w, http.StatusBadRequest, "product not found")
		ch.log.Error().Msg("product not found")
		return
	}
	session, _ := store.Get(r, "sushil")
	session.Options.MaxAge = 5 * 60
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
		ch.log.Error().Msg("url not exsist ")
		return
	}
	session, _ := store.Get(r, "sushil")
	strCart := session.Values["cart"].(string)
	var cart []database.Item
	json.Unmarshal([]byte(strCart), &cart)

	index := exists(uint(id), cart)
	if index >= 0 && cart[index].Quantity > 1 {
		cart[index].Quantity--
	} else if index >= 0 {
		cart = remove(cart, index)
	} else {
		responseError(w, http.StatusBadRequest, "Item not in the cart")
	}

	bytesCart, _ := json.Marshal(cart)
	session.Values["cart"] = string(bytesCart)
	session.Save(r, w)
	responseJson(w, http.StatusOK, "product removed from the cart successfully")
}
