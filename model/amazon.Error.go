package model

type ErrorResponse struct {
	Xmlns     string `xml:"xmlns,attr"`
	Error     Error  `xml:"Error"`
	RequestID string `xml:"RequestID"`
}

type Error struct {
	Type    string `xml:"Type"`
	Code    string `xml:"Code"`
	Message string `xml:"Message"`
}
