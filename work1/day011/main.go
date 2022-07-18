package main

import (
	"fmt"
)

func main() {
	var equation string //输入的方程式
	var a [100]int      //整型数组
	var num int         //常数项之和
	var sum int         //系数项之和
	var mid int         //分割等号两边
	var index int

	l := 0  //数组位置
	pd := 1 //判断正负
	fmt.Println("请输入您想要求解的方程式")
	fmt.Scanf("%s\n", &equation)
	//equation = "6a-5+1=2-2a"
	for k := 0; k < len(equation); k++ {
		if equation[k] == '+' {
			l++
			index++
			pd = 1

		}
		if equation[k] == '-' {
			index++
			l++
			pd = -1
		}
		if equation[k] >= '0' && equation[k] <= '9' {
			index++
			if a[l] == 0 {
				a[l] = int(equation[k]-'0') * pd
			} else {
				a[l] = 10*a[l] + int(equation[k]-'0')*pd
			}
		}
		if equation[k] >= 'a' && equation[k] <= 'z' {
			index++
			if a[l] != 0 {
				sum += a[l]
				a[l] = 0
			} else {
				sum += pd
			}
			l--
		}
		//等号后面的全部都要反过来 用 mid 存储位置方便后面计算
		if equation[k] == '=' {

			mid = l
			l++
			break
		}

	}
	for j := index; j < len(equation); j++ {
		if equation[j] == '+' {
			l++
			pd = 1

		}
		if equation[j] == '-' {
			l++
			pd = -1
		}
		if equation[j] >= '0' && equation[j] <= '9' {

			if a[l] == 0 {
				a[l] = int(equation[j]-'0') * pd
			} else {
				a[l] = 10*a[l] + int(equation[j]-'0')*pd
			}
		}
		if equation[j] >= 'a' && equation[j] <= 'z' {

			if a[l] != 0 {
				sum -= a[l]
				a[l] = 0
			} else {
				sum -= pd
			}
			l--
		}
	}

	//统计常数项
	for i := 0; i < len(a); i++ {
		if i <= mid {
			num -= a[i]
		} else {
			num += a[i]
		}
	}
	//计算(保留三位小数)
	var result float32
	result = float32(num) / float32(sum)
	fmt.Printf("最终结果为%0.3f\n", result)
}
