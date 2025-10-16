package calc

func Sum(a, b int) int {
	return a + b
}

// lembrar que para uma função poder ser utilizada em outro arquivo, é necessária ser pública e portanto começar coma a letra maiúscula.

var DentroDoPacoteCalc string = "Esta variável vem de dentro do pacote calc, no arquivo sum.go"

// a questão a primeira letra em miúscula também é válida para variaveis públicas e privadas como o caso à cima
