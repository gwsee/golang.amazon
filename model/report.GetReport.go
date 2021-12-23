package model

// AmazonEnvelope 是 _GET_XML_ALL_ORDERS_DATA_BY_LAST_UPDATE_ 返回的数据结构
type AmazonEnvelope struct {
	Xmlns string `xml:"-xmlns:xsi"`
	Xsi   string `xml:"-xsi:noNamespaceSchemaLocation"`
	// Header interface{} `xml:"Header"`
	MessageType string    `xml:"MessageType"`
	Message     []message `xml:"Message"`
}
type message struct {
	Order order `xml:"Order"`
}
type order struct {
	AmazonOrderID   string          `xml:"AmazonOrderID"`
	MerchantOrderID string          `xml:"MerchantOrderID"`
	PurchaseDate    string          `xml:"PurchaseDate"`
	LastUpdatedDate string          `xml:"LastUpdatedDate"`
	OrderStatus     string          `xml:"OrderStatus"`
	SalesChannel    string          `xml:"SalesChannel"`
	FulfillmentData fulfillmentData `xml:"FulfillmentData"`
	OrderItem       orderItem       `xml:"OrderItem"`
}
type fulfillmentData struct {
	FulfillmentChannel string  `xml:"FulfillmentChannel"`
	ShipServiceLevel   string  `xml:"ShipServiceLevel"`
	Address            Address `xml:"Address"`
}
type orderItem struct {
	AmazonOrderItemCode string    `xml:"AmazonOrderItemCode"`
	ASIN                string    `xml:"ASIN"`
	SKU                 string    `xml:"SKU"`
	ItemStatus          string    `xml:"ItemStatus"`
	ProductName         string    `xml:"ProductName"`
	Quantity            string    `xml:"Quantity"`
	ItemPrice           itemPrice `xml:"ItemPrice"`
}
type itemPrice struct {
	Component []component `xml:"Component"`
}
type component struct {
	Type   string `xml:"Type"`
	Amount amount `xml:"Amount"`
}
type amount struct {
	Currency string `xml:"-currency"`
	Text     string `xml:"#text"`
}

//-------------------美国
//商品报告

// GetFlatFileOpenListingsData _GET_FLAT_FILE_OPEN_LISTINGS_DATA_ 可售商品报告 （“库存报告”）
type GetFlatFileOpenListingsData struct {
}

// GetMerchantListingsAllData GET_MERCHANT_LISTINGS_ALL_DATA 所有商品报告
type GetMerchantListingsAllData struct {
	ItemName          string `json:"item-name"`
	ItemDescription   string `json:"item-description"`
	ListingId         string `json:"listing-id"`
	SellerSku         string `json:"seller-sku"`
	Price             string `json:"price"`
	Quantity          string `json:"quantity"`
	OpenDate          string `json:"open-date"`
	ImageUrl          string `json:"image-url"`
	ItemIsMarketplace string `json:"item-is-marketplace"`
	ProductIdType     string `json:"product-id-type"`

	ZshopShippingFee       string `json:"zshop-shipping-fee"`
	ItemNote               string `json:"item-note"`
	ItemCondition          string `json:"item-condition"`
	ZshopCateGory1         string `json:"zshop-category1"`
	ZshopBrowsePath        string `json:"zshop-browse-path"`
	ZshopStorefrontFeature string `json:"zshop-storefront-feature"`
	Asin1                  string `json:"asin1"`
	Asin2                  string `json:"asin2"`
	Asin3                  string `json:"asin3"`

	WillShipInternationally string `json:"will-ship-internationally"`
	ExpeditedShipping       string `json:"expedited-shipping"`
	ZshopBoldface           string `json:"zshop-boldface"`
	ProductId               string `json:"product-id"`
	BidForFeaturedPlacement string `json:"bid-for-featured-placement"`
	AddDelete               string `json:"add-delete"`
	PendingQuantity         string `json:"pending-quantity"`

	FulfillmentChannel    string `json:"fulfillment-channel"`
	MerchantShippingGroup string `json:"merchant-shipping-group"`
	Status                string `json:"status"`
}

