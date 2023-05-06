package myfs

import (
	"context"
	"strings"

	"github.com/winfsp/cgofuse/fuse"
	"gocloud.dev/blob"
)

type CloudFileSystem struct {
	fuse.FileSystemBase
	bucket *blob.Bucket
}

func (cf *CloudFileSystem) Getattr(path string, stat *fuse.Stat_t, fh uint64) (errc int) {
	if path == "/" {
		stat.Mode = fuse.S_IFDIR | 0555
		return 0
	}
	ctx := context.Background()
	name := strings.TrimLeft(path, "/")
	a, err := cf.bucket.Attributes(ctx, name)
	if err != nil {
		_, err := cf.bucket.Attributes(ctx, name+"/")
		if err != nil {
			return -fuse.ENOENT
		}
		stat.Mode = fuse.S_IFDIR | 0555
	} else {
		stat.Mode = fuse.S_IFREG | 0444
		stat.Size = a.Size
		stat.Mtim = fuse.NewTimespec(a.ModTime)
	}
	stat.Nlink = 1
	return 0
}

func (cf *CloudFileSystem) Readdir(path string, fill func(name string, stat *fuse.Stat_t, ofst int64) bool, ofst int64, fh uint64) (errc int) {
	ctx := context.Background()
	fill(".", nil, 0)
	fill("..", nil, 0)
	prefix := strings.TrimLeft(path, "/")
	if prefix != "" {
		prefix = prefix + "/"
	}
	i := cf.bucket.List(&blob.ListOptions{
		Prefix:    prefix,
		Delimiter: "/",
	})
	for {
		o, err := i.Next(ctx)
		if err != nil {
			break
		}
		key := o.Key[len(prefix):]
		if len(key) == 0 {
			continue
		}
		fill(strings.TrimRight(key, "/"), nil, 0)
	}
	return 0
}

func (cf *CloudFileSystem) Read(path string, buff []byte, ofst int64, fh uint64) (n int) {
	name := strings.TrimLeft(path, "/")
	ctx := context.Background()
	reader, err := cf.bucket.NewRangeReader(
		ctx, name, ofst, int64(len(buff)), nil)
	if err != nil {
		return
	}
	defer reader.Close()

	n, _ = reader.Read(buff)
	return
}
