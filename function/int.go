package function

import "strconv"

func AbsI(x int) (int, error) {
	if x < 0 {
		return -x, nil
	}
	return x, nil
}

func PlusI(x, y int) (int, error) {
	return x + y, nil
}

func MinusI(x, y int) (int, error) {
	return x - y, nil
}

func MulI(x, y int) (int, error) {
	return x * y, nil
}

func DivI(x, y int) (int, error) {
	return x / y, nil
}

func StrI(x int) (*string, error) {
	result := strconv.Itoa(x)
	return &result, nil
}
