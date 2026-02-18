package main

var s = []int{}            // slices permitem que sejam criado vazio e sem pré determinar seu tamanho
var s2 = []int{1, 5, 7, 9} // mas també permitem que sejá pré definido
var s3 = [2]int{1, 5}      // também pré detyerminar tamanho caso necessário

func main() {
	s = append(s, 1) // basta adicionar depois elementos ao slice
	s = append(s, 2)

	for i := 0; i < len(s); i++ {
		println(s[i])
	}

	println("=======---------")

	for i := 0; i < len(s2); i++ {
		println(s2[i])
	}

	println("=======---------")

	for i := 0; i < len(s3); i++ {
		println(s3[i])
	}
}
