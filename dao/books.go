package dao

import (
	"book-library-api/dto"
	"errors"
)

var BooksList = []dto.Book{
	{Id: "1", Author: "Bruce Banner", Price: 456, Quantity: 3},
	{Id: "2", Author: "Tony Stark", Price: 567, Quantity: 8},
	{Id: "3", Author: "Steve Rogers", Price: 523, Quantity: 6},
	{Id: "4", Author: "Natasha Romanoff", Price: 424, Quantity: 2},
}

func FindBookById(id string) (*dto.Book, error) {
	for i, b := range BooksList {
		if b.Id == id {
			return &BooksList[i], nil
		}
	}
	return nil, errors.New("Book Not found")
}
