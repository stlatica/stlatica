package auth

import (
	"context"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/kvs"
	"time"
)

// SignType 公開鍵だからばら撒いていいじゃん
// SignType 使用可能な署名アルゴリズムに対応した鍵に制限
type SignType interface {
	ecdsa.PrivateKey | ed25519.PrivateKey | rsa.PrivateKey
}

// PWAuth パスワードを用いたユーザー認証
type PWAuth[T SignType] interface {
	Auth(key T) (AccessToken[T], error)
}

// TokenAuth トークンを用いたユーザー認証
type TokenAuth[T SignType] interface {
	Verify(ctx context.Context, client kvs.Client, alg jwa.SignatureAlgorithm, key *T, lifetime int64) bool
	QuickVerify(ctx context.Context, client kvs.Client, alg jwa.SignatureAlgorithm, key *T, lifetime int64) bool
}

// UserCredential 認証に必要な情報を格納
type UserCredential[T SignType] struct {
	id        string
	pw        string
	createdBy int64
	alg       jwa.SignatureAlgorithm
	key       *T
}

// accessTokenEvidence トークン生成に必要な情報を格納
type accessTokenEvidence[T SignType] struct {
	id        string
	createdBy int64
	sessionId string
	podName   string
	alg       jwa.SignatureAlgorithm
	// key       *T
}

// AccessToken トークンを格納
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
func (u *UserCredential[T]) Auth(ctx context.Context, client kvs.Client, podName string, alg jwa.SignatureAlgorithm, key *T) (AccessToken[T], error) {
	storedPw := u.requestStoredPW(ctx, client, u.id)
	// PW check
	if !u.verifyPW(storedPw) {
		return AccessToken[T]{[]byte("")}, errors.New("invalid id or pw")
	}
	u.setTime()

	ate := u.toTokenEvidence(podName)
	accessToken := ate.makeJWS(alg, key)
	sessionId := uuid.New().String()

	saveTokens(ctx, client, sessionId, &accessToken)

	return accessToken, nil
}

// requestStoredPW ユーザーIDをキーにして保存されたPWを取得
func (u *UserCredential[T]) requestStoredPW(ctx context.Context, client kvs.Client, key string) string {
	storedPW, err := client.GetValue(ctx, key)
	if err != nil {
		panic(err)
	}
	return storedPW
}

// verifyPW パスワードの検証
func (u *UserCredential[T]) verifyPW(storedPw string) bool {
	return u.pw == storedPw
}

// setTime トークン生成時刻の設定
func (u *UserCredential[T]) setTime() {
	u.createdBy = time.Now().Unix()
}

// toTokenEvidence accessTokenEvidenceの生成
func (u *UserCredential[T]) toTokenEvidence(podName string) accessTokenEvidence[T] {
	return accessTokenEvidence[T]{id: u.id, createdBy: u.createdBy, sessionId: uuid.New().String(), podName: podName, alg: u.alg}
}

// serialize トークン生成に必要な情報をシリアライズ
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

// DeserializeAccessToken トークン生成に必要な情報をデシリアライズ
func (a *accessTokenEvidence[T]) DeserializeAccessToken(token []byte) error {
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

	return nil
}

// サーバで使用される署名アルゴリズムは1つに限る
// makeJWS 署名付きトークンの生成
func (a *accessTokenEvidence[T]) makeJWS(alg jwa.SignatureAlgorithm, key *T) AccessToken[T] {
	jsonData := a.serialize()
	token, err := jws.Sign(jsonData, alg, key)
	if err != nil {
		panic(err)
	}
	return AccessToken[T]{token: token}
}

// Export トークンのエクスポート
func (a *AccessToken[T]) Export() []byte {
	return a.token
}

// QuickVerify 高速なトークンの検証、保存されたトークンと比較を行わない
func (a *AccessToken[T]) QuickVerify(ctx context.Context, client kvs.Client, alg jwa.SignatureAlgorithm, key *T, lifetime int64) bool {
	jwt, err := jws.Verify(a.token, alg, key)
	if err != nil {
		panic(err)
		return false
	}

	// デシリアライズ
	ate := accessTokenEvidence[T]{id: "", createdBy: 9e18, sessionId: "", podName: ""}
	err = ate.DeserializeAccessToken(jwt)
	if err != nil {
		panic(err)
	}
	// 有効期限の確認とトークンの再発行を試行
	if ate.verifyExpirationDate(lifetime) {
		return true
	} else {
		// トークンの再発行を試行
		if a.refresh(ctx, client, ate, key) {
			return true
		} else {
			return false
		}
	}

}

