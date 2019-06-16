package cp

// Parse csv file from simulator
//    (Camera images + Steerangle + throttle + Speed )

import (
  "bufio"
  "encoding/csv"
  "os"
  "io"
  "strconv"
  "log"
  "strings"
  "fmt"
)

type trainingData struct {
  ImagePath string
  Speed float64
  Throttle float64
  Steering_angle float64
}

func GetDataFromCSV(file string) []trainingData {
  var l []trainingData
  csvFile, _ := os.Open(file)
  r := csv.NewReader(bufio.NewReader(csvFile))
  for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
    steering_angle,err := strconv.ParseFloat(record[3], 64)
    if err != nil {
        log.Println("Error when parse steering angle value")
    }
    throttle,err := strconv.ParseFloat(record[4], 64)
    if err != nil {
        log.Println("Error when parse throttle value")
    }
    speed,err := strconv.ParseFloat(record[6], 64)
    if err != nil {
        log.Println("Error when parse speed value")
    }
    l = append(l, trainingData{ImagePath: record[0],Speed: speed, Throttle: throttle, Steering_angle: steering_angle})
	}

return l
}

func ParseTestingFile(testset_filename string) ([]float64, []map[int]float64) {
  file, err := os.Open(testset_filename)
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  all_test_feature := []map[int]float64{}
  target := []float64{}
  for scanner.Scan() {
      test_feature := make(map[int]float64)
      for idv,value := range strings.Split(scanner.Text()," "){
          if idv > 0 {
            id , _ := strconv.Atoi(strings.Split(value,":")[0])
            val , _ := strconv.ParseFloat(strings.Split(value,":")[1], 64)
            test_feature[id] = val
          } else {
            d,_ := strconv.ParseFloat(value, 64)
            target = append(target,d)
          }
      }
      all_test_feature = append(all_test_feature,test_feature)
      fmt.Print(".")
  }
  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }
  fmt.Println("*")
  return target,all_test_feature
}
