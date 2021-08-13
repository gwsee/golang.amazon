package model

type ListOrderItemsResponse struct {
	Xmlns                string               `xml:"xmlns,attr"`
	ListOrderItemsResult ListOrderItemsResult `xml:"ListOrderItemsResult"`
	ResponseMetadata     ResponseMetadata     `xml:"ResponseMetadata"`
}

type ListOrderItemsResult struct {
	NextToken     string     `xml:"NextToken"`
	AmazonOrderId string     `xml:"AmazonOrderId"`
	OrderItems    OrderItems `xml:"OrderItems"`
}
type OrderItems struct {
	OrderItem []OrderItem `xml:"OrderItem"`
}
type OrderItem struct {
	ASIN                       string       `xml:"ASIN"`                       //商品的亚马逊标准识别号 (ASIN)
	OrderItemId                string       `xml:"OrderItemId"`                //亚马逊定义的订单商品识别号。
	SellerSKU                  string       `xml:"SellerSKU"`                  //商品的卖家 SKU。
	Title                      string       `xml:"Title"`                      //商品名称。
	QuantityOrdered            string       `xml:"QuantityOrdered"`            //下面两个子元素的父元素 (两个子元素的类型为：xs:string)：
	QuantityShipped            string       `xml:"QuantityShipped"`            //已配送的商品数量。
	ItemPrice                  Money        `xml:"ItemPrice"`                  //订单商品的售价。
	ShippingPrice              Money        `xml:"ShippingPrice"`              //运费
	ScheduledDeliveryEndDate   string       `xml:"ScheduledDeliveryEndDate"`   //订单预约送货上门的开始日期（目的地时区）。日期格式为 ISO 8601。
	ScheduledDeliveryStartDate string       `xml:"ScheduledDeliveryStartDate"` //订单预约送货上门的终止日期（目的地时区）。日期格式为 ISO 8601。
	CODFee                     Money        `xml:"CODFee"`                     //COD 服务费用
	CODFeeDiscount             Money        `xml:"CODFeeDiscount"`             //货到付款费用的折扣。
	GiftMessageText            string       `xml:"GiftMessageText"`            //买家提供的礼品消息。
	GiftWrapPrice              Money        `xml:"GiftWrapPrice"`              //商品的礼品包装金额。
	GiftWrapLevel              string       `xml:"GiftWrapLevel"`              //买家指定的礼品包装等级。
	PromotionIds               PromotionIds `xml:"PromotionIds"`               //PromotionId 元素列表。
	ConditionId                string       `xml:"ConditionId"`                //商品的状况。
	ConditionSubtypeId         string       `xml:"ConditionSubtypeId"`         //商品的子状况。
	ConditionNote              string       `xml:"ConditionNote"`              //卖家描述的商品状况
	ItemTax                    Money        `xml:"ItemTax"`                    //商品价格的税费。
	ShippingTax                Money        `xml:"ShippingTax"`                //运费的税费。
	GiftWrapTax                Money        `xml:"GiftWrapTax"`                //礼品包装金额的税费。
	ShippingDiscount           Money        `xml:"ShippingDiscount"`           //运费的折扣
	PromotionDiscount          Money        `xml:"PromotionDiscount"`          //报价中的全部促销折扣总计。
	InvoiceData                InvoiceData  `xml:"InvoiceData"`                //发票信息（仅适用于中国）。
}

type PromotionIds struct {
	PromotionId string `xml:"PromotionId"`
}
type InvoiceData struct {
	InvoiceRequirement           string `xml:"InvoiceRequirement"`           //发票要求信息。
	BuyerSelectedInvoiceCategory string `xml:"BuyerSelectedInvoiceCategory"` //买家在下订单时选择的发票类目信息。
	InvoiceTitle                 string `xml:"InvoiceTitle"`                 //买家指定的发票抬头。
	InvoiceInformation           string `xml:"InvoiceInformation"`           //发票信息。
}
