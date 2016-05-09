package service

import (
	"dao"
	"encoding/json"
	"fmt"
	"strconv"
	"vo"
)

func ServiceSeckilling(userid ,productid string) *vo.ReturnMsg {
	message := &vo.ReturnMsg{0, ""} //返回消息
	if !vo.Flag {
		message.SetErrno(0)
		message.SetErrMsg("秒杀失败")
		return message
	}
	if value, err := dao.RedisPoolOne.Get(vo.Product_Pre + vo.Product1_Query_Name); err == nil {
		if count, _ := strconv.Atoi(string(value)); count >= vo.Product1_Max_Num {
			vo.Flag = false
			message.SetErrno(0)
			message.SetErrMsg("秒杀失败")
		} else {
			entry := &vo.QueueEntry{"", "", ""}
			if (productid == vo.Product1_Query_Name || productid == vo.Product2_Query_Name || productid == vo.Product3_Query_Name) {
				entry.SetProductid(vo.Product_Pre + productid)
			} else {
				fmt.Println("errMsg:", "productid不存在")
				panic("productid不存在")
			}
			entry.SetUserid(userid)
			if entry.GetUserid() == "" || entry.GetProductid() == "" {
				fmt.Println("errMsg:", "参数错误")
				panic("参数错误")
			}
			if str, err := json.Marshal(entry); err == nil {
				redisError := dao.RedisPoolOne.RPush("list", string(str))
				if redisError != nil {
					fmt.Println("errMsg:", redisError)
					panic(redisError.Error())
				}
				message.SetErrno(0)
				message.SetErrMsg("秒杀中，请稍后查询")
			} else {
				fmt.Println("errMsg:", err)
				panic(err.Error())
			}
		}
	} else {
		fmt.Println("errMsg:", err)
		panic(err.Error())
	}
	return message
}
