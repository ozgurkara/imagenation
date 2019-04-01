package imageresizer

import (
	"log"
	"path/filepath"
	"store-image-resizer/src/configutil"
	"strings"

	"github.com/disintegration/imaging"
)

func CreateResizedImages(imagePath string) []string {
	var result []string
	sizes := getImageSizes()
	for i := 0; i < len(sizes); i++ {
		line := sizes[i]
		createdImagePath := resizeImage(line.size, line.width, line.height, filepath.Base(imagePath), imagePath)
		result = append(result, createdImagePath)
	}

	return result
}

func resizeImage(size string, width, height uint, imageName string, originalImagePath string) string {
	// Open a test image.
	src, err := imaging.Open(originalImagePath)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Resize the cropped image to width = 200px preserving the aspect ratio.
	src = imaging.Resize(src, int(width), 0, imaging.Lanczos)

	// Create a blurred version of the image.
	var createdImageName = getImageName(imageName, size)
	createdImagePath := configutil.GetTempFolderPath() + createdImageName

	// Save the resulting image as JPEG.
	err = imaging.Save(src, createdImagePath)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

	return createdImagePath
}

func getImageName(imageName, size string) string {
	nameWithOutExtension := strings.TrimSuffix(imageName, filepath.Ext(imageName))
	var createdImageName = nameWithOutExtension + "_" + size + filepath.Ext(imageName)

	return createdImageName
}

func getImageSizes() []imageSize {
	imageSizes := []imageSize{
		{size: "xs", width: 120, height: 82},
		{size: "s", width: 240, height: 164},
		{size: "m", width: 360, height: 246},
	}

	return imageSizes
}

type imageSize struct {
	size   string
	width  uint
	height uint
}
