package binance

import (
	"context"
	"encoding/json"
)

type BrokerGetFutuRecentRecord struct {
	c           *Client
	futuresType int
	startTime   int64
	endTime     int64
	page        *int
	pageSize    *int
}

type BrokerGetFutuRecentRecordItem struct {
	SubaccountId string `json:"subaccountId"`
	Income       string `json:"income"`
	Asset        string `json:"asset"`
	Symbol       string `json:"symbol"`
	TradeId      int64  `json:"tradeId"`
	InTime       int64  `json:"time"`
}

func (s *BrokerGetFutuRecentRecord) FuturesType(futuType int) *BrokerGetFutuRecentRecord {
	s.futuresType = futuType
	return s
}

func (s *BrokerGetFutuRecentRecord) StartTime(sTime int64) *BrokerGetFutuRecentRecord {
	s.startTime = sTime
	return s
}

func (s *BrokerGetFutuRecentRecord) EndTime(eTime int64) *BrokerGetFutuRecentRecord {
	s.endTime = eTime
	return s
}

func (s *BrokerGetFutuRecentRecord) Page(p int) *BrokerGetFutuRecentRecord {
	s.page = &p
	return s
}
func (s *BrokerGetFutuRecentRecord) PageSize(ps int) *BrokerGetFutuRecentRecord {
	s.pageSize = &ps
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerGetFutuRecentRecord) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {

	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	r.setParam("futuresType", s.futuresType)
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
func (s *BrokerGetFutuRecentRecord) Do(ctx context.Context, opts ...RequestOption) (res []BrokerGetFutuRecentRecordItem, err error) {

	data, err := s.createReqData(ctx, "/sapi/v1/broker/rebate/futures/recentRecord", opts...)
	if err != nil {
		return nil, err
	}

	var ret []BrokerGetFutuRecentRecordItem

	err = json.Unmarshal(data, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
