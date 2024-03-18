package utils

const (
	USD = "USD"
	KRW = "KRW"
	EUR = "EUR"	
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, KRW, EUR:
		return true
	}
	return false
}