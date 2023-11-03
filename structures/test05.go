package main

import "fmt"

type User struct {
	Login    string
	Email    string
	Password string
}

type Client struct {
	User
	Name string
	Age  int
}

func main() {
	user := User{
		Login:    "Admin",
		Email:    "antonio.santana.contato@gmail.com",
		Password: "admin123",
	}

	client := Client{
		user,
		"Antônio",
		18,
	}

	fmt.Printf(
		"My name is %s and I have %v years old! \nContact me: %s\n",
		client.Name, client.Age, user.Email,
	)

	fmt.Printf(
		"\nMy very security acess is:\nLogin: %s\nPassword: %s\n",
		user.Login, user.Password,
	)
}
