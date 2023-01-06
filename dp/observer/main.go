package main

import (
	"fmt"
	"math/rand"
)

type numberGenerator struct {
	observers []observer
}

func (ng *numberGenerator) AddObserver(observer observer) {
	ng.observers = append(ng.observers, observer)
}

func (ng *numberGenerator) notifyObservers() []int {
	result := make([]int, 0, len(ng.observers))
	for _, observer := range ng.observers {
		result = append(result, observer.update())
	}
	return result
}

type randomNumberGenerator struct {
	*numberGenerator
}

func NewRandomNumberGenerator() *randomNumberGenerator {
	return &randomNumberGenerator{&numberGenerator{}}
}

type number interface {
	getNumber() int
}

func (rng *randomNumberGenerator) getNumber() int {
	return rand.Intn(50)
}

func (rng *randomNumberGenerator) Execute() []int {
	return rng.notifyObservers()
}

type observer interface {
	update() int
}

type DigitObserver struct {
	generator number
}

func (do *DigitObserver) update() int {
	return do.generator.getNumber()
}

func main() {
	fmt.Println("observer")

	rng := NewRandomNumberGenerator()

	o1 := &DigitObserver{rng}
	o2 := &DigitObserver{rng}

	rng.AddObserver(o1)
	rng.AddObserver(o2)

	fmt.Println(rng.Execute())
}
