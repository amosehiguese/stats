package descriptive

import (
	"math"
	"testing"
)

func TestMedian(t *testing.T){
	// Test an empty slice
	data := []float64{}

	expected := 0.0

	if actual := Median(data); actual != expected{
		t.Errorf("Median(%v) = %v; expected %v", data, actual, expected)
	}

	// Test a slice of one element
	data = []float64{1}

	expected = 1

	if actual := Median(data); actual != expected {
		t.Errorf("Median(%v) = %v; expected %v", data, actual, expected)
	}

	// Test with a slice of NaN values
	data = []float64{math.NaN(), math.NaN()}

	expected = math.NaN()

	if actual := Median(data); math.IsNaN(actual) != math.IsNaN(expected) {
		t.Errorf("Median(%v) = %v; expected %v", data, actual, expected)
	}

	// Test with an odd-length slice
	data = []float64{1, 2, 3, 4, 5}

	expected = 3.0 

	if actual := Median(data); actual != expected {
		t.Errorf("Median(%v) = %v; expected %v", data, actual, expected)
	}


	// Test with an even-length slice
	data = []float64{1, 2, 3, 4, 5, 6}

	expected = 3.5 

	if actual := Median(data); actual != expected {
		t.Errorf("Median(%v) = %v; expected %v", data, actual, expected)
	}
}