package main
import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Boby []byte
}

func (p *Page) save() error  {
	filename := p.Title + ".txt"
	return  ioutil.WriteFile(filename,p.Boby,0600)
}

func loadPage(title string)  *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return  &Page{Title:title,Boby:boby}
}