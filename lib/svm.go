package cp

// Implementaytion of SVM machine learning

import (
 "github.com/ewalker544/libsvm-go"
 "fmt"
 "log"
 "strconv"
 "math"
)


func PredictFrom(inputfeatures map[int]float64,modelFile string) float64 {
  model := libSvm.NewModelFromFile(modelFile) //import model
  return model.Predict(inputfeatures) //get class from input
}

func Train(inputFilePath string, outputFilePath string)  {
  param := libSvm.NewParameter()      // Create a parameter object with default values
  param.KernelType = libSvm.POLY      // Use the polynomial kernel
  model := libSvm.NewModel(param)     // Create a model object from the parameter attributes
  // Create a problem specification from the training data and parameter attributes
  problem, err := libSvm.NewProblem(inputFilePath, param)
  if err != nil {
    fmt.Errorf("Fail to open file")
  }
  log.Println(problem.ProblemSize())
  model.Train(problem)                // Train the model from the problem specification
  model.Dump(outputFilePath)             // Dump the model into a user-specified file
}

func test(){

}

func CreateLabel(steering_angle float64, throttle float64, step float64) float64{
  //dec := steering_angle - math.Round(steering_angle)
  //bottomround := steering_angle-dec
  rounded_throttle := math.Round(throttle*8)/10
  steering_angle = steering_angle*100
  if steering_angle > 0 {

    return (math.Floor(steering_angle/step)*step)+rounded_throttle
  } else
  {
    return (math.Ceil(steering_angle/step)*step)-rounded_throttle
  }
}

func DecodeLabel(class float64) (float64,float64) {
  return math.Floor(class)/100,math.Round((class-math.Floor(class))*10)/10
}

func TrainEntry(class float64, features []uint8,speed float64) string {
  entry := fmt.Sprintf("%.1f",class)
  lastid := 0
  for id,value := range features {
    entry = entry+" "+strconv.Itoa(id+1)+":"+strconv.Itoa(int(value))
    lastid = id
  }
  entry = entry+" "+strconv.Itoa(lastid+2)+":"+fmt.Sprintf("%f",speed)
  return entry+"\n"
}
