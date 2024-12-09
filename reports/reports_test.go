package reports

import (
	"fmt"
	"testing"
)

func TestCountSafe(t *testing.T) {
	input := []struct{
		nums []int
		want0 bool
		want1 bool
	}{
		// provided
		{[]int{7, 6, 4, 2, 1}, true, true},
		{[]int{1, 2, 7, 8, 9}, false, false},
		{[]int{9, 7, 6, 2, 1}, false, false},
		{[]int{1, 3, 2, 4, 5}, false, true},
		{[]int{8, 6, 4, 4, 1}, false, true},
		{[]int{1, 3, 6, 7, 9}, true, true},

		// regression
		{[]int{11, 10, 11, 13, 16}, false, true},
	}

	for i, test := range input {
		t.Run(fmt.Sprintf("test %v  - part 1", i), func(t *testing.T) {
			got := isSafe(test.nums, false)
			if (got != test.want0) {
				t.Errorf("want %v got %v, had %v", test.want0, got, test.nums)
			}
		})

		t.Run(fmt.Sprintf("test %v  - part 2", i), func(t *testing.T) {
			got := isSafe(test.nums, true)
			if (got != test.want1) {
				t.Errorf("want %v got %v, had %v", test.want1, got, test.nums)
			}
		})
	}
}
