package fansible
// http://squarism.com/2014/10/13/yaml-go/

import (
  _ "fmt"
  yaml "gopkg.in/yaml.v2"
  "io/ioutil"
)

const (
  file string = "./top.yml"
)

type Top map[string][]string
var t Top

func InitTop() {
  content, err := ioutil.ReadFile(file)
  if err != nil {
    panic(err)
  }

  if err = yaml.Unmarshal([]byte(content), &t); err != nil {
    panic(err)
  }
}

func GetTop() Top {
  return t
}
