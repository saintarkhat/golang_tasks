package main

import "fmt"

func main() {
	var a, b int

	var arrayOfArrays [8][2]int
	for i := range arrayOfArrays { //assign
		fmt.Scanln(&a, &b)
		arrayOfArrays[i] = [...]int{a, b}
	}

	// Считаем суммы x и y каждого ферзя

	var arraySums = make([]int, 8)
	for i := range arraySums {
		arraySums[i] = arrayOfArrays[i][0] + arrayOfArrays[i][1]
	}
	distinct := make(map[int]bool)
	for _, v := range arraySums {
		distinct[v] = true
	}
	cntL := len(distinct)

	// Считаем Х ферзей

	var arrayX = make([]int, 8)
	for i := range arrayX {
		arrayX[i] = arrayOfArrays[i][0]
	}
	distinct = make(map[int]bool)
	for _, v := range arrayX {
		distinct[v] = true
	}
	cntX := len(distinct)

	// Считаем Y ферзей

	var arrayY = make([]int, 8)
	for i := range arrayY {
		arrayY[i] = arrayOfArrays[i][1]
	}
	distinct = make(map[int]bool)
	for _, v := range arrayY {
		distinct[v] = true
	}
	cntY := len(distinct)

	// если один из них меньше 8, то ферзи пересекаются

	if cntL < 8 || cntX < 8 || cntY < 8 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}