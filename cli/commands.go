package cli

import (
  "strings"
  "fmt"
  )

type CommandCallBack func(cli *ICLInterface, command string) (bool)

type Command struct {
  Trigger string
  Description string
  Usage string
  CallBack CommandCallBack
}

func NewCommand(trigger, description, usage string, callback CommandCallBack) (command *Command){
  return &Command{
    Trigger: trigger,
    Description: description,
    Usage: usage,
    CallBack: callback,
  }
}

func getCommands() (cmds map[string]*Command) {
  cmds = map[string]*Command{}
  cmds[cmdHelp] = NewCommand(cmdHelp, descHelp, usageHelp, help)
  cmds[cmdFile] = NewCommand(cmdFile, descFile, usageFile, file)
  cmds[cmdFunction] = NewCommand(cmdFunction, descFunction, usageFunction, function)
  cmds[cmdStop] = NewCommand(cmdStop, descStop, usageStop, stop)
  return
}

// Start of command handling functions

const cmdHelp = "help"
const descHelp = "Lists all the commands."
const usageHelp = "help | help [command]"

func help(cli *ICLInterface, command string) (bool){
  args := strings.Split(command, " ")

  if len(args) > 2 {
    fmt.Println("Usage: ", usageHelp)
  }

  if len(args) == 2 {
    cmd := cli.cmds[args[1]]
    fmt.Println("Command: ", cmd.Trigger)
    fmt.Println("Description: ", cmd.Description)
    fmt.Println("Usage: ", cmd.Usage)
  }

  if len(args) == 1 {
    fmt.Println("Commands:")
    for cmd, _ := range cli.cmds {
      fmt.Println(" - ", cmd)
    }
  }
  return true
}


const cmdFile = "file"
const descFile = "Gets all files, or gets all the functions in file if you give it an argument"
const usageFile = "file | file [file.go]"

func file(cli *ICLInterface, command string) (bool){
  args := strings.Split(command, " ")

  if len(args) > 2 {
    fmt.Println("Usage: ", usageFile)
  }

  if len(args) == 1 {
    fmt.Println("Files:")
    for k, _ := range cli.Project.Files {
      fmt.Println(" -", k)
    }
  }

  if len(args) == 2 {
    fmt.Println("File", args[1])
    for k, v := range cli.Project.Functions {
      if k.Filename == args[1] {
        fmt.Println(" -", k.LineNumber, v)
      }
    }
  }
  return true
}



const cmdFunction  = "function"
const descFunction = "Gets all functions, or gets all the instances of the function in your project"
const usageFunction = "function | function [function name]"

func function(cli *ICLInterface, command string) (bool){
  args := strings.Split(command, " ")

  if len(args) > 2 {
    fmt.Println("Usage:", usageFunction)
  }

  if len(args) == 1 {
    fmt.Println("Functions:")
    for _, v := range cli.Project.Functions {
      fmt.Println(" -", v)
    }
  }

  if len(args) == 2 {
    for k, v := range cli.Project.Functions {
      if strings.Contains(v, args[1]) {
        fmt.Println(v)
        fmt.Println("  -Filename:", k.Filename)
        fmt.Println("  -Line Number:",k.LineNumber)
      }
    }
  }
  return true
}


const cmdStop = "stop"
const descStop = "Stops the program gracefully."
const usageStop = "stop"

func stop(cli *ICLInterface, command string) (bool){
  args := strings.Split(command, " ")

  if len(args) > 1 {
    fmt.Println(usageStop)
  }

  if len(args) == 1 {
    fmt.Println("Stopping gomap. Goodbye :D")
    return false
  }
  return true
}
