package database

type CartProduct struct {
	ProductID   uint   `json:"productId"`
	ProductName string `json:"productName"`
	CategoryID  uint   `json:"categoryId"`
	Price       string `json:"price"`
	InventoryID uint   `json:"inventoryId"`
}

type RequestProduct struct {
	ProductID   uint      `json:"productId,omitempty"`
	ProductName string    `json:"productName,omitempty"`
	Category    Category  `json:"category,omitempty"`
	Price       string    `json:"price,omitempty"`
	Inventory   Inventory `json:"inventory,omitempty"`
}

type ResponseProduct struct {
	ProductID   uint
	ProductName string
	CategoryID  uint
	Price       string
	Quantity    uint
}

type Category struct {
	CategoryName string `json:"categoryName"`
}

type ResponseCategory struct {
	ProductID    uint
	CategoryCode uint   `json:"categoryCode"`
	CategoryName string `json:"categoryName"`
}

type Item struct {
	Product  ResponseProduct `json:"product"`
	Quantity uint            `json:"quantity"`
}

type Cart struct {
	Items    []Item  // cart can have many product
	Subtotal float64 `json:"subtotal"`
}

type Inventory struct {
	ProductID uint `json:"ProductId"`
	Quantity  uint `json:"quantity"`
}

type ResponseInventory struct {
	ProductID       uint
	QuantityInStock uint `json:"quantityInStock"`
}
