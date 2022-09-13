package main

import (
	"fmt"
	"time"
)

func main() {
	var dateStr string
	_, err := fmt.Scan(&dateStr)
	if err != nil {
		panic(err)
	}
	//dateStr = "1986-04-16T05:20:00+06:00"

	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		panic(err)
	}

	fmt.Println(date.Format(time.UnixDate))
}