// Verify 保存されたトークンとの比較を含めたトークンの検証
func (a *AccessToken[T]) Verify(ctx context.Context, client kvs.Client, alg jwa.SignatureAlgorithm, key *T, lifetime int64) bool {
	// 署名の検証
	jwt, err := jws.Verify(a.token, alg, key)
	if err != nil {
		panic(err)
		return false
	}

	// デシリアライズ
	ate := accessTokenEvidence[T]{id: "", createdBy: 9e18, sessionId: "", podName: ""}
	err = ate.DeserializeAccessToken(jwt)
	if err != nil {
		panic(err)
	}
	// 有効期限の確認
	if !ate.verifyExpirationDate(lifetime) {
		return false
	}

	// 有効期限の確認, 保存されたトークンの比較とトークンの再発行を試行
	if ate.verifyExpirationDate(lifetime) {
		storedToken := ate.requestStoredToken(ctx, client)
		return string(storedToken) == string(a.token)
	} else {
		// トークンの再発行を試行
		return a.refresh(ctx, client, ate, key)
	}

	return true
}

// refresh トークンの再発行
func (a *AccessToken[T]) refresh(ctx context.Context, client kvs.Client, ate accessTokenEvidence[T], key *T) bool {
	// リフレッシュトークンの検証
	_, err := ate.requestRefreshToken(ctx, client, a.Export())
	if err != nil {
		panic(err)
		return false
	}
	// トークンの再発行
	a.token = ate.regenerate(ctx, client, key)
	saveTokens(ctx, client, ate.sessionId, a)
	return true
}

// saveTokens トークンをDBに保存
func saveTokens[T SignType](ctx context.Context, client kvs.Client, sessionId string, accessToken *AccessToken[T]) {
	// save access token
	err := client.SetValue(ctx, sessionId, string(accessToken.Export()))
	if err != nil {
		panic(err)
	}
	// save refresh token
	// TODO: 有効期限の設定, redisのexpireを利用予定
	refreshToken := uuid.New()
	tmp := append([]byte(sessionId), accessToken.Export()...)
	key := sha256.Sum256(tmp)
	err = client.SetValue(ctx, string(key[:]), refreshToken.String())
	if err != nil {
		panic(err)
	}
}

// verifyExpirationDate トークンの有効期限の確認
func (a *accessTokenEvidence[T]) verifyExpirationDate(lifetime int64) bool {
	return time.Now().Unix() < a.createdBy+lifetime
}

// requestRefreshToken 保存されたリフレッシュトークンの要求
func (a *accessTokenEvidence[T]) requestRefreshToken(ctx context.Context, client kvs.Client, accessToken []byte) (string, error) {
	tmp := append([]byte(a.sessionId), accessToken...)
	key := sha256.Sum256(tmp)
	refreshToken, err := client.GetValue(ctx, string(key[:]))
	if err != nil {
		panic(err)
		return "", err
	}

	return refreshToken, nil
}

// setNewTime トークン生成時刻の更新
func (a *accessTokenEvidence[T]) setNewTime() {
	a.createdBy = time.Now().Unix()
}

// setNewSessionId セッションIDの更新
func (a *accessTokenEvidence[T]) setNewSessionId() {
	a.sessionId = uuid.New().String()
}

// regenerate トークンの再発行
func (a *accessTokenEvidence[T]) regenerate(ctx context.Context, client kvs.Client, key *T) []byte {
	// リフレッシュトークンの検証は必要？
	// トークンの情報の更新
	a.setNewTime()
	a.setNewSessionId()
	// トークンの再発行
	newToken := a.makeJWS(a.alg, key)

	return newToken.Export()
}

func (a *accessTokenEvidence[T]) requestStoredToken(ctx context.Context, client kvs.Client) []byte {
	storedToken, err := client.GetValue(ctx, a.sessionId)
	if err != nil {
		panic(err)
	}
	return []byte(storedToken)
}
