package binance

import (
	"context"
	"encoding/json"

	"github.com/adshao/go-binance/v2/common"
)

/*
	sub account trans (Future)
*/

type BrokerTransFutu struct {
	c           *Client
	asset       string //코인명
	futuresType string //선물 타입 ( USDT , COIN)
	amount      string //수량
	fromID      string
	toID        string
}

type BrokerTransFutuREQ struct {
	Success      bool   `json:"success"`
	TxnId        int64  `json:"txnId"`
	ClientTranId string `json:"clientTranId"`
}

func (s *BrokerTransFutu) Asset(assetid string) *BrokerTransFutu {
	s.asset = assetid
	return s
}

func (s *BrokerTransFutu) FuturesType(futuresType string) *BrokerTransFutu {
	s.futuresType = futuresType
	return s
}

func (s *BrokerTransFutu) Amount(amount string) *BrokerTransFutu {
	s.amount = amount
	return s
}

func (s *BrokerTransFutu) FromID(fromID string) *BrokerTransFutu {
	s.fromID = fromID
	return s
}

func (s *BrokerTransFutu) ToID(toID string) *BrokerTransFutu {
	s.toID = toID
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerTransFutu) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, reqStatus int, err error) {

	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	if s.futuresType == "" {
		s.futuresType = "USDT"
	}
	m := params{
		"asset":       s.asset,
		"futuresType": s.futuresType,
		"amount":      s.amount,
		"fromId":      s.fromID,
		"toId":        s.toID,
	}

	r.setFormParams(m)
	data, reqStatus, err = s.c.callAPI_Core(ctx, r, opts...)
	if err != nil {
		return []byte{}, reqStatus, err
	}
	return data, reqStatus, nil
}

// Do sends the request.
func (s *BrokerTransFutu) Do_Svr(ctx context.Context, opts ...RequestOption) (res *BrokerTransFutuREQ, reqErrCode *common.APIError, reqStatus int, err error) {
	data, status, err := s.createReqData(ctx, "/sapi/v1/broker/transfer/futures", opts...)
	if err != nil {
		return nil, nil, status, err
	}

	if status >= 400 {
		apiErr := new(common.APIError)
		json.Unmarshal(data, apiErr)
		return nil, apiErr, status, err
	}

	//fmt.Println(string(data))
	res = new(BrokerTransFutuREQ)
	err = json.Unmarshal(data, res)

	if err != nil {
		return nil, nil, status, err
	}

	return res, nil, status, nil
}
