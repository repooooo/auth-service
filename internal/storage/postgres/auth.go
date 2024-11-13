package postgres

import (
	"context"
)

func (s *Storage) Login(
	ctx context.Context,
	username string,
	password string,
) (
	bool,
	string,
	string,
	error,
) {
	const operation = "storage.postgres.Login"

	return false, "", "", nil
}

func (s *Storage) Logout(
	ctx context.Context,
	token string,
) (
	success bool,
	message string,
	err error,
) {
	const operation = "storage.postgres.Logout"

	return false, "", nil
}
