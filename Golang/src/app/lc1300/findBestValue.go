package lc1300

import "math"

func FindBestValue(arr []int, target int) int {
	arrLen := len(arr)
	minNum := int(target / arrLen)
	var diffNum int = 10001
	for {
		sum := SunArr(arr, minNum)
		newDiff := int(math.Abs(float64(target - sum)))
		if diffNum <= newDiff {
			return minNum - 1
		}
		diffNum = newDiff
		minNum += 1
	}

}

func SunArr(arr []int, num int) int {
	sum := 0
	for _, i := range arr {
		if i > num {
			sum += num
		} else {
			sum += i
		}
	}
	return sum
}
