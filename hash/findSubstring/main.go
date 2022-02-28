package main

func main() {

}

// s = "barfoothefoobarman", words = ["foo","bar"]

func findSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	wordsLen := 0
	wordsMap := make(map[string]int)
	for _, word := range words {
		wordsMap[word]++
		wordsLen += wordLen
	}

	var deque []string
	//var unorderMap map[string]int

	for i := 0; i+wordsLen <= len(s); i += wordLen {
		// 移出对头
		if len(deque) != 0 {

		}
	}
	return nil
}

func equalsMap(map1 map[string]int, map2 map[string]int) bool {
	for k1, v1 := range map1 {
		if v2, exist := map2[k1]; !exist || v1 != v2 {
			return false
		}
	}
	for k1, v1 := range map2 {
		if v2, exist := map1[k1]; !exist || v1 != v2 {
			return false
		}
	}
	return true
}
