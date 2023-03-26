package descriptive

import "testing"

func TestMean(t *testing.T) {
    data := []float64{1.0, 3.0, 5.0}
    expected := 3.0
    result := Mean(data)
    if result != expected {
        t.Errorf("Mean of %v was incorrect, got: %v, want: %v.", data, result, expected)
    }
}

func TestMeanEmptyList(t *testing.T) {
    data := []float64{}
    expected := 0.0
    result := Mean(data)
    if result != expected {
        t.Errorf("Mean of %v was incorrect, got: %v, want: %v.", data, result, expected)
    }
}

func TestMeanSingleValue(t *testing.T) {
    data := []float64{42.0}
    expected := 42.0
    result := Mean(data)
    if result != expected {
        t.Errorf("Mean of %v was incorrect, got: %v, want: %v.", data, result, expected)
    }
}
