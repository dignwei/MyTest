package main

import (
	"controller"
	"log"
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	http.HandleFunc("/zaixianshang/queryUserSeckillingInfo", controller.QueryUserSeckillingInfo)       //设置访问的路由
	http.HandleFunc("/zaixianshang/seckilling", controller.Seckilling)                                 //设置访问的路由
		http.HandleFunc("/zaixianshang/queryProductSeckillingInfo", controller.QueryProductSeckillingInfo) //设置访问的路由
	err := http.ListenAndServe(":9090", nil)                                                           //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
