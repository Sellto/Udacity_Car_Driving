package action

import (
  "image"
  "fmt"
  "log"
  "os"
  _ "image/jpeg"
  "../lib"
  "strconv"
)






func GetHOGFeature(data cp.HOGdata, csvFile string,sample_pourcent int, test_pourcent float64,filename string){
  traindata := cp.GetDataFromCSV(csvFile)
  f,_ := os.Create(fmt.Sprintf(filename+".training"))
  t,_ := os.Create(fmt.Sprintf(filename+".testing"))
  defer f.Close()
  i := 0.0
  for id,value := range traindata {
    if id%(100/sample_pourcent) == 0 {
      infile, err := os.Open(value.ImagePath)
      if err != nil {
          panic(err)
      }
      src, _, err := image.Decode(infile)
      if err != nil {
        panic(err)
        infile.Close()
      }
      hog, err := cp.ComputeHoG4All(src, data)
			if err != nil {
				log.Fatal(err)
			}
      for _,v := range hog {
        if v != 0 && value.Steering_angle != 0 {
          i = i+float64(test_pourcent/100)
          if i < 1 {
            f.WriteString(cp.TrainEntryHOG(value.Steering_angle,hog))
          } else {
            t.WriteString(cp.TrainEntryHOG(value.Steering_angle,hog))
            i = 0.0
          }
          break
        }
      }
    }
  }
}



func GetGSFeature(csvFile string,sample_pourcent int, test_pourcent float64,filename string){
  traindata := cp.GetDataFromCSV(csvFile)
  f,_ := os.Create(fmt.Sprintf(filename+".training"))
  t,_ := os.Create(fmt.Sprintf(filename+".testing"))
  defer f.Close()
  i := 0.0
  for id,value := range traindata {
    if id%(100/sample_pourcent) == 0 {
      infile, err := os.Open(value.ImagePath)
      if err != nil {
          panic(err)
      }
      src, _, err := image.Decode(infile)
      if err != nil {
        panic(err)
        infile.Close()
      }
      listpix := cp.GetImageGSFeature(src)
      for _,v := range listpix {
        if v != 0 && value.Steering_angle != 0 {
          i = i+float64(test_pourcent/100)
          if i < 1 {
            f.WriteString(cp.TrainEntry(value.Steering_angle,listpix))
          } else {
            t.WriteString(cp.TrainEntry(value.Steering_angle,listpix))
            i = 0.0
          }
          break
        }
      }
    }
  }
}

func GetBWFeature(csvFile string,sample_pourcent int, test_pourcent float64,filename string){
  traindata := cp.GetDataFromCSV(csvFile)
  f,_ := os.Create(fmt.Sprintf(filename+".training"))
  t,_ := os.Create(fmt.Sprintf(filename+".testing"))
  defer f.Close()
  i := 0.0
  for id,value := range traindata {
    if id%(100/sample_pourcent) == 0 {
      infile, err := os.Open(value.ImagePath)
      if err != nil {
          panic(err)
      }
      src, _, err := image.Decode(infile)
      if err != nil {
        panic(err)
        infile.Close()
      }
      listpix := cp.GetImageBWFeature(src)
      for _,v := range listpix {
        if v != 0 && value.Steering_angle != 0 {
          i = i+float64(test_pourcent/100)
          if i < 1 {
            f.WriteString(cp.TrainEntry(value.Steering_angle,listpix))
          } else {
            t.WriteString(cp.TrainEntry(value.Steering_angle,listpix))
            i = 0.0
          }
          break
        }
      }
    }
  }
}


func GetCustGSFeature(csvFile string,sample_pourcent int, test_pourcent float64,filename string){
  traindata := cp.GetDataFromCSV(csvFile)
  f,_ := os.Create(fmt.Sprintf(filename+".training"))
  t,_ := os.Create(fmt.Sprintf(filename+".testing"))
  defer f.Close()
  i := 0.0
  for id,value := range traindata {
    if id%(100/sample_pourcent) == 0 {
      infile, err := os.Open(value.ImagePath)
      if err != nil {
          panic(err)
      }
      src, _, err := image.Decode(infile)
      if err != nil {
        panic(err)
        infile.Close()
      }
      listpix := cp.GetImageCustGSFeature(src)
      for _,v := range listpix {
        if v != 0 && value.Steering_angle != 0 {
          i = i+float64(test_pourcent/100)
          if i < 1 {
            f.WriteString(cp.TrainEntry(value.Steering_angle,listpix))
          } else {
            t.WriteString(cp.TrainEntry(value.Steering_angle,listpix))
            i = 0.0
          }
          break
        }
      }
    }
  }
}
//
// func GetPictures(csvFile string){
//   traindata := GetDataFromCSV(csvFile)
//   for id,value := range traindata {
//     infile, err := os.Open(value.ImagePath)
//       if err != nil {
//           panic(err)
//       }
//       src, _, err := image.Decode(infile)
//       if err != nil {
//         panic(err)
//         infile.Close()
//       }
//       GetRight(src,id)
//     }
// }



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
