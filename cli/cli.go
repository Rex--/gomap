package cli

import (
  "bufio"
  "os"
  "fmt"
  "strings"
  )

type IProject struct {
  Name string
  Path string
  Functions map[FuncsKey]string
  Files map[string]string
}

type FuncsKey struct {
  Filename   string
  FilePath   string
  LineNumber int
}

type ICLInterface struct {
  Project *IProject
  cmds map[string]*Command
}

func NewICLInterface(project *IProject) (cli *ICLInterface) {
  c := getCommands()
  cli = &ICLInterface{
    Project: project,
    cmds: c,
  }
  return
}

func (cli *ICLInterface)Start() {
  reader := bufio.NewReader(os.Stdin)
  cntnu := true
  var command string
  for {
    command = prompt(reader)
    cntnu = cli.evaluate(command)
    if cntnu == false {
      break
    }
  }
}

func prompt(reader *bufio.Reader) (command string) {
  fmt.Print("> ")
  command, _ = reader.ReadString('\n')
  command = strings.TrimSuffix(command, "\n")
  return
}

func (cli *ICLInterface)evaluate(command string) (cntnu bool){
  commandList := strings.Split(command, " ")
  trigger := commandList[0]
  cmd, ok := cli.cmds[trigger]
  if !ok {
    fmt.Println("Unknown Command.")
    cntnu = true
  } else {
    cntnu = cmd.CallBack(cli, command)
  }
  return
}
