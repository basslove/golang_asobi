package main

import "fmt"

type Display interface {
	GetColumns() int
	GetRows() int
	GetRowText(row int) string
	Show(display Display) string
}

type baseDisplay struct {
	Display
}

func (bd *baseDisplay) Show(display Display) string {
	text := ""
	for i := 0; i < display.GetRows(); i++ {
		text += display.GetRowText(i)
	}
	return text
}

type StringDisplay struct {
	*baseDisplay
	text string
}

func NewStringDisplay(text string) *StringDisplay {
	return &StringDisplay{&baseDisplay{}, text}
}

func (sd *StringDisplay) GetColumns() int {
	return len(sd.text)
}

func (sd *StringDisplay) GetRows() int {
	return 1
}

func (sd *StringDisplay) GetRowText(row int) string {
	if row == 0 {
		return sd.text
	} else {
		return ""
	}
}

type border struct {
	*baseDisplay
	display Display
}

type SideBorder struct {
	*border
	borderChar string
}

func NewSideBorder(display Display, borderChar string) *SideBorder {
	return &SideBorder{&border{display: display}, borderChar}
}

func (sb *SideBorder) GetColumns() int {
	return len(sb.borderChar)*2 + sb.display.GetColumns()
}

func (sb *SideBorder) GetRows() int {
	return sb.display.GetRows()
}

func (sb *SideBorder) GetRowText(row int) string {
	return sb.borderChar + sb.display.GetRowText(row) + sb.borderChar
}

func main() {
	fmt.Println("decorator")

	d1 := NewStringDisplay("hello world")
	result := d1.Show(d1)
	fmt.Println(result)

	d2 := NewSideBorder(d1, "####")
	result = d2.Show(d2)
	fmt.Println(result)
}
