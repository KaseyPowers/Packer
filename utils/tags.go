package utils;

import (
  "fmt"
  "strconv"
  "strings"
)

type Tag struct {
  Major int
  Minor int
  Patch int
  PreRelease string
  PreReleaseVersion int
}

func MakeTag(input string) *Tag {
  clean := stripTagString(input)
  sections := splitTag(clean)
  output := &Tag{}
  for key, val := range(sections) {
    switch key {
    case 0:
      output.Major = getUint(val)
      // fmt.Println("Major: ", val)
    case 1:
      output.Minor = getUint(val)
      // fmt.Println("Minor: ", val)
    case 2:
      output.Patch = getUint(val)
      // fmt.Println("Patch: ", val)
    case 3:
      output.PreRelease = val
      // fmt.Println("PreRelease: ", val)
    case 4:
      output.PreReleaseVersion = getUint(val)
      // fmt.Println("PreReleaseVersion: ", val)
    }
  }
  return output
}

//Custom print function
func (t Tag) String() string {
  start := "v%d.%d.%d"
  vals := []interface{}{
    t.Major,
    t.Minor,
    t.Patch,
  }
  if t.PreRelease != "" {
    start += "-%s.%d"
    vals = append(vals, t.PreRelease, t.PreReleaseVersion )
  }
  return fmt.Sprintf(start, vals...)
}

// Comparisons

//true if empty
func (a *Tag) EmptyTag() bool {
  return a.Major == 0 && a.Minor == 0 && a.Patch == 0 && a.PreRelease == "" && a.PreReleaseVersion == 0
}

//negative if a less than b, 0 if equal, positive if a greater than b
func (a *Tag) Compare(b *Tag) int {

  compare := []int{
    a.Major - b.Major,
    a.Minor - b.Minor,
    a.Patch - b.Patch,
    comparePreRelease(a.PreRelease, b.PreRelease),
    a.PreReleaseVersion - b.PreReleaseVersion,
  }
  for _, val := range compare {
    if val != 0 {
      return val
    }
  }
  return 0
}

func (a *Tag) Equals(b *Tag) bool {
  result := true
  if !a.EmptyTag() &&  !b.EmptyTag() {
    result = a.Compare(b) == 0
  }
  // fmt.Printf("%v == %v: %v\n", a, b, result)
  return result
}
func (a *Tag) GTE(b *Tag) bool {
  result := true
  if !a.EmptyTag() &&  !b.EmptyTag() {
    result = a.Compare(b) >= 0
  }
  // fmt.Printf("%v >= %v: %v\n", a, b, result)
  return result
}
func (a *Tag) GT(b *Tag) bool {
  result := true
  if !a.EmptyTag() &&  !b.EmptyTag() {
    result = a.Compare(b) > 0
  }
  // fmt.Printf("%v > %v: %v\n", a, b, result)
  return result
}
func (a *Tag) LTE(b *Tag) bool {
  result := true
  if !a.EmptyTag() &&  !b.EmptyTag() {
    result = a.Compare(b) <= 0
  }
  // fmt.Printf("%v <= %v: %v\n", a, b, result)
  return result
}
func (a *Tag) LT(b *Tag) bool {
  result := true
  if !a.EmptyTag() &&  !b.EmptyTag() {
    result = a.Compare(b) < 0
  }
  // fmt.Printf("%v < %v: %v\n", a, b, result)
  return result
}

func comparePreRelease(a, b string) int {
  var aVal, bVal = 0, 0
  var aSet, bSet = false, false
  setVals := []string {
    "alpha",
    "beta",
  }
  for key, val := range setVals {
    if val == a {
      aVal = key
      aSet = true
    }
    if val == b {
      bVal = key
      bSet = true
    }
  }
  output := 0
  if aSet && bSet {
    output = aVal - bVal
  }
  return output
}

//Functions for Constructor:
func stripTagString(input string) string {
  longPrefix := "tags/"
  shortPrefix := "v"
  startIndex := strings.Index(input, longPrefix)
  if startIndex < 0 {
    startIndex = 0
  } else {
    startIndex += len(longPrefix)
  }
  endIndex := strings.IndexAny(input, "^")
  if endIndex < 0 {
    endIndex = len(input)
  }
  output := input[startIndex: endIndex]
  if strings.HasPrefix(output,shortPrefix) {
    output = output[ len(shortPrefix):]
  }
  return output
}

func splitTag(input string) []string {
  var output []string
  for _, val := range(strings.Split(input, ".")) {
    for _, subVal := range(strings.Split(val, "-")) {
      output = append(output, subVal)
    }
  }
  return output
}

func getUint(in string) int {
  if in == "" {
    return 0
  }
  output, err := strconv.ParseUint(in, 10, 0)
  if err != nil {
    Error(err)
  }
  return int(output)
}
