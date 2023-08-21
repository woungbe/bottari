package binance

import (
	"context"
	"encoding/json"

	"github.com/adshao/go-binance/v2/common"
)

/*
	브로커/ 선물 계좌 정보 조회
*/

type BrokerGetFutureAssetInfo struct {
	c            *Client
	futuresType  int     //	YES	1:USDT Margined Futures, 2:COIN Margined Futures
	subAccountId *string //계좌번호
	page         *int
	pageSize     *int
}

type BrokerGetFutureAssetInfoReq struct {
	Timestamp int64                                `json:"timestamp"`
	Data      []BrokerGetFutureAssetInfoRecordItem `json:"data"`
}

//
type BrokerGetFutureAssetInfoRecordItem struct {
	FuturesEnable                     bool   `json:"FuturesEnable"`                     // "futuresEnable": true, // if enable futures
	SubAccountId                      string `json:"subAccountId"`                      // "subAccountId": "367537027503425913",
	TotalInitialMarginOfUsdt          string `json:"totalInitialMarginOfUsdt"`          //"totalInitialMarginOfUsdt": "0.03558521",  //  initial margin
	TotalMaintenanceMarginOfUsdt      string `json:"totalMaintenanceMarginOfUsdt"`      //"totalMaintenanceMarginOfUsdt": "0.02695000", // maintenance margin
	TotalWalletBalanceOfUsdt          string `json:"totalWalletBalanceOfUsdt"`          //"totalWalletBalanceOfUsdt": "8.23222312", //  wallet balance
	TotalUnrealizedProfitOfUsdt       string `json:"totalUnrealizedProfitOfUsdt"`       //"totalUnrealizedProfitOfUsdt": "-0.78628370", //  unrealized profit
	TotalMarginBalanceOfUsdt          string `json:"totalMarginBalanceOfUsdt"`          //"totalMarginBalanceOfUsdt": "8.23432343",  // margin balance
	TotalPositionInitialMarginOfUsdt  string `json:"totalPositionInitialMarginOfUsdt"`  //"totalPositionInitialMarginOfUsdt": "0.33683000", // position initial margin
	TotalOpenOrderInitialMarginOfUsdt string `json:"totalOpenOrderInitialMarginOfUsdt"` //"totalOpenOrderInitialMarginOfUsdt": "0.00000000" // open order initial margin
}

func (s *BrokerGetFutureAssetInfo) FuturesType(futuresType int) *BrokerGetFutureAssetInfo {
	s.futuresType = futuresType
	return s
}

func (s *BrokerGetFutureAssetInfo) SubAccountId(subAccountId string) *BrokerGetFutureAssetInfo {
	s.subAccountId = &subAccountId
	return s
}
func (s *BrokerGetFutureAssetInfo) Page(p int) *BrokerGetFutureAssetInfo {
	s.page = &p
	return s
}
func (s *BrokerGetFutureAssetInfo) PageSize(ps int) *BrokerGetFutureAssetInfo {
	s.pageSize = &ps
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerGetFutureAssetInfo) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, reqStatus int, err error) {

	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	r.setParam("futuresType", s.futuresType)

	if s.subAccountId != nil {
		r.setParam("subAccountId", *s.subAccountId)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.pageSize != nil {
		r.setParam("size", *s.pageSize)
	}

	data, reqStatus, err = s.c.callAPI_Core(ctx, r, opts...)
	if err != nil {
		return []byte{}, reqStatus, err
	}
	return data, reqStatus, nil
}

// Do sends the request.
func (s *BrokerGetFutureAssetInfo) Do_Svr(ctx context.Context, opts ...RequestOption) (res *BrokerGetFutureAssetInfoReq, reqErrCode *common.APIError, reqStatus int, err error) {
	data, status, err := s.createReqData(ctx, "/sapi/v2/broker/subAccount/futuresSummary", opts...)
	if err != nil {
		return nil, nil, status, err
	}
	res = new(BrokerGetFutureAssetInfoReq)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, nil, status, err
	}
	return res, nil, status, nil
}
