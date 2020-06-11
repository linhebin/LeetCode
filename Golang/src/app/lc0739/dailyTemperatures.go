package lc0739

func DailyTemperatures(T []int) []int {
	DSlice := make([]int, 0)
	TLen := len(T)
	for i := 0; i < TLen; i++ {
		flag := true
		for j := 0; j < TLen-i; j++ {
			if T[i] < T[i+j] {
				DSlice = append(DSlice, j)
				flag = false
				break
			}
		}
		if flag {
			DSlice = append(DSlice, 0)
		}
	}
	return DSlice
}
