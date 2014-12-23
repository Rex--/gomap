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
  cmds[cmdExit] = NewCommand(cmdExit, descExit, usageExit, exit)
  cmds[exitAlias] = NewCommand(exitAlias, descExit, usageExit, exit)
  cmds[cmdStat] = NewCommand(cmdStat, descStat, usageStat, stat)
  cmds[cmdSearch] = NewCommand(cmdSearch, descSearch, usageSearch, search)
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
      fmt.Println(" -", cmd)
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


const cmdExit = "exit"
const descExit = "Exits the program gracefully."
const usageExit = "exit | quit"
const exitAlias = "quit"

func exit(cli *ICLInterface, command string) (bool){
  args := strings.Split(command, " ")

  if len(args) > 1 {
    fmt.Println("Usage:", usageExit)
  }

  if len(args) == 1 {
    fmt.Println("Stopping gomap. Goodbye :D")
    return false
  }
  return true
}

const cmdStat = "stat"
const descStat = "Gets the stats about your project"
const usageStat = "stat"

func stat(cli *ICLInterface, command string) (bool) {
  args := strings.Split(command, " ")

  if len(args) > 1 {
    fmt.Println("Usage:", usageStat)
  }

  if len(args) == 1 {
    fmt.Println("Project: ", cli.Project.Name)
    fmt.Println("Path to Project: ", cli.Project.Path)
    fmt.Println("Total files in project: ", cli.Project.Stats.FileCount)
    fmt.Println("Total functions in project: ", cli.Project.Stats.FuncCount)
    fmt.Println("Total line count in all: ", cli.Project.Stats.LineCount)
  }
  return true
}


const cmdSearch = "search"
const descSearch = "Searches for the string in all files and prints out the location."
const usageSearch = "search [string]"

func search(cli *ICLInterface, command string) (bool) {
  args := strings.Split(command, " ")

  if len(args) != 2 {
    fmt.Println("Usage:", usageSearch)
  }

  if len(args) == 2 {
    lines, err := cli.Project.SearchInFiles(args[1])
    if err != nil {
      fmt.Println("Error:", err.Error())
    }
    for k, v := range lines {
      fmt.Println(v)
      fmt.Println("  -Filename:", k.Filename)
      fmt.Println("  -Line Number:", k.LineNumber)
    }
  }
  return true
}
