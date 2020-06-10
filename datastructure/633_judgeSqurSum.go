package datastructure

// int(6.9)==6
// %.0f 四舍五入
func judgeSquareSum(c int) bool {
	i, j := 0, squr(c)

	for i <= j {
		result := i*i + j*j
		if result == c {
			return true
		} else if result > c {
			j--
		} else {
			i++
		}
	}
	return false
}

func squr(a int) int {
	b := a
	for b*b > a {
		b = (b + a/b) / 2
	}
	return b
}
