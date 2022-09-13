package main

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
)

func main() {
	const FileName = "task.zip"

	zf, err := zip.OpenReader(FileName)
	if err != nil {
		panic(err)
	}
	defer zf.Close()

	for _, file := range zf.File {
		f, err := file.Open()
		if err != nil {
			continue
		}
		defer f.Close()

		lines, _ := csv.NewReader(f).ReadAll()
		if len(lines) > 1 {
			fmt.Println(lines[4][2])
		}
	}
}
