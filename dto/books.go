package dto

type Book struct {
	Id       int    `json: "id"`
	Author   string `json: "author"`
	Price    uint   `json: "price"`
	Quantity int    `json: "quantity"`
}

var BooksList = []Book{
	{Id: 1, Author: "Bruce Banner", Price: 456, Quantity: 3},
	{Id: 2, Author: "Tony Stark", Price: 567, Quantity: 8},
	{Id: 3, Author: "Steve Rogers", Price: 523, Quantity: 6},
	{Id: 4, Author: "Natasha Romanoff", Price: 424, Quantity: 2},
}
