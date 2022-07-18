package main

import "fmt"

func main() {
	var n int
	var before []byte
	//var after string
	fmt.Println("输入本题的n")
	fmt.Scanf("%d\n", &n)
	fmt.Println("请输入原文")
	fmt.Scanf("%s\n", &before)

	for i := 0; i < len(before); i++ {
		if before[i] >= 'a' && before[i] <= byte(int('z')-n) {
			before[i] = byte(int(before[i]) + n)
		} else {
			before[i] = byte(int('a') + (n - (int('z') - int(before[i]) + 1)))
		}
	}
	for j := 0; j < len(before); j++ {
		fmt.Print(string(before[j]))
	}

}
