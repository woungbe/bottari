package binance

import (
	"context"
)

/*
	코인 지갑정보 전체 가져오기
*/

type GetAllWalletCoinInfoService struct {
	c *Client
}

// Deposit represents a single deposit entry.
type DepositListItem struct {
	InsertTime   int64   `json:"insertTime"`
	Amount       float64 `json:"amount"`
	Coin         string  `json:"coin"`
	Address      string  `json:"address"`
	AddressTag   string  `json:"addressTag"`
	NetworkName  string  `json:"network"`
	TxID         string  `json:"txId"`
	Status       int     `json:"status"`
	ConfirmTimes string  `json:"confirmTimes"`
}

// GetAllWalletCoinInfoService 코인 정보 조회
func (s *GetAllWalletCoinInfoService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/capital/config/getall",
		secType:  secTypeSigned,
	}
	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

/*
	계좌 일별 스낵샵 조회
*/

func (s *GetAccountSnapshotService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/accountSnapshot",
		secType:  secTypeSigned,
	}
	r.setParam("type", s.accountType)

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

/*
	입금 관련 명령
*/

type GetSAPIDepositsAddressService struct {
	c           *Client
	asset       string
	networkname *string //네트웍 이름
}

func (s *GetSAPIDepositsAddressService) Asset(v string) *GetSAPIDepositsAddressService {
	s.asset = v
	return s
}

func (s *GetSAPIDepositsAddressService) NetworkName(v string) *GetSAPIDepositsAddressService {
	s.networkname = &v
	return s
}

// StartUserStreamService 라이센스 키 요청
func (s *GetSAPIDepositsAddressService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/capital/deposit/address",
		secType:  secTypeSigned,
	}

	r.setParam("coin", s.asset)

	if s.networkname != nil {
		r.setParam("network", *s.networkname)
	}

	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

/*
	입금 내역 리스트 가져오기
*/

type GetSAPIDepositsHistory struct {
	c         *Client
	asset     *string //코인명
	status    *int    //상태 0(0:pending,6: credited but cannot withdraw, 1:success)
	startTime *int64  //조회 시작 시간 기본으로 90일 전임
	endTime   *int64  //조회 종료 시간 기본값은 현재
	offset    *int    //defalult 0 => 페이지위치인듯한데 아직모르겠음
	limit     *int    //
}

// Asset sets the asset parameter.
func (s *GetSAPIDepositsHistory) Asset(asset string) *GetSAPIDepositsHistory {
	s.asset = &asset
	return s
}

// Status sets the status parameter.
func (s *GetSAPIDepositsHistory) Status(status int) *GetSAPIDepositsHistory {
	s.status = &status
	return s
}

// StartTime sets the startTime parameter.
// If present, EndTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *GetSAPIDepositsHistory) StartTime(startTime int64) *GetSAPIDepositsHistory {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
// If present, StartTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *GetSAPIDepositsHistory) EndTime(endTime int64) *GetSAPIDepositsHistory {
	s.endTime = &endTime
	return s
}

func (s *GetSAPIDepositsHistory) Offset(offset int) *GetSAPIDepositsHistory {
	s.offset = &offset
	return s
}
func (s *GetSAPIDepositsHistory) Limit(limit int) *GetSAPIDepositsHistory {
	s.limit = &limit
	return s
}

// GetSAPIDepositsHistory 입금리스트 조회
func (s *GetSAPIDepositsHistory) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/capital/deposit/hisrec",
		secType:  secTypeSigned,
	}

	if s.asset != nil {
		r.setParam("coin", *s.asset)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

/*
	출금 내역 리스트 가져오기
*/

type GetSAPIWithdrawHistory struct {
	c         *Client
	asset     *string //코인명
	status    *int    //상태 0(0:pending,6: credited but cannot withdraw, 1:success)
	startTime *int64  //조회 시작 시간 기본으로 90일 전임
	endTime   *int64  //조회 종료 시간 기본값은 현재
	offset    *int    //defalult 0 => 페이지위치인듯한데 아직모르겠음
	limit     *int    //
}

type SAPIWithdraw struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	ApplyTime       string `json:"applyTime"`
	Coin            string `json:"coin"`
	ID              string `json:"id"`
	WithdrawOrderID string `json:"withdrawOrderID"`
	Network         string `json:"network"`
	TransferType    int    `json:"transferType"`
	Status          int    `json:"status"`
	TransactionFee  string `json:"transactionFee"`
	ConfirmNo       int    `json:"confirmNo"`
	TxID            string `json:"txId"`
	AddressTag      string `json:"addressTag"`
}

