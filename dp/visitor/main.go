package main

import (
	"fmt"
	"strconv"
)

type Visitor interface {
	VisitFile(file *file) string
	VisitDir(directory *directory) string
}

type Element interface {
	Accept(visitor Visitor) string
}

type entry interface {
	Element
	GetName() string
	GetSize() int
	Add(entry entry)
}

type initEntry struct {
	entry
	name string
}

func (ie *initEntry) GetName() string {
	return ie.name
}

func (ie *initEntry) print(entry entry) string {
	return entry.GetName() + " (" + strconv.Itoa(entry.GetSize()) + ")\n"
}

type file struct {
	*initEntry
	size int
}

func (f *file) GetSize() int {
	return f.size
}

func (f *file) Add(entry entry) {}

func (f *file) Accept(visitor Visitor) string {
	return visitor.VisitFile(f)
}

type directory struct {
	*initEntry
	dir []entry
}

func (d *directory) GetSize() int {
	size := 0
	for _, dir := range d.dir {
		size += dir.GetSize()
	}
	return size
}

func (d *directory) Add(entry entry) {
	d.dir = append(d.dir, entry)
}

func (d *directory) Accept(visitor Visitor) string {
	return visitor.VisitDir(d)
}

type listVisitor struct {
	currentDir string
}

func (lv *listVisitor) VisitFile(file *file) string {
	return lv.currentDir + "/" + file.print(file)
}

func (lv *listVisitor) VisitDir(directory *directory) string {
	saveDir := lv.currentDir
	list := lv.currentDir + "/" + directory.print(directory)
	lv.currentDir += "/" + directory.GetName()

	for _, dir := range directory.dir {
		list += dir.Accept(lv)
	}
	lv.currentDir = saveDir
	return list
}

func NewFile(name string, size int) *file {
	return &file{
		initEntry: &initEntry{name: name},
		size:      size,
	}
}
func NewDirectory(name string) *directory {
	return &directory{initEntry: &initEntry{name: name}}
}

func main() {
	rootDir := NewDirectory("root")
	usrDir := NewDirectory("usr")
	fileA := NewFile("new-A", 1)

	rootDir.Add(usrDir)
	rootDir.Add(fileA)

	fileB := NewFile("new-B", 2)
	usrDir.Add(fileB)

	fmt.Println(rootDir.Accept(&listVisitor{}))
}
