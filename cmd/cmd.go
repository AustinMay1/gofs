package cmd

import (
	"flag"
	
	"github.com/AustinMay1/gofs/project"
)

type Command struct {
	fs *flag.FlagSet  	// create
	name string		 	// Todo_CLI
	path string     	// /home/user/Dev/go_projs/name
	framework string 	// react
	ghName string      //  githubUser1
}

func(c *Command) Name() string {
	return c.fs.Name()
}

func(c *Command) Init(args []string) error {
	return c.fs.Parse(args)
}

func(c *Command) Run() error {
	project.CreateProject(c.path, c.name, c.ghName, c.framework)
	return nil
}

func Create() *Command {
	create := &Command {
		fs: flag.NewFlagSet("create", flag.ContinueOnError),
	}

	create.fs.StringVar(&create.name, "name", "GOFS-Project", "project name")
	create.fs.StringVar(&create.path, "dest", "./", "project destination")
	create.fs.StringVar(&create.framework, "type", "react", "frontend framework")
	create.fs.StringVar(&create.ghName, "username", "gh-user", "github username for go.mod")

	return create
}