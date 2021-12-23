package service

import (
	"fmt"

	"git.zx-tech.net/pengfeng/amazon.seller/http"
	"git.zx-tech.net/pengfeng/amazon.seller/model"
)

func init() {
	http.SaveActionLog = true
	headers = map[string]string{
		"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
		//"Host":                "mws.amazonservices.com",
		"x-amazon-user-agent": "ZhiXin/amazonErp (Language=Golang; Host=erp.test.mundossp.com)",
	}
	formType = "x-www-form-urlencoded"
}
func RequestReport() (obj model.RequestReportResponse, err error) {
	var a AmazonAuth
	res := a.Get("suyuchen-炮灰")
	mp := make(map[string]string)
	mp["ReportType"] = "_GET_MERCHANT_LISTINGS_DATA_BACK_COMPAT_" //必须字段
	mp["MarketplaceIdList.Id.1"] = res.MarketplaceID              //非必须
	mp["Merchant"] = res.SellerId
	err = CommonHandle("POST", "RequestReport", "/", res.MarketplaceID, res, mp, "xml", &obj)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(obj.RequestReportResult.ReportRequestInfo.ReportRequestId)
	}
	return
}

func GetReportList() (obj model.GetReportListResponse, err error) {
	var a AmazonAuth
	res := a.Get("suyuchen-炮灰")
	mp := make(map[string]string)
	mp["Merchant"] = res.SellerId
	mp["ReportTypeList.Type.1"] = "_GET_MERCHANT_LISTINGS_DATA_BACK_COMPAT_" //可以不用
	mp["MaxCount"] = "1"                                                     //可以不用
	err = CommonHandle("POST", "GetReportList", "/", res.MarketplaceID, res, mp, "xml", &obj)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(obj.GetReportListResult.ReportInfo[0].ReportId)
	}
	return
}

func GetReport() (list []model.GetMerchantListingsAllData, err error) {
	var a AmazonAuth
	res := a.Get("北美-david")
	mp := make(map[string]string)
	mp["Merchant"] = res.SellerId
	mp["ReportId"] = "33923649872018870"
	err = CommonHandle("POST", "GetReport", "/", res.MarketplaceID, res, mp, "text", &list)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(len(list))
		fmt.Println(list[0].ProductId)
		i := 0
		mp := make(map[string]int)
		mpA := make(map[string][]string)
		//STR1[:14]
		for _, v := range list {
			if v.ProductId != "" {
				i = i + 1
				//fmt.Println(i)
				//fmt.Println(i,"sku",v.SellerSku,"Asin1",v.Asin1,"Asin2",v.Asin2,"Asin3",v.Asin3)
				mp[v.SellerSku] = mp[v.SellerSku] + 1
				if len(v.SellerSku) > 14 {
					mpA[v.SellerSku[:14]] = append(mpA[v.SellerSku[:14]], v.SellerSku)
				} else {
					mpA[v.SellerSku] = []string{v.SellerSku}
				}

			}
		}
		//for k,v:=range mp{
		//	if v>1{
		//		fmt.Println(k,"......................",v)
		//	}else{
		//		fmt.Println(k)
		//	}
		//}
		total := 0
		for k, v := range mpA {
			total = total + 1
			fmt.Println(k, len(v))
			for kk, vv := range v {
				fmt.Println(kk, "------------>", vv)
			}
		}
		fmt.Println("总数为：", total)
	}
	return
}

func GetReport12() (list []model.GetMerchantListingsAllData, err error) {
	var a AmazonAuth
	res := a.Get("北美订单测试")
	mp := make(map[string]string)
	mp["Merchant"] = res.SellerId
	mp["ReportId"] = "33081521624018849"
	err = CommonHandle("POST", "GetReport", "/", res.MarketplaceID, res, mp, "text", &list)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(len(list))
		fmt.Println(list[0].ProductId)
		i := 0
		mp := make(map[string]int)
		mpA := make(map[string][]string)
		//STR1[:14]
		for _, v := range list {
			if v.ProductId != "" {
				i = i + 1
				//fmt.Println(i)
				//fmt.Println(i,"sku",v.SellerSku,"Asin1",v.Asin1,"Asin2",v.Asin2,"Asin3",v.Asin3)
				mp[v.SellerSku] = mp[v.SellerSku] + 1
				if len(v.SellerSku) > 14 {
					mpA[v.SellerSku[:14]] = append(mpA[v.SellerSku[:14]], v.SellerSku)
				} else {
					mpA[v.SellerSku] = []string{v.SellerSku}
				}

			}
		}
		//for k,v:=range mp{
		//	if v>1{
		//		fmt.Println(k,"......................",v)
		//	}else{
		//		fmt.Println(k)
		//	}
		//}
		total := 0
		for k, v := range mpA {
			total = total + 1
			fmt.Println(k, len(v))
			for kk, vv := range v {
				fmt.Println(kk, "------------>", vv)
			}
		}
		fmt.Println("总数为：", total)
	}
	return
}
func GetReportById(shop string, reportId string, data interface{}) (err error) {
	var a AmazonAuth
	res := a.Get(shop)
	mp := make(map[string]string)
	mp["Merchant"] = res.SellerId
	mp["ReportId"] = reportId
	err = CommonHandle("POST", "GetReport", "/", res.MarketplaceID, res, mp, "text", data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
