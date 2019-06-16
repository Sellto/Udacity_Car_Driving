package cp

import (
  "math"
  "image"
  //"fmt"
)

type HOGdata struct {
  ResizeX int
  ResizeY  int
  CellSize int
  BlockSize  int
  Epsilon  float64
  Orientation int
}

type Cell struct {
	Hist []float64
}

func NewCell(data HOGdata) Cell {
	cell := Cell{}
	cell.Hist = make([]float64, data.Orientation)
	return cell
}


func computeCellHist(grayImg *image.Gray,data HOGdata) [][]Cell {
	fu, fv := 0.0, 0.0
	cells := make([][]Cell, (data.ResizeX / data.CellSize))
	for ci := 0; ci < (data.ResizeX / data.CellSize); ci++ {
		cells[ci] = make([]Cell, (data.ResizeY / data.CellSize))
	}
	for cy := 0; cy < (data.ResizeY / data.CellSize); cy++ {
		for cx := 0; cx < (data.ResizeX / data.CellSize); cx++ {
			cells[cx][cy] = NewCell(data)
		}
	}

	for y := 0; y < data.ResizeY; y++ {
		for x := 0; x < data.ResizeX; x++ {
			stride := grayImg.Stride
			if x == 0 {
				fu = float64(int(grayImg.Pix[(y*stride)+(x+1)]) - int(grayImg.Pix[(y*stride)+(x)]))
			} else if x == data.ResizeX-1 {
				fu = float64(int(grayImg.Pix[(y*stride)+(x)]) - int(grayImg.Pix[(y*stride)+(x-1)]))
			} else {
				fu = float64(int(grayImg.Pix[(y*stride)+(x+1)]) - int(grayImg.Pix[(y*stride)+(x-1)]))

			}

			if y == 0 {
				fv = float64(int(grayImg.Pix[((y+1)*stride)+x]) - int(grayImg.Pix[((y)*stride)+x]))
			} else if y == data.ResizeY-1 {
				fv = float64(int(grayImg.Pix[((y)*stride)+x]) - int(grayImg.Pix[((y-1)*stride)+x]))
			} else {
				fv = float64(int(grayImg.Pix[((y+1)*stride)+x]) - int(grayImg.Pix[((y-1)*stride)+x]))
			}

			m := math.Sqrt(fu*fu + fv*fv)
			theta := 0.0
			if fu != 0.0 || fv != 0.0 {
				theta = (math.Atan(fv/fu) * 180.0 / math.Pi) + 90// 0 - 180.0
			}
			// Cell histogram
			bin := int(theta / float64(180/data.Orientation))
			if bin == data.Orientation {
				bin -= 1
			}

			cells[int(x/data.CellSize)][int(y/data.CellSize)].Hist[bin] += m
		}

	}

	return cells
}

func computeBlockNorm(cells [][]Cell,data HOGdata) []float64 {
	hogidx := 0
	cellnumx, cellnumy := data.ResizeX/data.CellSize, data.ResizeY/data.CellSize
	blocknumx, blocknumy := (data.ResizeX/data.CellSize)-data.BlockSize+1, (data.ResizeY/data.CellSize)-data.BlockSize+1
	hog := make([]float64, blocknumx*blocknumy*data.BlockSize*data.BlockSize*data.Orientation)
	for cy := 0; cy < cellnumy; cy++ {
		for cx := 0; cx < cellnumx; cx++ {
			if cx+data.BlockSize-1 >= cellnumx || cy+data.BlockSize-1 >= cellnumy {
				continue
			}

			v := blockL2Norm(cx, cy, cellnumx, cells,data)
			// block normalization
			for iny := 0; iny < data.BlockSize; iny++ {
				for inx := 0; inx < data.BlockSize; inx++ {
					for b := 0; b < data.Orientation; b++ {
						val := cells[cx][cy].Hist[b]
						hog[hogidx] = val / math.Sqrt(v+data.Epsilon*data.Epsilon)
						hogidx++
					}
				}
			}
		}
	}

	return hog
}

func blockL2Norm(cx, cy, cellnumx int, cells [][]Cell, data HOGdata) float64 {
	v := 0.0
	for iny := 0; iny < data.BlockSize ; iny++ {
		for inx := 0; inx < data.BlockSize ; inx++ {
			for b := 0; b < data.Orientation ; b++ {
				val := cells[cx+inx][cy+iny].Hist[b]
				v += (val * val)
			}
		}
	}

	return v
}


func ComputeHoG4All(img image.Image,data HOGdata) ([]float64, error) {
	grayImg := ProcessHOGImg(img,data)
	cells := computeCellHist(grayImg,data)
	hog := computeBlockNorm(cells,data)
	return hog, nil
}

func ComputeHoG4SideImg(img image.Image,data HOGdata) ([]float64, error) {
	grayImg := ToGrayscale(img)
	cells := computeCellHist(grayImg,data)
	hog := computeBlockNorm(cells,data)
	return hog, nil
}
