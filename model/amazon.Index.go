package model

type ResponseMetadata struct {
	RequestId string `xml:"RequestId"`
}

type IndexParams struct {
	Action           string `json:"Action"`           //您要对端点执行的操作，如 GetFeedSubmissionResult 操作。
	SignatureMethod  string `json:"SignatureMethod"`  //用于计算签名的 HMAC 哈希算法。HmacSHA256 和 HmacSHA1 都是支持的哈希算法，但亚马逊建议使用 HmacSHA256。
	SignatureVersion string `json:"SignatureVersion"` //当前使用的签名版本。这是亚马逊MWS 特定的信息，它告诉亚马逊MWS 您使用哪种算法来生成构成签名基础的字符串。对于亚马逊MWS，该值目前为 SignatureVersion=2。
	Timestamp        string `json:"Timestamp"`        //每个请求都必须包含请求的时戳。根据所用的 API 函数，您可以向请求提供一个过期日期和时间来代替时戳。 2008-06-26T18:12:21.078Z
	Version          string `json:"Version"`          //所调用的 API 部分的版本。
}
