package config

//const commodity = "" //商品
//const order = "" //订单
//const order_tracking = "" //订单追踪
//const order_waiting = "" //等待中订单
//const performance = "" //业绩
//const settlement = "" //结算
//const amazon_logistics = "" //亚马逊物流(FBA)
//const amazon_product_advertising = "" //亚马逊商品广告
//const sales_tax = "" //销售税费
//const classification_tree = "" //分类树

type AmazonSellerConfig struct{}

func (a AmazonSellerConfig) GetReportType(ty string) (reportTag string) {
	switch ty {

	case "commodity_saleable_goods_report(inventory_report)":
		//"可售商品报告 （“库存报告”）":
		reportTag = "_GET_FLAT_FILE_OPEN_LISTINGS_DATA_"
		//制表符分隔的、文本文件格式的可售商品报告，包含 SKU、ASIN、价格和数量输入项。针对商城和卖家平台卖家。

	case "commodity_saleable_goods_report":
		//可售商品报告
		reportTag = "_GET_MERCHANT_LISTINGS_DATA_BACK_COMPAT_"
		//制表符分隔的、文本文件格式的可售商品报告。

	case "commodity_seller_commodity_report(available_commodity_report)":
		//卖家商品报告 （“在售商品报告”）
		reportTag = "_GET_MERCHANT_LISTINGS_DATA_"
		//制表符分隔的、文本文件格式的详细在售商品报告。针对商城和卖家平台卖家。

	case "amazon_logistics_fba_managed_inventory":
		//亚马逊物流管理库存
		reportTag = "_GET_FBA_MYI_UNSUPPRESSED_INVENTORY_DATA_"
		//制表符分隔的文本文件。该报告中包含在售（非存档）库存的当前详情，如状况、数量和销售量等。报告内容接近实时更新。仅针对亚马逊物流卖家。针对商城和卖家平台卖家。
	}
	return
}

func (a AmazonSellerConfig) GetReportURL(id string) (url string) {
	switch id {

	// 北美(North America)地区
	case "A2Q3Y263D00KWC": //巴西
		url = "https://mws.amazonservices.com"
	case "A2EUQ1WTGCTBG2": //加拿大
		url = "https://mws.amazonservices.ca"
	case "A1AM78C64UM0Y8": //墨西哥
		url = "https://mws.amazonservices.com.mx"
	case "ATVPDKIKX0DER": //美國
		url = "https://mws.amazonservices.com"

	// 欧洲(Europe)地区
	case "A2VIGQ35RCS4UG": //阿拉伯联合酋长国（U.A.E.）
		url = "https://mws.amazonservices.ae"
	case "A1PA6795UKMFR9": //德國
		url = "https://mws-eu.amazonservices.com"
	case "ARBP9OOSHTCHU": //埃及
		url = "https://mws-eu.amazonservices.com"
	case "A1RKKUPIHCS9HS": //西班牙
		url = "https://mws-eu.amazonservices.com"
	case "A13V1IB3VIYZZH": //法國
		url = "https://mws-eu.amazonservices.com"
	case "A1F83G8C2ARO7P": //英國
		url = "https://mws-eu.amazonservices.com"
	case "A21TJRUUN4KGV": //印度
		url = "https://mws.amazonservices.in"
	case "APJ6JRA9NG5V4": //意大利
		url = "https://mws-eu.amazonservices.com"
	case "A1805IZSGTT6HS": //荷兰
		url = "https://mws-eu.amazonservices.com"
	case "A1C3SOZRARQ6R3": //波兰
		url = "https://mws-eu.amazonservices.com"
	case "A17E79C6D8DWNP": //沙特阿拉伯
		url = "https://mws-eu.amazonservices.com"
	case "A2NODRKZP88ZB9": //瑞典
		url = "https://mws-eu.amazonservices.com"
	case "A33AVAJ2PDY3EV": //土耳其
		url = "https://mws-eu.amazonservices.com"

		//远东(Far East)地区
	case "A19VAU5U5O7RUS": //新加坡
		url = "https://mws-fe.amazonservices.com"
	case "A39IBJ37TRP1C6": //澳大利亞
		url = "https://mws.amazonservices.com.au"
	case "A1VC38T7YXB528": //日本
		url = "https://mws.amazonservices.jp"

	}
	return
}
