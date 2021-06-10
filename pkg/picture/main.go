package picture

import (
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/hawx/img/crop"
	"hawx.me/code/img/utils"
	"image"
	"log"
	"os"
)

func getSizePicture(filePath string) (width int, height int) {
	if reader, err := os.Open(filePath); err == nil {
		defer reader.Close()
		im, _, err := image.DecodeConfig(reader)
		if err != nil {
			return
		}
		return im.Width, im.Height
	}
	return
}

func openPicture(path string) image.Image {
	im, err := gg.LoadImage(path)
	if err != nil {
		log.Fatal(err)
	}
	return im
}

func getCirclePicture(image image.Image, diameter int) image.Image{
	return crop.Circle(image, diameter, utils.Centre)
}

func getAvatar(path string, operation string, width int, height int) (image.Image, int, int) {
	im := openPicture(path)
	Wim, Him := getSizePicture(path)
	if width > 0 && height > 0 {
		im = imaging.Fill(im, width, height, imaging.Center, imaging.Lanczos)
		Wim = width
		Him = height
	}

	if operation == "circle" {
		diameter := Wim
		if Wim > Him {
			diameter = Him
		}

		im = getCirclePicture(im, diameter)
	}

	return im, Wim, Him
}

func getBackgroundPicture(path string, width int, height int) image.Image {
	im := openPicture(path)
	return imaging.Fill(im, width, height, imaging.Center, imaging.Lanczos)
}
