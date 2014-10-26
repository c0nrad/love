package main

import (
  "fmt"
  "time"
)

func main() {
  LearnFile("./data/lady_lazarus");
  //dumpBrain()

  currentIdea := findIdea("I");

  for {
    fmt.Println(currentIdea.Word);
    currentIdea = NextIdea(currentIdea);
    time.Sleep(1000 * time.Millisecond)
  }
}
