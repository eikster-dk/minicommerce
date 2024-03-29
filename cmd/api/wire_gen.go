// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"cloud.google.com/go/firestore"
	"context"
	firestore2 "github.com/eikc/minicommerce/pkg/firestore"
	"github.com/eikc/minicommerce/pkg/http"
	"github.com/eikc/minicommerce/pkg/storage"
	"google.golang.org/api/option"
)

// Injectors from wire.go:

func NewServer(ctx context.Context, bucketURL storage.BucketURL, projectID string, opts ...option.ClientOption) (*http.Server, error) {
	client, err := firestore.NewClient(ctx, projectID, opts...)
	if err != nil {
		return nil, err
	}
	downloadableService := firestore2.NewDownloadableService(client)
	storageStorage := storage.NewStorage(bucketURL)
	server := http.NewServer(downloadableService, storageStorage)
	return server, nil
}