// GetMerchantListingsDataBackCompat _GET_MERCHANT_LISTINGS_DATA_BACK_COMPAT_ 可售商品报告
type GetMerchantListingsDataBackCompat struct {
	ItemName          string `json:"item-name"`
	ItemDescription   string `json:"item-description"`
	ListingId         string `json:"listing-id"`
	SellerSku         string `json:"seller-sku"`
	Price             string `json:"price"`
	Quantity          string `json:"quantity"`
	OpenDate          string `json:"open-date"`
	ImageUrl          string `json:"image-url"`
	ItemIsMarketplace string `json:"item-is-marketplace"`
	ProductIdType     string `json:"product-id-type"`

	ZshopShippingFee       string `json:"zshop-shipping-fee"`
	ItemNote               string `json:"item-note"`
	ItemCondition          string `json:"item-condition"`
	ZshopCateGory1         string `json:"zshop-category1"`
	ZshopBrowsePath        string `json:"zshop-browse-path"`
	ZshopStorefrontFeature string `json:"zshop-storefront-feature"`
	Asin1                  string `json:"asin1"`
	Asin2                  string `json:"asin2"`
	Asin3                  string `json:"asin3"`

	WillShipInternationally string `json:"will-ship-internationally"`
	ExpeditedShipping       string `json:"expedited-shipping"`
	ZshopBoldface           string `json:"zshop-boldface"`
	ProductId               string `json:"product-id"`
	BidForFeaturedPlacement string `json:"bid-for-featured-placement"`
	AddDelete               string `json:"add-delete"`
	PendingQuantity         string `json:"pending-quantity"`

	FulfillmentChannel    string `json:"fulfillment-channel"`
	MerchantShippingGroup string `json:"merchant-shipping-group"`
}

// GetMerchantListingsData _GET_MERCHANT_LISTINGS_DATA_ 卖家商品报告 （“在售商品报告”） //32263349385018828
type GetMerchantListingsData struct {
}

//订单报告

// GetFlatFileActionableOrderData _GET_FLAT_FILE_ACTIONABLE_ORDER_DATA_ 待处理订单报告
type GetFlatFileActionableOrderData struct {
}

// _GET_ORDERS_DATA_  XML 文件格式的、计划的订单报告

type GetOrdersData struct {
}

// 32251962507018827

// GetFlatFileOrdersData _GET_FLAT_FILE_ORDERS_DATA_  请求或计划的、文本文件格式的订单报告 //制表符分隔的、文本文件格式的订单报告，可请求生成或计划生成。报告显示过去 60 天内的订单。针对商城和卖家平台卖家=
type GetFlatFileOrdersData struct {
	OrderId              string `json:"order-id"`
	OrderItemId          string `json:"order-item-id"`
	PurchaseDate         string `json:"purchase-date"`
	PaymentsDate         string `json:"payments-date"`
	BuyerEmail           string `json:"buyer-email"`
	BuyerName            string `json:"buyer-name"`
	BuyerPhoneNumber     string `json:"buyer-phone-number"`
	Sku                  string `json:"sku"`
	ProductName          string `json:"product-name"`
	QuantityPurchased    string `json:"quantity-purchased"`
	Currency             string `json:"currency"`
	ItemPrice            string `json:"item-price"`
	ItemTax              string `json:"item-tax"`
	ShippingPrice        string `json:"shipping-price"`
	ShippingTax          string `json:"shipping-tax"`
	ShipServiceLevel     string `json:"ship-service-level"`
	RecipientName        string `json:"recipient-name"`
	ShipAddress1         string `json:"ship-address-1"`
	ShipAddress2         string `json:"ship-address-2"`
	ShipAddress3         string `json:"ship-address-3"`
	ShipCity             string `json:"ship-city"`
	ShipState            string `json:"ship-state"`
	ShipPostalCode       string `json:"ship-postal-code"`
	ShipCountry          string `json:"ship-country"`
	ShipPhoneNumber      string `json:"ship-phone-number"`
	DeliveryStartDate    string `json:"delivery-start-date"`
	DeliveryEndDate      string `json:"delivery-end-date"`
	DeliveryTimeZone     string `json:"delivery-time-zone"`
	DeliveryInstructions string `json:"delivery-Instructions"`
}

// _GET_CONVERGED_FLAT_FILE_ORDER_REPORT_DATA_  文本文件格式的订单报告 //制表符分隔的、文本文件格式的订单报告，可请求生成或计划生成。仅针对商城卖家 //32247699927018827

