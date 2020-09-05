package main

import (
	algorithms "Algorithms"
	"fmt"
)

func main() {
	fmt.Println(algorithms.IsInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(algorithms.IsInterleave("aabcc", "dbbca", "aadbbbaccc"))
	fmt.Println(algorithms.IsInterleave("", "", ""))
	fmt.Println(algorithms.IsInterleave("cbcccbabbccbbcccbbbcabbbabcababbbbbbaccaccbabbaacbaabbbc", "abcbbcaababccacbaaaccbabaabbaaabcbababbcccbbabbbcbbb", "abcbcccbacbbbbccbcbcacacbbbbacabbbabbcacbcaabcbaaacbcbbbabbbaacacbbaaaabccbcbaabbbaaabbcccbcbabababbbcbbbcbb"))
}
