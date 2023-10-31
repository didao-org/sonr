package didwallets

import (
	"github.com/sonrhq/sonr/internal/crypto"
	"github.com/sonrhq/sonr/internal/mpc"
	"github.com/sonrhq/sonr/pkg/didcommon"
)

// Method is the DID method for Sonr Wallet Actor DIDs
const EthereumMethod = didcommon.Method("sonr")

type EthereumAccount struct {
	method    didcommon.Method
	ID        didcommon.Identifier
	Resources []didcommon.DIDResource
	acc       *mpc.AccountV1
	kss       mpc.KeyshareSet
}

// NewEthereumAccount creates a new Ethereum Wallet Actor DID
func NewEthereumAccount(key didcommon.SecretKey) (didcommon.WalletAccount, error) {
	ct := crypto.ETHCoinType
	m := didcommon.Method(ct.DIDMethod())
	acc, pks, err := mpc.GenerateV2("primary", ct)
	if err != nil {
		return nil, err
	}
	id := didcommon.Identifier(acc.Address)
	pbz, err := acc.Marshal()
	if err != nil {
		return nil, err
	}
	pr, err := id.AddResource("public", pbz)
	if err != nil {
		return nil, err
	}
	m.SetKey(id.String(), string(pbz))
	privBz, err := pks.EncryptUserKeyshare(key)
	if err != nil {
		return nil, err
	}
	encBz, err := privBz.Marshal()
	if err != nil {
		return nil, err
	}
	_, err = id.AddResource("private", encBz)
	if err != nil {
		return nil, err
	}
	return &EthereumAccount{
		method:    m,
		ID:        id,
		Resources: []didcommon.DIDResource{pr},
		acc:       acc,
		kss:       pks,
	}, nil
}

// ResolveEthereumAccount resolves a Sonr Wallet Actor DID
func ResolveEthereumAccount(didString string, key didcommon.SecretKey) (didcommon.WalletAccount, error) {
	ct := crypto.ETHCoinType
	m := didcommon.Method(ct.DIDMethod())

	id := didcommon.Identifier(didString)

	// Get public resource
	pubResource, err := id.FetchResource("public")
	if err != nil {
		return nil, err
	}
	acc := &mpc.AccountV1{}
	if err := acc.Unmarshal(pubResource); err != nil {
		return nil, err
	}
	// Get private resource and decrypt it
	privResource, err := id.FetchResource("private")
	if err != nil {
		return nil, err
	}
	epks := &mpc.EncKeyshareSet{}
	if err := epks.Unmarshal(privResource); err != nil {
		return nil, err
	}
	kss, err := epks.DecryptUserKeyshare(key)
	if err != nil {
		return nil, err
	}
	return &EthereumAccount{
		method: m,
		ID:     id,
		acc:    acc,
		kss:    kss,
	}, nil
}

// Address returns the address of the account
func (a *EthereumAccount) Address() string {
	return a.acc.Address
}

// Info returns the account data
func (a *EthereumAccount) Info() *crypto.AccountData {
	return a.acc.GetAccountData()
}

// Method returns the DID method
func (a *EthereumAccount) Method() didcommon.Method {
	return a.method
}

// Sign signs a message with the account
func (a *EthereumAccount) Sign(msg []byte) ([]byte, error) {
	return a.kss.Sign(msg)
}

// PublicKey returns the public key of the account
func (a *EthereumAccount) PublicKey() (*crypto.Secp256k1PubKey, error) {
	return a.acc.PublicKey(), nil
}

// Type returns the type of the account
func (a *EthereumAccount) Type() string {
	return "secp256k1"
}

// Verify verifies a signature
func (a *EthereumAccount) Verify(msg []byte, sig []byte) (bool, error) {
	return a.acc.Verify(msg, sig)
}