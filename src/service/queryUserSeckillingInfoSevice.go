package service

import (
	"dao"
	"strconv"
	"vo"
	"fmt"
)

/* 
 *   用户查询自己的秒杀结果接口
 *
 * errno  : 1
 *        status :2 参数错误
 * errno  : 0
 *        status :0  秒杀还未开始 ,商品被卖了
 *        status :1  秒杀成功    , 成功秒杀到，redis中查询到goodsId
 *        status :2  秒杀失败    , 没有秒杀到，redis中未查询到goodsId，且商品已经被卖完
 *        status :3  在秒杀中    , 没有秒杀到，redis中未查询到goodsId，但是商品还未卖完
 *
 */
func QueryUserSeckillingInfo(userid ,productid string) *vo.ResultPersonMsg {
	retMessage := &vo.ResultPersonMsg{0, "", ""}
	if userid == "" || productid == "" {
		fmt.Println("errMsg:", "参数错误")
		panic("参数错误")
	}
	if (productid != vo.Product1_Query_Name ) {
		fmt.Println("errMsg:", "productid不存在")
		panic("productid不存在")
	}
	value, _ := dao.RedisPoolOne.Get(vo.Product_Pre + vo.Product1_Query_Name)
	countGoodsSold, _ := strconv.Atoi(string(value))
	if 0 == countGoodsSold {
		retMessage.SetErrno(0)
		retMessage.SetStatus("0")
		retMessage.SetGoodsId("秒杀未开始...")
		return retMessage
	}
	goodsid, _ := dao.RedisPoolOne.Get(userid)
	if goodsid != "" { // 秒杀成功
		retMessage.SetErrno(0)
		retMessage.SetStatus("1")
		retMessage.SetGoodsId(goodsid)
	} else { // 秒杀失败
		if countGoodsSold < vo.Product1_Max_Num {
			retMessage.SetErrno(0)
			retMessage.SetStatus("3")
			retMessage.SetGoodsId("在秒杀中，请稍后查询...")
			return retMessage
		}
		retMessage.SetErrno(0)
		retMessage.SetStatus("2")
		retMessage.SetGoodsId("秒杀失败,未秒杀到商品")
	}
	return retMessage
}
