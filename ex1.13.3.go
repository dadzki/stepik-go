package main

import "fmt"

func main() {
	var seconds int
	fmt.Scan(&seconds)

	if seconds < 0 && seconds > 86399 {
		return
	}

	fmt.Printf("It is %v hours %v minutes.", seconds/3600, seconds/60%60)
}
