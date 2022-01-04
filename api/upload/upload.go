package upload

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	BucketEnv = "BUCKET"
)

type Upload struct{}

func New() *Upload {
	return &Upload{}
}

func (u *Upload) Default(ctx context.Context, key string) (string, error) {
	bucket, ok := os.LookupEnv(BucketEnv)
	if !ok {
		return "", fmt.Errorf("bucket not set")
	}
	if key == "" {
		return "", fmt.Errorf("key not specified in request")
	}
	url, err := uploadURL(bucket, key)
	return url, err
}

func uploadURL(bucket, key string) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return "", err
	}
	client := s3.NewPresignClient(s3.NewFromConfig(cfg))
	poi := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	pou, err := client.PresignPutObject(context.Background(), poi)
	if err != nil {
		return "", err
	}
	return pou.URL, nil
}
