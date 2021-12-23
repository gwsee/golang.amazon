package service

import (
	"fmt"
	"time"

	"git.zx-tech.net/pengfeng/amazon.seller/model"
)

func ListOrders() (orders []model.Order, err error) {
	var a AmazonAuth
	res := a.Get("北美-JOHNWOOD")
	mp := make(map[string]string)
	mp["LastUpdatedAfter"] = time.Now().AddDate(0, 0, -20).UTC().Format(time.RFC3339) // "2021-06-30T16:00:00Z"//time.Now().AddDate(0,0,-20).UTC().Format(time.RFC3339) //必须字段
	mp["MarketplaceId.Id.1"] = res.MarketplaceID                                      //非必须
	mp["SellerId"] = res.SellerId
	mp["Version"] = "2013-09-01"
	var obj model.ListOrdersResponse
	err = CommonHandle("POST", "ListOrders", "/Orders/2013-09-01", res.MarketplaceID, res, mp, "xml", &obj)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("订单长度为", len(obj.ListOrdersResult.Orders.Order))
		orders = obj.ListOrdersResult.Orders.Order
	}
	return
}

func ListOrdersByNextToken() (orders []model.Order, err error) {
	var a AmazonAuth
	res := a.Get("北美-JOHNWOOD")
	mp := make(map[string]string)
	mp["SellerId"] = res.SellerId
	mp["Version"] = "2013-09-01"
	mp["NextToken"] = "qDBg6XgXnFF/YYuZQbKv1QafEPREmauvizt1MIhPYZZooPAPAiTP5ILWXerqCA7hggphpTfE/HJRBUDTTfkLYd8ytOVgN7d/KyNtf5fepe0z86MVYTeRn/84Iho8M6XprTqurxw/vfKS68vI7Bqws+weLxD7b1CV0+mtVKXnPdgGvqGqJURbBr8yRpqBNBEzllhc4XvGqg5e0zIeaOVNezxWEXvdeDL7Zg54wKetRXLniKjk2yjCAzr1Jptpxugh04A7P4TjeGcwE7+t4NjazyEZY027dXAVTSGshRBy6ZTthhfBGj6Dwh9DV8iyDXzG25news0zC1zKK3EWjgCoR9sHkzT6ertl3CH/fKi8/jvmoZh76tkAUNtc4fGQa1T1W5Hbv3f/4h9oGHvB6FqwTkrDiHqb8mBE0YT7h6RiajS56ChWWFgK767I8qrcn7mc"
	var obj2 model.ListOrdersByNextTokenResponse
	err = CommonHandle("POST", "ListOrdersByNextToken", "/Orders/2013-09-01", res.MarketplaceID, res, mp, "xml", &obj2, "AM", "D", "C")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("订单长度为", len(obj2.ListOrdersByNextTokenResult.Orders.Order))
		orders = obj2.ListOrdersByNextTokenResult.Orders.Order
	}
	nextToken := obj2.ListOrdersByNextTokenResult.NextToken
	for {
		time.Sleep(5 * time.Second)
		fmt.Println("根据nexttoken请求数据", len(orders))
		if nextToken == "" {
			break
		}
		var obj1 model.ListOrdersByNextTokenResponse
		mp["NextToken"] = nextToken
		err = CommonHandle("POST", "ListOrdersByNextToken", "/Orders/2013-09-01", res.MarketplaceID, res, mp, "xml", &obj1)
		if err != nil {
			break
		} else {
			nextToken = obj1.ListOrdersByNextTokenResult.NextToken
			time.Sleep(time.Minute)
			orders = append(orders, obj1.ListOrdersByNextTokenResult.Orders.Order...)
		}
	}
	fmt.Println(len(orders))
	return
}
func GetOrder() (obj model.GetOrderResponse, err error) {
	var a AmazonAuth
	res := a.Get("北美订单测试")
	mp := make(map[string]string)
	mp["AmazonOrderId.Id.1"] = "112-0296089-0616242" //必须
	mp["SellerId"] = res.SellerId
	mp["Version"] = "2013-09-01"
	err = CommonHandle("POST", "GetOrder", "/Orders/2013-09-01", res.MarketplaceID, res, mp, "xml", &obj)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(obj.GetOrderResult.Orders.Order[0].OrderStatus)
	}
	return
}
func ListOrderItems() (obj model.ListOrderItemsResponse, err error) {
	var a AmazonAuth
	res := a.Get("北美订单测试")
	mp := make(map[string]string)
	mp["AmazonOrderId"] = "112-0296089-0616242" //必须
	mp["SellerId"] = res.SellerId
	mp["Version"] = "2013-09-01"
	err = CommonHandle("POST", "ListOrderItems", "/Orders/2013-09-01", res.MarketplaceID, res, mp, "xml", &obj)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(obj.ListOrderItemsResult.OrderItems.OrderItem[0].SellerSKU)
	}
	return
}
