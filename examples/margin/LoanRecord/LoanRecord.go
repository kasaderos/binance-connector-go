package main

import (
	"context"
	"fmt"

	binance_connector "github.com/kasaderos/binance-connector-go"
)

func main() {
	LoanRecord()
}

func LoanRecord() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// LoanRecordService - /sapi/v1/margin/loan
	loanRecord, err := client.NewLoanRecordService().Asset("BTC").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(loanRecord))
}
