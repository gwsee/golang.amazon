package model

type ListMarketplaceParticipationsByNextTokenResponse struct {
	Xmlns                                          string                              `xml:"xmlns,attr"`
	ListMarketplaceParticipationsByNextTokenResult ListMarketplaceParticipationsResult `xml:"ListMarketplaceParticipationsByNextTokenResult"`
	ResponseMetadata                               ResponseMetadata                    `xml:"ResponseMetadata"`
}
