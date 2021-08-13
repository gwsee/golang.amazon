package model

//type RequestReportParam struct {
//	ReportType        string `json:"ReportType"`
//	StartDate         string `json:"StartDate"`
//	EndDate           string `json:"EndDate"`
//	ReportOptions     string `json:"ReportOptions"`
//	MarketplaceIdList string `json:"MarketplaceIdList"` // 请求参数不在日本和中国使用。
//}

type RequestReportResponse struct {
	Xmlns               string              `xml:"xmlns,attr"`
	RequestReportResult RequestReportResult `xml:"RequestReportResult"`
	ResponseMetadata    ResponseMetadata    `xml:"ResponseMetadata"`
}
type RequestReportResult struct {
	ReportRequestInfo ReportRequestInfo `xml:"ReportRequestInfo"`
}
type ReportRequestInfo struct {
	ReportRequestId        string `xml:"ReportRequestId"`
	ReportType             string `xml:"ReportType"`
	StartDate              string `xml:"StartDate"`
	EndDate                string `xml:"EndDate"`
	Scheduled              string `xml:"Scheduled"`
	SubmittedDate          string `xml:"SubmittedDate"`
	ReportProcessingStatus string `xml:"ReportProcessingStatus"`
	GeneratedReportId      string `xml:"GeneratedReportId"`
	StartedProcessingDate  string `xml:"StartedProcessingDate"`
	CompletedDate          string `xml:"CompletedDate"`
}
