package auth

import "context"

type Authenticator interface {
	CreateUser(ctx context.Context, email, password string) (*CreateUserResponse, error)
	GetToken(ctx context.Context, email, password string) (*SignInResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (*RefreshTokenResponse, error)
	GetAccountInfo(ctx context.Context, idToken string) (*AccountInfoResponse, error)
}
