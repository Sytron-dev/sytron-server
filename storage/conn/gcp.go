package conn

import (
	"context"
	"os"

	"cloud.google.com/go/storage"
	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/api/option"
)

func initGCP() *storage.Client {
	client, err := storage.NewClient(context.TODO(), option.WithCredentialsFile("gcp_keys.json"))
	if err != nil {
		log.Error(err)
	}

	return client
}

var StorageClient = initGCP()

var (
	CMSBucketName   = os.Getenv("CMS_BUCKET")
	CMSBucketHandle = StorageClient.Bucket(CMSBucketName)
)
