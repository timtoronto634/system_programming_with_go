package myfs

import (
	"github.com/winfsp/cgofuse/fuse"
	"gocloud.dev/blob"
)

type CloudFileSystem struct {
	fuse.FileSystemBase
	bucket *blob.Bucket
}
