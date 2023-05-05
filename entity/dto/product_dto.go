package dto

type Product struct {
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
