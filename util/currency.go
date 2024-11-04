package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

func IsSupportedCurrency(currency string) bool {
	supportedCurrencies := map[string]bool{
		USD: true,
		EUR: true,
		CAD: true,
	}

	return supportedCurrencies[currency]
}
