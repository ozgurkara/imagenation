package imageresizer

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	"imagenation/src/configutil"
	"log"
	"os"
	"path/filepath"
	"strings"
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
	var createdImageName = getImageName(imageName, size)

	file, err := os.Open(originalImagePath)
	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	m := resize.Resize(width, height, img, resize.Lanczos3)

	createdImagePath := configutil.GetTempFolderPath() + createdImageName
	out, err := os.Create(createdImagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new tmp to file
	jpeg.Encode(out, m, nil)

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
