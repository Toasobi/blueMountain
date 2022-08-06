package main

import "fmt"

func isValid(s string) bool {
	//首先字符串判断长度
	if len(s)%2 == 1 {
		return false
	}
	if len(s) == 0 {
		return true
	}
	stack := make([]byte, len(s))
	var index int = -1 //栈索引
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			index++
			stack[index] = s[i]
		}
		if s[i] == ')' {
			if index == -1 {
				return false
			}
			if stack[index] == '(' {
				stack[index] = 0
				index--
			} else {
				return false
			}
		}
		if s[i] == ']' {
			if index == -1 {
				return false
			}
			if stack[index] == '[' {
				stack[index] = 0
				index--
			} else {
				return false
			}
		}
		if s[i] == '}' {
			if index == -1 {
				return false
			}
			if stack[index] == '{' {
				stack[index] = 0
				index--
			} else {
				return false
			}
		}
	}

	return index == -1
}

func main() {
	s := "[]{}()"
	b := isValid(s)
	fmt.Println(b)

}
