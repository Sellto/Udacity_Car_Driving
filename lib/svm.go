package cp

// Implementaytion of SVM machine learning

import (
 "github.com/ewalker544/libsvm-go"
 "fmt"
 "strconv"
)




func TrainEntry(class float64, features []float64) string {
  entry := fmt.Sprintf("%.0f",class*1000)
  for id,value := range features {
    if value != 0 {
      entry = entry+" "+strconv.Itoa(id+1)+":"+strconv.Itoa(int(value))
    }
  }
  return entry+"\n"
}

func TrainEntryBinaryClass(class float64, features []float64) string {
  entry := "-1"
  if class > 0 {
      entry =  "1"
  }
  for id,value := range features {
    if value != 0 {
      entry = entry+" "+strconv.Itoa(id+1)+":"+strconv.Itoa(int(value))
    }
  }
  return entry+"\n"
}

func BinaryClass(class float64, features []float64) string {
  entry := "-1"
  if class > 0 {
      entry =  "+1"
  }
  for id,value := range features {
    if value != 0 {
      entry = entry+" "+strconv.Itoa(id+1)+":"+strconv.Itoa(int(value*1000))
    }
  }
  return entry+"\n"
}

func TrainEntryHOG(class float64, features []float64) string {
  entry := fmt.Sprintf("%.0f",class*1000)
  for id,value := range features {
    if value != 0 {
      entry = entry+" "+strconv.Itoa(id+1)+":"+strconv.Itoa(int(value*1000))
    }
  }
  return entry+"\n"
}

func PredictFrom(inputfeatures map[int]float64,modelFile string) float64 {
  model := libSvm.NewModelFromFile(modelFile) //import model
  return model.Predict(inputfeatures) //get class from input
}

func PredictionError(model_filename string,target []float64, all_test_feature []map[int]float64 ) []float64 {
  error := []float64{}
  for id,value := range all_test_feature {
    prediction := PredictFrom(value,model_filename)
    error = append(error,target[id]-prediction)
    fmt.Print(".")
  }
  fmt.Println("*")
  return error
}
