package main

import "fmt"

type item interface {
	toString() string
}

var _ item = &mdLink{}
var _ item = &mdTray{}

type link interface {
	item
}

type tray interface {
	item
	AddToTray(item item)
}

type baseTray struct {
	tray []item
}

func (bt *baseTray) AddToTray(item item) {
	bt.tray = append(bt.tray, item)
}

type page interface {
	AddToContent(item item)
	Output() string
}

type basePage struct {
	content []item
}

func (bp *basePage) AddToContent(item item) {
	bp.content = append(bp.content, item)
}

type factory interface {
	CreateLink(caption, url string) link
	CreateTray(caption string) tray
	CreatePage(title, author string) page
}

type mdLink struct {
	caption, url string
}

func (ml *mdLink) toString() string {
	return "[" + ml.caption + "](" + ml.url + ")"
}

type mdTray struct {
	baseTray
	caption string
}

func (mt *mdTray) toString() string {
	tray := "- " + mt.caption + "\n"
	for _, item := range mt.tray {
		tray += item.toString() + "\n"
	}
	return tray
}

type mdPage struct {
	basePage
	title, author string
}

func (mp *mdPage) Output() string {
	content := "title: " + mp.title + "\n"
	content += "author: " + mp.author + "\n"
	for _, item := range mp.content {
		content += item.toString() + "\n"
	}
	return content
}

type MdFactory struct{}

func (mf *MdFactory) CreateLink(caption, url string) link {
	return &mdLink{caption, url}
}

func (mf *MdFactory) CreateTray(caption string) tray {
	return &mdTray{caption: caption}
}

func (mf *MdFactory) CreatePage(title, author string) page {
	return &mdPage{title: title, author: author}
}

func main() {
	fmt.Println("abstract factory")
	factory := MdFactory{}

	usYahoo := factory.CreateLink("Yahoo!", "http://www.yahoo.com")
	jaYahoo := factory.CreateLink("Yahoo!Japan", "http://www.yahoo.co.jp")

	tray := factory.CreateTray("Yahoo!")
	tray.AddToTray(usYahoo)
	tray.AddToTray(jaYahoo)

	page := factory.CreatePage("Title", "Author")
	page.AddToContent(tray)

	output := page.Output()

	expect := "title: Title\nauthor: Author\n- Yahoo!\n[Yahoo!](http://www.yahoo.com)\n[Yahoo!Japan](http://www.yahoo.co.jp)\n\n"

	if output != expect {
		fmt.Printf("Expect output to %s, but %s\n", expect, output)
	}
}
