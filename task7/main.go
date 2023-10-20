package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strconv"
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
	for i, word := range words {
		if math.Mod(float64(i), 2) != 0 {
			c, _ := strconv.Atoi(word)
			m[words[i-1]] += c
		}
	}

	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}
