package service

import (
	"strings"

	"git.zx-tech.net/pengfeng/amazon.seller/config"
	"git.zx-tech.net/pengfeng/amazon.seller/http"
	"git.zx-tech.net/pengfeng/amazon.seller/model"
)

//通用方法  more 第一个位置是 文件字符串的编码转换格式

func CommonHandle(method, action, uri, placeID string, auth model.AuthObject, m1 map[string]string, ret string, res interface{}, more ...string) (err error) {
	i := initParam(action)
	m2 := AmazonAuth(auth).Params(i, m1)
	str := AmazonAuth(auth).Sort(m2)
	var c config.AmazonSellerConfig
	url := c.GetReportURL(placeID)
	host := strings.Replace(url, "https://", "", 1)
	sign := AmazonAuth(auth).Sign(method, host, uri, str)
	m2["Signature"] = sign
	head:=make(map[string]string)
	head["HOST"]=host
	if headers!=nil{
		for k,v:=range headers{
			head[k]=v
		}
	}
	ans, err := http.Action(method, url+uri, formType, m2, head, ret)
	if err != nil {
		return
	}
	if ret == "" || ret == "xml" {
		err = http.TransXmlToStruct(ans, res)
	} else if ret == "text" {
		language := "GBK"
		if len(more) > 0 {
			language = more[0]
		}
		mp := http.ReadText(ans, language)
		err = http.TransMapToStruct(mp, res)
	}
	return
}
