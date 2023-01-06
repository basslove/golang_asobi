package main

import (
	"fmt"
	"math/rand"
)

const (
	GUU = iota
	CHO
	PAA
)

var hands []*hand

func init() {
	hands = []*hand{
		&hand{GUU},
		&hand{CHO},
		&hand{PAA},
	}
}

type hand struct {
	handValue int
}

func getHand(handValue int) *hand {
	return hands[handValue]
}

func (ha *hand) IsStrongerThan(h *hand) bool {
	return ha.fight(h) == 1
}

func (ha *hand) IsWeakerThan(h *hand) bool {
	return ha.fight(h) == -1
}

func (ha *hand) fight(h *hand) int {
	if ha == h {
		return 0
	} else if (ha.handValue+1)%3 == h.handValue {
		return 1
	} else {
		return -1
	}
}

type strategy interface {
	NextHand() *hand
	study(win bool)
}

type winningStrategy struct {
	seed     int64
	won      bool
	prevHand *hand
}

func (ws *winningStrategy) NextHand() *hand {
	if !ws.won {
		ws.prevHand = getHand(rand.Intn(3))
	}
	return ws.prevHand
}

func (ws *winningStrategy) study(win bool) {
	ws.won = win
}

type Player struct {
	Name                           string
	Strategy                       strategy
	wincount, losecount, gamecount int
}

func (p *Player) NextHand() *hand {
	return p.Strategy.NextHand()
}

func (p *Player) Win() {
	p.wincount++
	p.gamecount++
}

func (p *Player) Lose() {
	p.losecount++
	p.gamecount++
}

func (p *Player) Even() {
	p.gamecount++
}

func main() {
	fmt.Println("strategy")

	player1 := Player{Name: "A", Strategy: &winningStrategy{seed: 10}}
	player2 := Player{Name: "B", Strategy: &winningStrategy{seed: 20}}

	hand1 := player1.NextHand()
	hand2 := player2.NextHand()

	if hand1.IsStrongerThan(hand2) {
		player1.Win()
		player2.Lose()
	} else if hand1.IsWeakerThan(hand2) {
		player1.Lose()
		player2.Win()
	} else {
		player1.Even()
		player2.Even()
	}
}
