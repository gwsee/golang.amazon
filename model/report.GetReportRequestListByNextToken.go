package model

type GetReportRequestListByNextToken struct {
	Xmlns                                 string                     `xml:"xmlns,attr"`
	ResponseMetadata                      ResponseMetadata           `xml:"ResponseMetadata"`
	GetReportRequestListByNextTokenResult GetReportRequestListResult `xml:"GetReportRequestListByNextTokenResult"`
}

//type GetReportRequestListByNextTokenResult struct {
//	NextToken         string              `xml:"NextToken"`
//	HasNext           string              `xml:"HasNext"`
//	ReportRequestInfo []ReportRequestInfo `xml:"ReportRequestInfo"`
//}
