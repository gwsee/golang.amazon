package model

type GetReportRequestListResponse struct {
	Xmlns                      string                     `xml:"xmlns,attr"`
	GetReportRequestListResult GetReportRequestListResult `xml:"GetReportRequestListResult"`
	ResponseMetadata           ResponseMetadata           `xml:"ResponseMetadata"`
}

type GetReportRequestListResult struct {
	NextToken         string              `xml:"NextToken"`
	HasNext           string              `xml:"HasNext"`
	ReportRequestInfo []ReportRequestInfo `xml:"ReportRequestInfo"`
}
