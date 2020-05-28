package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	var g = []
	g,err := ioutil.ReadFile("file/test")
	if err != nil {
		fmt.Println("read fail", err)
	}
}
