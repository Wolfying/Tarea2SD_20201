package main

import (
	"fmt"
	"math"
)

func makeRange(min, max int32) []int32 {
	a := make([]int32, max-min)
	for i := range a {
		a[i] = min + int32(i)
	}
	return a
}

func main() {
	var number int32 = 1
	a := makeRange(0, int32(math.Floor(float64(number)/float64(3))))
	b := makeRange(int32(math.Floor(float64(number)/float64(3))), int32(math.Floor(2*float64(number)/float64(3))))
	c := makeRange(int32(math.Floor(2*float64(number)/float64(3))), number)
	fmt.Println(a)
	fmt.Printf("largo de a es %d \n", len(a))
	fmt.Println(b)
	fmt.Printf("largo de b es %d \n", len(b))
	fmt.Println(c)
	fmt.Printf("largo de c es %d \n", len(c))
}
