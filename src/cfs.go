package main

import (
  "os"
  "github.com/codegangsta/cli"
  cmd "./cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "bcfs"
    app.HelpName = "bcfs"
    app.Version = "0.1 (github.com/stainboy/bcfs)"
    app.Author = "Miles Chen"
    app.Email = "stainboyx@hotmail.com"  
    app.Usage = "Manipulate baidu/weiyun cloud storage files"

    app.Commands = cmd.Commands

    app.Run(os.Args)
}