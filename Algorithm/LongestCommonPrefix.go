package Algorithm

func LongestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	var length int = len(strs[0])
	var n int
	for k, v := range strs {
		if length > len(v) {
			length = len(v)
			n = k
		}
	}
	var substr string
	runes := []rune(strs[n])
	for k, v1 := range runes {
		for _, v2 := range strs[0:] {
			if v1 != []rune(v2)[k] {
				return substr
			}
		}
		substr = substr + string(v1)
	}
	return substr
}
