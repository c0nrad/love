package main

import (
  "math/rand"
  "fmt"
  "time"
)

const (
  TriggerInfluence = 2
)

func NextIdea(currIdea *Idea) (*Idea) {
  links := findLinks(currIdea);

  if len(links) == 0 {
    return findIdea("I");
  }
  var totalSum int64 = 0;
  for _, link := range links {
    totalSum += link.Count;
  }
  rand.Seed(time.Now().Unix())
  nextPath := rand.Int63n(totalSum)
  currSum := int64(0);
  for _, link := range links {
    if currSum <= nextPath && nextPath < (currSum + link.Count) {
      return link.To;
    }
    currSum += link.Count;
  }

  fmt.Println("UH OH! This shouldn't happen", currIdea.Word, nextPath);
  return currIdea;
}

func TriggerIdea(idea *Idea, depth int) {
  if (depth == 0) {
    return;
  }

  links := findLinks(idea);

  if len(links) == 0 {
    return;
  }

  for _, link := range links {
    link.Count *= TriggerInfluence;
    TriggerIdea(link.To, depth - 1);
  }
}

func TriggerSentence(line string, depth int) {
  words := tokenize([]byte(line));

  for _, w := range words {
    idea := findIdea(string(w));
    TriggerIdea(idea, depth);
  }
}
