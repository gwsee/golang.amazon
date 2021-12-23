package model

type ListFinancialEventsResponse struct {
	Xmlns                     string                    `xml:"xmlns,attr"`
	ListFinancialEventsResult ListFinancialEventsResult `xml:"ListFinancialEventsResult"`
	ResponseMetadata          ResponseMetadata          `xml:"ResponseMetadata"`
}

type ListFinancialEventsResult struct {
	NextToken       string          `xml:"NextToken"`
	FinancialEvents FinancialEvents `xml:"FinancialEvents"`
}
type FinancialEvents struct {
	ProductAdsPaymentEventList []ProductAdsPaymentEvents `xml:"ProductAdsPaymentEventList"`
	ServiceFeeEventList        []ServiceFeeEvents        `xml:"ServiceFeeEventList"`
	RefundEventList            []ShipmentEvents          `xml:"RefundEventList"`
	ShipmentEventList          []ShipmentEvents          `xml:"ShipmentEventList"`
	CouponPaymentEventList     []CouponPaymentEvents     `xml:"CouponPaymentEventList"`
	AdjustmentEventList        []AdjustmentEvents        `xml:"AdjustmentEventList"`
}
type ProductAdsPaymentEvents struct {
	ProductAdsPaymentEvent []ProductAdsPaymentEvent `xml:"ProductAdsPaymentEvent"`
}
type ServiceFeeEvents struct {
	ServiceFeeEvent []ServiceFeeEvent `xml:"ServiceFeeEvent"`
}
type ServiceFeeEvent struct {
	AmazonOrderId string        `xml:"AmazonOrderId"`
	FeeDescription string        `xml:"FeeDescription"`
	SellerSKU string        `xml:"SellerSKU"`
	FeeList       FeeComponents `xml:"FeeList"`
}
type ProductAdsPaymentEvent struct {
	TransactionType string `xml:"transactionType"`
	InvoiceId       string `xml:"invoiceId"`
	PostedDate      string `xml:"postedDate"`

	BaseValue        Money `xml:"baseValue"`
	TaxValue         Money `xml:"taxValue"`
	TransactionValue Money `xml:"transactionValue"`
}
type ShipmentEvents struct {
	ShipmentEvent []ShipmentEvent `xml:"ShipmentEvent"`
}
type CouponPaymentEvents struct {
	CouponPaymentEvent []CouponPaymentEvent `xml:"CouponPaymentEvent"`
}
type AdjustmentEvents struct {
	AdjustmentEvent []AdjustmentEvent `xml:"AdjustmentEvent"`
}

type AdjustmentEvent struct {
	AdjustmentType string `xml:"AdjustmentType"`
	PostedDate     string `xml:"PostedDate"`

	AdjustmentItemList AdjustmentItems `xml:"AdjustmentItemList"`
	AdjustmentAmount   Money           `xml:"AdjustmentAmount"`
}
type AdjustmentItems struct {
	AdjustmentItem []AdjustmentItem `xml:"AdjustmentItem"`
}
type AdjustmentItem struct {
	Quantity           string `xml:"Quantity"`
	SellerSKU          string `xml:"SellerSKU"`
	ProductDescription string `xml:"ProductDescription"`

	PerUnitAmount Money `xml:"PerUnitAmount"`
	TotalAmount   Money `xml:"TotalAmount"`
}
type CouponPaymentEvent struct {
	PaymentEventId          string `xml:"PaymentEventId"`
	SellerCouponDescription string `xml:"SellerCouponDescription"`
	CouponId                string `xml:"CouponId"`
	ClipOrRedemptionCount   string `xml:"ClipOrRedemptionCount"`
	PostedDate              string `xml:"PostedDate"`

	TotalAmount     Money           `xml:"TotalAmount"`
	FeeComponent    FeeComponent    `xml:"FeeComponent"`
	ChargeComponent ChargeComponent `xml:"ChargeComponent"`
}

type ShipmentEvent struct {
	PostedDate      string `xml:"PostedDate"`
	AmazonOrderId   string `xml:"AmazonOrderId"`
	SellerOrderId   string `xml:"SellerOrderId"`
	MarketplaceName string `xml:"MarketplaceName"`

	ShipmentItemAdjustmentList ShipmentItemAdjustmentList `xml:"ShipmentItemAdjustmentList"`
	ShipmentItemList           ShipmentItemList           `xml:"ShipmentItemList"`
}
type ShipmentItemAdjustmentList struct {
	ShipmentItem []ShipmentItem `xml:"ShipmentItem"`
}
type ShipmentItemList struct {
	ShipmentItem []ShipmentItem `xml:"ShipmentItem"`
}

type ShipmentItem struct {
	SellerSKU                string                `xml:"SellerSKU"`
	OrderItemId              string                `xml:"OrderItemId"`
	QuantityShipped          string                `xml:"QuantityShipped"`
	OrderAdjustmentItemId    string                `xml:"OrderAdjustmentItemId"`
	ItemFeeList              FeeComponents         `xml:"ItemFeeList"`
	ItemFeeAdjustmentList    FeeComponents         `xml:"ItemFeeAdjustmentList"`
	ItemChargeList           ChargeComponents      `xml:"ItemChargeList"`
	ItemChargeAdjustmentList ChargeComponents      `xml:"ItemChargeAdjustmentList"`
	PromotionList            Promotions            `xml:"PromotionList"`
	PromotionAdjustmentList  Promotions            `xml:"PromotionAdjustmentList"`
	ItemTaxWithheldList      TaxWithheldComponents `xml:"ItemTaxWithheldList"`
}
type Promotions struct {
	Promotion []Promotion `xml:"Promotion"`
}
type Promotion struct {
	PromotionType   string `xml:"PromotionType"`
	PromotionId     string `xml:"PromotionId"`
	PromotionAmount Money  `xml:"PromotionAmount"`
}
type FeeComponents struct {
	FeeComponent []FeeComponent `xml:"FeeComponent"`
}
type ChargeComponents struct {
	ChargeComponent []ChargeComponent `xml:"ChargeComponent"`
}

type TaxWithheldComponents struct {
	TaxWithheldComponent []TaxWithheldComponent `xml:"TaxWithheldComponent"`
}

type ChargeComponent struct {
	ChargeType   string `xml:"ChargeType"`
	ChargeAmount Money  `xml:"ChargeAmount"`
}
type FeeComponent struct {
	FeeType   string `xml:"FeeType"`
	FeeAmount Money  `xml:"FeeAmount"`
}
type TaxWithheldComponent struct {
	TaxCollectionModel string           `xml:"TaxCollectionModel"`
	TaxesWithheld      ChargeComponents `xml:"TaxesWithheld"`
}
