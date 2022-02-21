package main

import "fmt"

func newGenericFunc[age int64 | float32](myage age) {
	fmt.Println(myage)
}

func main() {
	fmt.Println("Go Generics!!")

	var age int64 = 26
	var age2 float32 = 26.6

	newGenericFunc(age)
	newGenericFunc(age2)
}
