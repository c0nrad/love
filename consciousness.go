package main

import (
  "math/rand"
  "fmt"
  "time"
)

func NextIdea(currIdea *Idea) (*Idea) {
  links := findLinks(currIdea);

  if len(links) == 0 {
    return nil;
  }

  totalSum := 0;
  for _, link := range links {
    totalSum += link.Count;
  }
  rand.Seed(time.Now().Unix())
  nextPath := rand.Intn(totalSum)
  currSum := 0;
  for _, link := range links {
    if currSum <= nextPath && nextPath < (currSum + link.Count) {
      return link.To;
    }
    currSum += link.Count;
  }

  fmt.Println("UH OH! This shouldn't happen", currIdea.Word, nextPath);
  return currIdea;
}
