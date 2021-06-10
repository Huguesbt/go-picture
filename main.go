package main

import (
	"flag"
	"github.com/HuguesBt/go-picture/pkg/picture"
	"log"
)

var (
	orientation string
	texts       = []string{
		"Lorem ipsum",
		"Lorem",
		"Lorem ipsum dolor sit. Lorem ipsum dolor sit amet",
	}
	width           int
	height          int
	padding         int
	widthAvatar     int
	heightAvatar    int
	font            string
	fontSize        float64
	lineSpacing     float64
	background      string
	avatar          string
	output          string
	avatarOperation string
)

func init() {
	flag.StringVar(&orientation, "o", "horizontal", "orientation of card ( vertical | horizontal )")

	flag.IntVar(&width, "w", 1000, "width of card")
	flag.IntVar(&height, "h", 300, "height of card")
	flag.IntVar(&padding, "p", 40, "padding betwwen border and picture and texte; if orientation is horizontal, padding left + right; if orientation is vertaical, padding is all direction")

	flag.IntVar(&widthAvatar, "wa", 256, "width of avatar; if picture greater, resize it; if 0 get width of picture")
	flag.IntVar(&heightAvatar, "ha", 256, "height of avatar; if picture greater, resize it; if 0 get height of picture")

	flag.StringVar(&font, "f", "fonts/arial.ttf", "path to the font to used (.ttf)")
	flag.Float64Var(&fontSize, "fs", 48, "size of font")
	flag.Float64Var(&lineSpacing, "ls", 1, "space between line")

	flag.StringVar(&background, "b", "pictures/bg.jpg", "path to the background picture")
	flag.StringVar(&avatar, "a", "pictures/cat.jpg", "path to the avatar picture")
	flag.StringVar(&output, "op", "out.png", "path to the output picture")

	flag.StringVar(&avatarOperation, "ao", "circle", "operation apply on avatar picture ( circle )")

	flag.Parse()
}

func main() {
	if orientation == "horizontal" {
		picture.CreateHorizontalCard(
			background,
			avatar,
			output,
			width,
			height,
			padding,
			widthAvatar,
			heightAvatar,
			avatarOperation,
			font,
			fontSize,
			lineSpacing,
			texts,
		)
	} else if orientation == "vertical" {
		picture.CreateVerticalCard(
			background,
			avatar,
			output,
			width,
			height,
			padding,
			widthAvatar,
			heightAvatar,
			avatarOperation,
			font,
			fontSize,
			lineSpacing,
			texts,
		)
	} else {
		log.Println("Orientation not recognized; Choose between vertical or horizontal")
	}
}
