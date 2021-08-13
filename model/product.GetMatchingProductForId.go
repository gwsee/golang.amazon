package model

type GetMatchingProductForIdResponse struct {
	Xmlns                         string                          `xml:"xmlns,attr"`
	GetMatchingProductForIdResult []GetMatchingProductForIdResult `xml:"GetMatchingProductForIdResult"`
	ResponseMetadata              ResponseMetadata                `xml:"ResponseMetadata"`
}

type GetMatchingProductForIdResult struct {
	Asin     string   `xml:"ASIN,attr"`
	IdType   string   `xml:"IdType,attr"`
	Id       string   `xml:"Id,attr"`
	Status   string   `xml:"status,attr"`
	Products Products `xml:"Products"`
}
type Products struct {
	XmlnsNs2 string  `xml:"ns2,attr"`
	Product  Product `xml:"Product"`
}
