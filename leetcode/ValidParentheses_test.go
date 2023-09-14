package leetcode

import (
	"testing"
)

// TODO: 有辦法改寫成 List 逐一測試(parallel) 嗎?
/*
type testDSPair struct {
	data   string
	target bool
}

func TestIsValid(t *testing.T) {
	t.Parallel()
	s := "()"
	target := true

	testDS := []testDSPair{
		{"()", true},
		{"()[]{}", true},
	}

	for _, ds := range testDS {
		s := ds.data
		target := ds.target

		result := isValid(s)

		if result != target {
			t.Errorf("Expected true but got %v", result)
		}
	}

	result := isValid(s)
	if result != target {
		t.Errorf("Expected true but got %v", result)
	}
}
*/

func TestIsValid1(t *testing.T) {
	s := "()"
	target := true

	result := isValid(s)
	if result != target {
		t.Errorf("Expected true but got %v", result)
	}
}

func TestIsValid2(t *testing.T) {
	s := "()[]{}"
	target := true

	result := isValid(s)
	if result != target {
		t.Errorf("Expected true but got %v", result)
	}
}

func TestIsValid3(t *testing.T) {
	s := "(]"
	target := false

	result := isValid(s)
	if result != target {
		t.Errorf("Expected true but got %v", result)
	}
}

func TestIsValid4(t *testing.T) {
	s := "{[]}"
	target := true

	result := isValid(s)
	if result != target {
		t.Errorf("Expected true but got %v", result)
	}
}

func TestIsValid5(t *testing.T) {
	s := "(("
	target := false

	result := isValid(s)
	if result != target {
		t.Errorf("Expected true but got %v", result)
	}
}

func TestIsValid6(t *testing.T) {
	s := "["
	target := false

	result := isValid(s)
	if result != target {
		t.Errorf("Expected true but got %v", result)
	}
}

func TestIsValid7(t *testing.T) {
	s := "]["
	target := false

	result := isValid(s)
	if result != target {
		t.Errorf("Expected true but got %v", result)
	}
}

func TestIsValid8(t *testing.T) {
	s := "(){}}{"
	target := false

	result := isValid(s)
	if result != target {
		t.Errorf("Expected true but got %v", result)
	}
}

func TestIsValid9(t *testing.T) {
	s := "{}[{}]((){})(){}"
	target := true

	result := isValid(s)
	if result != target {
		t.Errorf("Expected true but got %v", result)
	}
}
