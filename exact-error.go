package main

import (
	"fmt"
	"encoding/csv"
	"log"
	"os"
	"math"
)

func roundFloat(val float64, precision uint) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(val*ratio) / ratio
}


func main() {
	data := [][]string{
		{"Input", "Output", "Error"},
	}

	for i:=1.0; i<=7.0; i+=0.00001 {
		input := roundFloat(i, 5)
		output := i / 6.0 * 100.0
		output = output / 100.0 * 6.0
		output = roundFloat(output, 5)
		diffError := input - output

		newSlice := []string{fmt.Sprintf("%v", input), fmt.Sprintf("%v", output), fmt.Sprintf("%v", diffError)}
		data = append(data, newSlice)
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