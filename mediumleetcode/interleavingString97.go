package mediumleetcode

//type FirstStringAttempt struct {
//	S1 string
//	S2 string
//	S3 string
//}
//
//var firstString FirstStringAttempt

var firstPart = "aabcc"
var secPart = "dbbca"
var trueState = "aadbbcbcac"
var falseState = "aadbbbaccc"

func IsInterleave(s1 string, s2 string, s3 string) bool {

	firstPartLen, secPartLen, trueStateLen := len(s1), len(s2), len(s3)
	// length compare
	if firstPartLen+secPartLen != trueStateLen {
		return false
	}
	// create bool table
	dpTable := make([][]bool, firstPartLen+1)
	for i := range dpTable {
		dpTable[i] = make([]bool, secPartLen+1)
	}

	dpTable[0][0] = true
	// create first column and row within table
	for i := 1; i <= firstPartLen; i++ {
		dpTable[i][0] = dpTable[i-1][0] && s1[i-1] == s3[i-1]
	}
	for j := 1; j <= secPartLen; j++ {
		dpTable[0][j] = dpTable[0][j-1] && s2[j-1] == s3[j-1]
	}
	// fill remain table
	for i := 1; i <= firstPartLen; i++ {
		for j := 1; j <= secPartLen; j++ {
			match1 := dpTable[i-1][j] && s1[i-1] == s3[i+j-1]
			match2 := dpTable[i][j-1] && s2[j-1] == s3[i+j-1]
			dpTable[i][j] = match1 || match2
		}

	}

	return dpTable[firstPartLen][secPartLen]
}
