package main

import (
	"PDD_Muti/data"
	"github.com/gin-gonic/gin"
	"strconv"
)

var DB map[int64]*data.Good
var AllOrderFilter map[int64]*data.Order

func init() {
	DB = make(map[int64]*data.Good)
	AllOrderFilter = make(map[int64]*data.Order)
	DBLoad()
}

func main() {
	go DBSave()

	r := gin.Default()
	r.Static("/static", "./html")
	r.GET("/GetGoods", GetGoods)
	r.GET("/GetOrders", GetOrders)
	r.GET("/AddOrder", AddOrder)
	r.GET("/SaveDB", SaveDB)
	r.Run(":80")
}
func SaveDB(c *gin.Context) {
	ClearEndTimeMsOrder()
	DoDBSave()
}

func GetGoods(c *gin.Context) {
	c.JSON(200, DB)
}

func GetOrders(c *gin.Context) {
	good_id, _ := strconv.ParseInt(c.Query("good_id"), 10, 64)
	if good, ok := DB[good_id]; ok {
		c.JSON(200, good.Orders)
	}
}

func AddOrder(c *gin.Context) {
	order_id, _ := strconv.ParseInt(c.Query("order_id"), 10, 64)
	if AllOrderFilter[order_id] != nil {
		c.JSON(201, "已经添加过了哦")
		return
	}
	order := &data.Order{OrderID: order_id}
	AllOrderFilter[order_id] = order
	good, err := GetGoodInfo(order)
	if err != nil {
		c.JSON(201, err)
		return
	}
	if good == nil {
		c.JSON(201, "服务器解析失败")
		return
	}
	if _, ok := DB[good.Goods_id]; !ok {
		DB[good.Goods_id] = good
	}
	good = DB[good.Goods_id]
	good.Orders = append(good.Orders, *order)
	c.JSON(200, good.Orders)
}
