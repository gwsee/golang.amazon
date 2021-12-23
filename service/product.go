package service

import (
	"encoding/json"
	"fmt"

	"git.zx-tech.net/pengfeng/amazon.seller/model"
)

func GetMatchingProduct() (err error) {
	var a AmazonAuth
	res := a.Get("北美订单测试")
	mp := make(map[string]string)
	mp["ASINList.ASIN.1"] = "B098XKKLYR"    //必须
	mp["ASINList.ASIN.2"] = "B08HQSG9X3"    //必须
	mp["ASINList.ASIN.3"] = "B08HQNQVJJ"    //必须
	mp["MarketplaceId"] = res.MarketplaceID //必须
	mp["SellerId"] = res.SellerId
	mp["Version"] = "2011-10-01"
	var obj model.GetMatchingProductResponse
	err = CommonHandle("POST", "GetMatchingProduct", "/Products/2011-10-01", res.MarketplaceID, res, mp, "xml", &obj)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		st, _ := json.Marshal(obj)
		fmt.Println(string(st))
		fmt.Println(len(obj.GetMatchingProductResult))
	}
	return
}
func GetMatchingProductForId() (err error) {
	var a AmazonAuth
	res := a.Get("北美-david")
	mp := make(map[string]string)
	/**
	PD135001WZF00000-1
	PD135001WZF00000-FBA
	PD135001WZF00001-FBA
	PD135001WZF
	PD135001
	 */
	mp["IdList.Id.1"] = "PD135001WZF00000-1" //必须
	mp["IdList.Id.2"] = "PD135001WZF00000-FBA" //必须
	mp["IdList.Id.3"] = "PD135001WZF00001-FBA" //必须
	mp["IdList.Id.4"] = "PD135001WZF" //必须
	mp["IdList.Id.5"] = "PD135001" //必须
	mp["MarketplaceId"] = res.MarketplaceID //必须
	mp["SellerId"] = res.SellerId
	mp["Version"] = "2011-10-01"
	mp["IdType"] = "SellerSKU"
	var obj model.GetMatchingProductForIdResponse
	err = CommonHandle("POST", "GetMatchingProductForId", "/Products/2011-10-01", res.MarketplaceID, res, mp, "xml", &obj)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		st, _ := json.Marshal(obj)
		fmt.Println(string(st))
		fmt.Println(len(obj.GetMatchingProductForIdResult))
	}
	return
}
func GetMyPriceForASIN() (err error) {
	var a AmazonAuth
	res := a.Get("北美订单测试")
	mp := make(map[string]string)
	mp["ASINList.ASIN.1"] = "B098XKKLYR"    //必须
	mp["ASINList.ASIN.2"] = "B08HQSG9X3"    //必须
	mp["ASINList.ASIN.3"] = "B08HQNQVJJ"    //必须
	mp["MarketplaceId"] = res.MarketplaceID //必须
	mp["SellerId"] = res.SellerId
	mp["Version"] = "2011-10-01"
	var obj model.GetMyPriceForASINResponse
	err = CommonHandle("POST", "GetMyPriceForASIN", "/Products/2011-10-01", res.MarketplaceID, res, mp, "xml", &obj)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		st, _ := json.Marshal(obj)
		fmt.Println(string(st))
		fmt.Println(len(obj.GetMyPriceForASINResult))
	}
	return
}
