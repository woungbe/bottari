package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sangx2/upbit"
)

func main() {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 환경 변수 읽기
	AccessKey := os.Getenv("AccessKey")
	SecretKey := os.Getenv("SecretKey")

	u := upbit.NewUpbit(AccessKey, SecretKey)

	/*
		// 어카운트 가져오는 func
		accounts, remaining, e := u.GetAccounts()
		if e != nil {
			fmt.Println("GetAccounts error : %s", e.Error())
		} else {
			fmt.Printf("GetAccounts[remaining:%+v]\n", *remaining)
			for _, account := range accounts {
				fmt.Printf("%+v\n", *account)
			}
		}
	*/

	// 마켓 정보 리스트

	// markets, remaining, e := u.GetMarkets()
	// if e != nil {
	// 	fmt.Println("GetMarkets error : %s", e.Error())
	// } else {
	// 	fmt.Printf("GetMarkets[remaining:%+v]\n", *remaining)
	// 	for _, market := range markets {
	// 		fmt.Printf("%+v\n", *market)
	// 	}
	// }

	// 현재가 가져오기

	// coinlist := []string{"KRW-BTC", "KRW-ETH"}
	// ticker, remaining, e := u.GetTickers(coinlist)
	// if e != nil {
	// 	fmt.Println("GetTickers error : %s", e.Error())
	// } else {
	// 	fmt.Printf("GetTickers[remaining:%+v]\n", *remaining)
	// 	for _, m_ticker := range ticker {
	// 		fmt.Printf("%+v\n", *m_ticker)
	// 	}
	// }

	// => reponse
	// {Market:KRW-BTC TradeDate:20230821 TradeTime:064935 TradeDateKst:20230821 TradeTimeKst:154935 OpeningPrice:3.5946e+07 HighPrice:3.5967e+07 LowPrice:3.563e+07 TradePrice:3.5697e+07 PrevClosingPrice:3.5949e+07 Change:FALL ChangePrice:252000 ChangeRate:0.0070099307 SignedChangePrice:-252000 SignedChangeRate:-0.0070099307 TradeVolume:0.05017214 AccTradePrice:2.839714914420053e+10 AccTradePrice24h:6.889685402997546e+10 AccTradeVolume:794.00595579 AccTradeVolume24h:1920.57723089 Highest52WeekPrice:4.1569e+07 Highest52WeekDate:2023-06-30 Lowest52WeekPrice:2.07e+07 Lowest52WeekDate:2022-12-30 Timestamp:1692600575504}
	// {Market:KRW-ETH TradeDate:20230821 TradeTime:064934 TradeDateKst:20230821 TradeTimeKst:154934 OpeningPrice:2.311e+06 HighPrice:2.314e+06 LowPrice:2.288e+06 TradePrice:2.289e+06 PrevClosingPrice:2.312e+06 Change:FALL ChangePrice:23000 ChangeRate:0.0099480969 SignedChangePrice:-23000 SignedChangeRate:-0.0099480969 TradeVolume:0.43665399 AccTradePrice:4.39582891054803e+09 AccTradePrice24h:1.387685038537933e+10 AccTradeVolume:1910.78155822 AccTradeVolume24h:6028.19712899 Highest52WeekPrice:2.795e+06 Highest52WeekDate:2023-04-16 Lowest52WeekPrice:1.5e+06 Lowest52WeekDate:2022-12-30 Timestamp:1692600574884}

	// type Ticker struct {
	// 	Market             string  `json:"market"`
	// 	TradeDate          string  `json:"trade_date"`
	// 	TradeTime          string  `json:"trade_time"`
	// 	TradeDateKst       string  `json:"trade_date_kst"`
	// 	TradeTimeKst       string  `json:"trade_time_kst"`
	// 	OpeningPrice       float64 `json:"opening_price"`
	// 	HighPrice          float64 `json:"high_price"`
	// 	LowPrice           float64 `json:"low_price"`
	// 	TradePrice         float64 `json:"trade_price"`
	// 	PrevClosingPrice   float64 `json:"prev_closing_price"`
	// 	Change             string  `json:"change"`
	// 	ChangePrice        float64 `json:"change_price"`
	// 	ChangeRate         float64 `json:"change_rate"`
	// 	SignedChangePrice  float64 `json:"signed_change_price"`
	// 	SignedChangeRate   float64 `json:"signed_change_rate"`
	// 	TradeVolume        float64 `json:"trade_volume"`
	// 	AccTradePrice      float64 `json:"acc_trade_price"`
	// 	AccTradePrice24h   float64 `json:"acc_trade_price_24h"`
	// 	AccTradeVolume     float64 `json:"acc_trade_volume"`
	// 	AccTradeVolume24h  float64 `json:"acc_trade_volume_24h"`
	// 	Highest52WeekPrice float64 `json:"highest_52_week_price"`
	// 	Highest52WeekDate  string  `json:"highest_52_week_date"`
	// 	Lowest52WeekPrice  float64 `json:"lowest_52_week_price"`
	// 	Lowest52WeekDate   string  `json:"lowest_52_week_date"`
	// 	Timestamp          int64   `json:"timestamp"`
	// }

	// 네트워크 확인
	// https://api.upbit.com/v1/status/wallet

	// ([]*service.Wallet, *model.Remaining, error)

	// u.GetTickers(coinlist)
	wallets, remaining, e := u.GetWalletStatus()
	if e != nil {
		fmt.Println("GetWalletStatus error : %s", e.Error())
	} else {
		fmt.Println("GetTikers[reamining:%+v]\n", *remaining)
		for _, wallet := range wallets {
			val, err := PrintJSON(*wallet)
			if err != nil {
				fmt.Println("err : ", err)
			}
			// fmt.Println("%+v\n", *wallet)
			fmt.Println(string(val))
		}
	}

	/*
		type Wallet struct {
			Currency    string `json:"currency"`
			WalletState string `json:"wallet_state"`

			BlockState     string `json:"block_state"`
			BlockHeight    int64  `json:"block_height"`
			BlockUpdatedAt string `json:"block_updated_at"`
		}
	*/
}

func PrintJSON(wallet interface{}) ([]byte, error) {
	byteJSON, err := json.Marshal(wallet)
	// byteJSON, err := json.MarshalIndent(wallet, "", "	")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return byteJSON, err
}
