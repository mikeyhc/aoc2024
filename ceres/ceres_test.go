package ceres

import (
	"testing"
)

func TestCountXmas(t *testing.T) {
	input := [][]byte{
		{'.', '.', 'X', '.', '.', '.'},
		{'.', 'S', 'A', 'M', 'X', '.'},
		{'.', 'A', '.', '.', 'A', '.'},
		{'X', 'M', 'A', 'S', '.', 'S'},
		{'.', 'X', '.', '.', '.', '.'}}
	want := 4
	got := CountXmas(input)

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestCountMasX(t *testing.T) {
	input := [][]byte{
		[]byte(".M.S......"),
		[]byte("..A..MSMS."),
		[]byte(".M.S.MAA.."),
		[]byte("..A.ASMSM."),
		[]byte(".M.S.M...."),
		[]byte(".........."),
		[]byte("S.S.S.S.S."),
		[]byte(".A.A.A.A.."),
		[]byte("M.M.M.M.M."),
		[]byte("..........")}
	want := 9
	got := CountMasX(input)

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
