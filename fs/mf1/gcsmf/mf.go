package gcsmf

import (
	"context"

	"github.com/googleapis/google-cloud-go-testing/storage/stiface"
)

const (
	deFileMode = 0755
	gsPrefix   = "gs://"
)

type Mf struct {
	context       context.Context
	client        stiface.Client
	buckets       map[string]stiface.BucketHandle
	rawGcsObjects map[string]*GcsFile
}
