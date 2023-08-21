package binance

import (
	"context"
	"encoding/json"
)

type BrokerSubAccountSpotSummary struct {
	c            *Client
	subAccountId string
	page         *int64
	pagesize     *int64
}

type BrokerSubAccountSpotSummaryItem struct {
	SubAccountId      string `json:"subAccountId"`
	TotalBalanceOfBtc string `json:"totalBalanceOfBtc"`
}

type BrokerSubAccountSpotSummaryRequest struct {
	Timestamp int64                             `json:"timestamp"`
	Data      []BrokerSubAccountSpotSummaryItem `json:"data"`
}

func (s *BrokerSubAccountSpotSummary) SubAccountId(accID string) *BrokerSubAccountSpotSummary {
	s.subAccountId = accID
	return s
}

func (s *BrokerSubAccountSpotSummary) Page(page int64) *BrokerSubAccountSpotSummary {
	s.page = &page
	return s
}

func (s *BrokerSubAccountSpotSummary) Pagesize(pagesize int64) *BrokerSubAccountSpotSummary {
	s.pagesize = &pagesize
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerSubAccountSpotSummary) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {

	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	if s.subAccountId != "" {
		r.setParam("subAccountId", s.subAccountId)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.pagesize != nil {
		r.setParam("size", *s.pagesize)
	}

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do sends the request.
func (s *BrokerSubAccountSpotSummary) Do(ctx context.Context, opts ...RequestOption) (res *BrokerSubAccountSpotSummaryRequest, err error) {

	res = new(BrokerSubAccountSpotSummaryRequest)
	data, err := s.createReqData(ctx, "/sapi/v1/broker/subAccount/spotSummary", opts...)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *BrokerSubAccountSpotSummary) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {

	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/subAccount/spotSummary",
		secType:  secTypeSigned,
	}

	if s.subAccountId != "" {
		r.setParam("subAccountId", s.subAccountId)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.pagesize != nil {
		r.setParam("size", *s.pagesize)
	}

	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//-----------------------------------------------------------

type BrokerSubAccountFutureSummary struct {
	c            *Client
	subAccountId string
	page         *int64
	pagesize     *int64
}

type BrokerSubAccountFutureSummaryItem struct {
	FuturesEnable                     bool   `json:"futuresEnable"`
	SubAccountId                      string `json:"subAccountId"`
	TotalInitialMarginOfUsdt          string `json:"totalInitialMarginOfUsdt"`          //": "0.03558521",  //  initial margin
	TotalMaintenanceMarginOfUsdt      string `json:"totalMaintenanceMarginOfUsdt"`      //": "0.02695000", // maintenance margin
	TotalWalletBalanceOfUsdt          string `json:"totalWalletBalanceOfUsdt"`          //": "8.23222312", //  wallet balance
	TotalUnrealizedProfitOfUsdt       string `json:"totalUnrealizedProfitOfUsdt"`       //": "-0.78628370", //  unrealized profit
	TotalMarginBalanceOfUsdt          string `json:"totalMarginBalanceOfUsdt"`          //": "8.23432343",  // margin balance
	TotalPositionInitialMarginOfUsdt  string `json:"totalPositionInitialMarginOfUsdt"`  //": "0.33683000", // position initial margin
	TotalOpenOrderInitialMarginOfUsdt string `json:"totalOpenOrderInitialMarginOfUsdt"` //": "0.00000000" // open order initial margin
}

type BrokerSubAccountFutureSummaryRequest struct {
	Timestamp int64                               `json:"timestamp"`
	Data      []BrokerSubAccountFutureSummaryItem `json:"data"`
}

func (s *BrokerSubAccountFutureSummary) SubAccountId(accID string) *BrokerSubAccountFutureSummary {
	s.subAccountId = accID
	return s
}

func (s *BrokerSubAccountFutureSummary) Page(page int64) *BrokerSubAccountFutureSummary {
	s.page = &page
	return s
}

func (s *BrokerSubAccountFutureSummary) Pagesize(pagesize int64) *BrokerSubAccountFutureSummary {
	s.pagesize = &pagesize
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerSubAccountFutureSummary) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {

	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	if s.subAccountId != "" {
		r.setParam("subAccountId", s.subAccountId)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.pagesize != nil {
		r.setParam("size", *s.pagesize)
	}

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do sends the request.
func (s *BrokerSubAccountFutureSummary) Do(ctx context.Context, opts ...RequestOption) (res *BrokerSubAccountFutureSummaryRequest, err error) {

	res = new(BrokerSubAccountFutureSummaryRequest)
	data, err := s.createReqData(ctx, "/sapi/v1/broker/subAccount/futuresSummary", opts...)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, res)
	if err != nil {
		return res, err
	}
	return res, nil
}

//----------------
type BrokerSubAccountFutureSummaryV2 struct {
	c            *Client
	subAccountId string
	futuresType  int //1:USDT Margined Futures, 2:COIN Margined Futures
	page         *int64
	pagesize     *int64
}

func (s *BrokerSubAccountFutureSummaryV2) SubAccountId(accID string) *BrokerSubAccountFutureSummaryV2 {
	s.subAccountId = accID
	return s
}
func (s *BrokerSubAccountFutureSummaryV2) FuturesType(futuresType int) *BrokerSubAccountFutureSummaryV2 {
	s.futuresType = futuresType
	return s
}

func (s *BrokerSubAccountFutureSummaryV2) Page(page int64) *BrokerSubAccountFutureSummaryV2 {
	s.page = &page
	return s
}

func (s *BrokerSubAccountFutureSummaryV2) Pagesize(pagesize int64) *BrokerSubAccountFutureSummaryV2 {
	s.pagesize = &pagesize
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerSubAccountFutureSummaryV2) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {

	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	r.setParam("futuresType", s.futuresType)
	if s.subAccountId != "" {
		r.setParam("subAccountId", s.subAccountId)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.pagesize != nil {
		r.setParam("size", *s.pagesize)
	}

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do sends the request.
func (s *BrokerSubAccountFutureSummaryV2) Do(ctx context.Context, opts ...RequestOption) (res *BrokerSubAccountFutureSummaryRequest, err error) {

	res = new(BrokerSubAccountFutureSummaryRequest)
	data, err := s.createReqData(ctx, "/sapi/v2/broker/subAccount/futuresSummary", opts...)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *BrokerSubAccountFutureSummaryV2) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v2/broker/subAccount/futuresSummary",
		secType:  secTypeSigned,
	}

	r.setParam("futuresType", s.futuresType)
	if s.subAccountId != "" {
		r.setParam("subAccountId", s.subAccountId)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.pagesize != nil {
		r.setParam("size", *s.pagesize)
	}

	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}
