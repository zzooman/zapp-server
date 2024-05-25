package token

import (
	"fmt"
	"strings"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/google/uuid"
)

// PasetoMaker는 Maker 인터페이스를 구현하는 구조체입니다.
type PasetoMaker struct {
    symmetricKey paseto.V4SymmetricKey
    implicit     []byte
}

func NewPasetoMaker() Maker {
    return &PasetoMaker{
        symmetricKey: paseto.NewV4SymmetricKey(),
        implicit:     []byte("my implicit nonce"),
    }
}

// CreateToken은 특정 사용자 이름과 기간에 대해 새로운 토큰을 생성합니다.
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
    // paseto 토큰 생성
    token := paseto.NewToken()
    // 토큰 ID용 uuid 생성
    tokenID, err := uuid.NewRandom()
    if err != nil {
        return "", err
    }
    // 토큰에 데이터 추가
    token.Set("id", tokenID.String())
    token.Set("username", username)
    token.SetIssuedAt(time.Now())
    token.SetExpiration(time.Now().Add(duration))
    tokenString := token.V4Encrypt(maker.symmetricKey, maker.implicit)

    return tokenString, nil
}

// VerifyToken은 토큰이 유효한지 확인합니다.
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
    parser := paseto.NewParser()
    parser.AddRule(paseto.NotExpired())

    // 토큰 파싱
    parsedToken, err := parser.ParseV4Local(maker.symmetricKey, token, maker.implicit)
    if err != nil {
        fmt.Println("Error parsing token:", err)
        if strings.Contains(err.Error(), "expired") {
            return nil, ErrExpiredToken
        }
        return nil, ErrInvalidToken
    }

    // 토큰에서 페이로드 생성
    payload, err := getPayloadFromToken(parsedToken)
    if err != nil {
        return nil, ErrInvalidToken
    }
    return payload, nil
}

func getPayloadFromToken(t *paseto.Token) (*Payload, error) {
    id, err := t.GetString("id")
    if err != nil {
        return nil, ErrInvalidToken
    }
    username, err := t.GetString("username")
    if err != nil {
        return nil, ErrInvalidToken
    }
    issuedAt, err := t.GetIssuedAt()
    if err != nil {
        return nil, ErrInvalidToken
    }
    expiredAt, err := t.GetExpiration()
    if err != nil {
        return nil, ErrInvalidToken
    }

    return &Payload{
        ID:        uuid.MustParse(id),
        Username:  username,
        IssuedAt:  issuedAt,
        ExpiredAt: expiredAt,
    }, nil
}

var _ Maker = (*PasetoMaker)(nil)
