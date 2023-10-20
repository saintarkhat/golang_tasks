package main

import (
	"fmt"
	"math"
)

func main() {
	var a, i float64
	fmt.Scan(&a)
	i = 3

	if math.Mod(a, 2) == 0 {
		fmt.Println(2)
	} else {
		for i <= math.Sqrt(a) {
			if math.Mod(a, i) == 0 {
				a = i
				break
			}
			i += 2
		}
		fmt.Println(int(a))
	}
}
