package main

import (
    "errors"
    "fmt"
    "os"

    "github.com/AustinMay1/gofs/cmd"
)

type Runner interface {
    Init([]string) error
    Run() error
    Name() string
}

func root(args []string) error {
    if len(args) < 1 {
        return errors.New("you must pass a sub-command")
    } 

    cmds := []Runner {
        cmd.Create(),
    }

    subcommand := os.Args[1]

    for _, cmd := range cmds {
        if cmd.Name() == subcommand {
            cmd.Init(os.Args[2:])
            return cmd.Run()
        }
    }

    return fmt.Errorf("unknown command: %s", subcommand)
}

func main() {
    if err := root(os.Args[1:]); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
