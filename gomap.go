package main

import (
  "os"
  "fmt"
  "bufio"
  "strings"
  "path/filepath"
  "github.com/Rex--/gomap/re-wwite/cli"
  )

// IProject interface holds "Functions", or a map of functions.
//  Base interface for each project
type IProject struct {
  Name string
  Path string
  Functions map[IProjectKey]string
}

// IProjectKey holds a filepath, filename, and a line number relating to each and
//    every function in the project tree
type IProjectKey struct {
  Filename   string
  FilePath   string
  LineNumber int
}

var project cli.IProject

func main() {
  // The empty project struct
  project = cli.IProject{}

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

  project.Functions = make(map[cli.FuncsKey]string)

  project.Files = make(map[string]string)

  err := filepath.Walk(project.Path, walkFunction)
  if err != nil {
    fmt.Println("Error walking directory:", err.Error())
  }

  cL := cli.NewICLInterface(&project)
  cL.Start()
}

// walkFunction called everytime the walf function comes across a file or directory
//  Basically just getFunctionsFromFiles() for every file while putting in in a map.
func walkFunction(path string, info os.FileInfo, err error) (error) {
  // Is a go source file
  if strings.HasSuffix(info.Name(), ".go") {
    getFunctionsFromFile(info.Name(), path)
    project.Files[info.Name()] = path
  }
  return nil
}


// getFunctionFromFiles takes a path as an argument and gets all the functions from
//    the file if it is good formatted go code.
func getFunctionsFromFile(filename, filepath string) {
  // Open source file and check to see if it could. Defer to close it.
  file, err := os.Open(filepath)
  if err != nil {
    fmt.Println("Error opening file:", err.Error())
  }
  defer func() {
    if err := file.Close(); err != nil {
      fmt.Println("Error closing file:", err.Error())
    }
  }()

  // Make a new reader for the file
  reader := bufio.NewReader(file)

  // Start at zero and count up each line until we get an "EOF" error that means its the
  //    *surprise**surprise* End of File!
  var bLine []byte
  var sLine string
  for c := 1; c != 0; c++ {
    bLine, _, err = reader.ReadLine()
    if err != nil {
      if err.Error() == "EOF"{
        break
      } else {
        fmt.Println("Error reading line:", err.Error())
      }
    }
    sLine = string(bLine[:])
    // Is a function
    if strings.HasPrefix(sLine, "func") {
      project.Functions[cli.FuncsKey{filename, filepath, c}] = sLine
    }
  }
}
