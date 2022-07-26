package main

import "fmt"

func timeRequiredToBuy(tickets []int, k int) int { //k为需要关注的索引
	var second int
	for tickets[k] != 0 {
		for i := 0; i < len(tickets); i++ {
			if tickets[k] == 0 {
				return second
			}
			if tickets[i] != 0 {
				second++
				tickets[i]--
			} else {
				continue
			}
		}

	}
	return second
}

//测试代码
func main() {
	var tickets = [4]int{5, 1, 1, 1}
	k := 0
	a := timeRequiredToBuy(tickets[:], k)
	fmt.Println(a)
}
