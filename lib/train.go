package cp

import (
 "github.com/ewalker544/libsvm-go"
 "fmt"
 "strconv"
)

func Train_C_SVC_POLY(inputFilePath string, degree int, gamma ,coef0 , eps , penality float64)  {
  param := libSvm.NewParameter()
  param.SvmType = libSvm.C_SVC
  param.KernelType = libSvm.POLY
  param.Degree = degree
  param.Gamma = gamma
  param.Coef0 = coef0
  param.Eps = eps
  param.C = penality

  model := libSvm.NewModel(param)
  problem, err := libSvm.NewProblem(inputFilePath, param)
  if err != nil {
    fmt.Errorf("Fail to open file")
  }

  model.Train(problem)
  outputfile := fmt.Sprintf("C_SVC_POLY_degree="+strconv.Itoa(degree)+
                            "_gamma="+strconv.FormatFloat(gamma, 'f', -1, 64)+
                            "_coef0="+strconv.FormatFloat(coef0, 'f', -1, 64)+
                            "_eps="+strconv.FormatFloat(eps, 'f', -1, 64)+
                            "_C="+strconv.FormatFloat(penality, 'f', -1, 64)+
                            ".model")
  model.Dump(outputfile)
}
