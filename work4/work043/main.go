package main

import (
	"fmt"
)

// //还是运用栈

// func calculate(s string) int {
// 	stack := make([]int, len(s))
// 	var index = 0  //栈索引
// 	var sum = 0    //最后要返回的值
// 	var index2 = 1 //开始方位
// 	//移除空格
// 	s1 := removeNil(s)
// 	//判断是否就是数字
// 	value, err := strconv.ParseInt(s1, 10, 64)
// 	if err == nil {
// 		return int(value)
// 	}

// 	top, _ := strconv.ParseInt(s1[0:1], 10, 64)
// 	stack[index] = int(top)
// 	index++

// 	for i := index2; i < len(s1); i++ {
// 		if s1[i] == '+' {
// 			slice := s1[i+1 : i+2]
// 			value, _ := strconv.ParseInt(slice, 10, 64)
// 			stack[index] = int(value)
// 			index++
// 		}
// 		if s1[i] == '-' {
// 			slice := s1[i+1 : i+2]
// 			before, _ := strconv.ParseInt(slice, 10, 64)
// 			after := before * -1
// 			stack[index] = int(after)
// 			index++
// 		}
// 		if s1[i] == '*' {
// 			slice := s1[i+1 : i+2]
// 			value, _ := strconv.ParseInt(slice, 10, 64)
// 			result := value * int64(stack[index-1])
// 			stack[index-1] = int(result)
// 		}
// 		if s1[i] == '/' {
// 			slice := s1[i+1 : i+2]
// 			value, _ := strconv.ParseInt(slice, 10, 64)
// 			result := int64(stack[index-1]) / value
// 			stack[index-1] = int(result)
// 		}

// 		continue
// 	}
// 	for i := 0; i < len(stack); i++ {
// 		sum += stack[i]
// 	}
// 	return sum
// }

// func removeNil(s string) string {
// 	s0 := make([]byte, len(s))
// 	var index = 0
// 	var erase = 0
// 	for i := 0; i < len(s); i++ {
// 		if s[i] != ' ' {
// 			s0[index] = s[i]
// 			index++
// 		} else {
// 			erase += 1
// 			continue
// 		}
// 	}
// 	s1 := make([]byte, len(s0)-erase)
// 	for i := 0; i < len(s1); i++ {
// 		s1[i] = s0[i]
// 	}
// 	return string(s1)
// }

//还是题解好
func calculate(s string) (ans int) {
	stack := []int{}
	preSign := '+'
	num := 0
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			preSign = ch
			num = 0
		}
	}
	for _, v := range stack {
		ans += v
	}
	return
}

func main() {
	s := "0-2"
	a := calculate(s)
	fmt.Println(a)
}
