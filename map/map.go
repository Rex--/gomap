package gomap

import (
  "bufio"
  "os"
  "strings"
  "fmt"
  "path/filepath"
  )

  // Project interface holds "Functions", or a map of functions.
  //  Base interface for each project
type Project struct {
  Name string
  Path string
  Functions map[FuncsKey]string
  Files map[string]string
}

  // IProjectKey holds a filepath, filename, and a line number relating to each and
  //    every function in the project tree
type FuncsKey struct {
  Filename   string
  FilePath   string
  LineNumber int
}


// walkFunction called everytime the walf function comes across a file or directory
//  Basically just getFunctionsFromFiles() for every file while putting in in a map.
// define a map to keep the files so we can return them.
var tempFiles map[string]string

func walkFunction(path string, info os.FileInfo, err error) (rerr error) {
  if err != nil {
    return err
  }

  // Is a go source file
  if strings.HasSuffix(info.Name(), ".go") {
    tempFiles[info.Name()] = path
  }
  return nil
}

func (project *Project)FilesInProject() (err error) {
  tempFiles = make(map[string]string)
  err = filepath.Walk(project.Path, walkFunction)
  if err != nil {
    fmt.Println("Error walking directory:", err.Error())
    return
  }
  project.Files = tempFiles
  return
}


// getFunctionFromFiles takes a path as an argument and gets all the functions from
//    the file if it is good formatted go code.
func (project *Project)FunctionsFromFiles() (err error) {
  project.Functions = make(map[FuncsKey]string)
  for k, v := range project.Files {
    file, err := os.Open(v)
    if err != nil {
      fmt.Println("Error opening file:", err.Error())
    }
    defer func() {
      if err := file.Close(); err != nil {
        fmt.Println("Error closing file:", err.Error())
      }
    }()

    reader := bufio.NewReader(file)

    var bLine []byte
    var sLine string
    for i := 1; i != 0; i++ {
      bLine, _, err = reader.ReadLine()
      if err != nil {
        if err.Error() == "EOF" {
          break
        } else {
          fmt.Println("Error reading file:", err.Error())
        }
      }
      sLine = string(bLine[:])

      // There is a function being declard on this line
      if strings.HasPrefix(sLine, "func") {
        project.Functions[FuncsKey{k, v, i}] = sLine
      }
    }
  }

  // Open source file and check to see if it could. Defer to close it.
  // file, err := os.Open(filepath)
  // if err != nil {
  //   fmt.Println("Error opening file:", err.Error())
  // }
  // defer func() {
  //   if err := file.Close(); err != nil {
  //     fmt.Println("Error closing file:", err.Error())
  //     }
  // }()
  //
  // // Make a new reader for the file
  // reader := bufio.NewReader(file)
  //
  // // Start at zero and count up each line until we get an "EOF" error that means its the
  // //    *surprise**surprise* End of File!
  // var bLine []byte
  // var sLine string
  // for c := 1; c != 0; c++ {
  // bLine, _, err = reader.ReadLine()
  // if err != nil {
  //   if err.Error() == "EOF"{
  //     break
  //     } else {
  //       fmt.Println("Error reading line:", err.Error())
  //     }
  //   }
  //   sLine = string(bLine[:])
  //     // Is a function
  //   if strings.HasPrefix(sLine, "func") {
  //     project.Functions[gomap.FuncsKey{filename, filepath, c}] = sLine
  //   }
  // }
  return nil
}
