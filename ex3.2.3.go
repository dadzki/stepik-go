package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\n")

	text = strings.Replace(text, " ", "", -1)
	text = strings.Replace(text, ",", ".", -1)
	numbers := strings.Split(text, ";")

	number1, err := strconv.ParseFloat(numbers[0], 64)
	if err != nil {
		panic(err)
	}

	number2, err := strconv.ParseFloat(numbers[1], 64)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%.4f\n", number1/number2)

}
