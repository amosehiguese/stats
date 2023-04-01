package descriptive

import "testing"

const (
    checkMark = "\u2713"
    ballotX = "\u2717"
)

func TestMean(t *testing.T) {
    data := []float64{1.0, 3.0, 5.0, 7.0}
    expected := 4.0
    result := Mean(data)

    t.Log("\t\tGiven the need to test Mean.")
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
