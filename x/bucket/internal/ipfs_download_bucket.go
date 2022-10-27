package internal

import (
	"context"
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func DownloadBucket(sh *shell.Shell, whereIs *bt.Bucket, address string) ([]bt.BucketItem, error) {
	files, err := sh.FilesLs(context.Background(), whereIs.GetPath(address))
	if err != nil {
		return nil, fmt.Errorf("Failed to list item at %s - %e", whereIs.GetPath(address), err)
	}
	items := make([]bt.BucketItem, 0)
	for _, file := range files {
		reader, err := sh.FilesRead(context.Background(), whereIs.GetPath(address, file.Name))
		if err != nil {
			return nil, fmt.Errorf("Failed to read item at %s - %e", whereIs.GetPath(address, file.Name), err)
		}
		defer reader.Close()
		item, err := bt.NewBucketItemFromReader(file.Name, reader)
		if err != nil {
			return nil, fmt.Errorf("Failed to create item at %s - %e", whereIs.GetPath(address, file.Name), err)
		}
		items = append(items, item)
	}
	return items, nil
}

func DownloadBucketAsync(sh *shell.Shell, whereIs *bt.Bucket, address string) (BucketDownloadProgressCallback, error) {
	files, err := sh.FilesLs(context.Background(), whereIs.GetPath(address))
	if err != nil {
		return nil, fmt.Errorf("Failed to read item at %s - %e", whereIs.GetPath(address), err)
	}

	return newBucketDownloader(sh, whereIs, address, files), nil
}

type BucketDownloadProgressCallback interface {
	Total() uint64
	Listen() <-chan uint64
	OnComplete() <-chan []bt.BucketItem
}

type bucketDownloadProgressCallback struct {
	total        uint64
	current      uint64
	bucket       *bt.Bucket
	progressChan chan uint64
	itemsChan    chan bt.BucketItem
	items        []bt.BucketItem
	completeChan chan []bt.BucketItem
}

func newBucketDownloader(sh *shell.Shell, whereIs *bt.Bucket, address string, files []*shell.MfsLsEntry) *bucketDownloadProgressCallback {
	total := uint64(0)
	for _, file := range files {
		total += file.Size
	}
	cb := &bucketDownloadProgressCallback{
		total:        total,
		current:      0,
		bucket:       whereIs,
		progressChan: make(chan uint64),
		itemsChan:    make(chan bt.BucketItem),
		items:        make([]bt.BucketItem, 0),
	}
	go cb.handleDownload(sh, whereIs, address, files)
	go cb.handleChannels()
	return cb
}

func (cb *bucketDownloadProgressCallback) handleDownload(sh *shell.Shell, whereIs *bt.Bucket, address string, files []*shell.MfsLsEntry) {
	for _, file := range files {
		reader, err := sh.FilesRead(context.Background(), whereIs.GetPath(address, file.Name))
		if err != nil {
			cb.itemsChan <- nil
			return
		}
		defer reader.Close()
		item, err := bt.NewBucketItemFromReader(file.Name, reader)
		if err != nil {
			cb.itemsChan <- nil
			return
		}
		cb.itemsChan <- item
		cb.progressChan <- cb.current + file.Size
	}
}

// This method is called on a routine and updates the array of items and current progress until the download is complete
func (cb *bucketDownloadProgressCallback) handleChannels() {
	for {
		select {
		case item := <-cb.itemsChan:
			cb.items = append(cb.items, item)
		case progress := <-cb.progressChan:
			cb.current = progress
			if cb.current == cb.total {
				cb.completeChan <- cb.items

				// Close all channels
				close(cb.progressChan)
				close(cb.itemsChan)
				close(cb.completeChan)
				return
			}
		}
	}

}

func (b *bucketDownloadProgressCallback) Listen() <-chan uint64 {
	return b.progressChan
}

func (b *bucketDownloadProgressCallback) Total() uint64 {
	return b.total
}

func (b *bucketDownloadProgressCallback) OnComplete() <-chan []bt.BucketItem {
	return b.completeChan
}