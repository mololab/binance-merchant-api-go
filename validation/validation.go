package validation

import "fmt"

func ValidBinanceSecretKey(binanceSecretKey string) error {

	if binanceSecretKey == "" {
		return fmt.Errorf("binance secret key: not valid")
	}

	return nil
}

func ValidBinanceApiKey(binanceApiKey string) error {

	if binanceApiKey == "" {
		return fmt.Errorf("binance api key: not valid")
	}

	return nil
}
