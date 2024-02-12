package sum

import (
	"testing"
)

func TestInts(t *testing.T) {
	// Case 1: 1 + 2 + 3 + 4 + 5 = 15
	s := Ints([]int{1, 2, 3, 4, 5}...)
	if s != 15 {
		t.Errorf("Sum of one to five should be 15, but we got: %v", s)
	}

	// Case 2: Sum of empty input should be 0
	s = Ints()
	if s != 0 {
		t.Errorf("Sum of nothing should be 0, but we got: %v", s)
	}

	// Case 3: 1 + (-1) = 0
	s = Ints([]int{1, -1}...)

	if s != 0 {
		t.Errorf("Sum of one and minus one should be 0, but we got: %v", s)
	}
}

func TestIntsUsingTableDrivenTest(t *testing.T) {
	testCases := []struct {
		name    string
		numbers []int
		sum     int
	}{
		{"one to five", []int{1, 2, 3, 4, 5}, 15},
		{"nothing", nil, 0},
		{"one and minus one", []int{1, -1}, 0},
	}

	// loop over each item in test cases
	for _, testCases := range testCases {
		t.Run(testCases.name, func(t *testing.T) {
			s := Ints(testCases.numbers...)

			if s != testCases.sum {
				t.Errorf("Sum of %s should be %v, but we got: %v", testCases.name, testCases.sum, s)
			}
		})
	}
}
