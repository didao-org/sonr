package bucket

import (
	"errors"
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

var (
	errContentNotFound = func(id string) error {
		if id != "" {
			return fmt.Errorf("could not find content with id: %s", id)
		}

		return errors.New("could not find content")
	}
)

/*
Wraps items within a bucket. Items will be one of the following
DID -> reference to another bucket (WhereIs)
CID -> reference to content (map[string]interface{})
*/
type BucketContent struct {
	Item        interface{}           `json:"item"`
	Id          string                `json:"id"`
	ContentType bt.ResourceIdentifier `json:"contentType"`
}
type bucketImpl struct {
	address      string
	whereIs      *bt.WhereIs
	contentCache map[string]*BucketContent
	shell        *shell.Shell
	rpcEndpoint  string
}

func New(
	address string,
	whereIs *bt.WhereIs,
	shell *shell.Shell,
	rpcEndpoint string) Bucket {

	return &bucketImpl{
		address:      address,
		whereIs:      whereIs,
		shell:        shell,
		contentCache: make(map[string]*BucketContent),
		rpcEndpoint:  rpcEndpoint,
	}
}
