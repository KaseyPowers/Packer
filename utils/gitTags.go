package utils;

import (
  "os/exec"
)

func GitTags(min, max string) []*Tag {
  cmd := getCommandHere("ls-remote","--tags")
  tags, err := exec.Command(cmd[0], cmd[1:]...).Output()
  if err != nil {
    Error(err)
  }
  return MakeTags(outputToLines(tags), min, max)
}

func MakeTags(lines []string, min, max string) []*Tag {
  var output []*Tag
  minT := MakeTag(min)
  maxT := MakeTag(max)

  for _, val := range(lines) {
    if len(val) > 0 {
      thisVal := MakeTag(val)
      if thisVal.GTE(minT) && thisVal.LTE(maxT) && ( len(output) == 0 || !thisVal.Equals(output[len(output)-1]) ) {
        output = append(output, thisVal)
      }
    }
  }
  return output
}
