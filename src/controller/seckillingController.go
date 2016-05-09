package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service"
	"vo"
	"strings"
)

var counter = 0

func Seckilling(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm() //解析参数，默认是不会解析的
	counter += 1
	fmt.Println(counter)
	message := &vo.ReturnMsg{0, ""}
	defer func() {//异常处理
		if err := recover(); err != nil {
			message.SetErrno(1)
			message.SetErrMsg("出错了！")
		}
		if jsonstr, jsonerr := json.Marshal(message); jsonerr == nil {
			fmt.Fprintf(resp, string(jsonstr))
		} else {
			fmt.Fprintf(resp, "json错误")
		}
	}()
	userid, productid := "", ""
	for key, value := range req.Form {
		if key == "userid" {
			userid = strings.Join(value, "")
		} else if key == "productid" {
			productid = strings.Join(value, "")
		}
		fmt.Print("key:", key)
		fmt.Println("  value:", strings.Join(value, ""))
	}
	message = service.ServiceSeckilling(userid, productid)
}
