package auth_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domainentities "github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	domainerrors "github.com/stlatica/stlatica/packages/backend/app/domains/errors"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	domainports "github.com/stlatica/stlatica/packages/backend/app/domains/users/ports"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/kvs"
	"github.com/stlatica/stlatica/packages/backend/app/repositories/dao"
	authusecase "github.com/stlatica/stlatica/packages/backend/app/usecases/auth"
	"golang.org/x/crypto/bcrypt"
)

const (
	testPassword        = "password"
	testPreferredUserID = "test-user"
	testUserID          = "01ARZ3NDEKTSV4RRFFQ69G5FAV"
)

func TestAuthenticateAccessToken(t *testing.T) {
	t.Parallel()

	useCase := newTestUseCase(t)
	tokenPair, err := useCase.Login(context.Background(), authusecase.LoginParams{
		PreferredUserID: testPreferredUserID,
		Password:        testPassword,
	})
	if err != nil {
		t.Fatalf("Login() error = %v", err)
	}

	userID, err := useCase.AuthenticateAccessToken(context.Background(), tokenPair.AccessToken)
	if err != nil {
		t.Fatalf("AuthenticateAccessToken() error = %v", err)
	}
	if userID != testUserID {
		t.Fatalf("AuthenticateAccessToken() userID = %q", userID)
	}
}

func TestRefreshInvalidatesOldRefreshToken(t *testing.T) {
	t.Parallel()

	useCase := newTestUseCase(t)
	tokenPair, err := useCase.Login(context.Background(), authusecase.LoginParams{
		PreferredUserID: testPreferredUserID,
		Password:        testPassword,
	})
	if err != nil {
		t.Fatalf("Login() error = %v", err)
	}

	refreshedTokenPair, err := useCase.Refresh(context.Background(), tokenPair.RefreshToken)
	if err != nil {
		t.Fatalf("Refresh() error = %v", err)
	}
	if refreshedTokenPair.RefreshToken == tokenPair.RefreshToken {
		t.Fatal("Refresh() did not rotate refresh token")
	}

	_, err = useCase.Refresh(context.Background(), tokenPair.RefreshToken)
	if err == nil {
		t.Fatal("Refresh() with stale refresh token error = nil")
	}

	assertDomainErrorType(t, err, domainerrors.DomainErrorTypeUnauthorized)
}

func TestRefreshReturnsUnauthorizedWhenSessionIsMissing(t *testing.T) {
	t.Parallel()

	useCase, sessionStore := newTestUseCaseWithSessionStore(t)
	tokenPair, err := useCase.Login(context.Background(), authusecase.LoginParams{
		PreferredUserID: testPreferredUserID,
		Password:        testPassword,
	})
	if err != nil {
		t.Fatalf("Login() error = %v", err)
	}

	sessionStore.values = map[string]string{}

	_, err = useCase.Refresh(context.Background(), tokenPair.RefreshToken)
	if err == nil {
		t.Fatal("Refresh() error = nil")
	}

	assertDomainErrorType(t, err, domainerrors.DomainErrorTypeUnauthorized)
}

func TestRefreshReturnsServiceUnavailableOnSessionStoreFailure(t *testing.T) {
	t.Parallel()

	useCase, sessionStore := newTestUseCaseWithSessionStore(t)
	tokenPair, err := useCase.Login(context.Background(), authusecase.LoginParams{
		PreferredUserID: testPreferredUserID,
		Password:        testPassword,
	})
	if err != nil {
		t.Fatalf("Login() error = %v", err)
	}

	sessionStore.getErr = errors.New("valkey timeout")

	_, err = useCase.Refresh(context.Background(), tokenPair.RefreshToken)
	if err == nil {
		t.Fatal("Refresh() error = nil")
	}

	assertDomainErrorType(t, err, domainerrors.DomainErrorTypeServiceUnavailable)
}

