package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func contains(s []int, n int) bool {
	for _, v := range s {
		if v == n {
			return true
		}
	}

	return false
}

func main() {
	line, _ := ioutil.ReadFile("input.txt")

	line1 := string(line[:])
	re := regexp.MustCompile(`\r?\n`)
	line1 = re.ReplaceAllString(line1, " ")

	strings := strings.Split(line1, " ")

	var ints []int
	for _, s := range strings {
		if len(s) > 0 {
			v, _ := strconv.Atoi(s)
			if contains(ints, v) {
				fmt.Println("YES")
			} else {
				ints = append(ints, v)
				fmt.Println("NO")
			}
		}
	}
}
