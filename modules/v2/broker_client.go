package binance

/*
	브로커API 추가 함 Add by Lee H.K ( 2021-0419)
*/

//NewBrokerCreateSubAccount 브로커API subAccount 생성
func (c *Client) NewBrokerCreateSubAccount() *BrokerCreateSubAccount {
	return &BrokerCreateSubAccount{c: c}
}

//NewBrokerGetSubAccountInfo SubAccount 정보 가져오기
func (c *Client) NewBrokerGetSubAccountInfo() *BrokerGetSubAccountInfo {
	return &BrokerGetSubAccountInfo{c: c}
}

//NewBrokerEnableMargin 마진 사용설정
func (c *Client) NewBrokerEnableMargin() *BrokerEnableMargin {
	return &BrokerEnableMargin{c: c}
}

//NewBrokerEnableFutures 선물 사용설정
func (c *Client) NewBrokerEnableFutures() *BrokerEnableFutures {
	return &BrokerEnableFutures{c: c}
}

//NewBrokerCreateAPIKey  API Key  생성
func (c *Client) NewBrokerCreateAPIKey() *BrokerCreateAPIKey {
	return &BrokerCreateAPIKey{c: c}
}

//NewBrokerDeleteAPIKey API Key 삭제
func (c *Client) NewBrokerDeleteAPIKey() *BrokerDeleteAPIKey {
	return &BrokerDeleteAPIKey{c: c}
}

//NewBrokerGetSubAPIKey Api Key 가져오기
func (c *Client) NewBrokerGetSubAPIKey() *BrokerGetSubAPIKey {
	return &BrokerGetSubAPIKey{c: c}
}

//NewBrokerChangeSubAPIPermission API 권한 변경
func (c *Client) NewBrokerChangeSubAPIPermission() *BrokerChangeSubAPIPermission {
	return &BrokerChangeSubAPIPermission{c: c}
}

//NewBrokerChangeSubAccountCommission 기본거래 수수료 조정
func (c *Client) NewBrokerChangeSubAccountCommission() *BrokerChangeSubAccountCommission {
	return &BrokerChangeSubAccountCommission{c: c}
}

//NewBrokerChangeUSDTfuturesCommission 선물 수수료 조정 ( 추가됌 )
func (c *Client) NewBrokerChangeUSDTfuturesCommission() *BrokerChangeUSDTfuturesCommission {
	return &BrokerChangeUSDTfuturesCommission{c: c}
}

//NewBrokerGetUSDTfuturesCommission 선물수수료 정보 조회
func (c *Client) NewBrokerGetUSDTfuturesCommission() *BrokerGetUSDTfuturesCommission {
	return &BrokerGetUSDTfuturesCommission{c: c}
}

//NewBrokerBrokerTransSpot 내부 코인 전송
func (c *Client) NewBrokerBrokerTransSpot() *BrokerTransSpot {
	return &BrokerTransSpot{c: c}
}

//NewBrokerGetFutuRecentRecord 선물 리베이트 히스토리 조회
func (c *Client) NewBrokerGetFutuRecentRecord() *BrokerGetFutuRecentRecord {
	return &BrokerGetFutuRecentRecord{c: c}
}

//NewBrokerGetSpotRecentRecord 현물 리베이트 히스토리 조회
func (c *Client) NewBrokerGetSpotRecentRecord() *BrokerGetSpotRecentRecord {
	return &BrokerGetSpotRecentRecord{c: c}
}

//NewBrokerSubAccountSpotSummary 현물 계좌 잔고 정보
func (c *Client) NewBrokerSubAccountSpotSummary() *BrokerSubAccountSpotSummary {
	return &BrokerSubAccountSpotSummary{c: c}
}

//NewBrokerSubAccountFutureSummary 선물 계좌 잔고 정보
func (c *Client) NewBrokerSubAccountFutureSummary() *BrokerSubAccountFutureSummary {
	return &BrokerSubAccountFutureSummary{c: c}
}
func (c *Client) NewBrokerSubAccountFutureSummaryV2() *BrokerSubAccountFutureSummaryV2 {
	return &BrokerSubAccountFutureSummaryV2{c: c}
}

//선물 전송
func (c *Client) NewBrokerTransFutu() *BrokerTransFutu {
	return &BrokerTransFutu{c: c}
}

//유저 선물 지갑정보 조회
func (c *Client) NewBrokerGetFutureAssetInfo() *BrokerGetFutureAssetInfo {
	return &BrokerGetFutureAssetInfo{c: c}
}

//유저 입금내역 조회
func (c *Client) NewBrokerGetDepositHistory() *BrokerGetDepositHistory {
	return &BrokerGetDepositHistory{c: c}
}

//내부 전송 내역 조회
func (c *Client) NewBrokerTransSpotHistory() *BrokerTransSpotHistory {
	return &BrokerTransSpotHistory{c: c}
}

//마스터 계좌의 서브 계정 정보 조회
func (c *Client) NewBrokerAccountInfo() *BrokerAccountInfo {
	return &BrokerAccountInfo{c: c}
}
