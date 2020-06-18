package wechat

type UnifyOrderReq struct {
	AppId         string `xml:"appid"`
	Body          string `xml:"body"`
	MchId         string `xml:"mch_id"`
	NonceStr      string `xml:"nonce_str"`
	NotifyUrl     string `xml:"notify_url"`
	TradeType     string `xml:"trade_type"`
	SbillCreateIp string `xml:"spbill_create_ip"`
	TotalFee      int    `xml:"total_fee"`
	OutTradeNo    string `xml:"out_trade_no"`
	Sign          string `xml:"sign"`
}

type UnifyOrderResp struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	AppId      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	PrepayId   string `xml:"prepay_id"`
	TradeType  string `xml:"trade_type"`
}
