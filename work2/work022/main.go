package main

import "fmt"

//dfs深搜

var arr = [10001]int{1} //最前面一个数设为1 因为后面的数不可能比1还小
var n int
var m int

func rprint(a int) {
	for i := 1; i < a; i++ {
		fmt.Print(arr[i])
		fmt.Print("+")
	}
	fmt.Println(arr[a])
}

//这里思路很关键
//先分为全为1的式子，之后往前一项设为2（该项），若数还有剩余 递归搜索后面位的数有无符合要求的可选数字
//有则打印，无则继续回溯使该项加一（因为不符合后面数一定为1）然后打印，再跳出一层递归往前一项继续找

func search(a int) {
	for i := arr[a-1]; i <= m; i++ {
		if i == n {
			continue
		}
		arr[a] = i
		m -= i
		if m == 0 {
			rprint(a)
		} else {
			search(a + 1)
		}
		m += i //回溯

	}
}

func main() {

	fmt.Println("请输入自然数")
	fmt.Scanf("%d\n", &n)
	m = n
	search(1)
}
