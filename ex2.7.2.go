package main

import (
	"fmt"
)

func main() {
	var str string

	fmt.Scan(&str)
	sufix := "*"
	for i, v := range str {
		if i == len(str)-1 {
			sufix = ""
		}
		fmt.Printf("%s%v", string(v), sufix)
	}
}
