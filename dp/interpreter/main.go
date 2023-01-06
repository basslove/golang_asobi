package main

import (
	"fmt"
	"strings"
)

type Node interface {
	Parse(context *context)
	ToString() string
}

type ProgramNode struct {
	commandListNode Node
}

func (pn *ProgramNode) Parse(context *context) {
	context.skipToken("program")
	pn.commandListNode = &commandListNode{}
	pn.commandListNode.Parse(context)
}

func (pn *ProgramNode) ToString() string {
	return "program: " + pn.commandListNode.ToString()
}

type commandListNode struct {
	list []Node
}

func (cln *commandListNode) Parse(context *context) {
	for {
		if context.currentToken == "end" {
			break
		} else {
			commandNode := &commandNode{}
			commandNode.Parse(context)
			cln.list = append(cln.list, commandNode)
		}
	}
}

func (cln *commandListNode) ToString() string {
	var str string
	for _, l := range cln.list {
		str += l.ToString()
	}
	return str
}

type commandNode struct {
	node Node
}

func (cn *commandNode) Parse(context *context) {
	cn.node = &primitiveCommandNode{}
	cn.node.Parse(context)
}

func (cn *commandNode) ToString() string {
	return cn.node.ToString()
}

type primitiveCommandNode struct {
	name string
}

func (pcn *primitiveCommandNode) Parse(context *context) {
	pcn.name = context.currentToken
	context.skipToken(pcn.name)
}

func (pcn *primitiveCommandNode) ToString() string {
	return pcn.name + " "
}

type context struct {
	text         string
	tokens       []string
	currentToken string
}

func NewContext(text string) *context {
	ctx := &context{
		text:   text,
		tokens: strings.Fields(text),
	}
	ctx.nextToken()
	return ctx
}

func (c *context) nextToken() string {
	if len(c.tokens) == 0 {
		c.currentToken = ""
	} else {
		c.currentToken = c.tokens[0]
		c.tokens = c.tokens[1:]
	}
	return c.currentToken
}

func (c *context) skipToken(token string) {
	c.nextToken()
}

func main() {
	fmt.Println("interpreter")

	node := ProgramNode{}
	node.Parse(NewContext("program go right end"))
	fmt.Println(node.ToString())
}
