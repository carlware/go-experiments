package fs

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type GCS struct {
	client      *storage.Client
	bucketName  string
	placeHolder string
}

func New(keyPath, bucketName string) (*GCS, error) {
	client, err := storage.NewClient(context.Background(), option.WithCredentialsFile(keyPath))
	if err != nil {
		return nil, err
	}
	return &GCS{
		client:      client,
		bucketName:  bucketName,
		placeHolder: "https://storage.googleapis.com/" + bucketName + "/",
	}, nil
}

func (g *GCS) Upload(ctx context.Context, reader io.Reader) (string, error) {
	name := NewURI()
	bucket := g.client.Bucket(g.bucketName)
	w := bucket.Object(name).NewWriter(ctx)
	if _, err := io.Copy(w, reader); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}
	return g.placeHolder + name, nil
}

func (g *GCS) Close() {
	g.client.Close()
}
