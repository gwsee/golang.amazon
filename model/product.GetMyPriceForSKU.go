package model

type GetMyPriceForSKUResponse struct {
	Xmlns                  string                     `xml:"xmlns,attr"`
	GetMyPriceForSKUResult []GetMatchingProductResult `xml:"GetMyPriceForSKUResult"`
	ResponseMetadata       ResponseMetadata           `xml:"ResponseMetadata"`
}
