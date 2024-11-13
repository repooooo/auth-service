package authgrpc

import (
	"context"
	authpb "github.com/repooooo/protos/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(
		ctx context.Context,
		username string,
		password string,
	) (
		success bool,
		message string,
		token string,
		err error,
	)
	Logout(
		ctx context.Context,
		token string,
	) (
		success bool,
		message string,
		err error,
	)
}

type serverAPI struct {
	authpb.UnimplementedAuthServiceServer
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	authpb.RegisterAuthServiceServer(gRPC, &serverAPI{auth: auth})
}

func (s *serverAPI) Login(
	ctx context.Context,
	request *authpb.LoginRequest,
) (*authpb.LoginResponse, error) {
	if err := validateLoginRequest(request); err != nil {
		return nil, err
	}

	success, message, token, err := s.auth.Login(
		ctx,
		request.GetUsername(),
		request.GetPassword(),
	)
	if err != nil {
		// TODO: ...
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &authpb.LoginResponse{
		Success: success,
		Message: message,
		Token:   token,
	}, nil
}

func (s *serverAPI) Logout(
	ctx context.Context,
	request *authpb.LogoutRequest,
) (*authpb.LogoutResponse, error) {
	if err := validateLogoutRequest(request); err != nil {
		return nil, err
	}

	success, message, err := s.auth.Logout(ctx, request.GetToken())
	if err != nil {
		// TODO: ...
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &authpb.LogoutResponse{
		Success: success,
		Message: message,
	}, nil
}

func validateLoginRequest(request *authpb.LoginRequest) error {
	if request.GetUsername() == "" {
		return status.Error(codes.InvalidArgument, "username is required")
	}

	if request.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}

	return nil
}

func validateLogoutRequest(request *authpb.LogoutRequest) error {
	if request.GetToken() == "" {
		return status.Error(codes.InvalidArgument, "token is required")
	}
	return nil
}
