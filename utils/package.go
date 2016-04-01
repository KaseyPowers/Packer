package utils;

import (
  "fmt"
  "os"
  // "os/exec"
  // "strings"
  "io/ioutil"
  "encoding/json"
)

type PackageJson struct {
  Name string
  Version string
  Repository RepositoryType
}
type RepositoryType struct {
  Type string
  Url string
}

type Package struct {
  Name string
  Version *Tag
  Url string
}

func makePkg(in PackageJson) *Package {
  return &Package{
    Name: in.Name,
    Version: MakeTag(in.Version),
    Url: in.Repository.Url,
  }
}

func CurrentDir() string {
  loc, err := os.Getwd()
  if err != nil {
    Error(err)
  }
  return loc
}

func GetPackage(dir string) {
  file, err := ioutil.ReadFile(dir+"/package.json")
  if err != nil {
    Error(err)
  }
  // fmt.Printf("%T: %s\n", file, string(file))
  var jsnpkg PackageJson
  json.Unmarshal(file, &jsnpkg)
  pkg := makePkg(jsnpkg)

  start := "Results:\n\tName: %s\n\tVersion: %s\n\tUrl: %s\n"
  vals := []interface{}{
    pkg.Name,
    pkg.Version,
    pkg.Url,
  }
  fmt.Printf(start, vals...)
}