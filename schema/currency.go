package schema

type CurrencyError struct {
	ERROR []CurrencyCode
}

type CurrencyCode struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
