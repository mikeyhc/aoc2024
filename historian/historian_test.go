package historian

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestDiffLists(t *testing.T) {
	list0 := []int{1, 2, 3, 3, 3, 4}
	list1 := []int{3, 3, 3, 4, 5, 9}
	expected := []int{2, 1, 0, 1, 2, 5}

	got := DiffLists(list0, list1)

	if slices.Compare(expected, got) != 0 {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestSimilarityScore(t *testing.T) {
	list0 := []int{1, 2, 3, 3, 3, 4}
	list1 := []int{3, 3, 3, 4, 5, 9}
	expected := 31
	got := SimilarityScore(list0, list1)
	assertIntEqual(t, expected, got)
}

func TestSumSlice(t *testing.T) {
	list := []int{2, 1, 0, 1, 2, 5}
	expected := 11
	got := SumSlice(list)
	assertIntEqual(t, expected, got)
}

func TestIntAbs(t *testing.T) {
	t.Run("test positive", func(t *testing.T) {
		assertIntEqual(t, 11, intAbs(11))
	})

	t.Run("test negative", func(t *testing.T) {
		assertIntEqual(t, 11, intAbs(-11))
	})

	t.Run("test zero", func(t *testing.T) {
		assertIntEqual(t, 0, intAbs(0))
	})
}

func assertIntEqual(t testing.TB, expected, got int) {
	t.Helper()

	if expected != got {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