type GetConvergedFlatFileOrderReportData struct {
	PaymentsStatus        string `json:"payments-status"`         //
	OrderId               string `json:"order-id"`                //
	OrderItemId           string `json:"order-item-id"`           //
	PaymentsDate          string `json:"payments-date"`           //
	PaymentsTransactionId string `json:"payments-transaction-id"` //
	ItemName              string `json:"item-name"`               //
	ListingId             string `json:"listing-id"`              //
	Sku                   string `json:"sku"`                     //
	Price                 string `json:"price"`                   //
	ShippingFee           string `json:"shipping-fee"`            //
	QuantityPurchased     string `json:"quantity-purchased"`      //
	TotalPrice            string `json:"total-price"`             //
	PurchaseDate          string `json:"purchase-date"`           //
	BatchId               string `json:"batch-id"`                //
	BuyerEmail            string `json:"buyer-email"`             //
	BuyerName             string `json:"buyer-name"`              //
	RecipientName         string `json:"recipient-name"`          //
	ShipAddress1          string `json:"ship-address-1"`          //
	ShipAddress2          string `json:"ship-address-2"`          //
	ShipCity              string `json:"ship-city"`               //
	ShipState             string `json:"ship-state"`              //
	ShipZip               string `json:"ship-zip"`                //
	ShipCountry           string `json:"ship-country"`            //
	SpecialComments       string `json:"special-comments"`        //
	Upc                   string `json:"upc"`                     //
	ShipMethod            string `json:"ship-method"`             //
}

//订单追踪报告

// GetFlatFileAllOrdersDataByLastUpdate _GET_FLAT_FILE_ALL_ORDERS_DATA_BY_LAST_UPDATE_  文本文件格式的、按最后更新日期排列的订单报告 --32236947858018827
type GetFlatFileAllOrdersDataByLastUpdate struct {
	AmazonOrderId         string `json:"amazon-order-id"`         //
	MerchantOrderId       string `json:"merchant-order-id"`       //
	PurchaseDate          string `json:"purchase-date"`           //
	LastUpdatedDate       string `json:"last-updated-date"`       //
	OrderStatus           string `json:"order-status"`            //
	FulfillmentChannel    string `json:"fulfillment-channel"`     //
	SalesChannel          string `json:"sales-channel"`           //
	OrderChannel          string `json:"order-channel"`           //
	Url                   string `json:"url"`                     //
	ShipServiceLevel      string `json:"ship-service-level"`      //
	ProductName           string `json:"product-name"`            //
	Sku                   string `json:"sku"`                     //
	Asin                  string `json:"asin"`                    //
	ItemStatus            string `json:"item-status"`             //
	Quantity              string `json:"quantity"`                //
	Currency              string `json:"currency"`                //
	ItemPrice             string `json:"item-price"`              //
	ItemTax               string `json:"item-tax"`                //
	ShippingPrice         string `json:"shipping-price"`          //
	ShippingTax           string `json:"shipping-tax"`            //
	GiftWrapPrice         string `json:"gift-wrap-price"`         //
	GiftWrapTax           string `json:"gift-wrap-tax"`           //
	ItemPromotionDiscount string `json:"item-promotion-discount"` //
	ShipPromotionDiscount string `json:"ship-promotion-discount"` //
	ShipCity              string `json:"ship-city"`               //
	ShipState             string `json:"ship-state"`              //
	ShipPostalCode        string `json:"ship-postal-code"`        //
	ShipCountry           string `json:"ship-country"`            //
	PromotionIds          string `json:"promotion-ids"`           //
}

// GetFlatFileAllOrdersDataByOrderDate _GET_FLAT_FILE_ALL_ORDERS_DATA_BY_ORDER_DATE_  文本文件格式的、按订单日期排列的订单报告
type GetFlatFileAllOrdersDataByOrderDate struct {
}

// GetXmlAllOrdersDataByLastUpdate _GET_XML_ALL_ORDERS_DATA_BY_LAST_UPDATE_ XML 格式的、按最后更新日期排列的订单报告 -- 32244707266018827
type GetXmlAllOrdersDataByLastUpdate struct {
}

// _GET_XML_ALL_ORDERS_DATA_BY_ORDER_DATE_  XML 格式的、按订单日期排列的订单报告

type GetXmlAllOrdersDataByOrderDate struct {
}

//等待中订单报告

// GetFlatFilePendingOrdersData _GET_FLAT_FILE_PENDING_ORDERS_DATA_  文本文件格式的等待中订单报告
type GetFlatFilePendingOrdersData struct {
}