func TestRefreshReturnsServiceUnavailableOnMalformedSessionPayload(t *testing.T) {
	t.Parallel()

	useCase, sessionStore := newTestUseCaseWithSessionStore(t)
	tokenPair, err := useCase.Login(context.Background(), authusecase.LoginParams{
		PreferredUserID: testPreferredUserID,
		Password:        testPassword,
	})
	if err != nil {
		t.Fatalf("Login() error = %v", err)
	}

	sessionStore.overrideValue = "{not-json"

	_, err = useCase.Refresh(context.Background(), tokenPair.RefreshToken)
	if err == nil {
		t.Fatal("Refresh() error = nil")
	}

	assertDomainErrorType(t, err, domainerrors.DomainErrorTypeServiceUnavailable)
}

func TestAuthenticateAccessTokenReturnsUnauthorizedWhenSessionIsMissing(t *testing.T) {
	t.Parallel()

	useCase, sessionStore := newTestUseCaseWithSessionStore(t)
	tokenPair, err := useCase.Login(context.Background(), authusecase.LoginParams{
		PreferredUserID: testPreferredUserID,
		Password:        testPassword,
	})
	if err != nil {
		t.Fatalf("Login() error = %v", err)
	}

	sessionStore.values = map[string]string{}

	_, err = useCase.AuthenticateAccessToken(context.Background(), tokenPair.AccessToken)
	if err == nil {
		t.Fatal("AuthenticateAccessToken() error = nil")
	}

	assertDomainErrorType(t, err, domainerrors.DomainErrorTypeUnauthorized)
}

func TestAuthenticateAccessTokenReturnsServiceUnavailableOnSessionStoreFailure(t *testing.T) {
	t.Parallel()

	useCase, sessionStore := newTestUseCaseWithSessionStore(t)
	tokenPair, err := useCase.Login(context.Background(), authusecase.LoginParams{
		PreferredUserID: testPreferredUserID,
		Password:        testPassword,
	})
	if err != nil {
		t.Fatalf("Login() error = %v", err)
	}

	sessionStore.getErr = errors.New("valkey unavailable")

	_, err = useCase.AuthenticateAccessToken(context.Background(), tokenPair.AccessToken)
	if err == nil {
		t.Fatal("AuthenticateAccessToken() error = nil")
	}

	assertDomainErrorType(t, err, domainerrors.DomainErrorTypeServiceUnavailable)
}

func TestLoginReturnsServiceUnavailableWhenSessionStoreSaveFails(t *testing.T) {
	t.Parallel()

	useCase, sessionStore := newTestUseCaseWithSessionStore(t)
	sessionStore.setErr = errors.New("valkey unavailable")

	_, err := useCase.Login(context.Background(), authusecase.LoginParams{
		PreferredUserID: testPreferredUserID,
		Password:        testPassword,
	})
	if err == nil {
		t.Fatal("Login() error = nil")
	}

	assertDomainErrorType(t, err, domainerrors.DomainErrorTypeServiceUnavailable)
}

func newTestUseCase(t *testing.T) authusecase.AuthenticationUseCase {
	useCase, _ := newTestUseCaseWithSessionStore(t)
	return useCase
}

func newTestUseCaseWithSessionStore(t *testing.T) (
	authusecase.AuthenticationUseCase,
	*inMemorySessionStore,
) {
	t.Helper()

	userID, err := types.NewUserIDFromString(testUserID)
	if err != nil {
		t.Fatalf("NewUserIDFromString() error = %v", err)
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(testPassword), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("GenerateFromPassword() error = %v", err)
	}

	user := &domainentities.User{
		UserBase: domainentities.UserBase{
			UserID:          userID,
			PreferredUserID: testPreferredUserID,
			MailAddress:     "test@example.com",
		},
	}
	credential := &domainentities.UserAuthCredential{
		UserAuthCredentialBase: domainentities.UserAuthCredentialBase{
			UserID:            userID,
			EncryptedPassword: string(passwordHash),
		},
	}

	sessionStore := newInMemorySessionStore()

	return authusecase.NewAuthenticationUseCase(
		&fakeUserDAO{user: user},
		&fakeUserAuthCredentialDAO{credential: credential},
		sessionStore,
		authusecase.Config{
			JWTSecret:       "test-secret",
			AccessTokenTTL:  time.Minute,
			RefreshTokenTTL: time.Hour,
		},
	), sessionStore
}

