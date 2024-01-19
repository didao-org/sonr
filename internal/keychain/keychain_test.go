package keychain_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sonrhq/sonr/internal/keychain"
)

func TestNewKeychain(t *testing.T) {
	kc, err := keychain.New(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, kc.Address)
}

func BenchmarkNewKeychain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kc, err := keychain.New(context.Background())
		if err != nil {
			b.Fail()
		}
		if len(kc.Address) == 0 {
			b.Fail()
		}
	}
}