package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {

	m := make(map[string]int)

	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	space := []byte{'\n'}
	splitExample := bytes.Split(content, space)

	for _, line := range splitExample {
		command := string(line)
		if len(command) > 0 {
			words := strings.Fields(command)
			if words[0] == "DEPOSIT" {
				sumDep, _ := strconv.Atoi(words[2])
				m[words[1]] += int(sumDep)
			} else if words[0] == "WITHDRAW" {
				sumWithdraw, _ := strconv.Atoi(words[2])
				m[words[1]] += -int(sumWithdraw)
			} else if words[0] == "TRANSFER" {
				sumTransfer, _ := strconv.Atoi(words[3])
				m[words[1]] += -int(sumTransfer)
				m[words[2]] += int(sumTransfer)
			} else if words[0] == "INCOME" {
				sumDep, _ := strconv.Atoi(words[1])
				for key, value := range m {
					if value > 0 {
						m[key] += value * sumDep / 100
					}
				}
			} else if words[0] == "BALANCE" {
				if val, ok := m[words[1]]; ok {
					fmt.Println(val)
				} else {
					fmt.Println("ERROR")
				}
			}
		}
	}
}
