package FirstOccurrence

func StrStr(haystack string, needle string) int {
	pointerHaystack := 0
	pointerNeedle := 0
	for pointerHaystack < len(haystack) {
		if haystack[pointerHaystack] != needle[pointerNeedle] {
			pointerHaystack = pointerHaystack - pointerNeedle + 1
			if pointerHaystack == len(haystack) {
				return -1
			}
			pointerNeedle = 0
		}
		if haystack[pointerHaystack] == needle[pointerNeedle] {
			pointerHaystack++
			pointerNeedle++
		}
		if pointerNeedle == len(needle) {
			return pointerHaystack - pointerNeedle
		}
	}
	return -1
}
