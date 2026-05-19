package onboarding

type Request struct {
	AccountCredentials `json:"account_credentials"`
	CustomerInfo       `json:"customer_info"`
	DeviceInfo         `json:"device_info"`
}

type AccountCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CustomerInfo struct {
	Name     string       `json:"name"`
	Document string       `json:"document"`
	Type     CustomerType `json:"customer_type"`
}

type DeviceInfo struct {
	IPAddr    string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
	DeviceID  string `json:"device_id"`
}

type CustomerType string

const (
	INDIVIDUAL CustomerType = "INDIVIDUAL"
	BUSINESS   CustomerType = "BUSINESS"
	SYSTEM     CustomerType = "SYSTEM"
)
