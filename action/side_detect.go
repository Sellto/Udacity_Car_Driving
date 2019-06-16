package action

import (
  "../lib"
  "os"
  "fmt"
  "image"
  "strconv"
)


func GetSideHOGFeatures(data cp.HOGdata, inputfolder, outputfile string, pourcent float64, sidenum, nonenum int) {
f,_ := os.Create(fmt.Sprintf(outputfile+".training"))
t,_ := os.Create(fmt.Sprintf(outputfile+".testing"))
defer f.Close()
defer t.Close()
j := 0.0
for i := 1 ; i < sidenum ; i++ {
  infile, err := os.Open(fmt.Sprintf(inputfolder+"side-"+strconv.Itoa(i)+".png"))
  if err != nil {
      panic(err)
  }
  src, _, err := image.Decode(infile)
  if err != nil {
    panic(err)
    infile.Close()
  }
  hog, err := cp.ComputeHoG4SideImg(src, data)
  j = j+float64(pourcent/100)
  if j < 1 {
      f.WriteString(cp.BinaryClass(1,hog))
  } else {
      t.WriteString(cp.BinaryClass(1,hog))
      j = 0.0
  }
}
  for i := 1 ; i < nonenum ; i++ {
    infile, err := os.Open(fmt.Sprintf(inputfolder+"none-"+strconv.Itoa(i)+".png"))
    if err != nil {
        panic(err)
    }
    src, _, err := image.Decode(infile)
    if err != nil {
      panic(err)
      infile.Close()
    }
    hog, err := cp.ComputeHoG4SideImg(src, data)
    j = j+float64(pourcent/100)
    if j < 1 {
        f.WriteString(cp.BinaryClass(-1,hog))
    } else {
        t.WriteString(cp.BinaryClass(-1,hog))
        j = 0.0
    }
}
}
