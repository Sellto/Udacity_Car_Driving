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
