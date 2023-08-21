package binance

import (
	"context"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// API 요청에대한 RAte

func (c *Client) callAPI_Core(ctx context.Context, r *request, opts ...RequestOption) (data []byte, reqStatus int, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, 500, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, 500, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header

	//c.debug("request: %#v", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, 500, err
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, 500, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the retured error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	//c.debug("response: %#v", res)
	//c.debug("response body: %s", string(data))
	//c.debug("response status code: %d", res.StatusCode)
	UW := res.Header.Get("X-MBX-USED-WEIGHT-1M")
	USAPIW := res.Header.Get("X-SAPI-USED-IP-WEIGHT-1M")
	U1mOW := res.Header.Get("X-MBX-ORDER-COUNT-1M")
	U10sOW := res.Header.Get("X-MBX-ORDER-COUNT-10S")
	dt := res.Header.Get("Date")
	tm, _ := time.Parse("Mon, 02 Jan 2006 15:04:05 MST", dt)
	//fmt.Println("X-MBX-USED-WEIGHT-1M", UW, tm.UTC().Unix(), dt)
	//fmt.Println(res.Header)
	c.ReqCount.ReqTime = tm.UTC().Unix()

	if UW != "" {
		c.ReqCount.ReqCount, _ = strconv.ParseInt(UW, 10, 64)
	}
	if U1mOW != "" {
		c.ReqCount.Order1MCount, _ = strconv.ParseInt(U1mOW, 10, 64)
	}
	if U10sOW != "" {
		c.ReqCount.Order10SCount, _ = strconv.ParseInt(U10sOW, 10, 64)
	}
	if USAPIW != "" {
		tmp, _ := strconv.ParseInt(USAPIW, 10, 64)
		if c.ReqCount.ReqCount < tmp {
			c.ReqCount.ReqCount = tmp
		}
	}
	return data, res.StatusCode, nil

	/*
		if res.StatusCode >= 400 {
			apiErr := new(common.APIError)
			e := json.Unmarshal(data, apiErr)
			if e != nil {
				c.debug("failed to unmarshal json: %s", e)
			}
			return nil, apiErr
		}
		return data, nil
	*/
}

//--------------------------------

// NewGetSAPIDepositsAddressService  입금주소 정보 조회
func (c *Client) NewGetSAPIDepositsAddressService() *GetSAPIDepositsAddressService {
	return &GetSAPIDepositsAddressService{c: c}
}

//GetAllWalletCoinInfoService 코인 지갑정보 전체 가져오기
func (c *Client) NewGetAllWalletCoinInfoService() *GetAllWalletCoinInfoService {
	return &GetAllWalletCoinInfoService{c: c}
}

//GetSAPIDepositsHistory 입금내역 리스트
func (c *Client) NewGetSAPIDepositsHistory() *GetSAPIDepositsHistory {
	return &GetSAPIDepositsHistory{c: c}
}

//GetSAPIWithdrawHistory 출금 내역 리스트
func (c *Client) NewGetSAPIWithdrawHistory() *GetSAPIWithdrawHistory {
	return &GetSAPIWithdrawHistory{c: c}
}

//CreateSAPIWithdrawService 출금
func (c *Client) NewCreateSAPIWithdrawService() *CreateSAPIWithdrawService {
	return &CreateSAPIWithdrawService{c: c}
}

//GetMyTradeList
func (c *Client) NewGetMyTradeList() *GetMyTradeList {
	return &GetMyTradeList{c: c}
}
