package main

var meuArray [3]int // array tem um tamanho e um único tipo pré definiddo

func main() {
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 7

	for i := 0; i < 3; i++ {
		println(meuArray[i])
	}
}
