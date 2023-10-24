package storage

import (
	"context"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"

	"sytron-server/helpers/logger"
)

func initGCP() *storage.Client {
	client, err := storage.NewClient(context.TODO(), option.WithCredentialsFile("gcp_keys.json"))
	if err != nil {
		logger.Handle(err, "Creating GCP client")
	}

	return client
}

var StorageClient = initGCP()

var cmsBucketName = os.Getenv("CMS_BUCKET")
var CMSBucketHandle = StorageClient.Bucket(cmsBucketName)
