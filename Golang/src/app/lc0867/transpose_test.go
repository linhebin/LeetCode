package lc0867_test

import (
	"app/lc0867"
	"fmt"
	"testing"
)

func TestTranspose(t *testing.T) {
	a := [][]int{{1, 2, 3}, {4, 5, 6}}
	b := lc0867.Transpose(a)
	fmt.Println(b)
}