// GetPendingOrdersData _GET_PENDING_ORDERS_DATA_  XML 格式的、等待中订单报告
type GetPendingOrdersData struct {
}

// GetConvergedFlatFilePendingOrdersData _GET_CONVERGED_FLAT_FILE_PENDING_ORDERS_DATA_  文本文件格式的综合等待中订单报告
type GetConvergedFlatFilePendingOrdersData struct {
}

//--------------日本

type GetMerchantListingsAllDataJp struct {
	ItemName          string `json:"商品名"`
	ItemDescription   string `json:""`
	ListingId         string `json:"出品ID"`
	SellerSku         string `json:"出品者SKU"`
	Price             string `json:"価格"`
	Quantity          string `json:"数量"`
	OpenDate          string `json:"出品日"`
	ImageUrl          string `json:""`
	ItemIsMarketplace string `json:""`
	ProductIdType     string `json:"商品IDタイプ"`

	ZshopShippingFee       string `json:""`
	ItemNote               string `json:"コンディション説明"`
	ItemCondition          string `json:"コンディション"`
	ZshopCateGory1         string `json:""`
	ZshopBrowsePath        string `json:""`
	ZshopStorefrontFeature string `json:""`
	Asin1                  string `json:""`
	Asin2                  string `json:""`
	Asin3                  string `json:""`

	WillShipInternationally string `json:"国外へ配送可"`
	ExpeditedShipping       string `json:"迅速な配送"`
	ZshopBoldface           string `json:""`
	ProductId               string `json:"商品ID"`
	BidForFeaturedPlacement string `json:""`
	AddDelete               string `json:""`
	PendingQuantity         string `json:"在庫数"`

	FulfillmentChannel    string `json:"フルフィルメント・チャンネル"`
	MerchantShippingGroup string `json:"merchant-shipping-group"`
	Status                string `json:"ステータス"`
}

// 亚马逊库存报告--亚马逊物流管理库存 - 存档  -- _GET_FBA_MYI_ALL_INVENTORY_DATA_（日文和英文一致，只是解析的时候需要用日文编码就行）
type GetFbaMyiAllInventoryData struct {
	Sku                         string `json:"sku"`
	Fnsku                       string `json:"fnsku"`
	Asin                        string `json:"asin"`
	ProductName                 string `json:"product-name"`
	Condition                   string `json:"condition"`
	YourPrice                   string `json:"your-price"`
	MfnListingExists            string `json:"mfn-listing-exists"`
	MfnFulfillableQuantity      string `json:"mfn-fulfillable-quantity"`
	AfnListingExists            string `json:"afn-listing-exists"`
	AfnWarehouseQuantity        string `json:"afn-warehouse-quantity"`
	AfnFulfillableQuantity      string `json:"afn-fulfillable-quantity"`
	AfnUnsellableQuantity       string `json:"afn-unsellable-quantity"`
	AfnReservedQuantity         string `json:"afn-reserved-quantity"`
	AfnTotalQuantity            string `json:"afn-total-quantity"`
	PerUnitVolume               string `json:"per-unit-volume"`
	AfnInboundWorkingQuantity   string `json:"afn-inbound-working-quantity"`
	AfnInboundShippedQuantity   string `json:"afn-inbound-shipped-quantity"`
	AfnInboundReceivingQuantity string `json:"afn-inbound-receiving-quantity"`
	AfnResearchingQuantity      string `json:"afn-researching-quantity"`
	AfnReservedFutureSupply     string `json:"afn-reserved-future-supply"`
	AfnFutureSupplyBuyable      string `json:"afn-future-supply-buyable"`
}

//亚马逊物流退货报告 _GET_FBA_FULFILLMENT_CUSTOMER_RETURNS_DATA_（日文和英文一致，只是解析的时候需要用日文编码就行）

type GetFbaFulfillmentCustomerReturnsData struct {
	ReturnedDate     string `json:"return-date"`
	OrderID          string `json:"order-id"`
	MerchantSKU      string `json:"sku"`
	ASIN             string `json:"asin"`
	FNSKU            string `json:"fnsku"`
	Title            string `json:"product-name"`
	Quantity         string `json:"quantity"`
	FC               string `json:"fulfillment-center-id"`
	Disposition      string `json:"detailed-disposition"`
	Reason           string `json:"reason"`
	Status           string `json:"status"`
	Lpn              string `json:"license-plate-number"`
	CustomerComments string `json:"customer-comments"`
}

