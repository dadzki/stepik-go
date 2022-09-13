package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i := 1; i <= 3; i++ {
		file, err := os.Create(strconv.Itoa(i) + "text.txt")
		if err != nil {
			panic(err)
		}
		fmt.Println(file.Name())
		defer file.Close()
	}

	os.Rename("2text.txt", "4text.txt")

	for i := 4; i >= 2; i-- {
		fmt.Println(strconv.Itoa(i) + "text.txt")
		os.Remove(strconv.Itoa(i) + "text.txt")
	}
}
