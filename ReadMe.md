curl

```
curl 'http://mobile.yangkeduo.com/pincard_ask.html?__rp_name=brand_amazing_price_group&group_order_id=2007363467928373305' \
  -H 'authority: mobile.yangkeduo.com' \
  -H 'accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9' \
  -H 'accept-language: zh-CN,zh;q=0.9' \
  -H 'cache-control: no-cache' \
  -H 'cookie: api_uid=Ck65RGMhdYh/cwBmFYBmAg==; _nano_fp=XpEjXqgqnqgjX0djno_xLotYFs6ogewiy3j5kQ2c; webp=1; dilx=lyG2PI9QYwQLBpyHPKhTL; jrpl=K1SxOS15bF2QLtbRNJt2aHw8tdcaImYb; njrpl=K1SxOS15bF2QLtbRNJt2aHw8tdcaImYb; PDDAccessToken=WKB4AXQYL52AKR5GMMQIBADLMCXA73ASJSAVAYIJLS72NRHU5LRQ110e9a9; pdd_user_id=8646712553; pdd_user_uin=2OR2ZROLQAR4QU4OUXXJ4SUY54_GEXDA; rec_list_brand_sale_mall=rec_list_brand_sale_mall_ThoP8h; pdd_vds=gaULRORQJtkQgyKOFIHyULMIkOHIHQgQMEktRGSLHEXiJnjiWtJORnMogyHb; rec_list_brand_amazing_price_group=rec_list_brand_amazing_price_group_SQnnix' \
  -H 'pragma: no-cache' \
  -H 'sec-ch-ua: "Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "Windows"' \
  -H 'sec-fetch-dest: document' \
  -H 'sec-fetch-mode: navigate' \
  -H 'sec-fetch-site: none' \
  -H 'sec-fetch-user: ?1' \
  -H 'upgrade-insecure-requests: 1' \
  -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36' \
  --compressed
```
返回html:
```
...
<script>window.rawData={"store":{"errorRedirectUrl":"\u002F","isServerRendered":true,"isFinishInitLoading":true,"query":{"group_order_id":"2007363467928373305"},"isIOSSystem":false,"immersiveStore":{"immersiveEnable":false,"navHeight":0,"statusBarHeight":0,"isInitedOnClient":false},"$service":{"reportMap":{"serviceName":"BrandSalesMallService"}},"bannerUrl":"","title":"","extData":{},"tabList":[],"activeIndex":0,"listIDSurfix":"rec_list_brand_sale_mall_3mH92B","loadMoreLock":false,"$serverPmmPerfLogger":{"apiAxiosPlugin":{}},"userID":8646712553,"hostname":"mobile.yangkeduo.com","webpEnable":true,"canCancelWebpOnIOS":true,"currentPlatform":"unknown","isNativePlatform":false,"isAndroidPlatform":false,"isNativeAndroidPlatform":false,"isIOSPlatform":false,"isNativeIOSPlatform":false,"isTinyNativePlatform":false,"isWeChatMiniProgram":false,"isAndroidWeChatPlatform":false,"isPureWeChatPlatform":false,"appVersion":""}};</script>
...
```


提取信息:

        {"goodsInfo":{
            "goodsId":298497243673,
            "hdThumbUrl":"https:\u002F\u002Fimg.pddpic.com\u002Fmms-material-img\u002F2022-10-12\u002F7c24b6e8fb2978e3f9e9ad0058811541.jpeg",
            "goodsName":"【5人团】清风原木抽纸110抽24包纸巾批发整箱面巾纸卫生纸抽纸",
            "linkUrl":"goods.html?goods_id=298497243673&page_from=51&_oc_brand_neigou_param=_8_44",
            "activityPriceWord":"5人拼单价",
            "activityPrice":"24.9",
            "flagSet":Object{...},
            "originPriceWord":"全网低价",
            "originPrice":"36.6",
            "isPriceHidden":false,
            "customerNumWording":"5人团",
            "salesStr":"2.3万位用户已成功拼团",
            "prec":"\"\"",
            "payQuantityPercent":11,
            "payQuantityPercentStr":"已售11%"
        },
        "extraInfo":{
            "activityStatus":0,
            "groupList":Array[1],
            "btnText":"邀请好友拼单",
            "btnStatus":1,
            "total":5,
            "focusText":"",
            "groupStatus":0,
            "isSelfGroup":true,
            "briefRuleText":"需拉满相应人数成团，团满发货，不满自动退款",
            "usePasswordShare":false,
            "isFreeGroup":false
        }}