package descriptive

import (
	"sort"
)

// Median returns the median value of a slice of float64
func Median(data []float64) float64 {
	if len(data) == 0 {
		return 0.0
	}

	// Check if data are not sorted
	if !sort.Float64sAreSorted(data){
		sort.Float64s(data)
	}

	// return middle element if slice is odd
	if len(data) % 2 == 1 {
		return data[len(data) / 2]
	}

	return (data[(len(data) / 2) - 1] + data[len(data) / 2]) / 2.0


}