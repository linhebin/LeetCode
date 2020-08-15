package lc0125

import (
	"fmt"
	"strings"
)

func IsPalindrome(s string) bool {
	r := ""
	l := ""
	for _, c := range s {
		fmt.Println(c)
		if (65 <= c && c <= 90) || (97 <= c && c <= 122) || (48 <= c && c <= 57) {
			cc := string(c)
			if 65 <= c && c <= 90 {
				cc = strings.ToLower(cc)
			}
			r = r + cc
			l = cc + l
		}
	}
	fmt.Println(r)
	fmt.Println(l)
	return r == l
}
