package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// CoinResponse represents the structure of the API response
// We only map the fields we need (current_price)
type CoinResponse struct {
	ID           string  `json:"id"`
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	CurrentPrice float64 `json:"current_price"`
	High24H      float64 `json:"high_24h"`
	Low24H       float64 `json:"low_24h"`
}

func main() {
	// 1. Define command-line flags
	// usage: ./go-crypto-cli -coin=bitcoin -currency=usd
	coinName := flag.String("coin", "bitcoin", "The cryptocurrency ID (e.g., bitcoin, ethereum, ripple)")
	currency := flag.String("currency", "usd", "The target currency (e.g., usd, eur, try)")
	flag.Parse()

	// 2. Build the API URL (Using CoinGecko Public API)
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/markets?vs_currency=%s&ids=%s", *currency, *coinName)

	// 3. Make the HTTP Request
	fmt.Printf("⏳ Fetching price for %s in %s...\n", *coinName, *currency)

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("❌ Error fetching data: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// 4. Read and Parse the Response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Error reading response: %v\n", err)
		os.Exit(1)
	}

	var coins []CoinResponse
	if err := json.Unmarshal(body, &coins); err != nil {
		fmt.Printf("❌ Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// 5. Display Result
	if len(coins) > 0 {
		c := coins[0]
		fmt.Println("---------------------------------")
		fmt.Printf("💰 Coin:         %s (%s)\n", c.Name, strings.ToUpper(c.Symbol))
		fmt.Printf("💲 Current Price: %.2f %s\n", c.CurrentPrice, strings.ToUpper(*currency))
		fmt.Printf("📈 24h High:      %.2f\n", c.High24H)
		fmt.Printf("📉 24h Low:       %.2f\n", c.Low24H)
		fmt.Println("---------------------------------")
	} else {
		fmt.Println("⚠️ Coin not found. Please check the coin ID (e.g., 'bitcoin' not 'BTC').")
	}
}
