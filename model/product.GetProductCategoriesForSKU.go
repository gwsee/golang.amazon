package model

type GetProductCategoriesForSKUResponse struct {
	Xmlns                            string                           `xml:"xmlns,attr"`
	GetProductCategoriesForSKUResult GetProductCategoriesForSKUResult `xml:"GetProductCategoriesForSKUResult"`
	ResponseMetadata                 ResponseMetadata                 `xml:"ResponseMetadata"`
}

type GetProductCategoriesForSKUResult struct {
	Self []ProductCategory `xml:"Self"`
}

type ProductCategory struct {
	ProductCategoryId   string            `xml:"ProductCategoryId"`
	ProductCategoryName string            `xml:"ProductCategoryName"`
	Parent              []ProductCategory `xml:"Parent"`
}
