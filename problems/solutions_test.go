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

func TestFibSeries(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"test case from question", 10, 10},
		{"random test case", 10000, 3382},
		{"main test case", 4000000, 4613732},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			received := fibSeries(testCase.input)
			if received != testCase.expected {
				t.Errorf("expected %d, got %d", testCase.expected, received)
			}
		})
	}
}

func TestLargestPrimeFactor(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"test case from question", 13195, 29},
		{"main test case", 600851475143, 6857},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			received := largestPrimeFactor(testCase.input)
			if received != testCase.expected {
				t.Errorf("expected %d, got %d", testCase.expected, received)
			}
		})
	}
}

func TestPalindromeProduct(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"single digit solution", 1, 9},
		{"test case from question", 2, 9009},
		{"main test case", 3, 906609},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			received := largestPalindromeProduct(testCase.input)
			if received != testCase.expected {
				t.Errorf("expected %d, got %d", testCase.expected, received)
			}
		})
	}
}
