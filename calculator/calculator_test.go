package calculator_test

import (
	"testing"

	"github.com/Zaressa/pivottech/calculator"
)

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
		t.Run("test", func(t *testing.T) {
			got := calculator.Add(val.a, val.b)
			if got != val.want {
				t.Errorf("got %d want %d", got, val.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	data := []struct {
		example string
		a       int
		b       int
		want    int
	}{

		{"exapmple1", 4, 2, 2},
		{"exapmple2", 6, 2, 4},
		{"exapmple3", -1, 2, -3},
	}
	for _, val := range data {
		t.Run(val.example, func(t *testing.T) {
			got := calculator.Subtract(val.a, val.b)
			if got != val.want {
				t.Errorf("got %d want %d", got, val.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {

	got := calculator.Multiply(2, 2)
	want := 4
	if got != want {
		t.Errorf("got %d want %d", got, want)

	}
}

func TestPow(t *testing.T) {
	data := []struct {
		example string
		x       float64
		y       float64
		want    float64
	}{
		{"example", 2, 3, 8},
		{"exapmple1", 3, 3, 27},
	}

	for _, val := range data {
		t.Run(val.example, func(t *testing.T) {
			got := calculator.Pow(val.x, val.y)
			if got != val.want {
				t.Errorf("got %f want %f", got, val.want)
			}
		})
	}
}

func TestDivideTable(t *testing.T) {
	data := []struct {
		example string
		a       int
		b       int
		want    int
	}{

		{"example 1", 4, 2, 2},
		{"example 2", 6, 2, 3},
		{"example 3", 2, 0, 0},
	}
	for _, val := range data {
		t.Run(val.example, func(t *testing.T) {
			got, _ := calculator.Divide(val.a, val.b)
			if got != val.want {
				t.Errorf("got %d want %d", got, val.want)
			}
		})

	}
}
