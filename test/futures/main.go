package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	bybit "github.com/wuhewuhe/bybit.go.api"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiKey    string
	ApiSecret string
	ApiName   string
}

func getAssestBalance(bybitClient *bybit.Client) {
	resp, err := bybitClient.NewAssetService(map[string]interface{}{
		"accountType": "UNIFIED",
		"coin":        "USDT",
	}).GetAllCoinsBalance(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("resp.code", resp.RetCode, resp.RetMsg, resp.RetExtInfo, resp.Result)
	dataJson, _ := json.Marshal(resp.Result)
	log.Println(string(dataJson))
}

// func walletBalance(bybitClient *bybit.Client) {
// 	resp, err := bybitClient.NewAccountService(map[string]interface{}{}).GetAccountInfo(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }

func transferAssetFromFundToContract(bybitClient *bybit.Client) {

	transferID := uuid.New().String()
	fmt.Println("transferID", transferID)
	resp, err := bybitClient.NewAssetService(map[string]interface{}{
		"transferId":      transferID,
		"coin":            "BTC",
		"amount":          1,
		"fromAccountType": "FUND",
		"toAccountType":   "UNIFIED",
	}).CreateUniversalTransfer(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	dataJson, _ := json.Marshal(resp)
	log.Println(string(dataJson))
}

func listOpenSpotOrders(bybitClient *bybit.Client) {

	// symbol := "BTCUSDT"
	resp, err := bybitClient.NewTradeService(map[string]interface{}{
		"category": "spot",
		"symbol":   "BTCUSDT",
	}).GetOpenOrders(context.Background())

	if err != nil {
		log.Fatal(err)
	}
	dataJson, _ := json.Marshal(resp.Result.(map[string]interface{})["list"])
	log.Println(string(dataJson))
}

func createSpotLimitOrder(bybitClient *bybit.Client) {
	category := "spot"
	symbol := "BTCUSDT"
	// isLeverage := "false"
	side := "Buy"
	orderType := "Limit"
	qty := "0.01"
	limitPrice := "60500"
	resp, err := bybitClient.NewPlaceOrderService(category, symbol, side, orderType, qty).Price(limitPrice).TimeInForce("GTC").Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	dataJson, _ := json.Marshal(resp)
	log.Println(string(dataJson))
}

func updateSpotLimitOrderWithTpSl(bybitClient *bybit.Client) {
	orderID := "1719636060331665664"
	orderLinkId := "1719636060331665665"
	// price := "60500"
	resp, err := bybitClient.NewTradeService(map[string]interface{}{
		"category":     "spot",
		"symbol":       "BTCUSDT",
		"orderId":      orderID,
		"orderLinkId":  orderLinkId,
		"triggerPrice": "60500",
		"takeProfit":   "62000",
		"stopLoss":     "59000",
		"price":        "60500",
	}).AmendOrder(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	dataJson, _ := json.Marshal(resp)
	log.Println(string(dataJson))
}
func cancelAllOpendOrder(bybitClient *bybit.Client) {
	resp, err := bybitClient.NewTradeService(map[string]interface{}{
		"category": "spot",
		"symbol":   "BTCUSDT",
	}).CancelAllOrders(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	dataJson, _ := json.Marshal(resp)
	log.Println(string(dataJson))

}

func setLeverage(bybitClient *bybit.Client) {
	resp, err := bybitClient.NewPositionService(map[string]interface{}{
		"category":     "linear",
		"symbol":       "BTCUSDT",
		"buyLeverage":  "6",
		"sellLeverage": "6",
	}).SetPositionLeverage(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	dataJson, _ := json.Marshal(resp)
	log.Println(string(dataJson))
}

func setLeverageAndIsolateMode(bybitClient *bybit.Client) {
	resp, err := bybitClient.NewPositionService(map[string]interface{}{
		"category":     "linear",
		"symbol":       "BTCUSDT",
		"tradeMode":    1,
		"buyLeverage":  "10",
		"sellLeverage": "10",
	}).SwitchPositionMargin(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	dataJson, _ := json.Marshal(resp)
	log.Println(string(dataJson))
}

func listPositions(bybitClient *bybit.Client) {
	resp, err := bybitClient.NewPositionService(map[string]interface{}{
		"category": "linear",
		"symbol":   "BTCUSDT",
	}).GetPositionList(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	dataJson, _ := json.Marshal(resp.Result)
	log.Println(string(dataJson))
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var conf Config = Config{
		ApiKey:    os.Getenv("BYBIT_API_KEY"),
		ApiSecret: os.Getenv("BYBIT_API_SECRECT"),
	}

	bybitClient := bybit.NewBybitHttpClient(conf.ApiKey, conf.ApiSecret, bybit.WithBaseURL(bybit.MAINNET))
	resp, err := bybitClient.NewGetRiskLimitService().Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	dataJson, _ := json.Marshal(resp)
	log.Println(string(dataJson))
	getAssestBalance(bybitClient)
	// setLeverage(bybitClient)
	// setLeverageAndIsolateMode(bybitClient)
	// transferAssetFromFundToContract(bybitClient)
	// createSpotLimitOrder(bybitClient)
	// updateSpotLimitOrderWithTpSl(bybitClient)
	// listOpenSpotOrders(bybitClient)
	listPositions(bybitClient)
}