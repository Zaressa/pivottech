package calc_test

import (
	"calc"
	"testing"
)

func TestAdd(t *testing.T) {
	got := calc.Add(2, 2)
	want := 4
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := calc.Subtract(4, 1)
	want := 3
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestMultiply(t *testing.T) {
	got := calc.Multiply(2, 2)
	want := 4
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDivide(t *testing.T) {
	got, _ := calc.Divide(4, 2)
	want := 2
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestAddTable(t *testing.T) {
	data := []struct {
		a    int
		b    int
		want int
	}{
		{6, 6, 12},
		{2, 2, 4},
		{8, 2, 10},
	}
	for _, val := range data {
		got := calc.Add(val.a, val.b)
		if got != val.want {
			t.Errorf("got %d want %d", got, val.want)
		}
	}
}
func TestSubtractTable(t *testing.T) {
	data := []struct {
		a    int
		b    int
		want int
	}{
		{4, 2, 2},
		{6, 2, 4},
		{8, 2, 6},
	}

	for _, val := range data {
		got := calc.Subtract(val.a, val.b)
		if got != val.want {
			t.Errorf("got %d want %d", got, val.want)
		}
	}
}
func TestMultiplyTable(t *testing.T) {
	data := []struct {
		a    int
		b    int
		want int
	}{
		{4, 2, 8},
		{6, 2, 12},
		{8, 2, 16},
	}
	for _, val := range data {
		got := calc.Multiply(val.a, val.b)
		if got != val.want {
			t.Errorf("got %d want %d", got, val.want)
		}
	}
}

// func TestDivideTable(t *testing.T){
// data := []struct {
// 	a int
// 	b int
// 	want int
// } {
// 	{4, 2, 2},
// 	{6, 2, 3},
// 	{8, 2, 4},
// }
// for _, val := range data{
// 	got, _ := calc.Divide(val.a/ val.b)
// 	if got != val.want {
// 		t.Errorf("got %d want %d", got, val.want)
// 	}
// }
// }
// func TestDividebyZero(t *testing.T){
// data := []struct {
// 	a int
// 	b int
// 	want int
// } {
// 	{4, 0, 0},
// 	{6, 0, 0},
// 	{8, 0, 0},
// }
// for _, val := range data{
// 	got := calc.TestDividebyZero(val.a, val.b)
// 	if got != val.want {
// 		t.Errorf("got %d want %d", got, val.want)
// 	}
// }
// }
