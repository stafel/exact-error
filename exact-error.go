package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	data := [][]string{
		{"a", "1"},
		{"b", "2"},
		{"c", "3"},
	}

	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.WriteAll(data)
}