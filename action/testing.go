package action

import (
  "math"
  "../lib"
)

func GetMeanSD (inputfile,model string) (float64,float64) {
  target,features :=  cp.ParseTestingFile(inputfile)
  l := cp.PredictionError(model,target,features)
  sum := 0.0
  sd := 0.0
  for _,value := range l {
    sum += value
  }
  mean := sum/float64(len(l))
  for _,value := range l {
     sd += math.Pow(value - mean, 2)
  }
  sd = math.Sqrt(sd/float64(len(l)))
  return mean,sd
}
func GetPourcent (inputfile,model string) float64 {
  target,features :=  cp.ParseTestingFile(inputfile)
  l := cp.PredictionError(model,target,features)
  goodpred := 0
  for _,value := range l {
     if value == 0 {
       goodpred++
     }
  }
  return float64(goodpred/len(l))
}
