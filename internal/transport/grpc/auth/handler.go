package authgrpc

import (
	"context"
	authpb "github.com/repooooo/protos/gen/go/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Login(ctx context.Context, request *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	if err := validateLoginRequest(request); err != nil {
		return nil, err
	}

	success, message, token, err := s.auth.Login(ctx, request.GetUsername(), request.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, ErrInternalServer)
	}

	return &authpb.LoginResponse{
		Success: success,
		Message: message,
		Token:   token,
	}, nil
}

func validateLoginRequest(request *authpb.LoginRequest) error {
	if request.GetUsername() == "" {
		return status.Error(codes.InvalidArgument, ErrUsernameRequired)
	}

	if request.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, ErrPasswordRequired)
	}

	return nil
}

func (s *serverAPI) Logout(ctx context.Context, request *authpb.LogoutRequest) (*authpb.LogoutResponse, error) {
	if err := validateLogoutRequest(request); err != nil {
		return nil, err
	}

	success, message, err := s.auth.Logout(ctx, request.GetToken())
	if err != nil {
		return nil, status.Error(codes.Internal, ErrInternalServer)
	}

	return &authpb.LogoutResponse{
		Success: success,
		Message: message,
	}, nil
}

func validateLogoutRequest(request *authpb.LogoutRequest) error {
	if request.GetToken() == "" {
		return status.Error(codes.InvalidArgument, ErrTokenRequired)
	}

	return nil
}
