package model

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var UsersMock = []User{
	{ID: "1", Name: "user1", Age: 20},
	{ID: "2", Name: "user2", Age: 21},
	{ID: "3", Name: "user3", Age: 22},
}
