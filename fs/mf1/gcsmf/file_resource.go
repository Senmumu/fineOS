package gcsmf

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/googleapis/google-cloud-go-testing/storage/stiface"
)

const (
	maxWriteSize = 10000
)

//gcs file resource is google cloud file file resource
type gcsFileResource struct {
	context  context.Context
	mf       *Mf
	name     string
	fileMode os.FileMode

	object         stiface.ObjectHandle
	currentGcsSize uint64
	writer         io.WriteCloser
	writeCloseTime time.Time

	reader        io.ReadCloser
	readCloseTime time.Time
	offset        uint64
	closed        bool
}

func (mfo *gcsFileResource) Close() error {
	mfo.closed = true
	return mfo.tryCloseIo()
}

func (mfo *gcsFileResource) tryCloseIo() error {
	if err := mfo.tryCloseReader(); err != nil {
		return fmt.Errorf("error when closing reader:%v", err)
	}
	if err := mfo.tryCloseWriter(); err != nil {
		return fmt.Errorf("error when closing writer: %v", err)
	}
	return nil
}

func (mfo *gcsFileResource) tryCloseReader() error {
	if mfo.reader == nil {
		return nil
	}
	if err := mfo.reader.Close(); err != nil {
		return err
	}
	mfo.reader = nil
	return nil
}

func (mfo *gcsFileResource) tryCloseWriter() error {
	if mfo.writer == nil {
		return nil
	}
	//to avoid partial writes
	if mfo.currentGcsSize > mfo.offset {
		currentFile, err := mfo.object.NewRangeReader(mfo.context, mfo.offset, -1)
	}
	if err := mfo.writer.Close(); err != nil {
		return err
	}
	mfo.writer = nil
	return nil
}
