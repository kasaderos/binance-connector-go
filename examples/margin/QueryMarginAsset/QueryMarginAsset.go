package main

import (
	"context"
	"fmt"

	binance_connector "github.com/kasaderos/binance-connector-go"
)

func main() {
	QueryMarginAsset()
}

func QueryMarginAsset() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// QueryMarginAssetService - /sapi/v1/margin/asset
	queryMarginAsset, err := client.NewQueryMarginAssetService().Asset("USDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(queryMarginAsset))
}
