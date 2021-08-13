package model

type ListMatchingProducts struct {
	Xmlns            string           `xml:"xmlns,attr"`
	ListOrdersResult ListOrdersResult `xml:"ListOrdersResult"`
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}
