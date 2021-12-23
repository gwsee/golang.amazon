package model

type ListMarketplaceParticipationsResponse struct {
	Xmlns                               string                              `xml:"xmlns,attr"`
	ListMarketplaceParticipationsResult ListMarketplaceParticipationsResult `xml:"ListMarketplaceParticipationsResult"`
	ResponseMetadata                    ResponseMetadata                    `xml:"ResponseMetadata"`
}

type ListMarketplaceParticipationsResult struct {
	NextToken          string             `xml:"NextToken"`
	HasNext            string             `xml:"HasNext"`
	ListParticipations ListParticipations `xml:"ListParticipations"`
	ListMarketplaces   ListMarketplaces   `xml:"ListMarketplaces"`
}

type ListParticipations struct {
	Participation []Participation `xml:"Participation"`
}
type ListMarketplaces struct {
	Marketplace []Marketplace `xml:"Marketplace"`
}
type Participation struct {
	MarketplaceId              string `xml:"MarketplaceId"`
	SellerId                   string `xml:"SellerId"`
	HasSellerSuspendedListings string `xml:"HasSellerSuspendedListings"`
}
type Marketplace struct {
	MarketplaceId       string `xml:"MarketplaceId"`
	DefaultCountryCode  string `xml:"DefaultCountryCode"`
	DomainName          string `xml:"DomainName"`
	Name                string `xml:"Name"`
	DefaultCurrencyCode string `xml:"DefaultCurrencyCode"`
	DefaultLanguageCode string `xml:"DefaultLanguageCode"`
}
