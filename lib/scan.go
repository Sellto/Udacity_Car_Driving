package cp

// Implementaytion of SVM machine learning

import (
 //"github.com/ewalker544/libsvm-go"
 //"fmt"
 //"strconv"
 "image/draw"
 "image"
 //"os"
 //"fmt"
 "image/color"
)
//ScanImage(64,32,"center.jpg",false)

func ScanImage2(data HOGdata,size,overlap int, src image.Image,show bool,model string)[][]float64 {
  // infile, err := os.Open(img)
  // if err != nil {
  //     panic(err)
  // }
  // src, _, err := image.Decode(infile)
  // if err != nil {
  //   panic(err)
  //   infile.Close()
  // }

  prediction := make([][]float64, 1)
  for i := range prediction {
    prediction[i] = make([]float64, int(80/(size-overlap)))
}

  incoming := image.NewRGBA(image.Rect(0, 0, 80, size))
  draw.Draw(incoming, incoming.Bounds(), src, image.Pt(180,90), draw.Src)
  //rect := image.NewRGBA(image.Rect(0, 0, size, size))
  for y := 0 ; y < 1 ; y++ {
     for x := 0 ; x < int(80/(size-overlap)) ; x++ {
       emptypic := image.NewRGBA(image.Rect(0, 0, size, size))
       draw.Draw(emptypic, emptypic.Bounds(), incoming, image.Pt(x*(size-overlap),y*(size-overlap)), draw.Src)
       hog, _ := ComputeHoG4SideImg(emptypic, data)
       convert := make(map[int]float64)
       for id,value := range hog {
         convert[id] = value*1000
       }
       prediction[y][x] = PredictFrom(convert,model)
       if show == true {
         if prediction[y][x] == 1 {
           for ry := 0 ; ry < size ; ry++ {
             for rx := 0 ; rx < size ; rx++ {
               if rx < 1 || rx > size-2 || ry < 1 || ry > size-2 {
                 incoming.Set(rx+x*(size-overlap),ry+y*(size-overlap),color.RGBA{0,255,0,255})
               }
             }
           }
         } else if prediction[y][x] == -1 {
            for ry := 0 ; ry < size ; ry++ {
              for rx := 0 ; rx < size ; rx++ {
                if rx < 1 || rx > size-2 || ry < 1 || ry > size-2 {
                  incoming.Set(rx+x*(size-overlap),ry+y*(size-overlap),color.RGBA{255,0,0,255})
                }
              }
            }
          }
       }
     }
  }
  if show == true {
    DisplayImage(incoming,"test.png")
  }
  return prediction
}
func ScanImage(data HOGdata,size,overlap int, src image.Image,show bool,model string)[][]float64 {
  // infile, err := os.Open(img)
  // if err != nil {
  //     panic(err)
  // }
  // src, _, err := image.Decode(infile)
  // if err != nil {
  //   panic(err)
  //   infile.Close()
  // }

  prediction := make([][]float64, 1)
  for i := range prediction {
    prediction[i] = make([]float64, int(160/(size-overlap))-1)
}

  incoming := image.NewRGBA(image.Rect(0, 0, 160, size))
  draw.Draw(incoming, incoming.Bounds(), src, image.Pt(160,70), draw.Src)
  //rect := image.NewRGBA(image.Rect(0, 0, size, size))
  for y := 0 ; y < 1 ; y++ {
     for x := 0 ; x < int(160/(size-overlap))-1 ; x++ {
       emptypic := image.NewRGBA(image.Rect(0, 0, size, size))
       draw.Draw(emptypic, emptypic.Bounds(), incoming, image.Pt(x*(size-overlap),y*(size-overlap)), draw.Src)
       hog, _ := ComputeHoG4SideImg(emptypic, data)
       convert := make(map[int]float64)
       for id,value := range hog {
         convert[id] = value*1000
       }
       prediction[y][x] = PredictFrom(convert,model)
       if show == true {
         if prediction[y][x] == 1 {
           for ry := 0 ; ry < size ; ry++ {
             for rx := 0 ; rx < size ; rx++ {
               if rx < 1 || rx > size-2 || ry < 1 || ry > size-2 {
                 incoming.Set(rx+x*(size-overlap),ry+y*(size-overlap),color.RGBA{0,255,0,255})
               }
             }
           }
         } else if prediction[y][x] == -1 {
            for ry := 0 ; ry < size ; ry++ {
              for rx := 0 ; rx < size ; rx++ {
                if rx < 1 || rx > size-2 || ry < 1 || ry > size-2 {
                  incoming.Set(rx+x*(size-overlap),ry+y*(size-overlap),color.RGBA{255,0,0,255})
                }
              }
            }
          }
       }
     }
  }
  if show == true {
    DisplayImage(incoming,"test.png")
  }
  return prediction
}

func ScanImageALL(data HOGdata,size,overlap int, src image.Image,show bool,model string)[][]float64 {
  prediction := make([][]float64, int(160/(size-overlap)-1))
  for i := range prediction {
    prediction[i] = make([]float64, int(320/(size-overlap)-1))
  }
  incoming := image.NewRGBA(image.Rect(0, 0, 320, 160))
  draw.Draw(incoming, incoming.Bounds(), src, image.Pt(0,0), draw.Src)
  //rect := image.NewRGBA(image.Rect(0, 0, size, size))
  for y := 0 ; y < int(160/(size-overlap))-1; y++ {
     for x := 0 ; x < int(320/(size-overlap))-1 ; x++ {
       emptypic := image.NewRGBA(image.Rect(0, 0, size, size))
       draw.Draw(emptypic, emptypic.Bounds(), incoming, image.Pt(x*(size-overlap),y*(size-overlap)), draw.Src)
       hog, _ := ComputeHoG4SideImg(emptypic, data)
       convert := make(map[int]float64)
       for id,value := range hog {
         convert[id] = value*1000
       }
       prediction[y][x] = PredictFrom(convert,model)
       if show == true {
         if prediction[y][x] == 1 {
           for ry := 0 ; ry < size ; ry++ {
             for rx := 0 ; rx < size ; rx++ {
               if rx < 1 || rx > size-2 || ry < 1 || ry > size-2 {
                 incoming.Set(rx+x*(size-overlap),ry+y*(size-overlap),color.RGBA{0,255,0,255})
               }
             }
           }
         }
       }
     }
  }
  if show == true {
    DisplayImage(incoming,"test2.png")
  }
  return prediction
}
