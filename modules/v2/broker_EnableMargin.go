package binance

import (
	"context"
	"encoding/json"
)

type BrokerEnableMargin struct {
	c            *Client
	subAccountId string //계좌번호
}

//BrokerEnableMarginREQ 마진 사용설정 리퀘스트
type BrokerEnableMarginREQ struct {
	SubaccountId string `json:"subaccountId"`
	EnableMargin bool   `json:"enableMargin"`
	UpdateTime   int64  `json:"updateTime"`
}

//계좌번호
func (s *BrokerEnableMargin) SubAccountID(accid string) *BrokerEnableMargin {
	s.subAccountId = accid
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerEnableMargin) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {

	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId": s.subAccountId,
		"margin":       true,
	}

	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do sends the request.
func (s *BrokerEnableMargin) Do(ctx context.Context, opts ...RequestOption) (res *BrokerEnableMarginREQ, err error) {

	data, err := s.createReqData(ctx, "/sapi/v1/broker/subAccount/margin", opts...)
	if err != nil {
		return nil, err
	}
	res = new(BrokerEnableMarginREQ)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
