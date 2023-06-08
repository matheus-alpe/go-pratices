package basics

// (x, y int) = (x int, y int)
func Add(x, y int) int {
	return x + y
}

func Swap(x, y string) (string, string) {
	return y, x
}

// named returns
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return
}
