package descriptive

import (
	"testing"
)

const (
    checkMark = "\u2713"
    ballotX = "\u2717"
)

func TestMean(t *testing.T) {
    data := []float64{1.0, 3.0, 5.0, 7.0}
    expected := 4.0
    result, _ := Mean(data)

    t.Log("Given the need to test Mean.")
    {

        if result != expected {
            t.Fatalf("Mean of %v was incorrect, got: %v, want: %v. %v", data, result, expected, ballotX)
        }
        t.Log("\t\tShould be able to get expected mean.", checkMark)
    }
}

func TestMeanEmptyList(t *testing.T) {
    data := []float64{}
    expected := 0.0
    result, _ := Mean(data)
    t.Log("Given the need to test Mean for an empty slice.")
    {
        if result != expected {
            t.Fatalf("Mean of %v was incorrect, got: %v, want: %v. %v", data, result, expected, ballotX)
        }
        t.Log("\t\tShould be able to get expected mean.", checkMark)
    }
}

func TestMeanSingleValue(t *testing.T) {
    data := []float64{42.0}
    expected := 42.0
    result,_ := Mean(data)
    t.Log("Given the need to test Mean for a slice of 1 element.")
    {
        if result != expected {
            t.Fatalf("Mean of %v was incorrect, got: %v, want: %v. %v", data, result, expected, ballotX)
        }
        t.Log("\t\tShould be able to get expected mean.", checkMark)        
    }
}
