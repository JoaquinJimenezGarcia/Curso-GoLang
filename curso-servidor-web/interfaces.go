package main

import "fmt"

type animal interface {
	toMove() string
}

type dog struct{}
type fish struct{}
type bird struct{}

func (dog) toMove() string {
	return "Im a dog and I walk alone"
}

func (fish) toMove() string {
	return "Im a fish and I swim alone"
}

func (bird) toMove() string {
	return "Im a bird and I fly alone"
}

func moveAnimal(animal animal) {
	fmt.Println(animal.toMove())
}

func main() {
	dog := dog{}
	fish := fish{}
	bird := bird{}

	moveAnimal(dog)
	moveAnimal(fish)
	moveAnimal(bird)
}
