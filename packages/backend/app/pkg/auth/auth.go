package auth

import (
	"context"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/kvs"
	"golang.org/x/crypto/bcrypt"
)

const unseenEpoch = 9e18

// SignType 使用可能な署名アルゴリズムに対応した鍵に制限
type SignType interface {
	ecdsa.PrivateKey | ed25519.PrivateKey | rsa.PrivateKey
}

// PWAuth パスワードを用いたユーザー認証
type PWAuth[T SignType] interface {
	Auth(ctx context.Context, client kvs.Client) (AccessToken[T], error)
}

// TokenAuth トークンを用いたユーザー認証
type TokenAuth[T SignType] interface {
	Verify(ctx context.Context, client kvs.Client,
		alg jwa.SignatureAlgorithm, key *T, lifetime int64) (isVerified bool, userID string, err error)
	QuickVerify(ctx context.Context, client kvs.Client,
		alg jwa.SignatureAlgorithm, key *T, lifetime int64) (isVerified bool, userID string, err error)
}

// userCredential 認証に必要な情報を格納
type userCredential[T SignType] struct {
	userID      string
	pw          []byte
	encryptedPW []byte
	createdBy   int64
	alg         jwa.SignatureAlgorithm
	key         *T
}

// accessTokenEvidence トークン生成に必要な情報を格納
type accessTokenEvidence[T SignType] struct {
	id        string
	createdBy int64
	sessionID string
	alg       jwa.SignatureAlgorithm
}

// AccessToken トークンを格納
type AccessToken[T SignType] struct {
	token []byte
}

// NewPWAuth 認証情報を格納
func NewPWAuth[T SignType](id string, pw string, encryptedPW string, alg jwa.SignatureAlgorithm, key *T) PWAuth[T] {
	return &userCredential[T]{
		userID:      id,
		pw:          []byte(pw),
		encryptedPW: []byte(encryptedPW),
		createdBy:   unseenEpoch,
		alg:         alg,
		key:         key,
	}
}

// NewAccessToken アクセストークンを格納
func NewAccessToken[T SignType](token []byte) AccessToken[T] {
	return AccessToken[T]{token: token}
}

// Auth ID, PWで認証を行い、アクセストークンを返す
func (u *userCredential[T]) Auth(ctx context.Context, client kvs.Client) (AccessToken[T], error) {
	// PW check
	err := bcrypt.CompareHashAndPassword(u.encryptedPW, u.pw)
	if err != nil {
		return AccessToken[T]{[]byte("")}, err
	}
	u.setCurrentTime()

	ate := u.toTokenEvidence()
	accessToken, err := ate.makeJWS(u.alg, u.key)
	if err != nil {
		return AccessToken[T]{[]byte("")}, err
	}
	sessionID := uuid.New().String()

	err = saveTokens(ctx, client, sessionID, &accessToken)
	if err != nil {
		return AccessToken[T]{[]byte("")}, err
	}

	return accessToken, nil
}

// setCurrentTime トークン生成時刻の設定
func (u *userCredential[T]) setCurrentTime() {
	u.createdBy = time.Now().Unix()
}

// toTokenEvidence accessTokenEvidenceの生成
func (u *userCredential[T]) toTokenEvidence() accessTokenEvidence[T] {
	return accessTokenEvidence[T]{id: u.userID, createdBy: u.createdBy,
		sessionID: uuid.New().String(), alg: u.alg}
}

// serialize トークン生成に必要な情報をシリアライズ
func (a *accessTokenEvidence[T]) serialize() ([]byte, error) {
	jsonData, err := json.Marshal(
		&struct {
			ID        string `json:"userID,omitempty"`
			CreatedBy int64  `json:"created_by,omitempty"`
			SessionID string `json:"session_id,omitempty"`
			Alg       string `json:"alg,omitempty"`
		}{
			ID:        a.id,
			CreatedBy: a.createdBy,
			SessionID: a.sessionID,
			Alg:       a.alg.String(),
		})
	return jsonData, err
}

// DeserializeAccessToken トークン生成に必要な情報をデシリアライズ
func (a *accessTokenEvidence[T]) DeserializeAccessToken(token []byte) error {
	tmp := &struct {
		ID        string `json:"userID,omitempty"`
		CreatedBy int64  `json:"created_by,omitempty"`
		SessionID string `json:"session_id,omitempty"`
		Alg       string `json:"alg,omitempty"`
	}{}
	err := json.Unmarshal(token, &tmp)
	if err != nil {
		return err
	}
	a.id = tmp.ID
	a.createdBy = tmp.CreatedBy
	a.sessionID = tmp.SessionID
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

	return nil
}

// サーバで使用される署名アルゴリズムは1つに限る
// makeJWS 署名付きトークンの生成
func (a *accessTokenEvidence[T]) makeJWS(alg jwa.SignatureAlgorithm, key *T) (AccessToken[T], error) {
	jsonData, err := a.serialize()
	if err != nil {
		return AccessToken[T]{[]byte("")}, err
	}
	token, err := jws.Sign(jsonData, alg, key)
	if err != nil {
		return AccessToken[T]{[]byte("")}, err
	}
	return AccessToken[T]{token: token}, nil
}

