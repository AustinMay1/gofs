package project

import (
	"fmt"
	"os"
	"os/exec"
)

func CreateProject(path string, name string, username string, fw string) {
	err := exec.Command("mkdir", path+name).Run()

	if err != nil {
		fmt.Printf("cmd.Run: %s failed: %s\n", err, err)
	}

	os.Chdir(path + name)

	err = exec.Command("go", "mod", "init", "github.com/"+username+"/"+name).Run()

	if err == nil {
		fmt.Printf("Project created at: %s", path+name)
	} else {
		fmt.Printf("Error: %s\n%s", err, err)
	}

	err = exec.Command("touch", "main.go").Run()

	if fw == "react" {
		err = exec.Command("npm", "create", "vite@latest", "client", "--", "--template", "react-ts").Run()
	}

}