type GetFbaFulfillmentCustomerReturnsDataJp struct {
	ReturnedDate     string `json:"Returned Date"`
	OrderID          string `json:"Order ID"`
	MerchantSKU      string `json:"Merchant SKU"`
	ASIN             string `json:"ASIN"`
	FNSKU            string `json:"FNSKU"`
	Title            string `json:"Title"`
	Quantity         string `json:"Quantity"`
	FC               string `json:"FC"`
	Disposition      string `json:"Disposition"`
	Reason           string `json:"Reason"`
	Lpn              string `json:"lpn"`
	CustomerComments string `json:"Customer comments"`
}

// _GET_FBA_FULFILLMENT_CURRENT_INVENTORY_DATA_ 亚马逊物流每日库存历史报告（日文和英文一致，只是解析的时候需要用日文编码就行）
type GetFbaFulfillmentCurrentInventoryData struct {
	SnapshotDate        string `json:"snapshot-date"`
	Fnsku               string `json:"fnsku"`
	Sku                 string `json:"sku"`
	ProductName         string `json:"product-name"`
	Quantity            string `json:"quantity"`
	FulfillmentCenterId string `json:"fulfillment-center-id"`
	DetailedDisposition string `json:"detailed-disposition"`
	Country             string `json:"country"`
}

// _GET_FBA_STORAGE_FEE_CHARGES_DATA_ 亚马逊库存费用报告（日文和英文一致，只是解析的时候需要用日文编码就行）--即月仓储费
type GetFbaStorageFeeChargesData struct {
	Asin                          string `json:"asin"`
	Fnsku                         string `json:"fnsku"`
	ProductName                   string `json:"product_name"`
	FulfillmentCenter             string `json:"fulfillment_center"`
	CountryCode                   string `json:"country_code"`
	LongestSide                   string `json:"longest_side"`
	MedianSide                    string `json:"median_side"`
	ShortestSide                  string `json:"shortest_side"`
	MeasurementUnits              string `json:"measurement_units"`
	Weight                        string `json:"weight"`
	WeightUnits                   string `json:"weight_units"`
	ItemVolume                    string `json:"item_volume"`
	VolumeUnits                   string `json:"volume_units"`
	ProductSizeTier               string `json:"product_size_tier"`
	AverageQuantityOnHand         string `json:"average_quantity_on_hand"`
	AverageQuantityPendingRemoval string `json:"average_quantity_pending_removal"`
	EstimatedTotalItemVolume      string `json:"estimated_total_item_volume"`
	MonthOfCharge                 string `json:"month_of_charge"`
	StorageRate                   string `json:"storage_rate"`
	Currency                      string `json:"currency"`
	EstimatedMonthlyStorageFee    string `json:"estimated_monthly_storage_fee"`
	DangerousGoodsStorageType     string `json:"dangerous_goods_storage_type"`
	EligibleForInventoryDiscount  string `json:"eligible_for_inventory_discount"`
	QualifiesForInventoryDiscount string `json:"qualifies_for_inventory_discount"`
	TotalIncentiveFeeAmount       string `json:"total_incentive_fee_amount"`
	BreakdownIncentiveFeeAmount   string `json:"breakdown_incentive_fee_amount"`
	AverageQuantityCustomerOrders string `json:"average_quantity_customer_orders"`
}

//_GET_FBA_FULFILLMENT_REMOVAL_ORDER_DETAIL_DATA_ 弃置费
type GetFbaFulfillmentRemovalOrderDetailData struct {
	RequestDate       string `json:"request-date"`
	OrderId           string `json:"order-id"`
	OrderType         string `json:"order-type"`
	OrderStatus       string `json:"order-status"`
	LastUpdatedDate   string `json:"last-updated-date"`
	Sku               string `json:"sku"`
	Fnsku             string `json:"fnsku"`
	Disposition       string `json:"disposition"`
	RequestedQuantity string `json:"requested-quantity"`
	CancelledQuantity string `json:"cancelled-quantity"`
	DisposedQuantity  string `json:"disposed-quantity"`
	ShippedQuantity   string `json:"shipped-quantity"`
	InProcessQuantity string `json:"in-process-quantity"`
	RemovalFee        string `json:"removal-fee"`
	Currency          string `json:"currency"`
}
