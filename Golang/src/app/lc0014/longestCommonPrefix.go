package lc0014


func LongestCommonPrefix(strs []string) string {
	comStr := ""
	for i := 1; i <= len(strs[0]); i++ {
		comStr = strs[0][0:i]
		for _, itemStr := range strs {
			if len(itemStr) < i || itemStr[0:i] != comStr {
				return comStr[0:i-1]
			}
		}
		if i == len(strs[0]){
			return comStr
}
}
	return ""
}