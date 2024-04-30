package image

import (
	"github.com/disintegration/imaging"
	"log"
)

func Yasuo() {
	// Open a test image.
	src, err := imaging.Open("E:\\git-code\\a.jpg")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Resize the cropped image to width = 200px preserving the aspect ratio.
	src = imaging.Resize(src, 1000, 0, imaging.Lanczos)

	// Save the resulting image as JPEG.
	err = imaging.Save(src, "E:\\git-code\\a-yasuo01.jpg")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}
