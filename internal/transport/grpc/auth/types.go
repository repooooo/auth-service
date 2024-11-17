package authgrpc

// ValidatorType is a custom type for mapping validator keys.
type ValidatorType string

const (
	// ValidatorTypeLogin is the key for Login request validation.
	ValidatorTypeLogin ValidatorType = "Login"
	// ValidatorTypeLogout is the key for Logout request validation.
	ValidatorTypeLogout ValidatorType = "Logout"
)
