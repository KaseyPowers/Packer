package utils;

import (
  "strings"
)

func getCommandHere(args ...string) []string{
  dir := CurrentDir()
  return getCommand(dir, args...)
}

func getCommand(dir string, args ...string) []string{
  output := []string{
    "git",
    "--git-dir=" + dir + "/.git/",
  }
  return append(output, args...)
}

func outputToLines(input []byte) []string {
  rawString := string(input)
  return strings.Split(rawString, "\n")
}