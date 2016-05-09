package vo

import (
	"util"
	"os"
	"strconv"
)

var (
	// redis 和 web Server的Ip地址，默认为localhost
	Ip string
	// redis 和 web Server的Ip地址，默认为9090
	Port string
	//  商品标识前缀
	Product_Pre string
	//  要卖的商品数量
	Product1_Max_Num int
	//  第一种商品productId
	Product1_Query_Name string
	//  第-种商品key 数据库标识
	Product1_Query_String string
    //  商品是否卖完的标识，优化使用
	Flag bool
    // 如上
	Product2_Max_Num int
	Product2_Query_Name string
	Product2_Query_String string

	Product3_Max_Num int
	Product3_Query_Name string
	Product3_Query_String string
)

func init() {
	//读取properties 文件中的各种配置参数
	myConfig := new(util.Config)
	path,_ := os.Getwd()
	path += "/conf/properties"
	myConfig.InitConfig(path)

	Ip = myConfig.Read("common", "ip")
	Port = myConfig.Read("common", "port")
	Product_Pre = myConfig.Read("common", "query_prefix")

	Product1_Max_Num, _ = strconv.Atoi(myConfig.Read("product_1", "max_num"))
	Product1_Query_Name = myConfig.Read("product_1", "query_name")
	Product1_Query_String = myConfig.Read("product_1", "total_query_name")

	Product2_Max_Num, _ = strconv.Atoi(myConfig.Read("product_2", "max_num"))
	Product2_Query_Name = myConfig.Read("product_2", "query_name")
	Product2_Query_String = myConfig.Read("product_2", "total_query_name")

	Product3_Max_Num, _ = strconv.Atoi(myConfig.Read("product_3", "max_num"))
	Product3_Query_Name = myConfig.Read("product_3", "query_name")
	Product3_Query_String = myConfig.Read("product_3", "total_query_name")

	Flag = true
}