package main

import (
  //"./lib"
  //"image"
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
)

func getTestFeature(filename string) [][]string {
  file, err := os.Open(filename)
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  all_test_feature := [][]string{}
  for scanner.Scan() {
      test_feature := []string{}
      for _,value := range strings.Split(scanner.Text(),":"){
          test_feature = append(test_feature,strings.Split(value," ")[0])
      }
      all_test_feature = append(all_test_feature,test_feature)
  }
  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }
  return all_test_feature
}



func main() {
alltest := getTestFeature("sample_steer.testing")
fmt.Println(alltest)
}
