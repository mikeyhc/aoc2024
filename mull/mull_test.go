package mull

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	want := []Command{
		{Mul, 2, 4},
		{Mul, 5, 5},
		{Mul, 11, 8},
		{Mul, 8, 5},
	}

	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	got := Parse([]byte(input))
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v, had %v", want, got, input)
	}
}

func TestRunCommands(t *testing.T) {
	want := []int{8, 25, 88, 40}
	input := []Command{
		{Mul, 2, 4},
		{Mul, 5, 5},
		{Mul, 11, 8},
		{Mul, 8, 5},
	}

	got, err := RunCommands(input)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v, had %v", want, got, input)
	}
}
