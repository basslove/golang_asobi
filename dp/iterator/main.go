package main

import (
	"errors"
	"fmt"
)

var ErrNoSuchElement = errors.New("no such element")

type Iterable interface {
	Iterator()
}

type Iterator interface {
	HasNext() bool
	Next() (*Book, error)
}

type Book struct {
	name string
}

func NewBook(name string) *Book {
	return &Book{name: name}
}

func (bs *Book) GetName() string {
	return bs.name
}

type BookShelf struct {
	books []*Book
	last  int
}

func NewBookShelf(maxSize int) *BookShelf {
	return &BookShelf{books: make([]*Book, 0, maxSize), last: 0}
}

func (bs *BookShelf) GetBookAt(index int) *Book {
	return bs.books[index]
}

func (bs *BookShelf) AppendBook(book Book) {
	bs.books = append(bs.books, &book)
	bs.last++
}

func (bs *BookShelf) GetLength() int {
	return bs.last
}

func (bs *BookShelf) Iterator() *BookShelfIterator {
	return NewBookShelfIterator(*bs)
}

type BookShelfIterator struct {
	bookShelf BookShelf
	index     int
}

func NewBookShelfIterator(bookShelf BookShelf) *BookShelfIterator {
	return &BookShelfIterator{bookShelf: bookShelf, index: 0}
}

func (bsi *BookShelfIterator) HasNext() bool {
	if bsi.index < bsi.bookShelf.GetLength() {
		return true
	} else {
		return false
	}
}

func (bsi *BookShelfIterator) Next() (*Book, error) {
	if !bsi.HasNext() {
		return nil, ErrNoSuchElement
	}
	book := bsi.bookShelf.GetBookAt(bsi.index)
	bsi.index++

	return book, nil
}

func main() {
	fmt.Println("iterator")

	bookShelf := NewBookShelf(4)
	bookShelf.AppendBook(*NewBook("no.1"))
	bookShelf.AppendBook(*NewBook("no.2"))
	bookShelf.AppendBook(*NewBook("no.3"))
	bookShelf.AppendBook(*NewBook("no.4"))

	var it Iterator = bookShelf.Iterator()
	for it.HasNext() {
		book, err := it.Next()
		if err != nil {
			panic(err)
		}
		if book != nil {
			fmt.Println(book.GetName())
		}
	}
	fmt.Println()
}
