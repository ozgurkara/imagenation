package storage

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"imagenation/src/configutil"
	"os"
	"path/filepath"
)

func DownloadImage(bucketName, itemKey string) string {
	fmt.Println(bucketName)
	fmt.Println(itemKey)
	if len(bucketName) == 0 || len(itemKey) == 0 {
		exitErrorf("Bucket and item names required\nUsage: %s bucket_name item_name", os.Args[0])
	}

	bucket := bucketName
	item := itemKey

	file, err := os.Create(configutil.GetTempFolderPath() + filepath.Base(item))
	if err != nil {
		exitErrorf("Unable to open file %q, %v", err)
	}

	defer file.Close()

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(configutil.GetConfig().Region)},
	)

	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	if err != nil {
		exitErrorf("Unable to download item %q, %v", item, err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

	return configutil.GetTempFolderPath() + filepath.Base(item)
}
