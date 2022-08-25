package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var number, digit int

	fmt.Scan(&number, &digit)

	fmt.Printf(strings.Replace(strconv.Itoa(number), strconv.Itoa(digit), "", -1))
}
