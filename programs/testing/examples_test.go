package calendar

import (
	"testing"
)

func add(x, y int) (int){
	return x + y
}

func TestAddPass(t *testing.T) {
	expected := 4
	actual := add(2, 2)
	if expected != add(2, 2) {
		t.Error("Expected:", expected, "Actual:", actual)
	}
}

func TestAddFail(t *testing.T) {
	expected := 5
	actual := add(2, 2)
	if expected != add(2, 2) {
		t.Error("Expected:", expected, "Actual:", actual)
	}
}