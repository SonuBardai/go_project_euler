package utils

import "testing"

func TestContains(t *testing.T) {
	items := []string{"problems", "inputs"}
	if !contains(items, "problems") {
		t.Errorf("contains(%v, %v) = false; want true", items, "problems")
	}
	if contains(items, "golang") {
		t.Errorf("contains(%v, %v) = true; want false", items, "golang")
	}
}

func TestAny(t *testing.T) {
	items := []string{"problems", "inputs"}
	candidates := []string{"problems", "inputs"}
	if !Any(items, candidates) {
		t.Errorf("Any(%v, %v) = false; want true", items, candidates)
	}
	not_candidates := []string{"golang"}
	if Any(items, not_candidates) {
		t.Errorf("Any(%v, %v) = true; want false", items, not_candidates)
	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"test_1", "abcd", "dcba"},
		{"test_2", "0000", "0000"},
		{"test_3", "Hello", "olleH"},
		{"test_4", "98989898", "89898989"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			received := Reverse(testCase.input)
			if received != testCase.expected {
				t.Errorf("expected %s, got %s", testCase.expected, received)
			}
		})
	}
}
