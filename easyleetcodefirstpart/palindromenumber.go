package easyleetcodefirstpart

var (
	somevalue1 = 121
	somevalue2 = -121
	somevalue3 = 10
)

func IsPalindromenumber(x int) bool {
	if x < 0 {
		return false
	}
	original := x
	reversal := 0
	for x > 0 {
		reversal = reversal*10 + x%10
		x /= 10
	}

	return original == reversal
}

//alternative
//func isPalindrome(x int) bool {
//	var n string= strconv.Itoa(x)
//	var i int = 0
//	var j int = len(n) -1
//	for i < j {
//		if n[i] != n[j]{
//			return false
//		}
//		i++;
//		j--;
//	}
//	return true
//}
