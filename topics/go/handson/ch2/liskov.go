package ch2

import "fmt"

type Adder interface {
	Add(int, int) int
}

func PrintSum(a, b int, adder Adder) {
	fmt.Printf("%d + %d = %d", a, b, adder.Add(a, b))
}

type Int struct {
}

func (Int) Add(a, b int) int {
	return a + b
}

type Double struct {
}

func (d Double) Add(a, b float64) float64 {
	return a + b

}

func runTest() {
	PrintSum(1, 2, Int{})

	PrintSum(1, 2, Double{})
}
