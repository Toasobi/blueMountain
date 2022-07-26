package main

var index = 0
var indexD = 0
var indexC = 0

type AnimalShelf struct {
	animals [][]int
}

func Constructor() AnimalShelf {
	animalShelf := &AnimalShelf{
		animals: make([][]int, 2000),
	}
	return *animalShelf
}

func (this *AnimalShelf) Enqueue(animal []int) {
	for i := 0; i < len(animal); i++ {
		this.animals[i] = animal
	}
}

func (this *AnimalShelf) DequeueAny() []int {
	for i := 0; i < len(this.animals); i++ {
		if this.animals[i][0] == index {
			return this.animals[i]
		}
	}
	index++
	this.DequeueAny()

	//否则
	arr := []int{-1, -1}
	return arr
}

func (this *AnimalShelf) DequeueDog() []int {
	for i := 0; i < len(this.animals); i++ {
		if this.animals[i][1] == 1 {
			if this.animals[i][0] == index {
				return this.animals[i]
			}
		} else {
			continue
		}
	}
	indexD++
	this.DequeueAny()

	//否则
	arr := []int{-1, -1}
	return arr
}

func (this *AnimalShelf) DequeueCat() []int {
	for i := 0; i < len(this.animals); i++ {
		if this.animals[i][1] == 0 {
			if this.animals[i][0] == index {
				return this.animals[i]
			}
		} else {
			continue
		}
	}
	indexC++
	this.DequeueAny()

	//否则
	arr := []int{-1, -1}
	return arr

}

/**
 * Your AnimalShelf object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Enqueue(animal);
 * param_2 := obj.DequeueAny();
 * param_3 := obj.DequeueDog();
 * param_4 := obj.DequeueCat();
 */
