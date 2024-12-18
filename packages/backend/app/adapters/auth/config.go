package auth

import (
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/auth"
)

// Config is the configuration for AuthAdapter.
type Config[T auth.SignType] struct {
	SignatureAlgorithm  jwa.SignatureAlgorithm
	SignaturePrivateKey *T
}
