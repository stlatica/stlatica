package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	domainentities "github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	domainerrors "github.com/stlatica/stlatica/packages/backend/app/domains/errors"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/kvs"
	"github.com/stlatica/stlatica/packages/backend/app/repositories/dao"
	"golang.org/x/crypto/bcrypt"
)

// AuthenticationUseCase is the interface for authenticating user.
type AuthenticationUseCase interface {
	// Login authenticates user and creates new tokens.
	Login(ctx context.Context, params LoginParams) (*LoginResult, error)
	// Refresh refreshes tokens from a refresh token.
	Refresh(ctx context.Context, refreshToken string) (*RefreshResult, error)
	// AuthenticateAccessToken authenticates an access token and returns the user ID.
	AuthenticateAccessToken(ctx context.Context, accessToken string) (string, error)
}

// Config is the authentication configuration.
type Config struct {
	JWTSecret       string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

// LoginParams is the input for login.
type LoginParams struct {
	MailAddress     string
	PreferredUserID string
	Password        string
}

// LoginResult is the output for login.
type LoginResult struct {
	PreferredUserID       string
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  time.Time
	RefreshTokenExpiresAt time.Time
}

// RefreshResult is the output for refresh.
type RefreshResult struct {
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  time.Time
	RefreshTokenExpiresAt time.Time
}

// NewAuthenticationUseCase returns AuthenticationUseCase.
func NewAuthenticationUseCase(
	userDAO dao.UserDAO,
	userAuthCredentialDAO dao.UserAuthCredentialDAO,
	sessionStore kvs.Client,
	config Config,
) AuthenticationUseCase {
	return &authenticationUseCase{
		userDAO:               userDAO,
		userAuthCredentialDAO: userAuthCredentialDAO,
		sessionStore:          sessionStore,
		config:                config,
	}
}

type authenticationUseCase struct {
	userDAO               dao.UserDAO
	userAuthCredentialDAO dao.UserAuthCredentialDAO
	sessionStore          kvs.Client
	config                Config
}

type sessionState struct {
	UserID         string `json:"user_id"`
	RefreshTokenID string `json:"refresh_token_id"`
	ExpiresAtUnix  int64  `json:"expires_at_unix"`
}

type tokenClaims struct {
	TokenType string `json:"typ"`
	SessionID string `json:"sid"`

	jwt.RegisteredClaims
}

func (u *authenticationUseCase) Login(ctx context.Context, params LoginParams) (*LoginResult, error) {
	if params.Password == "" {
		return nil, domainerrors.NewDomainError(
			errors.New("password is required"),
			domainerrors.DomainErrorTypeInvalidData,
			"invalid login request",
		)
	}
	if params.MailAddress == "" && params.PreferredUserID == "" {
		return nil, domainerrors.NewDomainError(
			errors.New("mail_address or preferred_user_id is required"),
			domainerrors.DomainErrorTypeInvalidData,
			"invalid login request",
		)
	}

	user, err := u.resolveLoginUser(ctx, params.MailAddress, params.PreferredUserID)
	if err != nil {
		return nil, err
	}

	credential, err := u.userAuthCredentialDAO.GetUserAuthCredential(ctx, user.GetUserID())
	if err != nil {
		return nil, u.toUnauthorizedError(err)
	}
	if err := bcrypt.CompareHashAndPassword(
		[]byte(credential.GetEncryptedPassword()),
		[]byte(params.Password),
	); err != nil {
		return nil, u.toUnauthorizedError(err)
	}

	tokenPair, err := u.issueTokenPair(ctx, user.GetUserID().String())
	if err != nil {
		return nil, err
	}

	return &LoginResult{
		PreferredUserID:       user.GetPreferredUserID(),
		AccessToken:           tokenPair.AccessToken,
		RefreshToken:          tokenPair.RefreshToken,
		AccessTokenExpiresAt:  tokenPair.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: tokenPair.RefreshTokenExpiresAt,
	}, nil
}

func (u *authenticationUseCase) Refresh(ctx context.Context, refreshToken string) (*RefreshResult, error) {
	if refreshToken == "" {
		return nil, u.toUnauthorizedError(errors.New("refresh token is required"))
	}

	claims, err := u.parseAndValidateToken(refreshToken, tokenTypeRefresh)
	if err != nil {
		return nil, u.toUnauthorizedError(err)
	}

	session, err := u.loadSession(ctx, claims.SessionID)
	if err != nil {
		return nil, u.toUnauthorizedError(err)
	}
	if err := u.validateSession(session, claims, true); err != nil {
		return nil, u.toUnauthorizedError(err)
	}

	tokenPair, err := u.issueTokenPairForSession(ctx, session.UserID, claims.SessionID)
	if err != nil {
		return nil, err
	}

	return &RefreshResult{
		AccessToken:           tokenPair.AccessToken,
		RefreshToken:          tokenPair.RefreshToken,
		AccessTokenExpiresAt:  tokenPair.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: tokenPair.RefreshTokenExpiresAt,
	}, nil
}

func (u *authenticationUseCase) AuthenticateAccessToken(ctx context.Context, accessToken string) (string, error) {
	if accessToken == "" {
		return "", u.toUnauthorizedError(errors.New("access token is required"))
	}

	claims, err := u.parseAndValidateToken(accessToken, tokenTypeAccess)
	if err != nil {
		return "", u.toUnauthorizedError(err)
	}

	session, err := u.loadSession(ctx, claims.SessionID)
	if err != nil {
		return "", u.toUnauthorizedError(err)
	}
	if err := u.validateSession(session, claims, false); err != nil {
		return "", u.toUnauthorizedError(err)
	}

	return claims.Subject, nil
}

const (
	tokenTypeAccess  = "access"
	tokenTypeRefresh = "refresh"
)

type issuedTokenPair struct {
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  time.Time
	RefreshTokenExpiresAt time.Time
}

func (u *authenticationUseCase) resolveLoginUser(
	ctx context.Context,
	mailAddress string,
	preferredUserID string,
) (*domainentities.User, error) {
	var (
		userByMail      *domainentities.User
		userByPreferred *domainentities.User
		err             error
	)

	if mailAddress != "" {
		userByMail, err = u.userDAO.GetUserByMailAddress(ctx, mailAddress)
		if err != nil {
			return nil, u.toUnauthorizedError(err)
		}
	}

	if preferredUserID != "" {
		userByPreferred, err = u.userDAO.GetUserByPreferredUserID(ctx, preferredUserID)
		if err != nil {
			return nil, u.toUnauthorizedError(err)
		}
	}

	if userByMail != nil && userByPreferred != nil &&
		userByMail.GetUserID().String() != userByPreferred.GetUserID().String() {
		return nil, domainerrors.NewDomainError(
			errors.New("mail_address and preferred_user_id point to different users"),
			domainerrors.DomainErrorTypeInvalidData,
			"invalid login request",
		)
	}

	if userByMail != nil {
		return userByMail, nil
	}
	return userByPreferred, nil
}

func (u *authenticationUseCase) issueTokenPair(ctx context.Context, userID string) (*issuedTokenPair, error) {
	sessionID := uuid.NewString()
	return u.issueTokenPairForSession(ctx, userID, sessionID)
}

func (u *authenticationUseCase) issueTokenPairForSession(
	ctx context.Context,
	userID string,
	sessionID string,
) (*issuedTokenPair, error) {
	now := time.Now()
	accessExpiresAt := now.Add(u.config.AccessTokenTTL)
	refreshExpiresAt := now.Add(u.config.RefreshTokenTTL)
	refreshTokenID := uuid.NewString()

	accessToken, err := u.signToken(tokenClaims{
		TokenType: tokenTypeAccess,
		SessionID: sessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(accessExpiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	})
	if err != nil {
		return nil, err
	}

	refreshToken, err := u.signToken(tokenClaims{
		TokenType: tokenTypeRefresh,
		SessionID: sessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        refreshTokenID,
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(refreshExpiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	})
	if err != nil {
		return nil, err
	}

	if err := u.saveSession(ctx, sessionID, sessionState{
		UserID:         userID,
		RefreshTokenID: refreshTokenID,
		ExpiresAtUnix:  refreshExpiresAt.Unix(),
	}); err != nil {
		return nil, err
	}

	return &issuedTokenPair{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessExpiresAt,
		RefreshTokenExpiresAt: refreshExpiresAt,
	}, nil
}

func (u *authenticationUseCase) signToken(claims tokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(u.config.JWTSecret))
}

func (u *authenticationUseCase) parseAndValidateToken(rawToken string, expectedType string) (*tokenClaims, error) {
	claims := &tokenClaims{}
	token, err := jwt.ParseWithClaims(rawToken, claims, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Method.Alg())
		}
		return []byte(u.config.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims.TokenType != expectedType {
		return nil, errors.New("invalid token type")
	}
	if claims.Subject == "" {
		return nil, errors.New("invalid token subject")
	}
	if claims.SessionID == "" {
		return nil, errors.New("invalid session id")
	}
	return claims, nil
}

func (u *authenticationUseCase) validateSession(
	session *sessionState,
	claims *tokenClaims,
	requireRefreshTokenID bool,
) error {
	if session.UserID != claims.Subject {
		return errors.New("session subject mismatch")
	}
	if requireRefreshTokenID && session.RefreshTokenID != claims.ID {
		return errors.New("refresh token mismatch")
	}
	if time.Now().After(time.Unix(session.ExpiresAtUnix, 0)) {
		return errors.New("session expired")
	}
	return nil
}

func (u *authenticationUseCase) saveSession(ctx context.Context, sessionID string, session sessionState) error {
	payload, err := json.Marshal(session)
	if err != nil {
		return err
	}
	return u.sessionStore.SetValue(ctx, sessionKey(sessionID), string(payload))
}

func (u *authenticationUseCase) loadSession(ctx context.Context, sessionID string) (*sessionState, error) {
	payload, err := u.sessionStore.GetValue(ctx, sessionKey(sessionID))
	if err != nil {
		return nil, err
	}

	var session sessionState
	if err := json.Unmarshal([]byte(payload), &session); err != nil {
		return nil, err
	}
	return &session, nil
}

func sessionKey(sessionID string) string {
	return fmt.Sprintf("auth:session:%s", sessionID)
}

func (u *authenticationUseCase) toUnauthorizedError(err error) error {
	return domainerrors.NewDomainError(err, domainerrors.DomainErrorTypeUnauthorized, "unauthorized")
}
