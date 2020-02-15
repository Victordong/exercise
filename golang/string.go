package main

func isPalindrome2(s string) bool {
	i := 0
	j := len(s) - 1
	var num1, num2 uint8
	for i < j {
		if s[i] <= 'Z' && s[i] >= 'A' {
			num1 = s[i] - 'A'
		} else if s[i] <= 'z' && s[i] >= 'a' {
			num1 = s[i] - 'a'
		} else if s[i] <= '9' && s[i] >= '0' {
			num1 = s[i] - '0' + 100
		} else {
			i += 1
			continue
		}
		if s[j] <= 'Z' && s[j] >= 'A' {
			num2 = s[j] - 'A'
		} else if s[j] <= 'z' && s[j] >= 'a' {
			num2 = s[j] - 'a'
		} else if s[j] <= '9' && s[j] >= '0' {
			num2 = s[j] - '0' + 100
		} else {
			j -= 1
			continue
		}
		if num1 == num2 {
			i += 1
			j -= 1
		} else {
			return false
		}
	}
	return true
}
