package data

type Good struct {
	Goods_id           int64   `json:"goods_id"`
	GoodsName          string  `json:"goodsName"`
	ActivityPrice      string  `json:"activityPrice"`
	CustomerNumWording string  `json:"customerNumWording"`
	Orders             []Order `json:"-"`
}

type Order struct {
	Goods_id  int64  `json:"goods_id"`
	OrderID   string `json:"orderID"`
	UserID    int64  `json:"userID"`
	EndTimeMs int64  `json:"endTimeMs"`
}
