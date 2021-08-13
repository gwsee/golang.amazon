package model

type GetReportListResponse struct {
	Xmlns               string              `xml:"xmlns,attr"`
	GetReportListResult GetReportListResult `xml:"GetReportListResult"`
	ResponseMetadata    ResponseMetadata    `xml:"ResponseMetadata"`
}

type ReportInfo struct {
	ReportType      string `xml:"ReportType"`
	Acknowledged    string `xml:"Acknowledged"`
	ReportId        string `xml:"ReportId"`
	ReportRequestId string `xml:"ReportRequestId"`
	AvailableDate   string `xml:"AvailableDate"`
}

type GetReportListResult struct {
	NextToken  string       `xml:"NextToken"`
	HasNext    string       `xml:"HasNext"`
	ReportInfo []ReportInfo `xml:"ReportInfo"`
}
