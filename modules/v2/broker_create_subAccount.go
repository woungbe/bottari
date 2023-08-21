package binance

import (
	"context"
	"encoding/json"
)

type BrokerCreateSubAccount struct {
	c *Client
}

//REQCreateSubAccount create subaccount request data
type REQCreateSubAccount struct {
	SubaccountId string `json:"subaccountId"`
	Email        string `json:"email"`
	Tag          string `json:"tag"`
}

func (s *BrokerCreateSubAccount) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil

}

// Do sends the request.
func (s *BrokerCreateSubAccount) Do(ctx context.Context, opts ...RequestOption) (res *REQCreateSubAccount, err error) {

	data, err := s.createReqData(ctx, "/sapi/v1/broker/subAccount", opts...)
	if err != nil {
		return nil, err
	}
	res = new(REQCreateSubAccount)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//=========================================================================================

/*
	subAccount 정보조회
*/

type BrokerGetSubAccountInfo struct {
	c            *Client
	subAccountId *string //계좌번호
}

//REQCreateSubAccount create subaccount request data
type BrokerGetSubAccountInfoREQ struct {
	SubaccountId          string `json:"subaccountId"`
	Email                 string `json:"email"`
	MakerCommission       string `json:"makerCommission"`
	TakerCommission       string `json:"takerCommission"`
	MarginMakerCommission string `json:"marginMakerCommission"`
	MarginTakerCommission string `json:"marginTakerCommission"`
	CreateTime            int64  `json:"createTime"`
}

//계좌번호
func (s *BrokerGetSubAccountInfo) SubAccountID(accid string) *BrokerGetSubAccountInfo {
	s.subAccountId = &accid
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerGetSubAccountInfo) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {

	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	if s.subAccountId != nil {
		r.setParam("subAccountId", *s.subAccountId)
	}

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do sends the request.
func (s *BrokerGetSubAccountInfo) Do(ctx context.Context, opts ...RequestOption) (res []BrokerGetSubAccountInfoREQ, err error) {

	data, err := s.createReqData(ctx, "/sapi/v1/broker/subAccount", opts...)
	if err != nil {
		return []BrokerGetSubAccountInfoREQ{}, err
	}
	res = make([]BrokerGetSubAccountInfoREQ, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []BrokerGetSubAccountInfoREQ{}, err
	}
	return res, nil
}
