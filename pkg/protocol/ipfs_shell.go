package protocol

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path"
	"path/filepath"

	"github.com/ipfs/go-datastore"
	shell "github.com/ipfs/go-ipfs-api"
)

// IPFSShell implements a protocol.IPFS featuring:
// 	- IPFS initialization and error fallback
//	- In-memory cache mechanism
type IPFSShell struct {
	*shell.Shell
	cache datastore.Datastore
}

func NewIPFSShell(url string, cacheStore datastore.Datastore) *IPFSShell {
	return &IPFSShell{Shell: shell.NewShell(url), cache: cacheStore}
}

func (i *IPFSShell) LookUpData(ctx context.Context, cid string, data interface{}) error {
	//timestamp := fmt.Sprintf("%d", time.Now().Unix())
	outPath := filepath.Join(os.TempDir(), cid+".txt")

	defer os.Remove(outPath)

	if err := i.Get(cid, outPath); err != nil {
		return err
	}

	buf, err := os.ReadFile(outPath)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(buf, &data); err != nil {
		return err
	}

	return err
}

func (i *IPFSShell) GetData(ctx context.Context, cid string) ([]byte, error) {
	key := datastore.NewKey(cid)
	data, err := i.cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if data != nil {
		return data, nil
	}

	outPath := path.Join("..", "..", "data")
	if err = i.Get(cid, outPath); err != nil {
		return nil, err
	}

	buf, err := os.ReadFile(outPath)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(buf, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (i *IPFSShell) PutData(ctx context.Context, data []byte) (string, error) {
	cidStr, err := i.Add(bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	err = i.cache.Put(ctx, datastore.NewKey(cidStr), data)
	if err != nil {
		return "", err
	}

	return cidStr, nil
}

func (i *IPFSShell) PinFile(ctx context.Context, cidstr string) error {
	panic("implement me")
}

func (i *IPFSShell) RemoveFile(ctx context.Context, cidstr string) error {
	panic("implement me")
}
