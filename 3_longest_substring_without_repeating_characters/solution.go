package __longest_substring_without_repeating_characters

func maxInt(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func lengthOfLongestSubstring(s string) int {
	status := map[uint8]bool{}
	left, right := 0, 0
	length := len(s)
	maxLen := 0
	for left < length && right < length {
		if _, ok := status[s[right]]; !ok {
			status[s[right]] = true
			right++
			maxLen = maxInt(maxLen, right-left)
		} else {
			delete(status, s[left])
			left++
		}
	}
	return maxLen
}