// Export トークンのエクスポート
func (a *AccessToken[T]) Export() []byte {
	return a.token
}

// QuickVerify 高速なトークンの検証、保存されたトークンと比較を行わない
func (a *AccessToken[T]) QuickVerify(ctx context.Context, client kvs.Client,
	alg jwa.SignatureAlgorithm, key *T, lifetime int64) (bool, error) {
	// 署名の検証
	jwt, err := jws.Verify(a.token, alg, key)
	if err != nil {
		return false, err
	}

	// デシリアライズ
	ate := accessTokenEvidence[T]{id: "", createdBy: unseenEpoch, sessionID: ""}
	err = ate.DeserializeAccessToken(jwt)
	if err != nil {
		return false, err
	}
	// 有効期限の確認
	if ate.verifyExpirationDate(lifetime) {
		return true, nil
	}
	// トークンの再発行を試行
	res, err := a.refresh(ctx, client, ate, key)
	if err != nil {
		return false, err
	}
	return res, nil
}

// Verify 保存されたトークンとの比較を含めたトークンの検証
func (a *AccessToken[T]) Verify(ctx context.Context, client kvs.Client,
	alg jwa.SignatureAlgorithm, key *T, lifetime int64) (isVerified bool, userID string, err error) {
	// 署名の検証
	jwt, err := jws.Verify(a.token, alg, key)
	if err != nil {
		return false, "", err
	}

	// デシリアライズ
	ate := accessTokenEvidence[T]{id: "", createdBy: unseenEpoch, sessionID: ""}
	err = ate.DeserializeAccessToken(jwt)
	if err != nil {
		return false, "", err
	}
	// 有効期限の確認
	if !ate.verifyExpirationDate(lifetime) {
		return false, "", nil
	}

	// 有効期限の確認, 保存されたトークンの比較とトークンの再発行を試行
	if ate.verifyExpirationDate(lifetime) {
		var storedToken []byte
		storedToken, err = ate.requestStoredToken(ctx, client)
		if err != nil {
			return false, "", err
		}
		return string(storedToken) == string(a.token), ate.id, nil
	}
	// トークンの再発行を試行
	var res bool
	res, err = a.refresh(ctx, client, ate, key)
	if err != nil {
		return false, "", err
	}
	return res, ate.id, nil
}

// refresh トークンの再発行
func (a *AccessToken[T]) refresh(ctx context.Context, client kvs.Client,
	ate accessTokenEvidence[T], key *T) (bool, error) {
	// リフレッシュトークンの検証
	_, err := ate.requestRefreshToken(ctx, client, a.Export())
	if err != nil {
		return false, err
	}
	// トークンの再発行
	a.token, err = ate.regenerate(key)
	if err != nil {
		return false, err
	}
	err = saveTokens(ctx, client, ate.sessionID, a)
	if err != nil {
		return false, err
	}
	return true, nil
}

// saveTokens トークンをDBに保存
func saveTokens[T SignType](ctx context.Context, client kvs.Client, sessionID string,
	accessToken *AccessToken[T]) error {
	// save access token
	err := client.SetValue(ctx, sessionID, string(accessToken.Export()))
	if err != nil {
		return err
	}
	// save refresh token
	// TODO: 有効期限の設定, redisのexpireを利用予定
	refreshToken := uuid.New()
	tmp := append([]byte(sessionID), accessToken.Export()...)
	key := sha256.Sum256(tmp)
	err = client.SetValue(ctx, string(key[:]), refreshToken.String())
	if err != nil {
		return err
	}
	return nil
}

// verifyExpirationDate トークンの有効期限の確認
func (a *accessTokenEvidence[T]) verifyExpirationDate(lifetime int64) bool {
	return time.Now().Unix() < a.createdBy+lifetime
}

// requestRefreshToken 保存されたリフレッシュトークンの要求
func (a *accessTokenEvidence[T]) requestRefreshToken(ctx context.Context, client kvs.Client,
	accessToken []byte) (string, error) {
	tmp := append([]byte(a.sessionID), accessToken...)
	key := sha256.Sum256(tmp)
	refreshToken, err := client.GetValue(ctx, string(key[:]))
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

// setNewTime トークン生成時刻の更新
func (a *accessTokenEvidence[T]) setNewTime() {
	a.createdBy = time.Now().Unix()
}

// setNewSessionID セッションIDの更新
func (a *accessTokenEvidence[T]) setNewSessionID() {
	a.sessionID = uuid.New().String()
}

// regenerate トークンの再発行
func (a *accessTokenEvidence[T]) regenerate(key *T) ([]byte, error) {
	// リフレッシュトークンの検証は必要？
	// トークンの情報の更新
	a.setNewTime()
	a.setNewSessionID()
	// トークンの再発行
	newToken, err := a.makeJWS(a.alg, key)
	if err != nil {
		return []byte(""), err
	}

	return newToken.Export(), nil
}

func (a *accessTokenEvidence[T]) requestStoredToken(ctx context.Context, client kvs.Client) ([]byte, error) {
	storedToken, err := client.GetValue(ctx, a.sessionID)
	if err != nil {
		return []byte(""), err
	}
	return []byte(storedToken), nil
}
