# 📈 Go Crypto CLI

A lightweight, fast, and terminal-based cryptocurrency price tracker built with **Golang**. This tool fetches real-time data from the CoinGecko API without requiring any external dependencies or API keys.

## 🚀 Features

* **Real-Time Data:** Fetches current price, 24h high, and 24h low statistics.
* **Zero Dependencies:** Built entirely using Go's Standard Library (`net/http`, `encoding/json`, `flag`).
* **Customizable:** Support for any cryptocurrency (Bitcoin, Ethereum, Ripple, etc.) and fiat currency (USD, EUR, TRY).
* **Robust Error Handling:** Includes timeouts and graceful error messages.

## 🛠 Usage

### 1. Default (Bitcoin / USD)
Simply run the program to get the latest Bitcoin price in USD.
```bash
go run main.go
```

### 2. Custom Coin & Currency
Use flags to specify the coin and target currency.

**Example: Ethereum in Turkish Lira**
```bash
go run main.go -coin=ethereum -currency=try
```

**Example: Ripple (XRP) in Euro**
```bash
go run main.go -coin=ripple -currency=eur
```

## 🔧 Technical Highlights

This project demonstrates the following Golang concepts:
* **CLI Flags:** Parsing command-line arguments.
* **HTTP Client:** Making external API requests with custom timeouts.
* **JSON Unmarshalling:** Parsing complex JSON responses into Go Structs.
* **Struct Tags:** Mapping JSON keys to specific struct fields.

## 📦 Installation

```bash
# Clone the repository
git clone https://github.com/aliemreipek/go-crypto-cli.git
cd go-crypto-cli

# Run
go run main.go
```

---
*Developed by [Ali Emre Ipek](https://github.com/aliemreipek)*
