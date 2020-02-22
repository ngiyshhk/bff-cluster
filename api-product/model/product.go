package model

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var ProductsMock = []Product{
	{ID: "1", Name: "product1", Price: 100},
	{ID: "2", Name: "product2", Price: 110},
	{ID: "3", Name: "product3", Price: 120},
	{ID: "4", Name: "product4", Price: 130},
	{ID: "5", Name: "product5", Price: 140},
}
