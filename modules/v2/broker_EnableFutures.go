package binance

import (
	"context"
	"encoding/json"
)

type BrokerEnableFutures struct {
	c            *Client
	subAccountId string //계좌번호
}

//BrokerEnableMarginREQ 마진 사용설정 리퀘스트
type BrokerEnableFuturesREQ struct {
	SubaccountId  string `json:"subaccountId"`
	EnableFutures bool   `json:"enableFutures"`
	UpdateTime    int64  `json:"updateTime"`
}

//계좌번호
func (s *BrokerEnableFutures) SubAccountID(accid string) *BrokerEnableFutures {
	s.subAccountId = accid
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerEnableFutures) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {

	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId": s.subAccountId,
		"futures":      true,
	}

	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do sends the request.
func (s *BrokerEnableFutures) Do(ctx context.Context, opts ...RequestOption) (res *BrokerEnableFuturesREQ, err error) {

	data, err := s.createReqData(ctx, "/sapi/v1/broker/subAccount/futures", opts...)
	if err != nil {
		return nil, err
	}
	res = new(BrokerEnableFuturesREQ)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
