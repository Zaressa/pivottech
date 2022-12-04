package calculator

import "errors"

func Add(a, b int) int {
	return a + b
}
func Subtract(a, b int) int {
	return b - a
}

func Multiply(a, b int) int {
	return a * b
}
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("calculator: cannot divide by 0, please try again")
	}
	return a / b, nil

}
func DividebyZero(a, b int) (c int, ok bool) {
	defer func() {
		recover() // recover from panic if one occured. Set err to nil otherwise.
	}()
	c = a / b
	return c, true
}
