package futures

import (
	"context"
)

// Do send request
func (s *ExchangeInfoService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/exchangeInfo",
		secType:  secTypeNone,
	}
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

func (s *ListPriceChangeStatsService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/ticker/24hr",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

// Do send request
func (s *ServerTimeService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/time",
	}
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

func (s *StartUserStreamService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/fapi/v1/listenKey",
		secType:  secTypeSigned,
	}
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

func (s *KeepaliveUserStreamService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "PUT",
		endpoint: "/fapi/v1/listenKey",
		secType:  secTypeSigned,
	}
	r.setFormParam("listenKey", s.listenKey)
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

func (s *CloseUserStreamService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "DELETE",
		endpoint: "/fapi/v1/listenKey",
		secType:  secTypeSigned,
	}
	r.setFormParam("listenKey", s.listenKey)
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

// ChangePositionModeService 헷지 모드 변경
func (s *ChangePositionModeService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/fapi/v1/positionSide/dual",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{
		"dualSidePosition": s.dualSide,
	})
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//GetPositionModeService 헷지모드 정보 가져오기
func (s *GetPositionModeService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/positionSide/dual",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{})
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//GetAccountService 선물 계좌 정보 가져옴 (발란스 , 포지션등)
func (s *GetAccountService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v2/account",
		secType:  secTypeSigned,
	}
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//GetBalanceService 선물 발란스 정보
func (s *GetBalanceService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v2/balance",
		secType:  secTypeSigned,
	}
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//ChangeMarginTypeService 마진타입 변경
func (s *ChangeMarginTypeService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/fapi/v1/marginType",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{
		"symbol":     s.symbol,
		"marginType": s.marginType,
	})
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//ChangeLeverageService 레버리지 변경
func (s *ChangeLeverageService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/fapi/v1/leverage",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{
		"symbol":   s.symbol,
		"leverage": s.leverage,
	})
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//GetPositionRiskService 포지션 정보 조회
func (s *GetPositionRiskService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v2/positionRisk",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//GetLeverageBracketService 레버리지 정보 조회
func (s *GetLeverageBracketService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/leverageBracket",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//GetIncomeHistoryService 트랜잭션 리스트 조회
func (s *GetIncomeHistoryService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/income",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.incomeType != "" {
		r.setParam("incomeType", s.incomeType)
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

//GetPositionMarginHistoryService 포지션 마진 변경 내역 조회
func (s *GetPositionMarginHistoryService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/positionMargin/history",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s._type != nil {
		r.setParam("type", *s._type)
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

//ListAccountTradeService 거래내역 조회
func (s *ListAccountTradeService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/userTrades",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.fromID != nil {
		r.setParam("fromID", *s.fromID)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//ListOrdersService 종목 주문 내역 조회
func (s *ListOrdersService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/allOrders",
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

//ListOpenOrdersService 미체결 주문 내역 조회
func (s *ListOpenOrdersService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/openOrders",
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
		endpoint: "/fapi/v1/order",
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":           s.symbol,
		"side":             s.side,
		"type":             s.orderType,
		"quantity":         s.quantity,
		"newOrderRespType": s.newOrderRespType,
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
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
	if s.workingType != nil {
		m["workingType"] = *s.workingType
	}
	if s.priceProtect != nil {
		m["priceProtect"] = *s.priceProtect
	}
	if s.activationPrice != nil {
		m["activationPrice"] = *s.activationPrice
	}
	if s.callbackRate != nil {
		m["callbackRate"] = *s.callbackRate
	}
	if s.closePosition != nil {
		m["closePosition"] = *s.closePosition
	}
	r.setFormParams(m)

	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//CancelOrderService 취소 주문
func (s *CancelOrderService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "DELETE",
		endpoint: "/fapi/v1/order",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setFormParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setFormParam("origClientOrderId", *s.origClientOrderID)
	}

	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//CancelAllOpenOrdersService 전체 주문 취소
func (s *CancelAllOpenOrdersService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "DELETE",
		endpoint: "/fapi/v1/allOpenOrders",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//UpdatePositionMarginService 격리 마진 수정
func (s *UpdatePositionMarginService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/fapi/v1/positionMargin",
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
		"amount": s.amount,
		"type":   s.actionType,
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	r.setFormParams(m)

	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

// 멀티 주문 취소
type CancelMultipleOpenOrdersService struct {
	c                     *Client
	symbol                string
	orderIdList           string
	origClientOrderIdList string
}

// Symbol set symbol
func (s *CancelMultipleOpenOrdersService) Symbol(symbol string) *CancelMultipleOpenOrdersService {
	s.symbol = symbol
	return s
}

func (s *CancelMultipleOpenOrdersService) OrderIdList(orderIdList string) *CancelMultipleOpenOrdersService {
	s.orderIdList = orderIdList
	return s
}

func (s *CancelMultipleOpenOrdersService) OrigClientOrderIdList(origClientOrderIdList string) *CancelMultipleOpenOrdersService {
	s.origClientOrderIdList = origClientOrderIdList
	return s
}

// Do send request
func (s *CancelMultipleOpenOrdersService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "DELETE",
		endpoint: "/fapi/v1/batchOrders",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.orderIdList != "" {
		r.setFormParam("orderIdList", s.orderIdList)
	}

	if s.origClientOrderIdList != "" {
		r.setFormParam("origClientOrderIdList", s.origClientOrderIdList)
	}
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

// Do send request
func (s *ListBookTickersService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {

	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/ticker/bookTicker",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

//멀티 신규 주문
type MultipleNewOrdersService struct {
	c               *Client
	batchOrdersList string
}

func (s *MultipleNewOrdersService) BatchOrdersList(orderInfoList string) *MultipleNewOrdersService {
	s.batchOrdersList = orderInfoList
	return s
}

// Do send request
func (s *MultipleNewOrdersService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/fapi/v1/batchOrders",
		secType:  secTypeSigned,
	}
	r.setFormParam("batchOrders", s.batchOrdersList)
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}
