package main

import "fmt"

type memento struct {
	money int
}

func (m *memento) getMoney() int {
	return m.money
}

type Game struct {
	Money int
}

func (g *Game) CreateMemento() *memento {
	return &memento{
		g.Money,
	}
}

func (g *Game) RestoreMemento(memento *memento) {
	g.Money = memento.getMoney()
}

func main() {
	fmt.Println("momento")

	game := &Game{100}
	memento := game.CreateMemento()

	game.Money = 0
	game.RestoreMemento(memento)

	fmt.Println(game.Money)
}
