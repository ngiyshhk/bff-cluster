package model

type Book struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var BooksMock = []Book{
	{ID: "1", Name: "user1", Price: 1000},
	{ID: "2", Name: "user2", Price: 1100},
	{ID: "3", Name: "user3", Price: 1200},
}
