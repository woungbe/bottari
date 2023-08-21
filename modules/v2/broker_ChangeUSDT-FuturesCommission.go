package binance

import (
	"context"
	"encoding/json"
)

/*
	subAccount 선물 수수료 조정
*/

type BrokerChangeUSDTfuturesCommission struct {
	c               *Client
	subAccountId    string //계좌번호
	symbol          string //종목 코드
	makerAdjustment int
	takerAdjustment int
}

type BrokerChangeUSDTfuturesCommissionREQ struct {
	SubaccountId    int64  `json:"subAccountId"`
	Symbol          string `json:"symbol"`
	MakerAdjustment int    `json:"makerAdjustment"`
	TakerAdjustment int    `json:"takerAdjustment"`
	MakerCommission int    `json:"makerCommission"`
	TakerCommission int    `json:"takerCommission"`
}

//계좌번호
func (s *BrokerChangeUSDTfuturesCommission) SubAccountID(accid string) *BrokerChangeUSDTfuturesCommission {
	s.subAccountId = accid
	return s
}

//심볼
func (s *BrokerChangeUSDTfuturesCommission) Symbol(scode string) *BrokerChangeUSDTfuturesCommission {
	s.symbol = scode
	return s
}

func (s *BrokerChangeUSDTfuturesCommission) MakerAdjustment(AddValue int) *BrokerChangeUSDTfuturesCommission {
	s.makerAdjustment = AddValue
	return s
}

func (s *BrokerChangeUSDTfuturesCommission) TakerAdjustment(AddValue int) *BrokerChangeUSDTfuturesCommission {
	s.takerAdjustment = AddValue
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerChangeUSDTfuturesCommission) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {

	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":    s.subAccountId,
		"symbol":          s.symbol,
		"makerAdjustment": s.makerAdjustment,
		"takerAdjustment": s.takerAdjustment,
	}

	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do sends the request.
func (s *BrokerChangeUSDTfuturesCommission) Do(ctx context.Context, opts ...RequestOption) (res *BrokerChangeUSDTfuturesCommissionREQ, err error) {

	data, err := s.createReqData(ctx, "/sapi/v1/broker/subAccountApi/commission/futures", opts...)
	if err != nil {
		return nil, err
	}
	res = new(BrokerChangeUSDTfuturesCommissionREQ)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//=========================================================================================

/*
	subAccount 선물 수수료 조회
*/

type BrokerGetUSDTfuturesCommission struct {
	c            *Client
	subAccountId string  //계좌번호
	symbol       *string //종목 코드
}

//계좌번호
func (s *BrokerGetUSDTfuturesCommission) SubAccountID(accid string) *BrokerGetUSDTfuturesCommission {
	s.subAccountId = accid
	return s
}

//심볼
func (s *BrokerGetUSDTfuturesCommission) Symbol(scode string) *BrokerGetUSDTfuturesCommission {
	s.symbol = &scode
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerGetUSDTfuturesCommission) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {

	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	r.setParam("subAccountId", s.subAccountId)
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do sends the request.
func (s *BrokerGetUSDTfuturesCommission) Do(ctx context.Context, opts ...RequestOption) (res []BrokerChangeUSDTfuturesCommissionREQ, err error) {

	data, err := s.createReqData(ctx, "/sapi/v1/broker/subAccountApi/commission/futures", opts...)
	if err != nil {
		return []BrokerChangeUSDTfuturesCommissionREQ{}, err
	}
	res = make([]BrokerChangeUSDTfuturesCommissionREQ, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []BrokerChangeUSDTfuturesCommissionREQ{}, err
	}
	return res, nil
}
