package model

type ListFinancialEventsByNextTokenResponse struct {
	Xmlns                                string                    `xml:"xmlns,attr"`
	ListFinancialEventsByNextTokenResult ListFinancialEventsResult `xml:"ListFinancialEventsByNextTokenResult"`
	ResponseMetadata                     ResponseMetadata          `xml:"ResponseMetadata"`
}
