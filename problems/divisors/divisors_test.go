package divisors

import "testing"

func TestGcd(t *testing.T) {
	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"t-1", 4, 6, 2},
		{"t-2", 100, 50, 50},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			received := Gcd(testCase.a, testCase.b)
			if received != testCase.expected {
				t.Errorf("expected %d, got %d", testCase.expected, received)
			}
		})
	}
}
