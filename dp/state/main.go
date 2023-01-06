package main

import (
	"fmt"
	"strings"
)

type State interface {
	DoClock(context context, hour int)
	DoUse(context context)
}

var dayInstance *dayState

type dayState struct {
}

func GetDayInstance() *dayState {
	if dayInstance == nil {
		dayInstance = &dayState{}
	}
	return dayInstance
}

func (sf *dayState) DoClock(context context, hour int) {
	if hour < 9 || 17 <= hour {
		context.ChangeState(GetNightInstance())
	}
}

func (sf *dayState) DoUse(context context) {
	context.RecordLog("day")
}

var nightInstance *nightState

type nightState struct {
}

func GetNightInstance() *nightState {
	if nightInstance == nil {
		nightInstance = &nightState{}
	}
	return nightInstance
}

func (sf *nightState) DoClock(context context, hour int) {
	if 9 <= hour && hour < 17 {
		context.ChangeState(GetDayInstance())
	}
}

func (sf *nightState) DoUse(context context) {
	context.RecordLog("night")
}

type context interface {
	SetClock(hour int)
	ChangeState(state State)
	RecordLog(log string)
}

type SafeFrame struct {
	State State
	logs  []string
}

func (sf *SafeFrame) SetClock(hour int) {
	sf.State.DoClock(sf, hour)
}

func (sf *SafeFrame) ChangeState(state State) {
	sf.State = state
}

func (sf *SafeFrame) RecordLog(log string) {
	sf.logs = append(sf.logs, log)
}

func (sf *SafeFrame) Use() {
	sf.State.DoUse(sf)
}

func (sf *SafeFrame) GetLog() string {
	return strings.Join(sf.logs, " ")
}

func main() {
	fmt.Println("state")

	safeFrame := &SafeFrame{State: GetDayInstance()}

	hours := []int{8, 9, 16, 17}
	for _, h := range hours {
		safeFrame.SetClock(h)
		safeFrame.Use()
	}

	fmt.Println(safeFrame.GetLog())
}
