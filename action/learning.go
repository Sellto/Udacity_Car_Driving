package action

import (
  "github.com/ewalker544/libsvm-go"
  "fmt"
)


func Learning(inputFilePath,outputFile string, param *libSvm.Parameter) {
    model := libSvm.NewModel(param)
    problem, err := libSvm.NewProblem(inputFilePath, param)
    if err != nil {
      fmt.Errorf("Fail to open file")
    }
    model.Train(problem)
    model.Dump(outputFile)
  }
