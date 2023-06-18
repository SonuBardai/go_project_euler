package problems

import "testing"

func TestSumOfMultiples(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"test case from question", 10, 23},
		{"main test case", 1000, 233168},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			sumReceived := sumOfMultiples(testCase.input)
			if sumReceived != testCase.expected {
				t.Errorf("expected %d, got %d", testCase.expected, sumReceived)
			}
		})
	}
}
