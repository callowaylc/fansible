package main

import(
  _ "fmt"
  a "github.com/callowaylc/fansible"
)

func init() {
  a.InitOptions()
  a.InitLogs()
  a.InitTop()
}

func main() {
  top := a.GetTop()
  for target, _ := range top {
    a.Logs("process target to merge roles", a.LogFields{
      "target": target,
    })
  }
}