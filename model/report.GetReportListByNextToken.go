package model

type GetReportListByNextTokenResponse struct {
	Xmlns                          string              `xml:"xmlns,attr"`
	GetReportListByNextTokenResult GetReportListResult `xml:"GetReportListByNextTokenResult"`
	ResponseMetadata               ResponseMetadata    `xml:"ResponseMetadata"`
}

//type GetReportListByNextTokenResult struct {
//	NextToken  string       `xml:"NextToken"`
//	HasNext    string       `xml:"HasNext"`
//	ReportInfo []ReportInfo `xml:"ReportInfo"`
//}
