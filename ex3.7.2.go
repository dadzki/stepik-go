package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var dateStr string
	dateStr, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}
	dateStr = strings.TrimSpace(dateStr)

	//dateStr = "2020-05-15 08:00:00"
	//dateStr = "2020-05-15 14:00:00"

	const dateLayout = "2006-01-02 15:04:05"

	date, err := time.Parse(dateLayout, dateStr)
	if err != nil {
		panic(err)
	}

	if date.Hour() >= 13 {
		date = date.Add(time.Hour * 24)
	}

	fmt.Println(date.Format(dateLayout))
}
