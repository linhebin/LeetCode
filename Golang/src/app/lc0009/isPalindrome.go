package lc0009

import "app/common"

func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	} else if x < 10 {
		return true
	}
	xStr := common.Int2String(x)
	xLen := len(xStr)
	xNewStr := ""
	for i := xLen - 1; i >= 0; i-- {
		xNewStr += xStr[i : i+1]
	}
	if xStr == xNewStr {
		return true
	}
	return false
}

func IsPalindrome2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	} else if x < 10 {
		return true
	}
	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}
	return x == revertedNumber || x == revertedNumber/10
}
