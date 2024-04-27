package auth

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"time"
)

type UserCredential struct {
	id        string
	pw        string
	createdBy int64
}

type accessTokenEvidence struct {
	id         string
	createdBy  int64
	randomness string
}

type AccessCredential struct {
	sessionId   string
	accessToken []byte
}

func NewUserCredentials(id string, pw string) UserCredential {
	return UserCredential{
		id:        id,
		pw:        pw,
		createdBy: 9e18,
	}
}

func saveTokens(sessionId string, accessToken []byte) {
	refreshToken := uuid.New()
	// TODO: save access token
	// TODO: save refresh token

}

func (u *UserCredential) Auth(key rsa.PrivateKey) (AccessCredential, error) {
	storedPw := u.requestStoredPW()
	// PW check
	if !u.verify(storedPw) {
		return AccessCredential{sessionId: "", accessToken: []byte("")}, errors.New("invalid pw")
	}
	u.setTime()

	ate := u.toTokenEvidence()
	accessToken := ate.makeJWS(key)
	sessionId := uuid.New().String()

	saveTokens(sessionId, accessToken)

	return AccessCredential{sessionId: sessionId, accessToken: accessToken}, nil
}

func (u *UserCredential) requestStoredPW() string {
	// TODO
	return ""
}

func (u *UserCredential) verify(storedPw string) bool {
	return u.pw == storedPw
}

func (u *UserCredential) setTime() {
	u.createdBy = time.Now().Unix()
}

func (u *UserCredential) toTokenEvidence() accessTokenEvidence {
	return accessTokenEvidence{id: u.id, createdBy: u.createdBy, randomness: uuid.New().String()}
}

func (a *accessTokenEvidence) serialize() []byte {
	jsonData, err := json.Marshal(*a)
	if err != nil {
		panic(err)
	}
	return jsonData
}

func (a *accessTokenEvidence) makeJWS(key rsa.PrivateKey) []byte {
	jsonData := a.serialize()
	token, err := jws.Sign(jsonData, jwa.ES256, key)
	if err != nil {
		panic(err)
	}
	return token
}
