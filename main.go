package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func BubbleSort[N Number](input []N) []N {
	n := len(input)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}

	return input
}

func newGenericFunc[age Number](myage age) {
	val := int(myage) + 1
	fmt.Println(val)
}

func main() {
	fmt.Println("Go Generics!!")

	listints := []int32{4, 1, 13, 54, 13, 6}
	listfloats := []float32{.4, .1, 1.3, 5.4, 13.2, .06}
	fmt.Println(BubbleSort(listints))
	fmt.Println(BubbleSort(listfloats))

	var age int64 = 26
	var age2 float32 = 26.6

	newGenericFunc(age)
	newGenericFunc(age2)
}
