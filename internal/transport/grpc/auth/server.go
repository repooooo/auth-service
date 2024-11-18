package authgrpc

import (
	"context"
	authpb "github.com/repooooo/protos/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// serverAPI implements the AuthServiceServer interface.
type serverAPI struct {
	authpb.UnimplementedAuthServiceServer
	auth       Auth
	validators ValidatorMap
}

// newServerAPI creates a new instance of serverAPI.
func newServerAPI(auth Auth) *serverAPI {
	return &serverAPI{
		auth:       auth,
		validators: NewValidatorMap(),
	}
}

// Register registers the AuthServiceServer with gRPC.
func Register(gRPC *grpc.Server, auth Auth) {
	authpb.RegisterAuthServiceServer(gRPC, newServerAPI(auth))
}

func (s *serverAPI) Login(ctx context.Context, request *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	if err := s.validateRequest(ctx, ValidatorTypeLogin, request); err != nil {
		return nil, err
	}

	success, message, token, err := s.auth.Login(ctx, request.GetUsername(), request.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &authpb.LoginResponse{
		Success: success,
		Message: message,
		Token:   token,
	}, nil
}

func (s *serverAPI) Logout(ctx context.Context, request *authpb.LogoutRequest) (*authpb.LogoutResponse, error) {
	if err := s.validateRequest(ctx, ValidatorTypeLogout, request); err != nil {
		return nil, err
	}

	success, message, err := s.auth.Logout(ctx, request.GetToken())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &authpb.LogoutResponse{
		Success: success,
		Message: message,
	}, nil
}

// validateRequest validates a request using the appropriate validator.
func (s *serverAPI) validateRequest(ctx context.Context, validatorType ValidatorType, req interface{}) error {
	validator, exists := s.validators[validatorType]
	if !exists {
		return status.Errorf(codes.Internal, "validator not found for type: %s", validatorType)
	}
	return validator.Validate(ctx, req)
}
