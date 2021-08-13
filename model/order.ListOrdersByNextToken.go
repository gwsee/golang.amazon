package model

type ListOrdersByNextTokenResponse struct {
	Xmlns                       string           `xml:"xmlns,attr"`
	ListOrdersByNextTokenResult ListOrdersResult `xml:"ListOrdersByNextTokenResult"`
	ResponseMetadata            ResponseMetadata `xml:"ResponseMetadata"`
}

//type ListOrdersByNextTokenResult struct {
//	NextToken         string   `xml:"NextToken"`
//	LastUpdatedBefore string   `xml:"LastUpdatedBefore"`
//	Orders            []Orders `xml:"Orders"`
//}
