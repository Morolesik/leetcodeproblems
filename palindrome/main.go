package main

import "fmt"

func main() {
	str := "babad"

	fmt.Println("Result = ", longestPalindrome(str))
}

func longestPalindrome(s string) string {
	k := 0
	m := 0
	for j := 0; j < len(s); j++ {
		for i := len(s); i >= j; i-- {
			if isPalindrom(s[j:i]) {
				if i-j > m-k {
					k = j
					m = i
				}
				//fmt.Println("k,m=", k, "m=", m, "slice=", s[k:m])
				break
				//return s[j : len(s)-i]
			}
		}
		if (len(s) - j) < (m - k) {
			break
		}
	}

	return s[k:m]
	//return ""

}

func isPalindrom(s1 string) bool {
	//fmt.Println(s1)
	if len(s1) == 1 {
		return true
	}

	if len(s1) == 0 {
		return true
	}
	for i := 0; i <= len(s1)/2; i++ {
		//fmt.Println("testPali", string(s1[i]), " test2 ", string(s1[len(s1)-i-1]))
		if s1[i] != s1[len(s1)-i-1] {
			return false
		}
	}
	return true

}
