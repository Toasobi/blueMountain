package main

import "fmt"

func main() {
	var cowX int //存储坐标
	var johnX int
	var cowY int
	var johnY int

	var cowF int //存储方向 1向上,2向右,3向下,4向左
	var johnF int

	var step int //步数

	var str string
	mapp := make([][]byte, 10)
	fmt.Println("请输入地图")
	for i := 0; i < 10; i++ {
		fmt.Scanf("%s\n", &str)
		mapp[i] = []byte(str)

	} //初始化地图

	cowF = 1
	johnF = 1

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if mapp[i][j] == 'F' {
				johnX = i
				johnY = j

			}
			if mapp[i][j] == 'C' {
				cowX = i
				cowY = j

			}

		}

	} //初始化数据

	for cowX != johnX || cowY != johnY {

		//先判断农夫
		if johnF == 1 {
			if johnX-1 < 0 || mapp[johnX-1][johnY] == '*' {
				johnF = 2
			} else {
				johnX -= 1
			}
		} else if johnF == 2 {
			if johnY+1 >= 10 || mapp[johnX][johnY+1] == '*' {
				johnF = 3
			} else {
				johnY += 1
			}

		} else if johnF == 3 {
			if johnX+1 >= 10 || mapp[johnX+1][johnY] == '*' {
				johnF = 4
			} else {
				johnX += 1
			}
		} else {
			if johnY-1 < 0 || mapp[johnX][johnY-1] == '*' {
				johnF = 1
			} else {
				johnY -= 1
			}
		}

		//再判断牛牛

		if cowF == 1 {
			if cowX-1 < 0 || mapp[cowX-1][cowY] == '*' {
				cowF = 2
			} else {
				cowX -= 1
			}
		} else if cowF == 2 {
			if cowY+1 >= 10 || mapp[cowX][cowY+1] == '*' {
				cowF = 3
			} else {
				cowY += 1
			}
		} else if cowF == 3 {
			if cowX+1 >= 10 || mapp[cowX+1][cowY] == '*' {
				cowF = 4
			} else {
				cowX += 1
			}
		} else {
			if cowY-1 < 0 || mapp[cowX][cowY-1] == '*' {
				cowF = 1
			} else {
				cowY -= 1
			}
		}

		step += 1
		if step >= 100000 {
			fmt.Println(0)
			break
		}

	}
	fmt.Println(step)
}
