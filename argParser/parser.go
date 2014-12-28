package parser

import (
  "flags"
  )

const cmdHelp = "help"
const descHelp = "Lists all the commands."
const usageHelp = "help | help [command]"

const cmdFile = "file"
const descFile = "Gets all files, or gets all the functions in file if you give it an argument"
const usageFile = "[file.go]"

const cmdFunction  = "function"
const descFunction = "Gets all functions, or gets all the instances of the function in your project"
const usageFunction = "[function name]"

const cmdAdd = "add"
const descAdd = "Adds a project to gomap"
const usageAdd = "[/path/to/project]"


func ParseArgs() (args map[string]string) {
  args := make(map[string]string)
  file :=     flags.String(cmdFile, usageFile, descFile)
  function := flags.String(cmdFunction, usageFunction, descFunction)
  add :=      flags.String(cmdAdd, usageAdd, descAdd)
  if *file != usageFile {
    args[cmdFile] = *file
  }
  if *function != usageFunction {
    args[cmdFunction] = *function
  }
  if *add != usageAdd {
    args[cmdAdd] = *add
  }
  return
}
