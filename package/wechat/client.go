package wechat

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func UnifiedOrder(uniReq *UnifyOrderReq) error {
	m := make(map[string]interface{})
	m["appid"] = uniReq.AppId
	m["body"] = uniReq.Body
	m["mch_id"] = uniReq.MchId
	m["notify_url"] = uniReq.NotifyUrl
	m["trade_type"] = uniReq.TradeType
	m["spbill_create_ip"] = uniReq.SbillCreateIp
	m["total_fee"] = uniReq.TotalFee
	m["out_trade_no"] = uniReq.OutTradeNo
	m["nonce_str"] = uniReq.NonceStr
	uniReq.Sign = PaySign(m, "wxpay_api_key")
	bytes_req, err := xml.Marshal(uniReq)
	if err != nil {
		return fmt.Errorf("以xml形式编码发送错误, 原因:", err)
	}
	str_req := string(bytes_req)
	//wxpay的unifiedorder接口需要http body中xmldoc的根节点是<xml></xml>这种，所以这里需要replace一下
	str_req = strings.Replace(str_req, "UnifyOrderReq", "xml", -1)
	bytes_req = []byte(str_req)

	//发送unified order请求.
	request, err := http.NewRequest("POST", "url", bytes.NewReader(bytes_req))
	if err != nil {
		return fmt.Errorf("New Http Request发生错误，原因:", err)
	}
	request.Header.Set("Accept", "application/xml")
	//这里的http header的设置是必须设置的.
	request.Header.Set("Content-Type", "application/xml;charset=utf-8")

	c := http.Client{}
	resp, _err := c.Do(request)
	if _err != nil {
		return fmt.Errorf("请求微信支付统一下单接口发送错误, 原因:", _err)
	}
	bodyData, err := ioutil.ReadAll(resp.Body)
	xmlResp := &UnifyOrderResp{}
	_err = xml.Unmarshal(bodyData, xmlResp)

	//处理return code.
	if xmlResp.ReturnCode == "FAIL" {
		return fmt.Errorf("微信支付统一下单不成功，原因:", xmlResp.ReturnMsg)
	}
	return nil
}
