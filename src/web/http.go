package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main()  {
	//生成clinet 参数为默认值
	clinet := &http.Client{}

	//生成要访问的url
	url := "http://10.0.0.200:9090"
	//提交请求
	reqest,err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	//处理返回结果
	reponse, _ := clinet.Do(reqest)

	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应处理
	stdout :=os.Stdout
	_, err = io.Copy(stdout, reponse.Body)

	//返回的状态码
	status := reponse.StatusCode
	fmt.Println(status)
}
