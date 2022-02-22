package util

// Consraints for all supported currency
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// IsSupportedCurency returns true if the currency is supported
func IsSupportedCurency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
