package model

type GetMyPriceForASINResponse struct {
	Xmlns                   string                     `xml:"xmlns,attr"`
	GetMyPriceForASINResult []GetMatchingProductResult `xml:"GetMyPriceForASINResult"`
	ResponseMetadata        ResponseMetadata           `xml:"ResponseMetadata"`
}
