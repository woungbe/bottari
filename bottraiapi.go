package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	binance "woungbe.bottrai/modules/v2"
)

func main() {

	er := godotenv.Load(".env")
	if er != nil {
		log.Fatal("Error loading .env file")
	}

	// 환경 변수 읽기
	apiKey := os.Getenv("BinanaceAccesskey")
	secretKey := os.Getenv("BinanaceSecretKey")
	fmt.Println("BinanaceAccesskey :", apiKey)
	fmt.Println("BinanaceSecretKey :", secretKey)

	client := binance.NewClient(apiKey, secretKey)
	// futuresClient := binance.NewFuturesClient(apiKey, secretKey)   // USDT-M Futures
	// deliveryClient := binance.NewDeliveryClient(apiKey, secretKey) // Coin-M Futures
	//client := binance.NewClient(apiKey, secretKey)

	// res, err := client.NewGetAssetDetailService().Asset("ETHUSDT").Do(context.Background())
	// if err != nil {
	// 	fmt.Println("err : ", err)
	// }
	// fmt.Println("res : ", res)

	retData, st, er := client.NewGetAllWalletCoinInfoService().Do_Core(context.Background())
	fmt.Println("retData : ", st)
	fmt.Println("retData : ", string(retData))

	retData

}
