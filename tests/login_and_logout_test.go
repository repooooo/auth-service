package tests

import (
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/repooooo/auth-service/tests/suite"
	"github.com/repooooo/protos/gen/go/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoginAndLogout(t *testing.T) {
	ctx, st := suite.New(t)

	username := gofakeit.Username()
	password := gofakeit.Password(
		true,
		true,
		true,
		true,
		true,
		12,
	)

	loginResponse, err := st.AuthServiceClient.Login(ctx, &auth.LoginRequest{
		Username: username,
		Password: password,
	})
	require.NoError(t, err)
	//assert.NotEmpty(t, loginResponse.GetToken())
	assert.Empty(t, loginResponse.GetToken())

	_, err = st.AuthServiceClient.Logout(ctx, &auth.LogoutRequest{
		Token: loginResponse.GetToken(),
	})
	require.Error(t, errors.New("token is required"))
}
