package main

import (
	"fmt"
	"net/http"
	"github.com/Unknwon/log"
	"io/ioutil"
)

//test
func main() {
	fmt.Println("this is a git test!")
	fmt.Println("add in dev-test4!")
	fmt.Println("add in dev-test3!")
	client := http.Client{}

	resp, err := client.Get("http://www.baidu.com")
	if err != nil {
		log.Info("get error, err:%v", err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Info("read body error, err:%v", err)
		return
	}
	fmt.Println(string(b))
}
