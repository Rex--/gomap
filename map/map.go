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
  Name       string
  Path       string
  Functions  map[FuncsKey]string
  Files      map[string]string
  Stats      *Stat
}

  // IProjectKey holds a filepath, filename, and a line number relating to each and
  //    every function in the project tree
type FuncsKey struct {
  Filename   string
  FilePath   string
  LineNumber int
}

  // Stat structure holds all the statistic information I want to keep(Or think of keeping)
type Stat struct {
  FileCount int
  FuncCount int
  LineCount int
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
func (project *Project)FunctionsFromFiles() (tFu, tFi, tL int, err error) {
  project.Functions = make(map[FuncsKey]string)
  for k, v := range project.Files {
    tFi++
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
      tL++
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
        tFu++
        project.Functions[FuncsKey{k, v, i}] = sLine
      }
    }
  }
  return
}
