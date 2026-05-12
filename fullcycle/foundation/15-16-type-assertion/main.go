package main

var a interface{} = "hello" // interface vazia permite uma liberdade de trabalho com tipos, mas perigoso

func main() {
	println(a)
	x := a.(string)
	println(x)
	res, ok := a.(int)
	print(res, ok)
}
