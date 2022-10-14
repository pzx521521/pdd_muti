package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"pdd_muti/assets"
	"pdd_muti/data"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	CST_PPD_PRE = "http://mobile.yangkeduo.com/"
	//一定要加 __rp_name=brand_amazing_price_group 否则会没有信息
	CST_PPD_PINCARD_ASK = CST_PPD_PRE + "pincard_ask.html?__rp_name=brand_amazing_price_group&group_order_id="
	CST_PPD_GROUP       = "group.html?group_order_id="
)

func GetCookie() string {
	pdd_user_id := "8646712553"
	pdd_user_uin := "2OR2ZROLQAR4QU4OUXXJ4SUY54_GEXDAv"
	//pdd_user_id pdd_user_uin 都是不会改变的
	//PDDAccessToken 每次的登录都会改变
	//由于pdd 不能同时用浏览器和手机使用
	//这里使用抓包, 抓的微信小程序的PDDAccessToken
	//是否有过期时间  暂时不知道
	//此文件不做上传, 防止账号泄漏
	PDDAccessToken := strings.Trim(string(assets.PDDAccessToken), "\n")

	cookies := []string{
		"api_uid=Ck65RGMhdYh/cwBmFYBmAg==",
		" _nano_fp=XpEjXqgqnqgjX0djno_xLotYFs6ogewiy3j5kQ2c",
		" webp=1",
		" dilx=lyG2PI9QYwQLBpyHPKhTL",
		" jrpl=K1SxOS15bF2QLtbRNJt2aHw8tdcaImYb",
		" njrpl=K1SxOS15bF2QLtbRNJt2aHw8tdcaImYb",
		" rec_list_brand_sale_mall=rec_list_brand_sale_mall_c0bsuu",
		" PDDAccessToken=" + PDDAccessToken,
		" pdd_user_id=" + pdd_user_id,
		" pdd_user_uin=" + pdd_user_uin,
		" pdd_vds=gaLlNnNEQyosGEonOEaOPsmwQsLGatoOymbGOEQOLGGmmaIsImytNyEGilPa",
		" rec_list_brand_amazing_price_group=rec_list_brand_amazing_price_group_TmIHQd",
	}
	ans := ""
	for _, cookie := range cookies {
		ans = ans + cookie + ";"
	}
	return ans
}
func getHtml(group_order_id int64) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", CST_PPD_PINCARD_ASK+strconv.FormatInt(group_order_id, 10), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("authority", "mobile.yangkeduo.com")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("cookie", GetCookie())
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyText, nil
}

func analsysHtml(html []byte, order *data.Order) (*data.Good, error) {
	m := make(map[string]interface{})
	good := &data.Good{}
	reg := regexp.MustCompile(`<script>window.rawData=(.*?);</script>`)
	matchs := reg.FindSubmatch(html)
	if len(matchs) < 2 {
		return nil, errors.New("analsysHtml error!")
	}
	match := matchs[1]
	err := json.Unmarshal(match, &m)
	if err != nil {
		return nil, err
	}
	store := m["store"].(map[string]interface{})
	if store == nil {
		return nil, errors.New("analsys store error!")
	}
	userID := store["userID"].(float64)
	order.UserID = int64(userID)
	endTimeMs := store["endTimeMs"].(float64)
	order.EndTimeMs = int64(endTimeMs)

	if err != nil {
		return nil, errors.New("analsys userID error!")
	}
	goodsInfo := store["goodsInfo"].(map[string]interface{})
	good.GoodsName = goodsInfo["goodsName"].(string)
	good.GoodsName = strings.Trim(good.GoodsName, " ")
	linkUrl := goodsInfo["linkUrl"].(string)
	//goods.html?goods_id=298497243673&page_from=51&_oc_brand_neigou_param=_8_44
	reg = regexp.MustCompile(`goods_id=(\d+)`)
	ans := reg.FindStringSubmatch(linkUrl)
	if len(ans) < 2 {
		return nil, errors.New("analsys goods_id error!")
	}
	good.Goods_id, err = strconv.ParseInt(ans[1], 10, 64)
	if err != nil {
		return nil, errors.New("analsys goods_id error!")
	}
	order.Goods_id = good.Goods_id
	good.ActivityPrice = goodsInfo["activityPrice"].(string)

	good.CustomerNumWording = goodsInfo["customerNumWording"].(string)

	return good, nil
}

func GetGoodInfo(order *data.Order) (*data.Good, error) {
	html, err := getHtml(order.OrderID)
	if err != nil {
		return nil, err
	}
	good, err := analsysHtml(html, order)
	if err != nil {
		return nil, err
	}
	return good, nil
}

func DBSave() {
	t := time.NewTicker(time.Hour)
	defer t.Stop()
	for {
		<-t.C
		ClearEndTimeMsOrder()
		DoDBSave()
	}
}

func ClearEndTimeMsOrder() {
	now := time.Now().Unix() * 1000
	for _, good := range DB {
		for i2 := len(good.Orders) - 1; i2 >= 0; i2-- {
			order := good.Orders[i2]
			if order.EndTimeMs < now {
				good.Orders = append(good.Orders[:i2], good.Orders[i2+1:]...)
			}
		}
	}
}

func DoDBSave() {
	buf := bytes.Buffer{}
	if err := gob.NewEncoder(&buf).Encode(DB); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	os.Remove("./db.db")
	os.WriteFile("./db.db", buf.Bytes(), 0666)
}

func DBLoad() {
	data, err := os.ReadFile("./db.db")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	decoder := gob.NewDecoder(bytes.NewBuffer(data))
	// 进行反序列化
	err = decoder.Decode(&DB)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%v\n", DB)
}
