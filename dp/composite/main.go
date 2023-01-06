package main

import (
	"fmt"
	"strconv"
)

type Entry interface {
	GetName() string
	GetSize() int
	PrintList(prefix string) string
	Add(entry Entry)
}

var _ Entry = &file{}
var _ Entry = &directory{}

type InitEntry struct {
	Entry
	name string
}

func (ie *InitEntry) GetName() string {
	return ie.name
}

func (ie *InitEntry) print(entry Entry) string {
	return entry.GetName() + " (" + strconv.Itoa(entry.GetSize()) + ")\n"
}

type file struct {
	*InitEntry
	size int
}

func (f *file) GetSize() int {
	return f.size
}

func (f *file) PrintList(prefix string) string {
	return prefix + "/" + f.print(f)
}

func (f *file) Add(entry Entry) {}

type directory struct {
	*InitEntry
	dir []Entry
}

func (d *directory) GetSize() int {
	size := 0
	for _, dir := range d.dir {
		size += dir.GetSize()
	}
	return size
}

func (d *directory) Add(entry Entry) {
	d.dir = append(d.dir, entry)
}

func (d *directory) PrintList(prefix string) string {
	list := prefix + "/" + d.print(d)
	for _, dir := range d.dir {
		list += dir.PrintList(prefix + "/" + d.GetName())
	}
	return list
}

func NewFile(name string, size int) *file {
	return &file{
		InitEntry: &InitEntry{name: name},
		size:      size,
	}
}
func NewDirectory(name string) *directory {
	return &directory{InitEntry: &InitEntry{name: name}}
}

func main() {
	fmt.Println("composite")

	rootDir := NewDirectory("root")
	usrDir := NewDirectory("usr")
	fileA := NewFile("A", 1)

	rootDir.Add(usrDir)
	rootDir.Add(fileA)

	fileB := NewFile("B", 2)
	usrDir.Add(fileB)

	result := rootDir.PrintList("")
	fmt.Println(result)
}
