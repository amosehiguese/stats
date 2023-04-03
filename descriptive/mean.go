package descriptive

import (
	"fmt"
)

// Mean calculates the mean of a slice of float64 values.
func Mean(data []float64) (float64, error) {
    if len(data) == 0 {
        err := fmt.Errorf("empty slice not allowed")
        return 0.0, err
    }
    var sum float64
    for _, x := range data {
        sum += x
    }
    return sum / float64(len(data)), nil
}
