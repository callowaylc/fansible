package fansible

import(
  "os"
  "fmt"
  _ "bytes"
	flags "github.com/jessevdk/go-flags"
)

type Options struct {
  Inventory string `default:"/etc/ansible/host" short:"i" long:"inventory-file" description:"specify inventory host path"`
  Command string `default:"/etc/ansible/host" short:"c" long:"command" description:"wraps an ansible-playbook command"`
  Top bool `long:"apply-top" description:"applies all configured plays defined in the top file"`
  Logs bool `long:"logs" description:"enables logging by fansible application to stderr"`
}
var o Options

func InitOptions() {
  p := flags.NewParser(&o, flags.Default)
  _, err := p.Parse()
  
  if flagsErr, ok := err.(*flags.Error); ok {
    switch flagsErr.Type {
    case flags.ErrHelp:
      os.Exit(0)  
    case flags.ErrTag:
      panic(err)
    default:
      fmt.Println(err)
      os.Exit(2)
    }
  
  } else if err != nil {
    panic(err)
  }

  // in the event of passing no arguments, we print help
  // message and exit 0
  if len(os.Args) == 1 {
    p.WriteHelp(os.Stdout)
    fmt.Println()
    os.Exit(0)
  }  
}

func GetOptions() Options {
  return o
}
