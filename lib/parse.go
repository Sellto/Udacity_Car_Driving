package cp

import (
  "bufio"
  "encoding/csv"
  "os"
  "io"
  "strconv"
  "log"
)

type trainingData struct {
  imagePath string
  speed float64
  throttle float64
  steering_angle float64
}

func getDataFromCSV(file string) []trainingData {
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
    l = append(l, trainingData{imagePath: record[0],speed: speed, throttle: throttle, steering_angle: steering_angle})
	}
  
return l
}
