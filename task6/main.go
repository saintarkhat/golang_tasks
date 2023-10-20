package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	m := make(map[string]int)

	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	myString := string(content[:])

	words := strings.Fields(myString)
	for _, word := range words {
		m[word] += 1
	}

	maxNum := 0
	var minKey string
	for i, val := range m {
		if maxNum == 0 {
			maxNum = m[i]
			minKey = i
		} else if val > maxNum {
			maxNum = val
			minKey = i
		} else if (val == maxNum) && (minKey > i) {
			minKey = i
		}
	}

	fmt.Println(minKey)

}
