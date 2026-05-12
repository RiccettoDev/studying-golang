package main

func main() {
	println("Linha 1")
	println("Linha 2")
	println("Linha 3")
	defer println("Linha 4") // defer faz com que esta execucão sejá a ultima do código
	println("Linha 5")
	println("Linha 6")
	println("Linha 7")
}
