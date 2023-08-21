package binance

import (
	"context"
	"encoding/json"

	"github.com/adshao/go-binance/v2/common"
)

/*
	sub account trans
*/

type BrokerTransSpot struct {
	c      *Client
	asset  string //코인명
	amount string //수량
	fromID *string
	toID   *string
}

type BrokerTransSpotREQ struct {
	TxnId        int64  `json:"txnId"`
	ClientTranId string `json:"clientTranId"`
}

func (s *BrokerTransSpot) Asset(assetid string) *BrokerTransSpot {
	s.asset = assetid
	return s
}

func (s *BrokerTransSpot) Amount(amount string) *BrokerTransSpot {
	s.amount = amount
	return s
}

func (s *BrokerTransSpot) FromID(fromID string) *BrokerTransSpot {
	s.fromID = &fromID
	return s
}

func (s *BrokerTransSpot) ToID(toID string) *BrokerTransSpot {
	s.toID = &toID
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerTransSpot) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, reqStatus int, err error) {

	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"asset":  s.asset,
		"amount": s.amount,
	}

	if s.fromID != nil {
		m["fromId"] = *s.fromID
	}
	if s.toID != nil {
		m["toId"] = *s.toID
	}

	r.setFormParams(m)
	data, reqStatus, err = s.c.callAPI_Core(ctx, r, opts...)
	if err != nil {
		return []byte{}, reqStatus, err
	}
	return data, reqStatus, nil
}

// Do sends the request.
func (s *BrokerTransSpot) Do_Svr(ctx context.Context, opts ...RequestOption) (res *BrokerTransSpotREQ, reqErrCode *common.APIError, reqStatus int, err error) {
	data, status, err := s.createReqData(ctx, "/sapi/v1/broker/transfer", opts...)
	if err != nil {
		return nil, nil, status, err
	}

	if status >= 400 {
		apiErr := new(common.APIError)
		json.Unmarshal(data, apiErr)
		return nil, apiErr, status, err
	}

	//fmt.Println(string(data))
	res = new(BrokerTransSpotREQ)
	err = json.Unmarshal(data, res)

	if err != nil {
		return nil, nil, status, err
	}

	return res, nil, status, nil
}

/*
	내부 전송 히스토리 조회
*/
type BrokerTransSpotHistory struct {
	c             *Client
	fromId        *string
	toId          *string
	clientTranId  *string
	showAllStatus *string //INIT,PROCESS,SUCCESS,FAILURE.
	startTime     *int64
	endTime       *int64
	page          *int64
	pagesize      *int64 //default 500, max 500

	/*
		fromId	STRING	NO
		toId	STRING	NO
		clientTranId	STRING	NO	client transfer id
		showAllStatus	ENUM	NO	true or false, default: false
		startTime	LONG	NO
		endTime	LONG	NO
		page	INT	NO
		limit	INT	NO	default 500, max 500
	*/
}

type BrokerTransSpotHistoryRESP struct {
	FromId       string `json:"fromId"`
	ToId         string `json:"toId"`
	Asset        string `json:"asset"`
	Qty          string `json:"qty"`
	TransTime    int64  `json:"time"`
	TxnId        string `json:"txnId"`
	ClientTranId string `json:"clientTranId"`
	Status       string `json:"status"`
}

func (s *BrokerTransSpotHistory) FromID(fromID string) *BrokerTransSpotHistory {
	s.fromId = &fromID
	return s
}

func (s *BrokerTransSpotHistory) ToID(toId string) *BrokerTransSpotHistory {
	s.toId = &toId
	return s
}

func (s *BrokerTransSpotHistory) ClientTranId(clientTranId string) *BrokerTransSpotHistory {
	s.clientTranId = &clientTranId
	return s
}

func (s *BrokerTransSpotHistory) ShowAllStatus(showAllStatus string) *BrokerTransSpotHistory {
	s.showAllStatus = &showAllStatus
	return s
}

func (s *BrokerTransSpotHistory) StartTime(startTime int64) *BrokerTransSpotHistory {
	s.startTime = &startTime
	return s
}

func (s *BrokerTransSpotHistory) EndTime(endTime int64) *BrokerTransSpotHistory {
	s.endTime = &endTime
	return s
}

func (s *BrokerTransSpotHistory) Page(page int64) *BrokerTransSpotHistory {
	s.page = &page
	return s
}

func (s *BrokerTransSpotHistory) Pagesize(pagesize int64) *BrokerTransSpotHistory {
	s.pagesize = &pagesize
	return s
}

//createReqData 전송할 데이터 생성
func (s *BrokerTransSpotHistory) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	/*
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
	*/

	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.pagesize != nil {
		r.setParam("limit", *s.pagesize)
	}

	if s.fromId != nil {
		r.setParam("fromId", *s.fromId)
	}

	if s.toId != nil {
		r.setParam("toId", *s.toId)
	}

	if s.clientTranId != nil {
		r.setParam("clientTranId", *s.clientTranId)
	}
	if s.showAllStatus != nil {
		r.setParam("showAllStatus", *s.showAllStatus)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	data, reqStatus, err = s.c.callAPI_Core(ctx, r, opts...)
	if err != nil {
		return []byte{}, reqStatus, err
	}
	return data, reqStatus, nil
}

func (s *BrokerTransSpotHistory) Do_Svr(ctx context.Context, opts ...RequestOption) (res []BrokerTransSpotHistoryRESP, reqErrCode *common.APIError, reqStatus int, err error) {
	data, status, err := s.createReqData(ctx, "/sapi/v1/broker/transfer", opts...)
	if err != nil {
		return nil, nil, status, err
	}

	if status >= 400 {
		apiErr := new(common.APIError)
		json.Unmarshal(data, apiErr)
		return nil, apiErr, status, err
	}
	var req []BrokerTransSpotHistoryRESP
	err = json.Unmarshal(data, &req)

	if err != nil {
		return nil, nil, status, err
	}

	return req, nil, status, nil
}
