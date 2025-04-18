package util

const (
	USD = "USD"
	EUR = "EUR"
)

// IsSupportedCurrency retorna verdadero si el tipo de moneda es soportado
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR:
		return true
	}
	return false
}
