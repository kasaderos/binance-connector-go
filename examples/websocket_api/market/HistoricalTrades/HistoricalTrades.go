package main

import (
	"context"
	"fmt"
	"log"

	binance_connector "github.com/kasaderos/binance-connector-go"
)

func main() {
	HistoricalTrades()
}

func HistoricalTrades() {
	client := binance_connector.NewWebsocketAPIClient("", "", "wss://testnet.binance.vision/ws-api/v3")
	err := client.Connect()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	defer client.Close()

	response, err := client.NewHistoricalTradesService().Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println(binance_connector.PrettyPrint(response))

	client.WaitForCloseSignal()
}
