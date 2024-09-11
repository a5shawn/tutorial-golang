package main

import "fmt"

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}
	// SumInts(ints)
	sumInts := SumIntsOrFloats(ints)
	fmt.Println("sum ints", sumInts)

	sumFloats := SumIntsOrFloats(floats)
	fmt.Println("sum float", sumFloats)
}

func SumInts(m map[string]int64) {
	var sum int64
	for _, v := range m {
		sum += v
	}
	fmt.Println("sum =", sum)
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}
