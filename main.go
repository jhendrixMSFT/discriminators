package main

import (
	"discriminators/farm"
	"fmt"
)

type myFish struct {
	farm.FishType
}

func main() {
	fishExample()
	salmonExample()
}

func fishExample() {
	fish := farm.GetFish()
	fmt.Printf("fish ID: %d\n", fish.GetFish().ID)
	switch x := fish.(type) {
	case *farm.Salmon:
		fmt.Printf("caught a %s salmon\n", x.Breed)
	case *farm.Tuna:
		fmt.Printf("the tuna has color %s\n", x.Color)
	case *farm.Fish:
		fmt.Printf("generic fish ID %d\n", x.ID)
	}
}

func salmonExample() {
	fish := farm.GetSalmon()
	fmt.Printf("salmon ID: %d\n", fish.GetFish().ID)
	switch x := fish.(type) {
	case *farm.Salmon:
		fmt.Printf("caught a %s salmon\n", x.Breed)
	case *farm.SmartSalmon:
		fmt.Printf("the smart salmon's IQ is %d\n", x.IQ)
	}
}
