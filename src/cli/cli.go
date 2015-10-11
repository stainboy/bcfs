package cli

import (
  "github.com/codegangsta/cli"
)

var Commands = []cli.Command {
    {
        Name: "login",
        Flags: []cli.Flag {
            cli.StringFlag {
                Name:  "oauth, o",
                Usage: "use weibo|qq|renren login",
            },
        },
        Usage:  "Login to your personal cloud storage",
        Action: login,
    },
    {
        Name: "ls",
        Flags: []cli.Flag {
            cli.BoolFlag {
                Name:  "recursive, R",
                Usage: "list subdirectories recursively",
            },
        },
        Usage:  "List information about the FILEs",
        Action: ls,
    },
    {
        Name: "mkdir",
        Flags: []cli.Flag {
            cli.BoolFlag {
                Name:  "parents, p",
                Usage: "no error if existing, make parent directories as needed",
            },
        },
        Usage:  "Create the DIRECTORY(ies), if they do not already exist",
        Action: mkdir,
    },
}