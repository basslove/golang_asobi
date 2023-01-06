package main

import (
	"fmt"
	"strconv"
)

type Command interface {
	Execute() string
}

type MacroCommand struct {
	commands []Command
}

func (mc *MacroCommand) Execute() string {
	var result string
	for _, command := range mc.commands {
		result += command.Execute() + "\n"
	}
	return result
}

func (mc *MacroCommand) Append(command Command) {
	mc.commands = append(mc.commands, command)
}

func (mc *MacroCommand) Undo() {
	if len(mc.commands) != 0 {
		mc.commands = mc.commands[:len(mc.commands)-1]
	}
}

func (mc *MacroCommand) Clear() {
	mc.commands = []Command{}
}

type Position struct {
	X, Y int
}

type DrawCommand struct {
	Position *Position
}

func (dc *DrawCommand) Execute() string {
	return strconv.Itoa(dc.Position.X) + "." + strconv.Itoa(dc.Position.Y)
}

func main() {
	fmt.Println("command")

	var macro = &MacroCommand{}

	macro.Append(&DrawCommand{&Position{1, 1}})
	macro.Append(&DrawCommand{&Position{2, 2}})
	fmt.Println(macro.Execute())

	macro.Undo()
	fmt.Println(macro.Execute())
}
