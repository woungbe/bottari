package binance

import (
	"context"
	"encoding/json"

	"github.com/adshao/go-binance/v2/common"
)

type BrokerGetDepositHistory struct {
	c            *Client
	subAccountId *string //하위계좌ID
	coin         *string
	status       *int //0(0:pending,6: credited but cannot withdraw, 1:success)
	startTime    *int64
	endTime      *int64
	limit        *int //Default：500
	offest       *int //Default：0
}

type BrokerGetDepositHistoryItem struct {
	SubaccountId  string `json:"subaccountId"`
	Address       string `json:"address"`
	AddressTag    string `json:"addressTag"`
	Amount        string `json:"amount"`
	Coin          string `json:"coin"`
	InsertTime    int64  `json:"insertTime"`
	Network       string `json:"network"`
	Status        int    `json:"status"`
	TxId          string `json:"txId"`
	SourceAddress string `json:"sourceAddress"`
	ConfirmTimes  string `json:"confirmTimes"`
}

func (s *BrokerGetDepositHistory) SubAccountId(subAccountId string) *BrokerGetDepositHistory {
	s.subAccountId = &subAccountId
	return s
}

func (s *BrokerGetDepositHistory) Coin(coin string) *BrokerGetDepositHistory {
	s.coin = &coin
	return s
}
func (s *BrokerGetDepositHistory) Status(status int) *BrokerGetDepositHistory {
	s.status = &status
	return s
}

func (s *BrokerGetDepositHistory) StartTime(startTime int64) *BrokerGetDepositHistory {
	s.startTime = &startTime
	return s
}
func (s *BrokerGetDepositHistory) EndTime(endTime int64) *BrokerGetDepositHistory {
	s.endTime = &endTime
	return s
}
func (s *BrokerGetDepositHistory) Limit(limit int) *BrokerGetDepositHistory {
	s.limit = &limit
	return s
}
func (s *BrokerGetDepositHistory) Offest(offest int) *BrokerGetDepositHistory {
	s.offest = &offest
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerGetDepositHistory) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, reqStatus int, err error) {

	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	if s.subAccountId != nil {
		r.setParam("subAccountId", *s.subAccountId)
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
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
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.offest != nil {
		r.setParam("offest", *s.offest)
	}

	data, reqStatus, err = s.c.callAPI_Core(ctx, r, opts...)
	if err != nil {
		return []byte{}, reqStatus, err
	}
	return data, reqStatus, nil
}

// Do sends the request.
func (s *BrokerGetDepositHistory) Do_Svr(ctx context.Context, opts ...RequestOption) (res []BrokerGetDepositHistoryItem, reqErrCode *common.APIError, reqStatus int, err error) {
	data, status, err := s.createReqData(ctx, "/sapi/v1/broker/subAccount/depositHist", opts...)
	if err != nil {
		return nil, nil, status, err
	}
	if status >= 400 && status < 500 {
		retErr := new(common.APIError)
		json.Unmarshal(data, &retErr)
		return nil, retErr, status, nil
	}

	var tmp []BrokerGetDepositHistoryItem
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, nil, status, err
	}
	return tmp, nil, status, nil
}
