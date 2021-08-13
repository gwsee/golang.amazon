package model

type GetServiceStatusResponse struct {
	Xmlns                  string                 `xml:"xmlns,attr"`
	GetServiceStatusResult GetServiceStatusResult `xml:"GetServiceStatusResult"`
	ResponseMetadata       ResponseMetadata       `xml:"ResponseMetadata"`
}

type GetServiceStatusResult struct {
	Status    string  `xml:"Status"`
	Timestamp string  `xml:"Timestamp"`
	MessageId string  `xml:"MessageId"`
	Messages  Message `xml:"Messages"`
}

type Message struct {
	Locale string `xml:"Locale"`
	Text   string `xml:"Text"`
}
