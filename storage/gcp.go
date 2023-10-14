package storage

import (
	"context"

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

var CMSBucketHandle = StorageClient.Bucket("stride-cms")
