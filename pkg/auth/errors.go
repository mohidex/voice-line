package auth

import "fmt"

type AuthError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AuthError) Error() string {
	return fmt.Sprintf("Firebase Error: Code %d, Message: %s", e.Code, e.Message)
}

func NewFirebaseError(code int, message string) *AuthError{
	return &AuthError{
		Code:    code,
		Message: message,
	}
}
