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

	for j:=10.0; j < 10000000; j = j * 10 {

		var stepSize float64
		stepSize = 1.0 / j

		roundingSize := uint(math.Log10(j)) // this is the precision, number of after comma numbers

		stepSize = roundFloat(stepSize, roundingSize)

		log.Print(fmt.Sprintf("Calculating grades with step size %v", stepSize))

		for i:=1.0; i<=7.0; i+=stepSize {
			input := i
			input = roundFloat(input, roundingSize)

			output := i / 6.0 * 100.0
			output = output / 100.0 * 6.0
			output = roundFloat(output, roundingSize)

			diffError := input - output

			if diffError != 0.0 {
			input := roundFloat(i, 5)
				log.Fatal(fmt.Sprintf("Found difference error at step size %v and rounding %v with input %v, output %v, error %v", stepSize, roundingSize, input, output, diffError))
			}

			newSlice := []string{fmt.Sprintf("%v", input), fmt.Sprintf("%v", output), fmt.Sprintf("%v", diffError)}
			data = append(data, newSlice)
		}
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