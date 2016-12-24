package main

import(
  _ "fmt"
  _ "io/ioutil"
  a "github.com/callowaylc/fansible"
)

func init() {
  a.InitOptions()
  a.InitLogs()
  a.InitTop()
}

func main() {
  top := a.GetTop()
  for target, plays := range top {
    a.Logs("process target to merge roles", a.LogFields{
      "target": target,
      "plays": plays,
    })


  }
}