package printqueue

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuildOrderingRules(t *testing.T) {
	input := []Pair{
		{47, 53},
		{97, 13},
		{97, 61},
		{97, 47},
		{75, 29},
		{61, 13},
		{75, 53},
		{29, 13},
		{97, 29},
		{53, 29},
		{61, 53},
		{97, 53},
		{61, 29},
		{47, 13},
		{75, 47},
		{97, 75},
		{47, 61},
		{75, 61},
		{47, 29},
		{75, 13},
		{53, 13},
	}

	want := map[int][]int{
		29:{13},
		47:{53, 13, 61, 29},
		53:{29, 13},
		61:{13, 53, 29},
		75:{29, 53, 47, 61, 13},
		97:{13, 61, 47, 29, 53, 75},
	}

	got := BuildOrderingRules(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestValidOrdering(t *testing.T) {
	rules := map[int][]int{
		29:{13},
		47:{53, 13, 61, 29},
		53:{29, 13},
		61:{13, 53, 29},
		75:{29, 53, 47, 61, 13},
		97:{13, 61, 47, 29, 53, 75},
	}

	tests := []struct{ input []int; want bool}{
		{[]int{75, 47, 61, 53, 29}, true},
		{[]int{97, 61, 53, 29, 13}, true},
		{[]int{75, 29, 13}, true},
		{[]int{75, 97, 47, 61, 53}, false},
		{[]int{61, 13, 29}, false},
		{[]int{97, 13, 75, 29, 47}, false},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("validate update - test %v", i), func(t *testing.T) {
			got := ValidOrdering(test.input, rules)
			if got != test.want {
				t.Errorf("want %v, got %v, had %v", test.want, got, test.input)
			}
		})
	}
}

func TestFindMiddle(t *testing.T) {
	tests := []struct{ input []int; want int}{
		{[]int{75, 47, 61, 53, 29}, 61},
		{[]int{97, 61, 53, 29, 13}, 53},
		{[]int{75, 29, 13}, 29},
		{[]int{75, 97, 47, 61, 53}, 47},
		{[]int{61, 13, 29}, 13},
		{[]int{97, 13, 75, 29, 47}, 75},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("find middle - test %v", i), func(t *testing.T) {
			got := FindMiddle(test.input)
			if got != test.want {
				t.Errorf("want %v, got %v, had %v", test.want, got, test.input)
			}
		})
	}
}

func TestReorder(t *testing.T) {
	rules := map[int][]int{
		29:{13},
		47:{53, 13, 61, 29},
		53:{29, 13},
		61:{13, 53, 29},
		75:{29, 53, 47, 61, 13},
		97:{13, 61, 47, 29, 53, 75},
	}

	tests := []struct{ input []int; want []int}{
		{[]int{75, 97, 47, 61, 53}, []int{97, 75, 47, 61, 53}},
		{[]int{61, 13, 29}, []int{61, 29, 13}},
		{[]int{97, 13, 75, 29, 47}, []int{97, 75, 47, 29, 13}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("reorder - test %v", i), func(t *testing.T) {
			got := Reorder(test.input, rules)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("want %v, got %v, had %v", test.want, got, test.input)
			}
		})
	}
}
