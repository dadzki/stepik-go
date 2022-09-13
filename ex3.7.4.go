package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	startDate := time.Unix(1589570165, 0)

	var text string
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	text = strings.TrimSpace(text)

	//text = "12 мин. 13 сек."
	//text = "71 мин. 8 сек."

	re := regexp.MustCompile(`(\d+) мин\. (\d+) сек\.`)
	items := re.FindStringSubmatch(text)
	if len(items) == 0 {
		panic("Invalid text format!")
	}

	minutes, _ := strconv.Atoi(items[1])
	seconds, _ := strconv.Atoi(items[2])
	d := time.Minute*time.Duration(minutes) + time.Second*time.Duration(seconds)
	fmt.Println(startDate.Add(d).UTC().Format(time.UnixDate))
}
