package authgrpc

import (
	"context"
	authpb "github.com/repooooo/protos/gen/go/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ValidatorType is a custom type for mapping validator keys.
type ValidatorType string

const (
	// ValidatorTypeLogin is the key for Login request validation.
	ValidatorTypeLogin ValidatorType = "Login"
	// ValidatorTypeLogout is the key for Logout request validation.
	ValidatorTypeLogout ValidatorType = "Logout"
)

// Validator defines the interface for request validation.
type Validator interface {
	Validate(ctx context.Context, req interface{}) error
}

// ValidatorMap is a wrapper for a map of validators.
type ValidatorMap map[ValidatorType]Validator

// NewValidatorMap creates and returns a new ValidatorMap.
func NewValidatorMap() ValidatorMap {
	return ValidatorMap{
		ValidatorTypeLogin:  &LoginValidator{},
		ValidatorTypeLogout: &LogoutValidator{},
	}
}

// LoginValidator validates LoginRequest.
type LoginValidator struct{}

func (v *LoginValidator) Validate(ctx context.Context, req interface{}) error {
	loginReq, ok := req.(*authpb.LoginRequest)
	if !ok {
		return status.Error(codes.InvalidArgument, "invalid request type for login")
	}

	if loginReq.GetUsername() == "" {
		return status.Error(codes.InvalidArgument, "username is required")
	}

	if loginReq.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}

	return nil
}

// LogoutValidator validates LogoutRequest.
type LogoutValidator struct{}

func (v *LogoutValidator) Validate(ctx context.Context, req interface{}) error {
	logoutReq, ok := req.(*authpb.LogoutRequest)
	if !ok {
		return status.Error(codes.InvalidArgument, "invalid request type for logout")
	}

	if logoutReq.GetToken() == "" {
		return status.Error(codes.InvalidArgument, "token is required")
	}

	return nil
}
