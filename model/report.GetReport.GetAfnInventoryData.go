package model

type GetAfnInventoryData struct {
	SellerSku              string `json:"seller-sku"`
	FulfillmentChannelSku  string `json:"fulfillment-channel-sku"`
	Asin                   string `json:"asin"`
	ConditionType          string `json:"condition-type"`
	WarehouseConditionCode string `json:"Warehouse-Condition-code"`
	QuantityAvailable      string `json:"Quantity Available"`
}
