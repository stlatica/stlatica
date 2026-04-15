package auth

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestAuthenticateAccessToken(t *testing.T) {
	t.Parallel()

	useCase := &authenticationUseCase{
		sessionStore: newInMemorySessionStore(),
		config: Config{
			JWTSecret:       "test-secret",
			AccessTokenTTL:  time.Minute,
			RefreshTokenTTL: time.Hour,
		},
	}

	tokenPair, err := useCase.issueTokenPair(context.Background(), "01ARZ3NDEKTSV4RRFFQ69G5FAV")
	if err != nil {
		t.Fatalf("issueTokenPair() error = %v", err)
	}

	userID, err := useCase.AuthenticateAccessToken(context.Background(), tokenPair.AccessToken)
	if err != nil {
		t.Fatalf("AuthenticateAccessToken() error = %v", err)
	}
	if userID != "01ARZ3NDEKTSV4RRFFQ69G5FAV" {
		t.Fatalf("AuthenticateAccessToken() userID = %q", userID)
	}
}

func TestRefreshInvalidatesOldRefreshToken(t *testing.T) {
	t.Parallel()

	useCase := &authenticationUseCase{
		sessionStore: newInMemorySessionStore(),
		config: Config{
			JWTSecret:       "test-secret",
			AccessTokenTTL:  time.Minute,
			RefreshTokenTTL: time.Hour,
		},
	}

	tokenPair, err := useCase.issueTokenPair(context.Background(), "01ARZ3NDEKTSV4RRFFQ69G5FAV")
	if err != nil {
		t.Fatalf("issueTokenPair() error = %v", err)
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