type fakeUserDAO struct {
	user *domainentities.User
}

var _ dao.UserDAO = (*fakeUserDAO)(nil)

func (d *fakeUserDAO) GetUser(_ context.Context, userID types.UserID) (*domainentities.User, error) {
	if d.user.GetUserID().String() != userID.String() {
		return nil, errors.New("user not found")
	}
	return d.user, nil
}

func (d *fakeUserDAO) CreateUser(
	_ context.Context,
	_ domainentities.UserBase,
) (*domainentities.User, error) {
	return nil, errors.New("unsupported")
}

func (d *fakeUserDAO) GetUserByPreferredUserID(
	_ context.Context,
	preferredUserID string,
) (*domainentities.User, error) {
	if d.user.GetPreferredUserID() != preferredUserID {
		return nil, errors.New("user not found")
	}
	return d.user, nil
}

func (d *fakeUserDAO) GetUserByMailAddress(_ context.Context, mailAddress string) (*domainentities.User, error) {
	if d.user.GetMailAddress() != mailAddress {
		return nil, errors.New("user not found")
	}
	return d.user, nil
}

func (d *fakeUserDAO) GetFollows(
	_ context.Context,
	_ domainports.FollowsGetParams,
) ([]*domainentities.User, error) {
	return nil, errors.New("unsupported")
}

func (d *fakeUserDAO) GetFollowers(
	_ context.Context,
	_ domainports.FollowersGetParams,
) ([]*domainentities.User, error) {
	return nil, errors.New("unsupported")
}

func (d *fakeUserDAO) FollowUser(_ context.Context, _ domainentities.UserRelationBase) error {
	return errors.New("unsupported")
}

type fakeUserAuthCredentialDAO struct {
	credential *domainentities.UserAuthCredential
}

var _ dao.UserAuthCredentialDAO = (*fakeUserAuthCredentialDAO)(nil)

func (d *fakeUserAuthCredentialDAO) GetUserAuthCredential(
	_ context.Context,
	userID types.UserID,
) (*domainentities.UserAuthCredential, error) {
	if d.credential.GetUserID().String() != userID.String() {
		return nil, errors.New("credential not found")
	}
	return d.credential, nil
}

func (d *fakeUserAuthCredentialDAO) CreateUserAuthCredential(
	_ context.Context,
	_ types.UserID,
	_ string,
) (*domainentities.UserAuthCredential, error) {
	return nil, errors.New("unsupported")
}

type inMemorySessionStore struct {
	values        map[string]string
	getErr        error
	setErr        error
	overrideValue string
}

func newInMemorySessionStore() *inMemorySessionStore {
	return &inMemorySessionStore{
		values: map[string]string{},
	}
}

func (s *inMemorySessionStore) SetValue(_ context.Context, key string, value string) error {
	if s.setErr != nil {
		return s.setErr
	}
	s.values[key] = value
	return nil
}

func (s *inMemorySessionStore) GetValue(_ context.Context, key string) (string, error) {
	if s.getErr != nil {
		return "", s.getErr
	}
	if s.overrideValue != "" {
		return s.overrideValue, nil
	}
	value, ok := s.values[key]
	if !ok {
		return "", kvs.ErrValueNotFound
	}
	return value, nil
}

func assertDomainErrorType(
	t *testing.T,
	err error,
	expected domainerrors.DomainErrorType,
) {
	t.Helper()

	var domainErr domainerrors.DomainError
	if !errors.As(err, &domainErr) {
		t.Fatalf("error does not implement DomainError: %v", err)
	}
	if domainErr.DomainErrorType() != expected {
		t.Fatalf("DomainErrorType() = %v, want %v", domainErr.DomainErrorType(), expected)
	}
}
