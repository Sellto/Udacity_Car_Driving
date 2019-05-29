package cp

// Implementaytion of SVM machine learning

import (
 "github.com/ewalker544/libsvm-go"
 "fmt"
 "strconv"
)


func PredictFrom(inputfeatures map[int]float64,modelFile string) float64 {
  model := libSvm.NewModelFromFile(modelFile) //import model
  return model.Predict(inputfeatures) //get class from input
}


func TrainEntry(class float64, features []uint8) string {
  entry := fmt.Sprintf("%.3f",class)
  for id,value := range features {
    entry = entry+" "+strconv.Itoa(id+1)+":"+strconv.Itoa(int(value))
  }
  return entry+"\n"
}
