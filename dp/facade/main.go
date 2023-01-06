package main

import "fmt"

var db = map[string]string{
	"sample1@gmail.com": "sample1",
	"sample2@gmail.com": "sample2",
}

type database struct {
}

func (d *database) getNameByMail(mail string) string {
	return db[mail]
}

type mdWriter struct {
}

func (mw *mdWriter) title(title string) string {
	return "# Welcome to " + title + "'s page!"
}

type PageMaker struct {
}

func (pm *PageMaker) MakeWelcomePage(mail string) string {
	database := database{}
	writer := mdWriter{}

	name := database.getNameByMail(mail)
	page := writer.title(name)

	return page
}

func main() {
	maker := PageMaker{}
	fmt.Println(maker.MakeWelcomePage("sample1@gmail.com"))
}
