package main

import (
	"fmt"
	"net/http"
	"github.com/Unknwon/log"
	"io/ioutil"
	"os"
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
	f, err := os.Open("/Users/klook/git/gotest/baidu.html")
	if err != nil {
		log.Info("open file error, err:%v", err)
		return
	}
	fmt.Println(f.Name())
	//err = f.Chmod(666)
	//if err != nil {
	//	log.Info("chmod error, err:%v", err)
	//	return
	//}
	err = ioutil.WriteFile(f.Name(), b, 0666)
	if err != nil {
		log.Info("write file error, err:%v", err)
		return
	}
	fmt.Println("add commit1 ")
	fmt.Println("add commit2 ")
	fmt.Println("add commit3 ")
	fmt.Println("debug commit1")
	fmt.Println("debug commit2")
	fmt.Println("debug commit3")
}
