package main

import (
  "io/ioutil"
  "log"
  "bytes"
)

func tokenize(data []byte) ([][]byte) {
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

func readFile(filename string) [][]byte {
  data, err := ioutil.ReadFile(filename);
  if err != nil {
    log.Fatal("Error reading file", err);
  }

  return tokenize(data);
}

func LearnFile(filename string) {
  words := readFile(filename)

  fromNode := findNode("I");
  var prevLink *Link = nil;

  for _, word := range words {
    toNode := findNode(string(word));
    prevLink = LearnLink(fromNode, toNode, prevLink);
    fromNode = prevLink.To;
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
