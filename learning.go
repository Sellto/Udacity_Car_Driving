package main

import (
  "./lib"
  //"image"
  //"os"
  //"fmt"
  //"time"
  //"log"
)

var trainingfile  = "/Users/selltom/OneDrive - Haute Ecole Léonard de Vinci/Ecam/5EO/ia/sample_steer.training"

func main() {
//Train_C_SVC_POLY(inputFilePath string, degree int, gamma int,coef0 int, eps float64, penality float64)
cp.Train_C_SVC_POLY(trainingfile,3,0.0,0,1e-3,0.1)
}
