package function

func NotB(x bool) (bool, error) {
	return !x, nil
}

func AndB(x, y bool) (bool, error) {
	return x && y, nil
}

func OrB(x, y bool) (bool, error) {
	return x || y, nil
}

func XorB(x, y bool) (bool, error) {
	return (x && !y) || (!x && y), nil
}

func NandB(x, y bool) (bool, error) {
	return !x || !y, nil
}