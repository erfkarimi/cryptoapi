package main 

import "github.com/gin-gonic/gin"
import "net/http"

type coin struct{
	ID			string `json:"id"`
	Name 		string `json:"name"`
	Symbol		string `json:"symbol"`	
	Price		string `json:"price"`
}

var coins = []coin{
	{
		ID:  "1", 
		Name: "Bitcoin",
		Symbol: "BTC",
		Price: "$16.633",
	},
	
	{	
		ID:  "2", 
		Name: "Ethereum",
		Symbol: "ETH",
	 	Price: "$1.196",
	},
	{
		ID:  "3",
		Name: "BNB",
		Symbol: "BNB",
		Price: "$243",
	},
	{
		ID:  "4",
		Name: "XRP",
		Symbol: "XRP",
		Price: "$0.35",
	},
	{
		ID:  "5",
		Name: "Cardano",
		Symbol: "ADA",
		Price: "$0.25",
	},
	{
		ID:  "6",
		Name: "Cosmos Hub",
		Symbol: "ATOM",
		Price: "$9.12",
	},
	{
		ID:  "7",
		Name: "Basic attention",
		Symbol: "BAT",
		Price: "$0.17",
	},
	{
		ID:  "8",
		Name: "Dogecoin",
		Symbol: "DOGE",
		Price: "$0.07",
	},
	{
		ID:  "9",
		Name: "Tether",
		Symbol: "USDT",
		Price: "$1.0",
	},
	{
		ID:  "10",
		Name: "Tron",
		Symbol: "TRX",
		Price: "$0.053",
	},
	
}

func getCoins(context *gin.Context){
	context.IndentedJSON(http.StatusOK, coins)
}

func main(){
	router := gin.Default()
	router.GET("/coins", getCoins)
	router.Run("127.0.0.1:9000")

}