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
  fmt.Print("Enter your project's name: ")
  proName, _ := inRd.ReadString('\n')
  proName = strings.TrimSuffix(proName, "\n") // Remove the newline off the end
  project.Name = proName
  fmt.Print("Enter your full project's path: ")
  proPath, _ := inRd.ReadString('\n')
  proPath = strings.TrimSuffix(proPath, "\n") // Remove the newline(\n) off the end
  if exists(proPath) {
    project.Path = proPath

    err := project.FilesInProject()
    if err != nil {
      fmt.Println("Error getting files in your project")
    }

    tFu, tFi, tL, err := project.FunctionsFromFiles()
    if err != nil {
      fmt.Println("Error getting functions in files")
    }

    project.Stats = &gomap.Stat{}
    project.Stats.FileCount = tFi
    project.Stats.FuncCount = tFu
    project.Stats.LineCount = tL

    cL := cli.NewICLInterface(&project)
    cL.Start()
  } else {
    fmt.Println("Path does not exist. Exiting...")
  }
}

//Thanks stackoverflow: http://stackoverflow.com/questions/10510691/how-to-check-whether-a-file-or-directory-denoted-by-a-path-exists-in-golang
func exists(path string) (bool) {
  _, err := os.Stat(path)
  if err == nil { return true }
  if os.IsNotExist(err) { return false }
  return false
}
