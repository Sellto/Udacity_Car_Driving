package cp

// Implementaytion of SVM machine learning 

import (
 "github.com/ewalker544/libsvm-go"
 "fmt"
 "log"
 "strconv"
)


func predictFrom(inputfeatures map[int]float64,modelFile string) float64 {
  model := libSvm.NewModelFromFile(modelFile) //import model
  return model.Predict(inputfeatures) //get class from input
}

func train(inputFilePath string, kernel int, outputFilePath string)  {
  param := libSvm.NewParameter()      // Create a parameter object with default values
  param.KernelType = kernel     // Use the polynomial kernel
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

func CreateLabel(steering_angle float64, throttle float64) string{

  return strconv.FormatFloat(steering_angle, 'E', -1, 64) + "," + strconv.FormatFloat(throttle, 'E', -1, 64)
}
