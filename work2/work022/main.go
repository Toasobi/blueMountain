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
