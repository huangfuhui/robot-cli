package main

import (
	"fmt"
	"os"
	"net/http"
	"net/url"
	"io/ioutil"
)

func main() {
	defer func() {
			if err := recover(); err != nil {
				fmt.Println("发生了一些未知错误...")
			}
		}()

	argsQuantity := len(os.Args)
	if argsQuantity < 2 {
		fmt.Println("请输入参数")
		return
	}

	command := ""
	for _, v := range os.Args[2:] {
		command += v + " "
	}

	// 拼装参数
	requestUrl, _ := url.Parse("")
	query := requestUrl.Query()
	query.Set("handset", os.Args[1])
	query.Set("command", command)
	requestUrl.RawQuery = query.Encode()

	// 发起请求
	resp, err := http.Get(requestUrl.String());
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	// 解析结果
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
 
 	// 输出结果
	fmt.Println(string(body))
}
