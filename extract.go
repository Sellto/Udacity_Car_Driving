package main

import (
  "./lib"
  "image"
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
)

var csv = "/Users/selltom/driving_log.csv"


func getFeature(csvFile string,sample_pourcent int, test_pourcent float64){
  traindata := cp.GetDataFromCSV(csvFile)
  f,_ := os.Create("sample_steer.training")
  t,_ := os.Create("sample_steer.testing")
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
      listpix := cp.GetImageFeature(src)
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


func getTestFeature(filename string) {
  file, err := os.Open(filename)
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
      test_feature := []string{}
      for _,value := range strings.Split(scanner.Text(),":"){
          test_feature = append(test_feature,strings.Split(value," ")[0])
      }
      fmt.Println(test_feature)
  }
  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }
}

func main() {
//getFeature(csv,20,20.0)
getTestFeature("sample_steer.testing")
}
