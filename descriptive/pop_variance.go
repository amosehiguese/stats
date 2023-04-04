package descriptive

import (
	"fmt"
	"log"
	"math"
)

func PopVar(data []float64) (float64, error) {
	n := float64(len(data))
	if n == 0 {
		err := fmt.Errorf("empty population not allowed")
		return 0.0, err
	}

	mean, err := Mean(data)
	if err != nil {
		log.Fatal(err)
	}

	var summation, variance float64


	for _, value := range data {
		dev := value - mean
		summation += math.Pow(dev, 2)
	}

	variance = summation / float64((n))

	return variance, nil
}
