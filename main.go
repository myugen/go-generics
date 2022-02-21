package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func newGenericFunc[age Number](myage age) {
	val := int(myage) + 1
	fmt.Println(val)
}

func main() {
	fmt.Println("Go Generics!!")

	var age int64 = 26
	var age2 float32 = 26.6

	newGenericFunc(age)
	newGenericFunc(age2)
}
