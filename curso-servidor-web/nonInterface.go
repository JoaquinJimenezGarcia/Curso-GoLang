package main

import "fmt"

type dog struct{}
type fish struct{}
type bird struct{}

func (dog) walk() string {
	return "Im a dog and I walk alone"
}

func (fish) swim() string {
	return "Im a fish and I swim alone"
}

func (bird) fly() string {
	return "Im a bird and I fly alone"
}

func (bird) walk() string {
	return "Im a bird and also I can walk"
}

func main() {
	dog := dog{}
	fish := fish{}
	bird := bird{}

	fmt.Println(dog.walk())
	fmt.Println(fish.swim())
	fmt.Println(bird.fly())
}
