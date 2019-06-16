package cp

import (
	  "image/color"
		"image"
		"math"
		"image/draw"
)



func ResetMatrix(x,y,o int) [][][]float64 {
matrix := make([][][]float64,y)
for i := range matrix {
    matrix[i] = make([][]float64, x)
    for j := range matrix[i]{
      matrix[i][j] = make([]float64, o)
      for k := range matrix[i][j] {
        matrix[i][j][k] = 0.0
      }
    }
}
return matrix
}

func BacktoMatrix(hog []float64, data HOGdata) [][][]float64{
cellnumx, cellnumy := data.ResizeX/data.CellSize, data.ResizeY/data.CellSize
matrix := ResetMatrix(cellnumx,cellnumy,data.Orientation)
idx := 0
for cy := 0 ; cy < cellnumy-data.BlockSize+1; cy++ {
  for cx := 0 ; cx < cellnumx-data.BlockSize+1; cx++ {
    for by := 0 ; by < data.BlockSize; by++ {
      for bx := 0 ; bx < data.BlockSize; bx++ {
        for o := 0 ; o < data.Orientation ;o++ {
          matrix[cy+by][cx+bx][o] += hog[idx]
          idx++
          }
        }
      }
    }
  }
  return matrix
}


func ShowCellHOG(v []float64) *image.RGBA{
  m := image.NewRGBA(image.Rect(0, 0, 40, 40))
  for k := 0 ; k < len(v) ; k++ {
      angle := math.Pi * (float64(k)/float64(len(v)))
      for r := -20; r < 20; r++ {
        x := float64(r)*math.Cos(angle)
        y := float64(r)*math.Sin(angle)
        m.Set(int(x)+20, int(y)+20, color.NRGBA64{uint16(65535),uint16(0),uint16(0),uint16(65535*v[k])})
      }
  }
return m
}




func Display (m [][][]float64,input image.Image) *image.RGBA{
  img := image.NewRGBA(image.Rect(0, 0, 40*len(m[0]), 40*len(m)))
  draw.Draw(img, img.Bounds(),input, image.Pt(0,0), draw.Src)
  for idj,i := range m {
    for idi,j := range i {
      hogCell := ShowCellHOG(j)
      draw.Draw(img, image.Rect(40*idi, 40*idj, 40+40*idi, 40+40*idj),hogCell, image.Pt(0,0), draw.Src)
    }
  }
  return img
}
