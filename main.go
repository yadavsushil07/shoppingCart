package main

import (
	"fmt"

	"github.com/yadavsushil07/shoppingCart/routes"
)

func main() {
	fmt.Println("Welcome to Shopping Cart, onliine shopping platform")
	routes.InitializeRoutes()
}
