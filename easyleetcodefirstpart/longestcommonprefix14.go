package easyleetcodefirstpart

import (
	"sort"
)

var firstexample = []string{"flower", "flow", "flight"}
var secondexample = []string{"dog", "racecar", "car"}

func LongestCommonPrefix(strs []string) string {

	sort.Strings(strs)
	firststr := string(strs[0])
	laststr := string(strs[len(strs)-1])
	var longeststr string = ""

	if len(strs) == 0 {
		return ""
	}
	if len(strs) > 0 {
		for i := 0; i < len(firststr); i++ {
			if i >= len(laststr) || (firststr[i]) != (laststr[i]) {
				break
			}
			longeststr += string(firststr[i])

		}
	}
	return longeststr
}

//func longestCommonPrefix(strs []string) string {
//	var char byte
//	//res := ""
//	for i := 0; ; i++ {
//		for k, str := range strs {
//			if len(str) <= i || (str[i] != char && k != 0) {
//				return strs[0][:i]
//			} else {
//				char = str[i]
//			}
//		}
//	}
//	return ""
//}
