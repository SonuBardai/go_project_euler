package palindrome
import "testing"

func TestNiner(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"2", 2, 99},
		{"3", 3, 999},
		{"4", 4, 9999},
		{"9", 9, 999999999},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			received := Niner(testCase.input)
			if received != testCase.expected {
				t.Errorf("expected %d, got %d", testCase.expected, received)
			}
		})
	}
}