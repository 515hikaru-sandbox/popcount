package popcount

import (
	"testing"
)

func testPopCount(t *testing.T, x uint64, expected int) {
	result := PopCount(x)
	if result != expected {
		t.Errorf("%d population is %d. but got=%d", x, expected, result)
	}
}

func testPopCount3(t *testing.T, x uint64, expected int) {
	result := PopCount3(x)
	if result != expected {
		t.Errorf("%d population is %d. but got=%d", x, expected, result)
	}
}

func TestPopCount(t *testing.T) {
	testPopCount(t, 1, 1)
	testPopCount(t, 2, 1)
	testPopCount(t, 3, 2)
	testPopCount(t, 1023, 10)
	testPopCount(t, 1024, 1)
}

func TestPopCount3(t *testing.T) {
	testPopCount3(t, 1, 1)
	testPopCount3(t, 2, 1)
	testPopCount3(t, 3, 2)
	testPopCount3(t, 1023, 10)
	testPopCount3(t, 1024, 1)
}
