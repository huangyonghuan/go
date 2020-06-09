package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func TestHandle(res http.ResponseWriter, req *http.Request) {
	entry_name := req.FormValue("admin")
	entry_password := req.FormValue("password")
	var s, i = conf()
	if entry_name != s || entry_password != i {
		//if admin != "admin" || password != "123456" {
		res.Write([]byte("Login Fail,Please Try Again!"))
		//}
	} else {
		_, _ = res.Write([]byte("test for web : " + entry_name + "	" + entry_password))
	}
}

func main() {
	port := "80"
	//log.Printf(os.Getwd())
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		t, err := template.ParseFiles("./html/login2.html")
		//t, err := template.ParseFiles("./web/index.html")
		//t, err :=template.ParseFiles("C:\\Users\\huawei\\go\\src\\login\\web\\index.html")
		if err != nil {
			log.Println("err")
		}
		t.Execute(res, nil)
	})
	http.HandleFunc("/Login", TestHandle)
	fmt.Println("start http server at:", port)
	http.ListenAndServe(":"+port, nil)
}
