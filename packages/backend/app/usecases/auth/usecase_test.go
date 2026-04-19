package auth_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domainentities "github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	domainports "github.com/stlatica/stlatica/packages/backend/app/domains/users/ports"
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
}

func newTestUseCase(t *testing.T) authusecase.AuthenticationUseCase {
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

	return authusecase.NewAuthenticationUseCase(
		&fakeUserDAO{user: user},
		&fakeUserAuthCredentialDAO{credential: credential},
		newInMemorySessionStore(),
		authusecase.Config{
			JWTSecret:       "test-secret",
			AccessTokenTTL:  time.Minute,
			RefreshTokenTTL: time.Hour,
		},
	)
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
	values map[string]string
}

func newInMemorySessionStore() *inMemorySessionStore {
	return &inMemorySessionStore{
		values: map[string]string{},
	}
}

func (s *inMemorySessionStore) SetValue(_ context.Context, key string, value string) error {
	s.values[key] = value
	return nil
}

func (s *inMemorySessionStore) GetValue(_ context.Context, key string) (string, error) {
	value, ok := s.values[key]
	if !ok {
		return "", errors.New("not found")
	}
	return value, nil
}
