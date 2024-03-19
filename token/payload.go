package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)


var (
	ErrExpiredToken = errors.New("token is expired")
	ErrInvalidToken = errors.New("token is invalid")
)
// Payload is a struct to hold the payload data
type Payload struct {
	ID 			uuid.UUID 	`json:"id"`
	Username 	string 		`json:"username"`
	IssuedAt 	time.Time 	`json:"issued_at"`
	ExpiredAt 	time.Time 	`json:"expired_at"`	
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID: tokenID,
		Username: username,
		IssuedAt: time.Now(),
	}
	payload.ExpiredAt = payload.IssuedAt.Add(duration)
	return payload, nil
}



// For jwt.Claims interface
func (payload *Payload) GetSubject() (string, error) {
	return payload.Username, nil
}
func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.IssuedAt), nil
}
func (payload *Payload) GetIssuer() (string, error) {
	return "zapp", nil
}
func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.IssuedAt), nil 
}
func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.ExpiredAt), nil
}
func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return jwt.ClaimStrings{}, nil
}
