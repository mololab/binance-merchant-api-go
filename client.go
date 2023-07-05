package binanceMerchant

import "github.com/mololab/binance-merchant-api-go/validation"

type Client struct {
	SecretKey string
	ApiKey    string
}

func New(binanceSecretKey string, binanceApiKey string) (client *Client, err error) {

	return createClient(binanceSecretKey, binanceApiKey)
}

func createClient(binanceSecretKey string, binanceApiKey string) (*Client, error) {

	if err := validation.ValidBinanceSecretKey(binanceSecretKey); err != nil {
		return nil, err
	}

	if err := validation.ValidBinanceApiKey(binanceApiKey); err != nil {
		return nil, err
	}

	return &Client{SecretKey: binanceSecretKey, ApiKey: binanceApiKey}, nil
}
