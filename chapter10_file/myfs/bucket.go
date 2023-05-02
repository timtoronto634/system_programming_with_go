package myfs

import (
	"context"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/gcsblob"
)

func Bucket() {
	ctx := context.Background()
	bucket, _ := blob.OpenBucket(ctx, "gs://my-bucket")
	defer bucket.Close()
	reader, err := bucket.NewReader(ctx, "my-file", nil)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
}
