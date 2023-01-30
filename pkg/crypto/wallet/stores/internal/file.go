package internal

import (
	"github.com/sonrhq/core/pkg/crypto/wallet"
	"github.com/sonrhq/core/pkg/crypto/wallet/accounts"
	vaultv1 "github.com/sonrhq/core/x/identity/types/vault/v1"
	bolt "go.etcd.io/bbolt"
)

type FileStore struct {
	accConfig *vaultv1.AccountConfig
	path      string
	db        *bolt.DB
	bucketKey []byte
	pwd       []byte
}

func NewFileStore(p string, pwd []byte, accCfg *vaultv1.AccountConfig) (wallet.Store, error) {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(p, 0600, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	ds := &FileStore{
		accConfig: accCfg,
		path:      p,
		db:        db,
		bucketKey: []byte(accCfg.DID()),
		pwd:       pwd,
	}
	return ds, nil
}

func (ds *FileStore) GetAccount(name string) (wallet.Account, error) {
	var acc wallet.Account
	ds.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(ds.bucketKey)
		v := b.Get([]byte(name))
		w, err := accounts.LoadFromBytes(v)
		if err != nil {
			return err
		}
		acc = w
		return nil
	})
	return acc, nil
}

func (ds *FileStore) PutAccount(w wallet.Account, name string) error {
	ds.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(ds.bucketKey)
		v, err := w.Marshal()
		if err != nil {
			return err
		}
		return b.Put([]byte(name), v)
	})
	return nil
}

// JWKClaims returns the JWKClaims for the store to be signed by the identity
func (ds *FileStore) JWKClaims(acc wallet.Account) (string, error) {
	return "", nil
}

// VerifyJWKClaims verifies the JWKClaims for the store
func (ds *FileStore) VerifyJWKClaims(claims string, acc wallet.Account) error {
	return nil
}
