package Longest_Common_Prefix

func LongestCommonPrefix(strs []string) string {
	var minLength int = len(strs[0])
	var pref string
	for _, str := range strs {
		if minLength > len(str) {
			minLength = len(str)
		}
	}

	for i := 0; i < minLength; i++ {
		var ch string = string(strs[0][i])
		for _, str := range strs {
			if ch != string(str[i]) {
				return pref
			}

		}
		pref += string(ch)
	}
	return pref
}
