package main

import (
  "./lib"
  //"fmt"
  "image"
  "os"
  "fmt"
)




func main() {
    traindata := cp.GetDataFromCSV("/Users/selltom/Downloads/Udacity/data4/driving_log.csv")
    //fmt.Println(traindata)
    f,_ := os.Create("data4.training")
    defer f.Close()
    //infile, _ := os.Open(traindata[0].ImagePath)
    //defer infile.Close()
    //src1, _, _ := image.Decode(infile)
    //list1 := cp.GetImageCropCenterFeature(src1)
    //f.WriteString(cp.TrainEntry(1.0,list1))
    for _,value := range traindata{
      infile, err := os.Open(value.ImagePath)
      //defer infile.Close()
      if err != nil {
          panic(err)
      }

      src, _, err := image.Decode(infile)
      if err != nil {
        // replace this with real error handling
        panic(err)
      infile.Close()
      }

    listpix := cp.GetImageCropCenterFeature(src)
    //fmt.Println(value.Steering_angle)
    f.WriteString(cp.TrainEntry(cp.CreateLabel(value.Steering_angle,value.Throttle,2),listpix,value.Speed))

}
fmt.Println("Begin Train")
cp.Train("/Users/selltom/Desktop/ia/data4.training","/Users/selltom/Desktop/ia/data4.model_POLY")
}
