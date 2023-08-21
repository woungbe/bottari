package binance

import (
	"context"
	"encoding/json"
)

/**


 */

type BrokerAccountInfo struct {
	c *Client
}

type BrokerAccountInfoRESP struct {
	MaxMakerCommission string `json:"maxMakerCommission"`
	MinMakerCommission string `json:"minMakerCommission"`
	MaxTakerCommission string `json:"maxTakerCommission"`
	MinTakerCommission string `json:"minTakerCommission"`
	SubAccountQty      int64  `json:"subAccountQty"`
	MaxSubAccountQty   int64  `json:"maxSubAccountQty"`
}

//createReqData 전송할 데이터 생성
func (s *BrokerAccountInfo) createReqData(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {

	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do sends the request.
func (s *BrokerAccountInfo) Do(ctx context.Context, opts ...RequestOption) (res *BrokerAccountInfoRESP, err error) {

	data, err := s.createReqData(ctx, "/sapi/v1/broker/info", opts...)
	if err != nil {
		return nil, err
	}
	res = new(BrokerAccountInfoRESP)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
