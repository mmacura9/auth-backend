package tokenutil

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/aead/chacha20poly1305"
)

type PasetoMaker struct {
	symmetricKey paseto.V4SymmetricKey
	implicit     []byte
}

func NewPasetoMaker(implicit string) (Maker, error) {
	if len(implicit) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("Invalid implicit part size: must have exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		symmetricKey: paseto.NewV4SymmetricKey(),
		implicit:     []byte(implicit),
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return "", payload, err
	}

	token, err := paseto.NewTokenFromClaimsJSON(payloadJson, nil)
	if err != nil {
		return "", payload, err
	}

	encryptedToken := token.V4Encrypt(maker.symmetricKey, maker.implicit)

	return encryptedToken, payload, err
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	parser := paseto.NewParser()
	parser.AddRule(paseto.NotExpired())
	parsedToken, err := parser.ParseV4Local(maker.symmetricKey, token, maker.implicit)
	if err != nil {
		if strings.Contains(err.Error(), "expired") {
			return nil, ErrorExpiredToken
		}
		return nil, ErrorInvalidToken
	}

	payload, err := getPayloadFromToken(parsedToken)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func getPayloadFromToken(t *paseto.Token) (*Payload, error) {
	payload := Payload{}
	err := json.Unmarshal(t.ClaimsJSON(), &payload)

	return &payload, err
}
