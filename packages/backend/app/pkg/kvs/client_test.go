package kvs

import (
	"errors"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestToGetValueErrorMapsRedisNilToErrValueNotFound(t *testing.T) {
	t.Parallel()

	err := toGetValueError(redis.Nil)
	if !errors.Is(err, ErrValueNotFound) {
		t.Fatalf("errors.Is(err, ErrValueNotFound) = false, err = %v", err)
	}
}

func TestToGetValueErrorPreservesOperationalErrors(t *testing.T) {
	t.Parallel()

	originalErr := errors.New("timeout")
	err := toGetValueError(originalErr)
	if !errors.Is(err, originalErr) {
		t.Fatalf("errors.Is(err, originalErr) = false, err = %v", err)
	}
}
