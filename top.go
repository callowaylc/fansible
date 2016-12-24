package fansible

import (
  _ "fmt"
  "io/ioutil"
  yaml "gopkg.in/yaml.v2"
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
