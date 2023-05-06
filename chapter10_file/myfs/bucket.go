package myfs

import (
	"context"
	"fmt"
	"os"

	"github.com/winfsp/cgofuse/fuse"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/s3blob"
)

func Bucket() {
	ctx := context.Background()
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s <bucket-name> <mount-point> etc...\n", os.Args[1])
		os.Exit(1)
	}
	bucket, err := blob.OpenBucket(ctx, os.Args[1])
	if err != nil {
		panic(err)
	}
	defer bucket.Close()
	cf := &CloudFileSystem{bucket: bucket}
	host := fuse.NewFileSystemHost(cf)
	host.Mount(os.Args[2], os.Args[3:])
}
