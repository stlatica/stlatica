package objectstorage

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/friendsofgo/errors"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
)

// Client is an interface for object storage client.
type Client interface {
	// MakeBucket creates a new bucket on object storage.
	MakeBucket(ctx context.Context, bucketName string) error
	// BucketExists checks if a bucket exists.
	BucketExists(ctx context.Context, bucketName string) (bool, error)
	// GetObject gets an object from object storage.
	GetObject(ctx context.Context, bucketName string, objectName string) (io.ReadCloser, error)
	// PutObject puts an object to object storage.
	PutObject(ctx context.Context, bucketName string, objectName string, reader io.Reader, contentType string) error
}

type client struct {
	orgClient *s3.Client
	appLogger *logger.AppLogger
}

// NewClient creates a new object storage client.
func NewClient(ctx context.Context, appLogger *logger.AppLogger, endpoint string,
	region string, accessKey string, secretKey string) Client {
	cred := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")

	cfg, err := config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(cred))
	if err != nil {
		panic(err)
	}

	s3Client := s3.NewFromConfig(cfg, func(options *s3.Options) {
		options.UsePathStyle = true
		options.BaseEndpoint = aws.String(endpoint)
		options.Region = region
	})

	return &client{
		orgClient: s3Client,
		appLogger: appLogger,
	}
}

func (c *client) MakeBucket(ctx context.Context, bucketName string) error {
	_, err := c.orgClient.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return err
	}
	c.appLogger.Info(ctx,
		fmt.Sprintf("Bucket created, bucket name: %s", bucketName), "", time.Time{}, "")
	return nil
}

func (c *client) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	_, err := c.orgClient.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	exists := true
	if err != nil {
		var apiError *types.NotFound
		if errors.As(err, &apiError) {
			exists = false
			err = nil
		}
	}
	return exists, err
}

func (c *client) GetObject(ctx context.Context, bucketName string, objectName string) (io.ReadCloser, error) {
	output, err := c.orgClient.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectName),
	})
	return output.Body, err
}

func (c *client) PutObject(ctx context.Context,
	bucketName string, objectName string, reader io.Reader, contentType string) error {
	_, err := c.orgClient.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(objectName),
		Body:        reader,
		ContentType: aws.String(contentType),
	})
	return err
}
