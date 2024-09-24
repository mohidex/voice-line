package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type FirebaseAuth struct {
	APIKey  string
	BaseURL string
}

func NewFirebaseAuth(apiKey, baseURL string) *FirebaseAuth {
	return &FirebaseAuth{
		APIKey:  apiKey,
		BaseURL: baseURL,
	}
}

func (f *FirebaseAuth) CreateUser(ctx context.Context, email, password string) (*CreateUserResponse, error) {
	url := fmt.Sprintf("%s/v1/accounts:signUp?key=%s", f.BaseURL, f.APIKey)
	body := map[string]string{
		"email":             email,
		"password":          password,
		"returnSecureToken": "true",
	}
	var createUserResponse CreateUserResponse
	err := f.makeRequest(ctx, url, body, &createUserResponse)
	if err != nil {
		return nil, err
	}
	return &createUserResponse, nil

}

func (f *FirebaseAuth) GetToken(ctx context.Context, email, password string) (*SignInResponse, error) {
	url := fmt.Sprintf("%s/v1/accounts:signInWithPassword?key=%s", f.BaseURL, f.APIKey)
	body := map[string]string{
		"email":             email,
		"password":          password,
		"returnSecureToken": "true",
	}
	var signInResponse SignInResponse
	if err := f.makeRequest(ctx, url, body, &signInResponse); err != nil {
		return nil, err
	}
	return &signInResponse, nil
}

func (f *FirebaseAuth) RefreshToken(ctx context.Context, refreshToken string) (*RefreshTokenResponse, error) {
	url := fmt.Sprintf("%s/v1/token?key=%s", f.BaseURL, f.APIKey)
	body := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
	}
	var refreshTokenResponse RefreshTokenResponse
	err := f.makeRequest(ctx, url, body, &refreshTokenResponse)
	if err != nil {
		return nil, err
	}
	return &refreshTokenResponse, nil
}

func (f *FirebaseAuth) GetAccountInfo(ctx context.Context, idToken string) (*AccountInfoResponse, error) {
	url := fmt.Sprintf("%s/v1/accounts:lookup?key=%s", f.BaseURL, f.APIKey)
	body := map[string]string{
		"idToken": idToken,
	}
	var firebaseResp FirebaseGetAccountInfoResponse
	err := f.makeRequest(ctx, url, body, &firebaseResp)
	if err != nil {
		return nil, err
	}

	if len(firebaseResp.Users) == 0 {
		return nil, fmt.Errorf("no user information found in response")
	}

	user := firebaseResp.Users[0]
	return NewAccountInfoResponse(user.Email, user.LocalID, user.EmailVerified, user.CreationTime, user.LastLoginTime), nil
}

func (f *FirebaseAuth) makeRequest(ctx context.Context, url string, body map[string]string, result interface{}) error {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("invalid request body: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var firebaseError AuthError
		if err := json.NewDecoder(resp.Body).Decode(&firebaseError); err == nil {
			return &firebaseError // Return Firebase-specific error
		}
		return fmt.Errorf("error response: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}

	return nil
}
