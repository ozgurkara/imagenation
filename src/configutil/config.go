package configutil

import (
	"encoding/json"
	"fmt"
	"os"
)

const ProductionConfig = "config.json"

type Config struct {
	Region                     string `json:"region"`
	Bucket                     string `json:"bucket"`
	StorageResizedImagesTarget string `json:"storage_resized_images_target"`
}

func IsEnvironmentAws() bool {
	isAws := os.Getenv("aws_environment")

	if len(isAws) <= 0 {
		return false
	}

	return true
}

func GetConfig() Config {
	if IsEnvironmentAws() {
		return loadConfigurationAwsLambda()
	} else {
		return loadConfiguration(ProductionConfig)
	}
}

func GetTempFolderPath() string {
	path := "tmp/"
	if IsEnvironmentAws() {
		return "/" + path
	}

	return path
}

func loadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}

func loadConfigurationAwsLambda() Config {
	region := os.Getenv("region")
	bucket := os.Getenv("bucket")
	storageResizedImagesTarget := os.Getenv("storage_resized_images_target")

	return Config{
		Region:                     region,
		Bucket:                     bucket,
		StorageResizedImagesTarget: storageResizedImagesTarget,
	}
}
