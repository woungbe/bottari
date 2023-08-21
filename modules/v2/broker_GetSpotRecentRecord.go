package binance

import (
	"context"
	"encoding/json"
)

type BrokerGetSpotRecentRecord struct {
	c            *Client
	subAccountId string
	startTime    int64
	endTime      int64
	page         *int
	pageSize     *int
}

type BrokerGetSpotRecentRecordItem struct {
	SubaccountId string `json:"subaccountId"`
	Income       string `json:"income"`
	Asset        string `json:"asset"`
	Symbol       string `json:"symbol"`
	TradeId      int64  `json:"tradeId"`
	InTime       int64  `json:"time"`
}

func (s *BrokerGetSpotRecentRecord) SubAccountId(accID string) *BrokerGetSpotRecentRecord {
	s.subAccountId = accID
	return s
}

func (s *BrokerGetSpotRecentRecord) StartTime(sTime int64) *BrokerGetSpotRecentRecord {
	s.startTime = sTime
	return s
}

func (s *BrokerGetSpotRecentRecord) EndTime(eTime int64) *BrokerGetSpotRecentRecord {
	s.endTime = eTime
	return s
}

func (s *BrokerGetSpotRecentRecord) Page(p int) *BrokerGetSpotRecentRecord {
	s.page = &p
	return s
}
func (s *BrokerGetSpotRecentRecord) PageSize(ps int) *BrokerGetSpotRecentRecord {
	s.pageSize = &ps
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerGetSpotRecentRecord) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {

	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	r.setParam("subAccountId", s.subAccountId)
	r.setParam("startTime", s.startTime)
	r.setParam("endTime", s.endTime)

	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.pageSize != nil {
		r.setParam("size", *s.pageSize)
	}

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do sends the request.
func (s *BrokerGetSpotRecentRecord) Do(ctx context.Context, opts ...RequestOption) (res []BrokerGetSpotRecentRecordItem, err error) {

	data, err := s.createReqData(ctx, "/sapi/v1/broker/rebate/recentRecord", opts...)
	if err != nil {
		return nil, err
	}

	var ret []BrokerGetSpotRecentRecordItem

	err = json.Unmarshal(data, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
