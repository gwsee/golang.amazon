package model

//AuthObject 注册为开发者后获取
type AuthObject struct {
	AWSAccessKeyId   string `json:"aws_access_key_id,omitempty"`
	SecretKey        string `json:"secret_key,omitempty"`
	AuthSellerObject `json:"auth_seller_object"`
}

//AuthSellerObject 第三方授权后获取
type AuthSellerObject struct {
	SellerId      string `json:"seller_id,omitempty"`      //商家授权后获取的卖家ID
	MWSAuthToken  string `json:"mws_auth_token,omitempty"` //第三方授权后的mwsTOKEN
	MarketplaceID string `json:"marketplace_id,omitempty"` //不同区域对应的市场ID不一致
}
