package fansible
// http://squarism.com/2014/10/13/yaml-go/

import (
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
)

type Top struct {
  Environments map[string][]string
}

func InitTop() {

}