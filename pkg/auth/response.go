package auth

type SignInResponse struct {
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}

func NewSignInResponse(idToken, refreshToken, expiresIn string) *SignInResponse {
	return &SignInResponse{
		IDToken:      idToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}
}

type RefreshTokenResponse struct {
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    string `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

func NewRefreshTokenResponse(idToken, refreshToken, expiresIn, tokenType string) *RefreshTokenResponse {
	return &RefreshTokenResponse{
		IDToken:      idToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
		TokenType:    tokenType,
	}
}

type CreateUserResponse struct {
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	Email        string `json:"email"`
	LocalID      string `json:"localId"`
	ExpiresIn    string `json:"expiresIn"`
}

func NewCreateUserResponse(idToken, refreshToken, expiresIn, email, localId string) *CreateUserResponse {
	return &CreateUserResponse{
		IDToken:      idToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
		Email:        email,
		LocalID:      localId,
	}
}

type AccountInfoResponse struct {
	Email         string `json:"email"`
	LocalID       string `json:"localId"`
	EmailVerified bool   `json:"emailVerified"`
	CreationTime  string `json:"createdAt"`
	LastLoginTime string `json:"lastLoginAt"`
}

func NewAccountInfoResponse(email, localID string, emailVerified bool, creationTime, lastLoginTime string) *AccountInfoResponse {
	return &AccountInfoResponse{
		Email:         email,
		LocalID:       localID,
		EmailVerified: emailVerified,
		CreationTime:  creationTime,
		LastLoginTime: lastLoginTime,
	}
}

type FirebaseGetAccountInfoResponse struct {
	Users []AccountInfoResponse `json:"users"`
}
