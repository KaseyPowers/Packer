package utils;

import (
  "log"
)

func Error(err ...interface{}) {
  log.Fatal(err)
}