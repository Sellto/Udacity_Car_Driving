package cp

import (
  "image/color"
  //"image/draw"
  "math"
  //"encoding/base64"
  "image"
  //"strings"
  //"fmt"
  _"image/jpeg"
  "github.com/nfnt/resize"
  "github.com/oliamb/cutter"
  "image/png"
  "os"
  "fmt"
  "image/draw"
)

func displayImage(img image.Image,name string){
  outfile, err := os.Create(name)
  if err != nil {
    panic(err)
  }
  defer outfile.Close()
  png.Encode(outfile, img)
}


func GetImageFeature(img image.Image) []uint8 {
      croppedImg, _ := cutter.Crop(img, cutter.Config{
        Width:  320,
        Height: 40,
        Anchor: image.Point{0, 100},
        Options: cutter.Copy,
      })
      //displayImage(croppedImg,"test.png")
      processed_img := toGrayscale2(resize.Resize(16, 4, croppedImg, resize.NearestNeighbor))
      //displayImage(processed_img,"test3.png")
      return processed_img.Pix
}

func GetImageCropCenterFeature(img image.Image) []uint8 {
      m := image.NewRGBA(image.Rect(0, 0, 160, 40))
      pt := image.Pt(240,100)
      rect := image.Rect(80,0,320,40)
      draw.Draw(m, rect, img, pt, draw.Src)
      pt = image.Pt(0,100)
      rect = image.Rect(0,0,80,40)
      draw.Draw(m, rect, img, pt, draw.Src)
      processed_img := toGrayscale2(resize.Resize(8, 4, m, resize.NearestNeighbor))
      //displayImage(processed_img,"test2.png")
      return processed_img.Pix
}

func toGrayscale(img image.Image) *image.Gray {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	grayScale := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{w, h}})
	for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
					imageColor := img.At(x, y)
					rr, gg, bb, _ := imageColor.RGBA()
          fmt.Println(imageColor.RGBA())
					r := math.Pow(float64(rr), 2.2)
					g := math.Pow(float64(gg), 2.2)
					b := math.Pow(float64(bb), 2.2)
					m := math.Pow(0.2125*r+0.7154*g+0.0721*b, 1/2.2)
					Y := uint16(m + 0.5)
					grayColor := color.Gray{uint8(Y >> 8)}
					grayScale.Set(x, y, grayColor)
			}
		}
    return grayScale
}

func toGrayscale2(img image.Image) *image.Gray {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	grayScale := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{w, h}})
  White := color.Gray{uint8(255)}
  Black := color.Gray{uint8(0)}
  for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
					imageColor := img.At(x, y)
					rr, _, bb, _ := imageColor.RGBA()
          //mean := (rr+gg+bb)/3
          if rr > bb+5000 {
            grayScale.Set(x, y, White)
          } else {
            grayScale.Set(x, y, Black)
          }
			}
		}

    return grayScale
}
