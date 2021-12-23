package model

type ListOrdersResponse struct {
	Xmlns            string           `xml:"xmlns,attr"`
	ListOrdersResult ListOrdersResult `xml:"ListOrdersResult"`
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

type ListOrdersResult struct {
	NextToken         string `xml:"NextToken"`
	LastUpdatedBefore string `xml:"LastUpdatedBefore"`
	Orders            Orders `xml:"Orders"`
}
type Orders struct {
	Order []Order `xml:"Order"`
}
type Order struct {
	ShipmentServiceLevelCategory   string                 `xml:"ShipmentServiceLevelCategory" json:"shipment_service_level_category,omitempty"` //订单的配送服务级别分类
	ShipServiceLevel               string                 `xml:"ShipServiceLevel" json:"ship_service_level,omitempty"`                          //货件服务水平。
	EarliestShipDate               string                 `xml:"EarliestShipDate" json:"earliest_ship_date,omitempty"`                          //您承诺的订单发货时间范围的第一天
	LatestShipDate                 string                 `xml:"LatestShipDate" json:"latest_ship_date,omitempty"`                              //您承诺的订单发货时间范围的最后一天
	MarketplaceId                  string                 `xml:"MarketplaceId" json:"marketplace_id,omitempty"`                                 //订单生成所在商城的匿名编码。
	SalesChannel                   string                 `xml:"SalesChannel" json:"sales_channel,omitempty"`                                   //订单中第一件商品的销售渠道。
	OrderChannel                   string                 `xml:"OrderChannel" json:"order_channel,omitempty"`                                   //订单中第一件商品的订单渠道。
	OrderType                      string                 `xml:"OrderType" json:"order_type,omitempty"`                                         //订单类型。
	BuyerEmail                     string                 `xml:"BuyerEmail" json:"buyer_email,omitempty"`                                       //买家的匿名电子邮件地址。
	FulfillmentChannel             string                 `xml:"FulfillmentChannel" json:"fulfillment_channel,omitempty"`                       //订单配送方式：亚马逊配送 (AFN) 或卖家自行配送 (MFN)。
	OrderStatus                    string                 `xml:"OrderStatus" json:"order_status,omitempty"`                                     //当前的订单状态。
	BuyerName                      string                 `xml:"BuyerName" json:"buyer_name,omitempty"`                                         //买家姓名。
	LastUpdateDate                 string                 `xml:"LastUpdateDate" json:"last_update_date,omitempty"`                              //订单的最后更新日期。
	PurchaseDate                   string                 `xml:"PurchaseDate" json:"purchase_date,omitempty"`                                   //创建订单的日期。
	NumberOfItemsShipped           string                 `xml:"NumberOfItemsShipped" json:"number_of_items_shipped,omitempty"`                 //已配送的商品数量
	NumberOfItemsUnshipped         string                 `xml:"NumberOfItemsUnshipped" json:"number_of_items_unshipped,omitempty"`             //未配送的商品数量。
	AmazonOrderId                  string                 `xml:"AmazonOrderId" json:"amazon_order_id,omitempty"`                                //亚马逊所定义的订单编码，格式为 3-7-7。
	SellerOrderId                  string                 `xml:"SellerOrderId" json:"seller_order_id,omitempty"`                                //卖家所定义的订单编码。
	PaymentMethod                  string                 `xml:"PaymentMethod" json:"payment_method,omitempty"`                                 //订单的主要付款方式。
	ShippedByAmazonTFM             bool                   `xml:"ShippedByAmazonTFM" json:"shipped_by_amazon_tfm,omitempty"`                     //指明订单配送方是否是亚马逊配送 (Amazon TFM) 服务。
	TFMShipmentStatus              string                 `xml:"TFMShipmentStatus" json:"tfm_shipment_status,omitempty"`                        //亚马逊 TFM订单的状态。仅当ShippedByAmazonTFM = True时返回。请注意：即使当 ShippedByAmazonTFM = True 时，如果您还没有创建货件，也不会返回
	CbaDisplayableShippingLabel    string                 `xml:"CbaDisplayableShippingLabel" json:"cba_displayable_shipping_label,omitempty"`   //卖家自定义的配送方式，属于Checkout by Amazon (CBA) 所支持的四种标准配送设置中的一种。有关更多信息，请参阅您所在商城 Amazon Payments 帮助中心的“设置灵活配送方式”主题。
	IsBusinessOrder                bool                   `xml:"IsBusinessOrder" json:"is_business_order,omitempty"`
	IsGlobalExpressEnabled         bool                   `xml:"IsGlobalExpressEnabled" json:"is_global_express_enabled,omitempty"`
	IsPrime                        bool                   `xml:"IsPrime" json:"is_prime,omitempty"`
	IsReplacementOrder             bool                   `xml:"IsReplacementOrder" json:"is_replacement_order,omitempty"`
	ReplacedOrderId                string                 `xml:"ReplacedOrderId" json:"replaced_order_id,omitempty"`
	IsSoldByAB                     bool                   `xml:"IsSoldByAB" json:"is_sold_by_ab,omitempty"`
	EarliestDeliveryDate           string                 `xml:"EarliestDeliveryDate" json:"earliest_delivery_date,omitempty"` //您承诺的订单送达时间范围的第一天
	IsPremiumOrder                 bool                   `xml:"IsPremiumOrder" json:"is_premium_order,omitempty"`
	IsISPU                         bool                   `xml:"IsISPU" json:"is_ispu,omitempty"`
	LatestDeliveryDate             string                 `xml:"LatestDeliveryDate" json:"latest_delivery_date,omitempty"` //您承诺的订单送达时间范围的最后一天
	PaymentMethodDetails           PaymentMethodDetail    `xml:"PaymentMethodDetails" json:"payment_method_details"`
	OrderTotal                     Money                  `xml:"OrderTotal" json:"order_total"`           //订单的总费用。
	ShippingAddress                Address                `xml:"ShippingAddress" json:"shipping_address"` //订单的配送地址。
	PaymentExecutionDetail         PaymentExecutionDetail `xml:"PaymentExecutionDetail" json:"payment_execution_detail"`
	DefaultShipFromLocationAddress Address                `xml:"DefaultShipFromLocationAddress" json:"default_ship_from_location_address"`
}

type Money struct {
	CurrencyCode string `xml:"CurrencyCode" json:"currency_code"` //,omitempty为什么会有这个 忘记了 先去掉
	Amount       string `xml:"Amount" json:"amount"`
	CurrencyAmount   string `xml:"CurrencyAmount" json:"currency_amount"`
}
type Address struct {
	Name                         string `xml:"Name"`
	AddressLine1                 string `xml:"AddressLine1"`
	AddressLine2                 string `xml:"AddressLine2"`
	AddressLine3                 string `xml:"AddressLine3"`
	City                         string `xml:"City"`
	AddressType                  string `xml:"AddressType"`
	IsAddressSharingConfidential bool   `xml:"isAddressSharingConfidential"`
	PostalCode                   string `xml:"PostalCode"`
	Country                      string `xml:"Country"`
	StateOrRegion                string `xml:"StateOrRegion"`
	Phone                        string `xml:"Phone"`
	State                        string `xml:"State"`
	CountryCode                  string `xml:"CountryCode"`
}
type PaymentExecutionDetail struct {
	PaymentExecutionDetailItem []PaymentExecutionDetailItem `xml:"PaymentExecutionDetailItem"`
}
type PaymentExecutionDetailItem struct {
	Payment       Money  `xml:"Payment"`
	PaymentMethod string `xml:"PaymentMethod"` //COD 订单的次级付款方式。
}

type PaymentMethodDetail struct {
	PaymentMethodDetail string `xml:"PaymentMethodDetail"`
}
