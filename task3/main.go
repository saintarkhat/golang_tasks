package main

import (
	"fmt"
	"strings"
)

func main() {
	var N, K int
	fmt.Scanln(&N, &K)

	var kegli = make([]string, N)
	//запоняем список кегель
	for i := 0; i < N; i++ {
		kegli[i] = "I"
	}
	//выбиваем кегли
	var l, r int
	for i := 0; i < K; i++ {
		fmt.Scanln(&l, &r)
		for j, _ := range kegli {
			if (j+1 >= l) && (j+1 <= r) {
				kegli[j] = "."
			}
		}
	}
	fmt.Println(strings.Join(kegli, ""))
}
