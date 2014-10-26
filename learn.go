package main

import (
  "io/ioutil"
  "log"
  "bytes"
)

func readFile(filename string) [][]byte {
  data, err := ioutil.ReadFile(filename);
  if err != nil {
    log.Fatal("Error reading file", err);
  }

  data = bytes.Replace(data, []byte("\n"), []byte(" "), -1);
  data = bytes.Replace(data, []byte("\t"), []byte(" "), -1);
  words := bytes.Split(data, []byte(" "));

  out := [][]byte{};
  for _, w := range words  {
    w = bytes.TrimSpace(w)
    if len(w) == 0 {
      continue;
    }
    out = append(out, w);
  }
  return out;
}

func LearnFile(filename string) {
  words := readFile(filename)

  prevWord:= "I";

  for _, word := range words {
    prevWord = LearnLink(prevWord, string(word)).To.Word;
  }
}

func LearnFiles(dir string) {
  files, err := ioutil.ReadDir(dir);
  if err != nil {
      log.Fatal("Error reading directory", err)
  }

  for _, file := range files {
    LearnFile(file.Name());
  }
}
