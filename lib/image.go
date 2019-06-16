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
  "github.com/disintegration/imaging"
  "image/png"
  "os"
  //"fmt"
  //"image/draw"
  //"strconv"
)


// Display processed image
func DisplayImage(img image.Image,name string){
  outfile, err := os.Create(name)
  if err != nil {
    panic(err)
  }
  defer outfile.Close()
  png.Encode(outfile, img)
}

// Return list pixels of cropped camera image with custom grayscale
func GetImageCustGSFeature(img image.Image) []float64 {
      croppedImg, _ := cutter.Crop(img, cutter.Config{
        Width:  240,
        Height: 60,
        Anchor: image.Point{40, 80},
        Options: cutter.Copy,
      })
      processed_img := ToCustomGrayscale(resize.Resize(24, 6, croppedImg, resize.NearestNeighbor))
      //processed_img = ToGrayscale(imaging.Blur(processed_img, 1.5))
      pixel := []float64{}
      for _,value := range processed_img.Pix {
        pixel = append(pixel,float64(value))
      }
      return pixel
}

// Return list pixels of cropped camera image with standard grayscale
func GetImageGSFeature(img image.Image) []float64 {
      croppedImg, _ := cutter.Crop(img, cutter.Config{
        Width:  240,
        Height: 60,
        Anchor: image.Point{40, 80},
        Options: cutter.Copy,
      })
      processed_img := ToGrayscale(resize.Resize(24, 6, croppedImg, resize.NearestNeighbor))
      pixel := []float64{}
      for _,value := range processed_img.Pix {
        pixel = append(pixel,float64(value))
      }
      return pixel
}

// Return list pixels of cropped camera image with black and white value
func GetImageBWFeature(img image.Image) []float64 {
      croppedImg, _ := cutter.Crop(img, cutter.Config{
        Width:  240,
        Height: 60,
        Anchor: image.Point{40, 80},
        Options: cutter.Copy,
      })
      processed_img := ToBW(resize.Resize(24, 6, croppedImg, resize.NearestNeighbor))
      pixel := []float64{}
      for _,value := range processed_img.Pix {
        pixel = append(pixel,float64(value))
      }
      return pixel
}


// Return the image with black and white transformation (custom treshold)
func ToBW(img image.Image) *image.Gray {
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
          if rr < bb+5000 {
            grayScale.Set(x, y, White)
          } else {
            grayScale.Set(x, y, Black)
          }
			}
		}
    return grayScale
}


// Return the image with standrd grayscale transformation
func ToGrayscale(imgSrc image.Image) *image.Gray {
	bounds := imgSrc.Bounds()
    w, h := bounds.Max.X, bounds.Max.Y
    grayScale := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{w, h}})
    for x := 0; x < w; x++ {
        for y := 0; y < h; y++ {
            imageColor := imgSrc.At(x, y)
            rr, gg, bb, _ := imageColor.RGBA()
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

// Return the image with custom treshold grayscale transformation
func ToCustomGrayscale(img image.Image) *image.Gray {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	grayScale := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{w, h}})
  for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
					imageColor := img.At(x, y)
					rr, _, bb, _ := imageColor.RGBA()
          if  rr < bb + 5000{
            grayScale.Set(x, y, color.Gray{uint8(0)})
          } else if rr - bb > 10000 {
						grayScale.Set(x, y, color.Gray{uint8(0)})
					} else {
            grayScale.Set(x, y, color.Gray{uint8((rr-bb)*1000)})
          }
			}
		}
    return grayScale
}

// Return grayscale image adaptated for HOG processing
func ProcessHOGImg(img image.Image,data HOGdata ) *image.Gray {
      croppedImg, _ := cutter.Crop(img, cutter.Config{
        Width:  320,
        Height: 60,
        Anchor: image.Point{0, 80},
        Options: cutter.Copy,
      })
      processed_img := ToCustomGrayscale(resize.Resize(uint(data.ResizeX), uint(data.ResizeY), croppedImg, resize.MitchellNetravali))
      return ToGrayscale(imaging.Blur(processed_img, 1.5))
}
