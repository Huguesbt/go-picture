package picture

import (
	"github.com/fogleman/gg"
	"image"
)

func CreateHorizontalCard(
	bg string,
	avatar string,
	output string,
	W int,
	H int,
	P int,
	Wavatar int,
	Havatar int,
	avatarOperation string,
	font string,
	fontSize float64,
	lineSpacing float64,
	textes []string,
) {
	var (
		dc  *gg.Context
		im  image.Image
		Wim int
		Him int
	)

	dc = gg.NewContext(W, H)

	if bg != "" {
		imBg := getBackgroundPicture(bg, dc.Width(), dc.Height())
		dc.DrawImage(imBg, 0, 0)
	}

	if avatar != "" {
		im, Wim, Him = getAvatar(avatar, avatarOperation, Wavatar, Havatar)
		dc.DrawImage(im, P, (H-Him)/2)
	}

	dc.Stroke()
	dc.SetRGB(1, 1, 1)

	if font != "" {
		if err := dc.LoadFontFace(font, fontSize); err != nil {
			panic(err)
		}
	}

	if len(textes) > 0 {
		xInit := float64(2*P + Wim)
		countTxt := len(textes)
		widthTxt := float64(W - Wim - (3 * P))

		var (
			heightLine float64
			widthLine  float64
		)

		for _, txt := range textes {
			widthLine, heightLine = dc.MeasureMultilineString(txt, lineSpacing)
			if widthLine > widthTxt {
				countTxt += int(widthLine / widthTxt)
			}
		}

		heightTxts := float64(countTxt) * heightLine * lineSpacing
		margin := (float64(H) - heightTxts) / 2

		for i, txt := range textes {
			dc.DrawStringWrapped(
				txt,
				xInit,
				margin+(float64(i)*heightLine*lineSpacing),
				0,
				0,
				widthTxt,
				lineSpacing,
				gg.AlignCenter,
			)
		}
	}

	dc.SavePNG(output)
}
