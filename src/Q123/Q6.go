
package main

import (
	"fmt"
)

func main() {
	a, b := f2(7, 2)
	fmt.Println(a, b)
}


func f2(a, b int) (c, d int) {
	if a < b {
		c, d = a, b
	} else {
		c, d = b, a
	}
	return
}
