package model

type User struct {
	Name    string
	Surname string
	Company string
	Age     string
	Email   string
	Pass    string
}

// mocks data
var Users = []User{
	{Name: "John",
		Surname: "Smith",
		Company: "Redis",
		Age:     "23",
		Email:   "john@redis.com",
		Pass:    "pass"},
	{Name: "John1",
		Surname: "Smith1",
		Company: "Redis1",
		Age:     "21",
		Email:   "john1@redis.com",
		Pass:    "pass1"},
	{Name: "John2",
		Surname: "Smith2",
		Company: "Redis2",
		Age:     "22",
		Email:   "john2@redis.com",
		Pass:    "pass2"},
}
