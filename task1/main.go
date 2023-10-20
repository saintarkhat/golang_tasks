package main

import "fmt"

func getSomeVars() string {
	fmt.Println("getSomeVars execution")
	return "getSomeVars result"
}

func main() {
	var a, b, c, d, x int
	fmt.Scan(&a, &b, &c, &d)
	nums := []int{a, b, c, d}
	x = int(^uint(0) >> 1)
	for _, num := range nums {
		if num < x {
			x = num
		}
	}
	fmt.Println(x)
}
