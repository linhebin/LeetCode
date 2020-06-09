package common

import "strconv"

func String2Int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func Int2String(i int) string {
	s := strconv.Itoa(i)
	return s
}
