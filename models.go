package main


type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

var users = []User{
	{Id: 1, Name: "Caleb", Age: 20},
	{Id: 2, Name: "Joshua", Age: 18},
}

