#GO-PICTURE

This project create a picture with background image + avatar + textes with custom ttf font
Output is out.png

Avatar and texte will centerized in heigth or width by orientation

Multi line is allowed

## cli

with default value:
`go run main.go`
or 
`go build . && ./picture`

config:
```
Usage of ./picture:
  -a string
        path to the avatar picture (default "pictures/cat.jpg")
  -ao string
        operation apply on avatar picture ( circle ) (default "circle")
  -b string
        path to the background picture (default "pictures/bg.jpg")
  -f string
        path to the font to used (.ttf) (default "fonts/arial.ttf")
  -fs float
        size of font (default 48)
  -h int
        height of card (default 300)
  -ha int
        height of avatar; if picture greater, resize it; if 0 get height of picture (default 256)
  -ls float
        space between line (default 1)
  -o string
        orientation of card ( vertical | horizontal ) (default "horizontal")
  -op string
        path to the output picture (default "out.png")
  -p int
        padding betwwen border and picture and texte; if orientation is horizontal, padding left + right; if orientation is vertaical, padding is all direction (default 40)
  -w int
        width of card (default 1000)
  -wa int
        width of avatar; if picture greater, resize it; if 0 get width of picture (default 256)

```

## module

after import

```
import github.com/HuguesBt/go-picture/pkg/picture

# For create an horizontal picture
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

# For create a vertical picture
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
```
