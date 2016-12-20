package main

import(
  "fmt"
  a "github.com/callowaylc/fansible"
)

func init() {
  a.InitOptions()
  a.InitTop()
}

func main() {
  top := a.GetTop()
  fmt.Println(top)
}