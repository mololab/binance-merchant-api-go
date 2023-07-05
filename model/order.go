package model

import "github.com/mololab/binance-merchant-api-go/constant"

type CreateOrderParams struct {
	Merchant        *Merchant          `json:"merchant,omitempty"`
	Env             Env                `json:"env"`             // required
	MerchantTradeNo string             `json:"merchantTradeNo"` // required
	OrderAmount     *float64           `json:"orderAmount,omitempty"`
	Currency        *constant.Currency `json:"currency,omitempty"` // Currency and fiatCurrency cannot be both null
	FiatAmount      *float64           `json:"fiatAmount,omitempty"`
	FiatCurrency    *string            `json:"fiatCurrency,omitempty"` // Currency and fiatCurrency cannot be both null
	Goods           Goods              `json:"goods"`                  // required
	Shipping        *Shipping          `json:"shipping,omitempty"`
}

type Merchant struct {
	SubMerchantId string `json:"subMerchantId"`
}

type Env struct {
	TerminalType  constant.TerminalType `json:"terminalType"` // required
	OSType        *constant.OSType      `json:"osType,omitempty"`
	OrderClientIp *string               `json:"orderClientIp,omitempty"`
	CookieId      *string               `json:"cookieId,omitempty"`
}

type Goods struct {
	GoodsType        constant.GoodsType `json:"goodsType"`        // required
	GoodsCategory    string             `json:"goodsCategory"`    // required
	ReferenceGoodsId string             `json:"referenceGoodsId"` // required
	GoodsName        string             `json:"goodsName"`        // required
	GoodsDetail      *string            `json:"goodsDetail,omitempty"`
	GoodsUnitAmount  *GoodsUnitAmount   `json:"goodsUnitAmount,omitempty"`
	GoodsQuantity    *string            `json:"goodsQuantity,omitempty"`
}

type GoodsUnitAmount struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

type Shipping struct{} // TODO: complete
