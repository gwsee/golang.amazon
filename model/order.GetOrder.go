package model

type GetOrderResponse struct {
	Xmlns            string           `xml:"xmlns,attr"`
	GetOrderResult   GetOrderResult   `xml:"GetOrderResult"`
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

type GetOrderResult struct {
	NextToken         string `xml:"NextToken"`
	LastUpdatedBefore string `xml:"LastUpdatedBefore"`
	Orders            Orders `xml:"Orders"`
}
