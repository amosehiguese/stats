package descriptive

// Mean calculates the mean of a slice of float64 values.
func Mean(data []float64) float64 {
    var sum float64
    for _, x := range data {
        sum += x
    }
    return sum / float64(len(data))
}
