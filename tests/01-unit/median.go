package unit

import (
	"fmt"
	"strconv"
)

func Median(numbers ...float64) float64 {
	total := 0.0

	for _, num := range numbers {
		total += num
	}

	median := total / float64(len(numbers))
	roundedMedian, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", median), 64)
	return roundedMedian
}
