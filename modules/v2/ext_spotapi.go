package binance

import (
	"context"
)

//서버 시간 체크
func (s *ServerTimeService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v3/time",
	}
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

// StartUserStreamService 라이센스 키 요청
func (s *StartUserStreamService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/api/v3/userDataStream",
		secType:  secTypeAPIKey,
	}
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//KeepaliveUserStreamService 라이센스키 연장 요청
func (s *KeepaliveUserStreamService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "PUT",
		endpoint: "/api/v3/userDataStream",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//CloseUserStreamService 라이센스 정지
func (s *CloseUserStreamService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "DELETE",
		endpoint: "/api/v3/userDataStream",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//GetAccountService 계좌 상세 정보 조회
func (s *GetAccountService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v3/account",
		secType:  secTypeSigned,
	}
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)

	return data, st, err
}

/*
	종목별 거래내역 조회
*/

type GetMyTradeList struct {
	c         *Client
	symbol    string //코인명
	startTime *int64 //조회 시작 시간
	endTime   *int64 //조회 종료 시간
	fromId    *int64 //	TradeId to fetch from. Default gets most recent trades.
	limit     *int   //
}

func (s *GetMyTradeList) Symbol(v string) *GetMyTradeList {
	s.symbol = v
	return s
}

func (s *GetMyTradeList) StartTime(startTime int64) *GetMyTradeList {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
func (s *GetMyTradeList) EndTime(endTime int64) *GetMyTradeList {
	s.endTime = &endTime
	return s
}

func (s *GetMyTradeList) FromId(fromId int64) *GetMyTradeList {
	s.fromId = &fromId
	return s
}

func (s *GetMyTradeList) Limit(limit int) *GetMyTradeList {
	s.limit = &limit
	return s
}

// GetMyTradeList 입금리스트 조회
func (s *GetMyTradeList) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v3/myTrades",
		secType:  secTypeSigned,
	}

	r.setParam("symbol", s.symbol)

	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	if s.fromId != nil {
		r.setParam("fromId", *s.fromId)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//All Orders 조회
func (s *ListOrdersService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v3/allOrders",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//미체결내역 조회
func (s *ListOpenOrdersService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v3/openOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//신규 주문
func (s *CreateOrderService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/api/v3/order",
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
		"side":   s.side,
		"type":   s.orderType,
	}
	if s.quantity != nil {
		m["quantity"] = *s.quantity
	}
	if s.quoteOrderQty != nil {
		m["quoteOrderQty"] = *s.quoteOrderQty
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.newClientOrderID != nil {
		m["newClientOrderId"] = *s.newClientOrderID
	}
	if s.stopPrice != nil {
		m["stopPrice"] = *s.stopPrice
	}
	if s.icebergQuantity != nil {
		m["icebergQty"] = *s.icebergQuantity
	}
	if s.newOrderRespType != nil {
		m["newOrderRespType"] = *s.newOrderRespType
	}
	r.setFormParams(m)
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)

	return data, st, err
}

//주문 취소
func (s *CancelOrderService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "DELETE",
		endpoint: "/api/v3/order",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setFormParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setFormParam("origClientOrderId", *s.origClientOrderID)
	}
	if s.newClientOrderID != nil {
		r.setFormParam("newClientOrderId", *s.newClientOrderID)
	}
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)

	return data, st, err
}

//전체 주문 취소
func (s *CancelOpenOrdersService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "DELETE",
		endpoint: "/api/v3/openOrders",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}
