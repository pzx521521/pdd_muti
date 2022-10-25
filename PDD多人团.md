# pdd多人团分享页面

## 背景:

pdd多人团,很烦尽管里面的东西很便宜, 但是很烦, 于是做了这个pdd分享界面, 附上源码

## 使用 

+ 打开[分享页面](https://pdd.parap.ml/pdd/) 
+ **如果想加入别人的订单**->搜索自己想要的东西->点击所有订单->点击显示二维码, 扫码即可
+ **如果想分享自己的订单**->分享带二维码的图片->添加订单->你将可以看到自己的订单出现在分享界面上
  + 可以通过手机摄像头直接扫面分享图片的二维码
  + 可以直接把图片拖拽到上传图片的窗口
  + 可以直接输入分享页面的订单id

## 原理分析

### 通过分享二维码获取产品信息及其结束时间

url:

```
http://mobile.yangkeduo.com/pincard_ask.html?__rp_name=brand_amazing_price_group&group_order_id=${group_order_id}
```

部分cookie:

```
	pdd_user_id 用户id 对于单个用户是不会变的
	pdd_user_uin 用户全局id 对于单个用户是不会变的
	PDDAccessToken 用户权健 每几天会变,有过期时间
```

### 存在问题

1. 貌似不是每一个订单都可以参与的, 大部分都是`活动未开放`
2. Cookie 过期的时间没做特殊测试, 由于pdd_app 使用ssl, 手机没有root,抓包没有抓到对应的信息, 采用的是wx小程序的cookie
3. bug肯定是有的

### 技术栈:

前端: vue + elementui

后端; golang 内存数据库

### [源码地址](https://github.com/pzx521521/pdd_muti/)及使用

+ dokcer 部署详见dockerfile

+ 修改 assets/PDDAccessToken 为自己的

+ 修改 `Util.go`中的pdd_user_id, pdd_user_uin  为自己的