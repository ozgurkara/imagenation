package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"imagenation/src/configutil"
	"imagenation/src/imageresizer"
	"imagenation/src/storage"
	"os"
	"path/filepath"
)

func HandleRequest(ctx context.Context, event storage.S3Event) (string, error) {
	itemKey := event.Records[0].S3.Object.Key
	if 0 >= len(itemKey) {
		return "item key is null", nil
	}

	generateImage(itemKey)

	return fmt.Sprintf("ItemKey %s!", itemKey), nil
}

func generateImage(itemKey string) {
	config := configutil.GetConfig()

	downloadedFilePath := storage.DownloadImage(config.Bucket, itemKey)
	createdImagePaths := imageresizer.CreateResizedImages(downloadedFilePath)

	for i := 0; i < len(createdImagePaths); i++ {
		line := createdImagePaths[i]
		fileName := filepath.Base(line)
		storage.UploadImage(line, config.StorageResizedImagesTarget+fileName)
	}

	removeImages(configutil.GetTempFolderPath())
}

func removeImages(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if configutil.IsEnvironmentAws() {
		lambda.Start(HandleRequest)
	} else {
		testGenerateImage()
	}
}

func testGenerateImage() {
	imageKeys := []string{
		"images/product/download.jpeg",
	}
	for i := 0; i < len(imageKeys); i++ {
		generateImage(imageKeys[i])
	}
}
