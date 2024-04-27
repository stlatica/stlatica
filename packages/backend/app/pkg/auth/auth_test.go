package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"testing"
)

func Test(t *testing.T) {
	// create token
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	uc := UserCredential{id: "0000", pw: "0000", createdBy: 0}
	te := uc.toTokenEvidence()
	at := te.makeJWS(*key).token
	fmt.Println(te)
	fmt.Println(string(at))

	// decode token
	newToken := accessTokenEvidence{}
	payload, err := jws.Verify(at, jwa.ES256, key)
	if err != nil {
		t.Error(err)
	}
	err = newToken.DeserializeAccessToken(payload)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(newToken)

	fmt.Println(te == newToken)

}
