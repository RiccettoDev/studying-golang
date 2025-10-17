package main

import "fmt"

type Client struct {
	Name string
	Age  int
}

func (c Client) Show() {
	fmt.Println("Exibindo cliente pelo método")
}

type ClientInternational struct {
	Client
	Country string
}

func main() {

	fmt.Println("--------------")
	fmt.Println("Structs")
	fmt.Println("--------------")

	client1 := Client{Name: "Eduardo", Age: 35}

	fmt.Println(client1)
	fmt.Println("--------------")

	fmt.Printf("Nome: %s Idade: %d\n", client1.Name, client1.Age)

	fmt.Println("--------------")
	fmt.Println("Herança")
	fmt.Println("--------------")
	client2 := ClientInternational{
		Client:  Client{Name: "Eduardo", Age: 35},
		Country: "Brasil",
	}

	fmt.Println(client2)
	fmt.Println("--------------")
	fmt.Printf("Nome: %s Idade: %d País: %s\n", client2.Name, client2.Age, client2.Country)

	fmt.Println("")
	client2.Show()
}
