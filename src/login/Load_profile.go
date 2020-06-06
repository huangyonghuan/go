package main

import (
	"fmt"
	"io/ioutil"
)
import "encoding/json"

type login_conf struct {
	Username string
	Password string
}
type Result struct {
	Version string
	Login   login_conf
}

func conf() (string, string) {

	// json序列化成变成对象文件-----------------------------------------------------begin
	bytes, err := ioutil.ReadFile("./conf/host.json")

	var r Result
	err = json.Unmarshal([]byte(bytes), &r)
	if err != nil {
		fmt.Printf("err was %v", err)
	}

	//fmt.Println(r.Version)
	//fmt.Println(r.Login.Username)
	//fmt.Println(r.Login.Password)
	return r.Login.Username, r.Login.Password
}
