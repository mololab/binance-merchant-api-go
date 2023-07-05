package binanceMerchant

func Opt[T any](t T) *T {
	return &t
}
