package http

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/id"
	"github.com/9299381/bingo/package/util"
	"github.com/9299381/bingo/package/wechat"
	"io/ioutil"
	"net/http"
)

func WeChatNotifyDecodeRequest(_ context.Context, req *http.Request) (interface{}, error) {

	bodyData, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	wxn := &WXPayNotify{}
	err = xml.Unmarshal(bodyData, wxn)
	if err != nil {
		return nil, err
	}
	data := util.Struct2Map(wxn)

	return &bingo.Request{
		Id:   id.New(),
		Data: data,
	}, nil
}

// HTTP返回数据编码函数
func WeChatNotifyEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-Requested-With,authToken")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Expose-Headers", "*")
	return json.NewEncoder(w).Encode(response)

}

type WXPayNotify struct {
	ReturnCode    string `xml:"return_code"`
	ReturnMsg     string `xml:"return_msg"`
	Appid         string `xml:"appid"`
	MchID         string `xml:"mch_id"`
	DeviceInfo    string `xml:"device_info"`
	NonceStr      string `xml:"nonce_str"`
	Sign          string `xml:"sign"`
	ResultCode    string `xml:"result_code"`
	ErrCode       string `xml:"err_code"`
	ErrCodeDes    string `xml:"err_code_des"`
	Openid        string `xml:"openid"`
	IsSubscribe   string `xml:"is_subscribe"`
	TradeType     string `xml:"trade_type"`
	BankType      string `xml:"bank_type"`
	TotalFee      int64  `xml:"total_fee"`
	FeeType       string `xml:"fee_type"`
	CashFee       int64  `xml:"cash_fee"`
	CashFeeType   string `xml:"cash_fee_type"`
	CouponFee     int64  `xml:"coupon_fee"`
	CouponCount   int64  `xml:"coupon_count"`
	CouponID0     string `xml:"coupon_id_0"`
	CouponFee0    int64  `xml:"coupon_fee_0"`
	TransactionID string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	Attach        string `xml:"attach"`
	TimeEnd       string `xml:"time_end"`
}

func wxpayVerifySign(needVerifyM map[string]interface{}, sign string) bool {
	signCalc := wechat.PaySign(needVerifyM, "API_KEY")

	if sign == signCalc {
		fmt.Println("签名校验通过!")
		return true
	}

	fmt.Println("签名校验失败!")
	return false
}
