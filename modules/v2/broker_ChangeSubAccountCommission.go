package binance

import (
	"context"
	"encoding/json"
)

/*
	subAccount 수수료 조정
*/

type BrokerChangeSubAccountCommission struct {
	c               *Client
	subAccountId    string //계좌번호
	makerCommission float64
	takerCommission float64
}

type BrokerChangeSubAccountCommissionREQ struct {
	SubaccountId    string `json:"subaccountId"`
	MakerCommission string `json:"makerCommission"`
	TakerCommission string `json:"takerCommission"`
}

//계좌번호
func (s *BrokerChangeSubAccountCommission) SubAccountID(accid string) *BrokerChangeSubAccountCommission {
	s.subAccountId = accid
	return s
}

//MakerCommission 마커 수수료
func (s *BrokerChangeSubAccountCommission) MakerCommission(makerfee float64) *BrokerChangeSubAccountCommission {
	s.makerCommission = makerfee
	return s
}

//MakerCommission 마커 수수료
func (s *BrokerChangeSubAccountCommission) TakerCommission(makerfee float64) *BrokerChangeSubAccountCommission {
	s.takerCommission = makerfee
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerChangeSubAccountCommission) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {

	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":    s.subAccountId,
		"makerCommission": s.makerCommission,
		"takerCommission": s.takerCommission,
	}

	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do sends the request.
func (s *BrokerChangeSubAccountCommission) Do(ctx context.Context, opts ...RequestOption) (res *BrokerChangeSubAccountCommissionREQ, err error) {

	data, err := s.createReqData(ctx, "/sapi/v1/broker/subAccountApi/commission", opts...)
	if err != nil {
		return nil, err
	}
	res = new(BrokerChangeSubAccountCommissionREQ)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
