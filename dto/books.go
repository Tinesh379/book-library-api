package dto

type Book struct {
	Id       string `json: "id"`
	Author   string `json: "author"`
	Price    uint   `json: "price"`
	Quantity int    `json: "quantity"`
}
