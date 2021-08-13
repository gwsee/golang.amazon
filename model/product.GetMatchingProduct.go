package model

type GetMatchingProductResponse struct {
	Xmlns                    string                     `xml:"xmlns,attr"`
	GetMatchingProductResult []GetMatchingProductResult `xml:"GetMatchingProductResult"`
	ResponseMetadata         ResponseMetadata           `xml:"ResponseMetadata"`
}

type GetMatchingProductResult struct {
	Asin      string  `xml:"ASIN,attr"`
	SellerSKU string  `xml:"SellerSKU,attr"`
	Status    string  `xml:"status,attr"`
	Product   Product `xml:"Product"`
}
type Product struct {
	XmlnsNs2      string        `xml:"ns2,attr"`
	Identifiers   Identifiers   `xml:"Identifiers"`
	AttributeSets AttributeSets `xml:"AttributeSets"`
	Relationships Relationships `xml:"Relationships"`
	SalesRankings SalesRankings `xml:"SalesRankings"`
	Offers        Offers        `xml:"Offers"`
}
type Identifiers struct {
	MarketplaceASIN MarketplaceASIN `xml:"MarketplaceASIN"`
}
type MarketplaceASIN struct {
	MarketplaceId string `xml:"MarketplaceId"`
	ASIN          string `xml:"ASIN"`
}

type AttributeSets struct {
	ItemAttributes ItemAttributes `xml:"ItemAttributes"`
}
type ItemAttributes struct {
	Lang                   string     `xml:"lang,attr"`
	Binding                string     `xml:"Binding"`
	Brand                  string     `xml:"Brand"`
	Label                  string     `xml:"Label"`
	Manufacturer           string     `xml:"Manufacturer"`
	ManufacturerMinimumAge Config     `xml:"ManufacturerMinimumAge"`
	PartNumber             string     `xml:"PartNumber"`
	ProductGroup           string     `xml:"ProductGroup"`
	ProductTypeName        string     `xml:"ProductTypeName"`
	Publisher              string     `xml:"Publisher"`
	SmallImage             SmallImage `xml:"SmallImage"`
	Studio                 string     `xml:"Studio"`
	Title                  string     `xml:"Title"`
}
type SmallImage struct {
	URL    string `xml:"URL"`
	Height Config `xml:"Height"`
	Width  Config `xml:"Width"`
}
type Config struct {
	Units string `xml:"Units,attr"`
	Text  string `xml:",innerxml"`
}

type Relationships struct {
	VariationParent Variation   `xml:"VariationParent"`
	VariationChild  []Variation `xml:"VariationChild"`
}
type Variation struct {
	Identifiers Identifiers `xml:"Identifiers"`
	Color       string      `xml:"Color"`
}
type SalesRankings struct {
	SalesRank []SalesRank `xml:"SalesRank"`
}
type SalesRank struct {
	ProductCategoryId string `xml:"ProductCategoryId"`
	Rank              string `xml:"Rank"`
}
type Offers struct {
	Offer Offer `xml:"Offer"`
}
type Offer struct {
	BuyingPrice        BuyingPrice `xml:"BuyingPrice"`
	RegularPrice       Money       `xml:"RegularPrice"`
	FulfillmentChannel string      `xml:"FulfillmentChannel"`
	ItemCondition      string      `xml:"ItemCondition"`
	ItemSubCondition   string      `xml:"ItemSubCondition"`
	SellerId           string      `xml:"SellerId"`
	SellerSKU          string      `xml:"SellerSKU"`
}
type BuyingPrice struct {
	LandedPrice  Money `xml:"LandedPrice"`
	ListingPrice Money `xml:"ListingPrice"`
	Shipping     Money `xml:"Shipping"`
}
