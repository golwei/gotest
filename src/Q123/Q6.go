
package main

import (
	"fmt"
)

func main() {
	a, b := f2(7, 2)
	fmt.Println(a, b)
}

func f1(a []float64) (avg float64) {
	sum := 0.0
	for _, v := range a {
		sum += v

	}
	avg = sum / float64(len(a))
	return
}
func f2(a, b int) (c, d int) {
	if a < b {
		c, d = a, b
	} else {
		c, d = b, a
	}
	return
}
