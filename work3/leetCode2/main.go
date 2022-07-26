package main

type RecentCounter struct {
	arr []int
}

func Constructor() RecentCounter {
	recentCounter := &RecentCounter{
		arr: make([]int, 100),
	}
	return *recentCounter
}

func (this *RecentCounter) Ping(t int) int {
	this.arr = append(this.arr, t)
	for this.arr[0] < t-3000 {
		this.arr = this.arr[1:]
	}
	return len(this.arr)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
