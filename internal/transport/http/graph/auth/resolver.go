package authgraph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	auth Auth
}

func NewResolver(auth Auth) *Resolver {
	return &Resolver{
		auth: auth,
	}
}
