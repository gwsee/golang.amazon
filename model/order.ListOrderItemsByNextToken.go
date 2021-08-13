package model

type ListOrderItemsByNextTokenResponse struct {
	Xmlns                           string               `xml:"xmlns,attr"`
	ListOrderItemsByNextTokenResult ListOrderItemsResult `xml:"ListOrderItemsByNextTokenResult"`
	ResponseMetadata                ResponseMetadata     `xml:"ResponseMetadata"`
}

//type ListOrderItemsByNextTokenResult struct {
//	NextToken     string     `xml:"NextToken"`
//	AmazonOrderId string     `xml:"AmazonOrderId"`
//	OrderItems    OrderItems `xml:"OrderItems"`
//}
