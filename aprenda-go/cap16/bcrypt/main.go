package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "123456"
	senhaErrada := "123"

	x, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(x))

	y := bcrypt.CompareHashAndPassword(x, []byte(senha))
	if y != nil {
		fmt.Println("Você digitou a senha errada")
	} else {
		fmt.Println("Você digitou a senha correta")
	}

	z := bcrypt.CompareHashAndPassword(x, []byte(senhaErrada))
	if z != nil {
		fmt.Println("Você digitou a senha errada")
	} else {
		fmt.Println("Você digitou a senha correta")
	}

	fmt.Println(y)
	fmt.Println(z)

}
