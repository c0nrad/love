package main

import (
  "fmt"
  "time"
)

func main() {
  LearnFile("./data/lady_lazarus");
  //LearnFile("./data/bell_jar");

  currentIdea := findIdea("I");

  TriggerSentence("Sylvia, do you love me?", 2);
  //dumpBrain();

  for {
    fmt.Print(currentIdea.Word, " ");
    currentIdea = NextIdea(currentIdea);
    time.Sleep(300 * time.Millisecond)
  }
}
