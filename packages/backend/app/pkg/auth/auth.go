package auth

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"time"
)

// 公開鍵だからばら撒いていいじゃん
// SignType 使用可能な署名アルゴリズムに対応した鍵に制限
type SignType interface {
	ecdsa.PrivateKey | ed25519.PrivateKey | rsa.PrivateKey
}

type PWAuth[T SignType] interface {
	Auth(key T) (AccessToken[T], error)
}

type TokenAuth[T SignType] interface {
	Verify(alg jwa.SignatureAlgorithm, key T) (AccessToken[T], error)
	QuickVerify(alg jwa.SignatureAlgorithm, key T) (bool, error)
}

// UserCredential 認証に必要な情報を格納
type UserCredential[T SignType] struct {
	id        string
	pw        string
	createdBy int64
	alg       jwa.SignatureAlgorithm
	key       *T
}

// トークン生成に必要な情報を格納
type accessTokenEvidence[T SignType] struct {
	id        string
	createdBy int64
	sessionId string
	podName   string
	alg       jwa.SignatureAlgorithm
	key       *T
}

type AccessToken[T SignType] struct {
	token []byte
}

func NewUserCredentials[T SignType](id string, pw string, alg jwa.SignatureAlgorithm, key *T) UserCredential[T] {
	return UserCredential[T]{
		id:        id,
		pw:        pw,
		createdBy: 9e18,
		alg:       alg,
		key:       key,
	}
}

func NewAccessToken[T SignType](token []byte) AccessToken[T] {
	return AccessToken[T]{token: token}
}

// Auth ID, PWで認証を行い、アクセストークンを返す
func (u *UserCredential[T]) Auth(podName string, alg jwa.SignatureAlgorithm, key T) (AccessToken[T], error) {
	storedPw := u.requestStoredPW()
	// PW check
	if !u.verifyPW(storedPw) {
		return AccessToken[T]{[]byte("")}, errors.New("invalid pw")
	}
	u.setTime()

	ate := u.toTokenEvidence(podName)
	accessToken := ate.makeJWS(alg, key)
	sessionId := uuid.New().String()

	saveTokens(sessionId, accessToken)

	return accessToken, nil
}

func (u *UserCredential[T]) requestStoredPW() string {
	// TODO
	return ""
}

func (u *UserCredential[T]) verifyPW(storedPw string) bool {
	return u.pw == storedPw
}

func (u *UserCredential[T]) setTime() {
	u.createdBy = time.Now().Unix()
}

func (u *UserCredential[T]) toTokenEvidence(podName string) accessTokenEvidence[T] {
	return accessTokenEvidence[T]{id: u.id, createdBy: u.createdBy, sessionId: uuid.New().String(), podName: podName, alg: u.alg, key: u.key}
}

func (a *accessTokenEvidence[T]) serialize() []byte {
	jsonData, err := json.Marshal(
		&struct {
			Id        string `json:"id,omitempty"`
			CreatedBy int64  `json:"created_by,omitempty"`
			SessionId string `json:"session_id,omitempty"`
			PodName   string `json:"pod_name,omitempty"`
			Alg       string `json:"alg,omitempty"`
		}{
			Id:        a.id,
			CreatedBy: a.createdBy,
			SessionId: a.sessionId,
			PodName:   a.podName,
			Alg:       a.alg.String(),
		})
	if err != nil {
		panic(err)
	}
	return jsonData
}

func (a *accessTokenEvidence[T]) DeserializeAccessToken(token []byte, key *T) error {
	tmp := &struct {
		Id        string
		CreatedBy int64
		SessionId string
		PodName   string
		Alg       string
	}{}
	err := json.Unmarshal(token, &tmp)
	if err != nil {
		panic(err)
		return err
	}
	a.id = tmp.Id
	a.createdBy = tmp.CreatedBy
	a.sessionId = tmp.SessionId
	a.podName = tmp.PodName
	switch tmp.Alg {
	case "EdDSA":
		a.alg = jwa.EdDSA
	case "ES256":
		a.alg = jwa.ES256
	case "HS256":
		a.alg = jwa.HS256
	default:
		a.alg = jwa.NoSignature
	}
	a.key = key

	return nil
}

// トークンの生成
// サーバで使用される署名アルゴリズムは1つに限る
func (a *accessTokenEvidence[T]) makeJWS(alg jwa.SignatureAlgorithm, key T) AccessToken[T] {
	jsonData := a.serialize()
	token, err := jws.Sign(jsonData, alg, key)
	if err != nil {
		panic(err)
	}
	return AccessToken[T]{token: token}
}

func (a *AccessToken[T]) Export() []byte {
	return a.token
}

func (a *AccessToken[T]) QuickVerify(alg jwa.SignatureAlgorithm, key T) bool {
	_, err := jws.Verify(a.token, alg, key)
	if err != nil {
		panic(err)
		return false
	}
	// TODO: check effective date

	return true
}

func (a *AccessToken[T]) Verify(alg jwa.SignatureAlgorithm, key T) AccessToken[T] {
	ate := accessTokenEvidence[T]{id: "", createdBy: 9e18, sessionId: "", podName: ""}
	jwt, err := jws.Verify(a.token, alg, key)
	if err != nil {
		return AccessToken[T]{[]byte("")}
	}
	// TODO: request stored access token
	_ = jwt

	return AccessToken[T]{}
}

func saveTokens[T SignType](sessionId string, accessToken AccessToken[T]) {
	refreshToken := uuid.New()
	// TODO: save access token
	// TODO: save refresh token
	_ = refreshToken

}
