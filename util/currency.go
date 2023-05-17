package util

const (
	USD = "USD"
	EUR = "EUR"
	RON = "RON"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, RON:
		return true
	}
	return false
}