// Asset sets the asset parameter.
func (s *GetSAPIWithdrawHistory) Asset(asset string) *GetSAPIWithdrawHistory {
	s.asset = &asset
	return s
}

// Status sets the status parameter.
func (s *GetSAPIWithdrawHistory) Status(status int) *GetSAPIWithdrawHistory {
	s.status = &status
	return s
}

// StartTime sets the startTime parameter.
// If present, EndTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *GetSAPIWithdrawHistory) StartTime(startTime int64) *GetSAPIWithdrawHistory {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
// If present, StartTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *GetSAPIWithdrawHistory) EndTime(endTime int64) *GetSAPIWithdrawHistory {
	s.endTime = &endTime
	return s
}

func (s *GetSAPIWithdrawHistory) Offset(offset int) *GetSAPIWithdrawHistory {
	s.offset = &offset
	return s
}
func (s *GetSAPIWithdrawHistory) Limit(limit int) *GetSAPIWithdrawHistory {
	s.limit = &limit
	return s
}

// GetSAPIWithdrawHistory 라이센스 키 요청
func (s *GetSAPIWithdrawHistory) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/capital/withdraw/history",
		secType:  secTypeSigned,
	}

	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

/*
	현물 <-> 선물
*/

// Do send request
func (s *FuturesTransferService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/futures/transfer",
		secType:  secTypeSigned,
	}
	m := params{
		"asset":  s.asset,
		"amount": s.amount,
		"type":   s.transferType,
	}
	r.setFormParams(m)

	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

// Do send request
func (s *ListFuturesTransferService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/futures/transfer",
		secType:  secTypeSigned,
	}
	r.setParams(params{
		"asset":     s.asset,
		"startTime": s.startTime,
	})
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}

	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}

/*
	출금
*/

type CreateSAPIWithdrawService struct {
	c                  *Client
	asset              string
	withdrawOrderID    *string
	network            *string
	address            string
	addressTag         *string
	amount             string
	transactionFeeFlag *bool
	name               *string
}

// Asset sets the asset parameter (MANDATORY).
func (s *CreateSAPIWithdrawService) Asset(v string) *CreateSAPIWithdrawService {
	s.asset = v
	return s
}

// WithdrawOrderID sets the withdrawOrderID parameter.
func (s *CreateSAPIWithdrawService) WithdrawOrderID(v string) *CreateSAPIWithdrawService {
	s.withdrawOrderID = &v
	return s
}

// Network sets the network parameter.
func (s *CreateSAPIWithdrawService) Network(v string) *CreateSAPIWithdrawService {
	s.network = &v
	return s
}

// Address sets the address parameter (MANDATORY).
func (s *CreateSAPIWithdrawService) Address(v string) *CreateSAPIWithdrawService {
	s.address = v
	return s
}

// AddressTag sets the addressTag parameter.
func (s *CreateSAPIWithdrawService) AddressTag(v string) *CreateSAPIWithdrawService {
	s.addressTag = &v
	return s
}

// Amount sets the amount parameter (MANDATORY).
func (s *CreateSAPIWithdrawService) Amount(v string) *CreateSAPIWithdrawService {
	s.amount = v
	return s
}

// TransactionFeeFlag sets the transactionFeeFlag parameter.
func (s *CreateSAPIWithdrawService) TransactionFeeFlag(v bool) *CreateSAPIWithdrawService {
	s.transactionFeeFlag = &v
	return s
}

// Name sets the name parameter.
func (s *CreateSAPIWithdrawService) Name(v string) *CreateSAPIWithdrawService {
	s.name = &v
	return s
}

// Do sends the request.
func (s *CreateSAPIWithdrawService) Do_Core(ctx context.Context, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/capital/withdraw/apply",
		secType:  secTypeSigned,
	}
	r.setParam("coin", s.asset)
	r.setParam("address", s.address)
	r.setParam("amount", s.amount)
	if v := s.withdrawOrderID; v != nil {
		r.setParam("withdrawOrderId", *v)
	}
	if v := s.network; v != nil {
		r.setParam("network", *v)
	}
	if v := s.addressTag; v != nil {
		r.setParam("addressTag", *v)
	}
	if v := s.transactionFeeFlag; v != nil {
		r.setParam("transactionFeeFlag", *v)
	}
	if v := s.name; v != nil {
		r.setParam("name", *v)
	}

	data, st, err := s.c.callAPI_Core(ctx, r, opts...)
	return data, st, err
}
