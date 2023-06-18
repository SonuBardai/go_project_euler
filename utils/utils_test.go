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
