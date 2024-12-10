package mull

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("simple parse", func(t *testing.T) {
		want := []Command{
			{Mul, 2, 4},
			{Mul, 5, 5},
			{Mul, 11, 8},
			{Mul, 8, 5},
		}
		input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
		got := Parse([]byte(input))
		assertSliceEqual(t, want, got)
	})

	t.Run("complex parse", func(t *testing.T) {
		want := []Command{
			{Mul, 2, 4},
			{Dont, 0, 0},
			{Mul, 5, 5},
			{Mul, 11, 8},
			{Do, 0, 0},
			{Mul, 8, 5},
		}
		input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
		got := Parse([]byte(input))
		assertSliceEqual(t, want, got)
	})
}

func TestRunCommands(t *testing.T) {
	input := []Command{
		{Mul, 2, 4},
		{Dont, 0, 0},
		{Mul, 5, 5},
		{Mul, 11, 8},
		{Do, 0, 0},
		{Mul, 8, 5},
	}

	t.Run("basic run", func(t *testing.T) {
		want := []int{8, 25, 88, 40}
		got, err := RunCommands(input, false)
		assertNoError(t, err)
		assertSliceEqual(t, want, got)
	})

	t.Run("conditional run", func(t *testing.T) {
		want := []int{8, 40}
		got, err := RunCommands(input, true)
		assertNoError(t, err)
		assertSliceEqual(t, want, got)
	})
}

func assertSliceEqual[T comparable](t testing.TB, want, got []T) {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("unexpected error: %v", got)
	}
}
