package service

import (
	"dao"
	"strings"
	"vo"
	"fmt"
)

//根据商品id查询成功秒杀的所有用户id和用户购买商品的具体编号
func ServiceQueryProductSeckillingInfo(productid string) *vo.ResultProductMsg {
	returnMsg := &vo.ResultProductMsg{0, nil}
	if (productid != vo.Product1_Query_Name) {
		fmt.Println("errMsg:", "productid不存在")
		panic("productid不存在")
	}

	productInfo, _ := dao.RedisPoolOne.LRange(vo.Product1_Query_String)

	if productInfo == nil {
		fmt.Println("errMsg:", "无法查询到结果")
		panic("无法查询到结果")
	}
	goodsList := []vo.KillEntry{}
	for _, entry := range productInfo {
		tmp := strings.Split(entry, "*")
		userid := tmp[0]
		goodsid := tmp[1]
		killEntry := vo.KillEntry{userid, goodsid}
		goodsList = append(goodsList, killEntry)
	}
	returnMsg.SetErrno(0)
	returnMsg.SetList(goodsList)
	return returnMsg

}
