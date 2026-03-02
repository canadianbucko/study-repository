package calc

import "errors"

type Authentic interface {
	Sum(arg1 float64, arg2 float64) (float64, error)
	Minus(float64, float64) (float64, error)
	Multiply(float64, float64) (float64, error)
	Divide(float64, float64) (float64, error)
}

type calcImpl struct{}

func New() Authentic {
	return calcImpl{}
}

func (c calcImpl) Sum(arg float64, arg2 float64) (float64, error) {
	if arg <= -1000 || arg >= 1000 || arg2 <= -1000 || arg2 >= 1000 {
		err := errors.New("аргументы вне поля возможностей")
		return 0, err
	}
	sum := arg + arg2
	return sum, nil
}
func (c calcImpl) Minus(arg float64, arg2 float64) (float64, error) {
	if arg <= -1000 || arg >= 1000 || arg2 <= -1000 || arg2 >= 1000 {
		err := errors.New("аргументы вне поля возможностей")
		return 0, err
	}
	minus := arg - arg2
	return minus, nil
}

func (c calcImpl) Multiply(arg float64, arg2 float64) (float64, error) {
	if arg <= -1000 || arg >= 1000 || arg2 <= -1000 || arg2 >= 1000 {
		err := errors.New("аргументы вне поля возможностей")
		return 0, err
	}
	multiply := arg * arg2
	return multiply, nil
}

func (c calcImpl) Divide(arg float64, arg2 float64) (float64, error) {
	if arg <= -1000 || arg >= 1000 || arg2 <= -1000 || arg2 >= 1000 {
		err := errors.New("аргументы вне поля возможностей")
		return 0, err
	}
	if arg2 == 0 {
		err := errors.New("деление на 0 нельзя")
		return 0, err
	}
	multiply := arg / arg2
	return multiply, nil
}
