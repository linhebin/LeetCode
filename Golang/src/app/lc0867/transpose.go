package lc0867

func Transpose(A [][]int) [][]int {
	newA := make([][]int, 0)
	for i := 0; i < len(A[0]); i++ {
		newList := make([]int, 0)
		for _, itemList := range A {
			newList = append(newList, itemList[i])
		}
		newA = append(newA, newList)
	}
	return newA
}
