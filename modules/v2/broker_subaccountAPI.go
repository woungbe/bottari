package binance

import (
	"context"
	"encoding/json"
)

/*
	subAccount API key 생성
*/
type BrokerCreateAPIKey struct {
	c            *Client
	subAccountId string //계좌번호
	canTrade     bool   //일반 거래 상태
	marginTrade  bool   //마진거래 상태
	futuresTrade bool   //선물 거래 상태
}

type BrokerCreateAPIKeyREQ struct {
	SubaccountId string `json:"subaccountId"`
	ApiKey       string `json:"apiKey"`
	SecretKey    string `json:"secretKey"`
	CanTrade     bool   `json:"canTrade"`
	MarginTrade  bool   `json:"marginTrade"`
	FuturesTrade bool   `json:"futuresTrade"`
}

//계좌번호
func (s *BrokerCreateAPIKey) SubAccountID(accid string) *BrokerCreateAPIKey {
	s.subAccountId = accid
	return s
}

//SpotTrade 일반 거래 상태
func (s *BrokerCreateAPIKey) SpotTrade(useFlg bool) *BrokerCreateAPIKey {
	s.canTrade = useFlg
	return s
}

//SpotTrade 마진 거래 상태
func (s *BrokerCreateAPIKey) MarginTrade(useFlg bool) *BrokerCreateAPIKey {
	s.marginTrade = useFlg
	return s
}

//SpotTrade 마진 거래 상태
func (s *BrokerCreateAPIKey) FuturesTrade(useFlg bool) *BrokerCreateAPIKey {
	s.futuresTrade = useFlg
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerCreateAPIKey) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {

	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId": s.subAccountId,
	}

	m["canTrade"] = s.canTrade
	m["marginTrade"] = s.marginTrade
	m["futuresTrade"] = s.futuresTrade

	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do sends the request.
func (s *BrokerCreateAPIKey) Do(ctx context.Context, opts ...RequestOption) (res *BrokerCreateAPIKeyREQ, err error) {

	data, err := s.createReqData(ctx, "/sapi/v1/broker/subAccountApi", opts...)
	if err != nil {
		return nil, err
	}
	res = new(BrokerCreateAPIKeyREQ)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//==============================================================================

/*
	subAccount API key 삭제
*/
type BrokerDeleteAPIKey struct {
	c                *Client
	subAccountId     string //계좌번호
	subAccountApiKey string //api key
}

//계좌번호
func (s *BrokerDeleteAPIKey) SubAccountID(accid string) *BrokerDeleteAPIKey {
	s.subAccountId = accid
	return s
}

func (s *BrokerDeleteAPIKey) SubAccountApiKey(apikey string) *BrokerDeleteAPIKey {
	s.subAccountApiKey = apikey
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerDeleteAPIKey) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {

	r := &request{
		method:   "DELETE",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":     s.subAccountId,
		"subAccountApiKey": s.subAccountApiKey,
	}

	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do sends the request.
func (s *BrokerDeleteAPIKey) Do(ctx context.Context, opts ...RequestOption) error {
	_, err := s.createReqData(ctx, "/sapi/v1/broker/subAccountApi", opts...)
	return err
}

//================================================================================

/*
	subAccount API key 가져오기
*/
type BrokerGetSubAPIKey struct {
	c            *Client
	subAccountId string //계좌번호
}

type BrokerGetSubAPIKeyREQ struct {
	SubaccountId string `json:"subaccountId"`
	ApiKey       string `json:"apiKey"`
	CanTrade     bool   `json:"canTrade"`
	MarginTrade  bool   `json:"marginTrade"`
	FuturesTrade bool   `json:"futuresTrade"`
}

//계좌번호
func (s *BrokerGetSubAPIKey) SubAccountID(accid string) *BrokerGetSubAPIKey {
	s.subAccountId = accid
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerGetSubAPIKey) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {

	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	//r.setParams()
	r.setParam("subAccountId", s.subAccountId)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do sends the request.
func (s *BrokerGetSubAPIKey) Do(ctx context.Context, opts ...RequestOption) (res []BrokerGetSubAPIKeyREQ, err error) {

	data, err := s.createReqData(ctx, "/sapi/v1/broker/subAccountApi", opts...)
	if err != nil {
		return []BrokerGetSubAPIKeyREQ{}, err
	}
	res = make([]BrokerGetSubAPIKeyREQ, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []BrokerGetSubAPIKeyREQ{}, err
	}
	return res, nil
}

//================================================================================

/*
	subAccount API  정보 수정
*/

type BrokerChangeSubAPIPermission struct {
	c                *Client
	subAccountId     string //계좌번호
	subAccountApiKey string //api key
	canTrade         bool   //일반 거래 상태
	marginTrade      bool   //마진거래 상태
	futuresTrade     bool   //선물 거래 상태

}
type BrokerChangeSubAPIPermissionREQ struct {
	SubaccountId string `json:"subaccountId"`
	ApiKey       string `json:"apiKey"`
	CanTrade     bool   `json:"canTrade"`
	MarginTrade  bool   `json:"marginTrade"`
	FuturesTrade bool   `json:"futuresTrade"`
}

//계좌번호
func (s *BrokerChangeSubAPIPermission) SubAccountID(accid string) *BrokerChangeSubAPIPermission {
	s.subAccountId = accid
	return s
}

func (s *BrokerChangeSubAPIPermission) SubAccountApiKey(apikey string) *BrokerChangeSubAPIPermission {
	s.subAccountApiKey = apikey
	return s
}

//SpotTrade 일반 거래 상태
func (s *BrokerChangeSubAPIPermission) SpotTrade(useFlg bool) *BrokerChangeSubAPIPermission {
	s.canTrade = useFlg
	return s
}

//SpotTrade 마진 거래 상태
func (s *BrokerChangeSubAPIPermission) MarginTrade(useFlg bool) *BrokerChangeSubAPIPermission {
	s.marginTrade = useFlg
	return s
}

//SpotTrade 마진 거래 상태
func (s *BrokerChangeSubAPIPermission) FuturesTrade(useFlg bool) *BrokerChangeSubAPIPermission {
	s.futuresTrade = useFlg
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerChangeSubAPIPermission) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {

	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":     s.subAccountId,
		"subAccountApiKey": s.subAccountApiKey,
	}

	m["canTrade"] = s.canTrade
	m["marginTrade"] = s.marginTrade
	m["futuresTrade"] = s.futuresTrade

	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do sends the request.
func (s *BrokerChangeSubAPIPermission) Do(ctx context.Context, opts ...RequestOption) (res *BrokerChangeSubAPIPermissionREQ, err error) {

	data, err := s.createReqData(ctx, "/sapi/v1/broker/subAccountApi/permission", opts...)
	if err != nil {
		return nil, err
	}
	res = new(BrokerChangeSubAPIPermissionREQ)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
