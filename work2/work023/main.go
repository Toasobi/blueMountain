package main

import "fmt"

//dfs深搜典型题

var n int = 3

var m int = 3

var x int = 1

var y int = 1

var ar = [][]int{}

var dx = []int{1, 1, -1, -1, 2, 2, -2, -2}
var dy = []int{2, -2, -2, 2, 1, -1, -1, 1}

func check(nx int, ny int) bool {
	if nx > m-1 || ny > n-1 || nx < 0 || ny < 0 {
		return false
	}
	return true
}

func dfs(x int, y int, sum int) { //传递原来的位置和走过的步数
	ar[x][y] = sum
	var nx int
	var ny int
	for i := 0; i < 8; i++ {
		nx = x + dx[i]
		ny = y + dy[i]

		if check(nx, ny) && (sum+1 < ar[nx][ny] || ar[nx][ny] == -1) {
			dfs(nx, ny, sum+1)
		}

	}

}

func main() {

	//接收数据

	fmt.Println("n:")
	fmt.Scanf("%d\n", &n)

	fmt.Println("m:")
	fmt.Scanf("%d\n", &m)

	fmt.Println("x:")
	fmt.Scanf("%d\n", &x)

	fmt.Println("y:")
	fmt.Scanf("%d\n", &y)

	//初始化二维数组
	ar = make([][]int, n)
	for i := 0; i < m; i++ {
		ar[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ar[i][j] = -1
		}
	}

	dfs(x-1, y-1, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(ar[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}

}
