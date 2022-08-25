package main

import "fmt"

func main() {
	var n, newNumb int
	fmt.Scan(&n)

	newNumb = n%10*100 + n/10%10*10 + n/100

	fmt.Println(newNumb)
}
