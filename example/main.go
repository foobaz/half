package main

import (
	"fmt"
	"github.com/foobaz/half"
	"math"
)

func main() {
	var h half.Float16

	h = half.From64(0.00007)
	fmt.Printf("minimum == %f -> %f\n", 0.00007, h.To32())

	h = half.From32(math.Ln2)
	fmt.Printf("log2 == %f -> %f\n", math.Ln2, h.To64())

	h = half.From64(math.Pi)
	fmt.Printf("pi == %f -> %f\n", math.Pi, h.To64())

	h = half.From32(65500)
	fmt.Printf("maximum == %f -> %f\n", 65500.0, h.To32())

	h = half.From64(math.NaN())
	fmt.Printf("nan == %f -> %f\n", math.NaN(), h.To32())
}
