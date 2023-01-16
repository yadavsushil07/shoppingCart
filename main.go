package main

import (
	"fmt"
	"time"

	"github.com/yadavsushil07/shoppingCart/routes"
)

func main() {
	fmt.Println("Welcome to Shopping Cart, onliine shopping platform")
	routes.Routes()
	time.Sleep(10 * time.Second)
}
