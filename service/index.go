package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"

	"git.zx-tech.net/pengfeng/amazon.seller/http"
	"git.zx-tech.net/pengfeng/amazon.seller/model"
)

type AmazonAuth model.AuthObject

var headers map[string]string
var formType string

func Init(l, a bool) {
	http.SaveActionLog = l
	http.SaveActionLogAns = a
}

/**
获取访问amazon接口的主要数据
Get
后期放入表中加密获取相应的数据
*/

//Get 获取amazon密钥等
func (a AmazonAuth) Get(ty string) (auth model.AuthObject) {
	//根据ty获取不同区域的 秘钥等
	return
}
func initParam(act string) (i model.IndexParams) {
	i.Action = act
	i.Version = "2009-01-01"
	i.SignatureMethod = "HmacSHA256"
	i.SignatureVersion = "2"
	i.Timestamp = time.Now().UTC().Format(time.RFC3339) // "2021-07-19T08:35:52Z"//
	return
}

//Header 亚马逊头
func (a AmazonAuth) Header() (m map[string]string) {
	return
}

//Params 对参数进行处理
func (a AmazonAuth) Params(ind model.IndexParams, mp map[string]string) map[string]string {
	m := make(map[string]string)
	m["AWSAccessKeyId"] = a.AWSAccessKeyId
	m["MWSAuthToken"] = a.MWSAuthToken
	//m["SellerId"] = a.SellerId
	m["Action"] = ind.Action
	m["SignatureMethod"] = ind.SignatureMethod
	m["SignatureVersion"] = ind.SignatureVersion
	m["Timestamp"] = ind.Timestamp
	m["Version"] = ind.Version
	if mp != nil {
		for k, v := range mp {
			m[k] = v
		}
	}
	return m
}

func (a AmazonAuth) Sort(mp map[string]string) (res string) {
	var sortSlice amazonSort
	for k, _ := range mp {
		sortSlice = append(sortSlice, k)
	}
	sort.Stable(sortSlice)
	var strArr []string
	for _, v := range sortSlice {
		val := mp[v]
		strArr = append(strArr, url.QueryEscape(v)+"="+url.QueryEscape(val))
	}
	res = strings.Join(strArr, "&")
	return
}
func (a AmazonAuth) Map(mpOld map[string]string) map[string]string {
	mpNew := make(map[string]string)
	for k, v := range mpOld {
		mpNew[url.QueryEscape(k)] = url.QueryEscape(v)
	}
	return mpNew
}

type amazonSort []string

func (a amazonSort) Len() int      { return len(a) }
func (a amazonSort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a amazonSort) Less(i, j int) bool {
	return a[i] < a[j]
}

//Sign 对参数进行签名
func (a AmazonAuth) Sign(method, host, uri, query string) (sign string) {
	sha256 := sha256.New
	hash := hmac.New(sha256, []byte(a.SecretKey))
	template := "%s\n%s\n%s\n%s"
	template = fmt.Sprintf(template, method, host, uri, query)
	hash.Write([]byte(template))
	sign = base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return
}
