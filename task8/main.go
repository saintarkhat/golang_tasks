package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func lagrange(x float64, f int) []string {
	var res []string
	for i := 0; i < 4; i++ {
		var s float64
		if i == 0 {
			s = float64(int(math.Sqrt(x)) - f)
		} else {
			s = float64(int(math.Sqrt(x)))
		}
		x = x - math.Pow(s, 2)
		if s > 0 {
			t := strconv.FormatInt(int64(s), 10)
			res = append(res, t)
		}
	}
	return res
}

func lagrangeTest(s []string, x float64) bool {
	var y float64
	for _, v := range s {
		vT, _ := strconv.Atoi(v)
		y += math.Pow(float64(vT), 2)
	}
	if x == y {
		return true
	} else {
		return false
	}
}

func main() {

	var x float64
	fmt.Scanln(&x)
	var res []string
	f := 0
	res = lagrange(x, f)
	for !lagrangeTest(res, x) {
		f++
		res = lagrange(x, f)
	}
	fmt.Println(strings.Join(res, " "))
}
