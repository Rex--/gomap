package main

import (
  "os"
  "fmt"
  "bufio"
  "strings"
  "github.com/Rex--/gomap/cli"
  "github.com/Rex--/gomap/map"
  )


var project gomap.Project

func main() {
  // The empty project struct
  project = gomap.Project{}

  //Get the Project info from the user
  inRd := bufio.NewReader(os.Stdin)
  fmt.Print("Enter your project's name:")
  proName, _ := inRd.ReadString('\n')
  proName = strings.TrimSuffix(proName, "\n") // Remove the newline off the end
  project.Name = proName
  fmt.Print("Enter your full project's path:")
  proPath, _ := inRd.ReadString('\n')
  proPath = strings.TrimSuffix(proPath, "\n") // Remove the newline(\n) off the end
  project.Path = proPath

  err := project.FilesInProject()
  if err != nil {
    fmt.Println("Error getting files in your project")
  }

  err = project.FunctionsFromFiles()
  if err != nil {
    fmt.Println("Error getting functions in files")
  }

  cL := cli.NewICLInterface(&project)
  cL.Start()
}
