package main

import (
	"fmt"
	"sort"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func getChild(matrix map[string][]string, point string, score map[string]int) map[string]int {
	val := matrix[point]
	if len(val) >= 1 {
		for _, children := range val {
			parScore := score[point]
			score[children] += 1 + parScore
			score = getChild(matrix, children, score)
		}
	}
	return score
}

func main() {
	var n int
	fmt.Scan(&n)
	var a, b string
	var child, parent []string
	matrixHierarchy := make(map[string][]string)
	matrixScore := make(map[string]int)
	for i := 0; i < n; i++ {
		fmt.Scanln(&b, &a)
		if !contains(child, b) {
			child = append(child, b)
			parent = append(parent, a)
			matrixHierarchy[a] = append(matrixHierarchy[a], b)
		}
	}
	var ancestor string
	for _, par := range parent {
		if !contains(child, par) {
			ancestor = par
			break
		}
	}
	matrixScore[ancestor] = 0

	matrixScore = getChild(matrixHierarchy, ancestor, matrixScore)

	keys := make([]string, 0, len(matrixScore))

	for k := range matrixScore {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, matrixScore[k])
	}
}
