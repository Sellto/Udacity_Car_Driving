package svm

import (
 "github.com/ewalker544/libsvm-go"
)


func predictFrom(inputfeatures []int,modelFile file) string {
  model := libSvm.NewModelFromFile(modelFile) //import model
  return model.Predict(features) //get class from input
}

func train(inputFilePath string, kernel int, outputFilePath sting)  {
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
