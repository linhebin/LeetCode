package lc0046

func TranslateNum(num int) int {
	if num < 10 {
		return 1
	}
	var res int
	if num%100 <= 25 && num%100 > 9 {
		res += TranslateNum(num / 100)
		res += TranslateNum(num / 10)
	} else {
		res += TranslateNum(num / 10)
	}
	return res
}
