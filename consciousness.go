package main

import (
  "math/rand"
  "fmt"
  "time"
)

const (
  TriggerInfluence = 2
  OriginalityCoefficent = 25
)

func init() {
  rand.Seed(time.Now().Unix());
}

func NextLink(currLink *Link) (*Link) {

  stayOnLink := rand.Int63n(100);
  if stayOnLink <= OriginalityCoefficent {
    return currLink.Next;
  }

  links := findLinks(currLink.To);

  if len(links) == 0 {
    return currLink;
  }

  var totalSum int64 = 0;
  for _, link := range links {
    totalSum += link.Count;
  }
  nextPath := rand.Int63n(totalSum)
  currSum := int64(0);
  for _, link := range links {
    if currSum <= nextPath && nextPath < (currSum + link.Count) {
      return link;
    }
    currSum += link.Count;
  }

  fmt.Println("UH OH! This shouldn't happen", currLink);
  return currLink;
}

var PrevNode *Link;

func NextThought(currLink *Link, maxLen int) (string, *Link) {
  out := string("");
  numWords := 0;
  for {
    out += currLink.From.Word + " ";
    word := currLink.From.Word
    if (word[len(word) - 1] == '.' || numWords >= maxLen) {
      return out, currLink;
    }

    currLink = NextLink(currLink);
    numWords += 1;
  }
}

func TriggerNode(node *Node, depth int) {
  if (depth == 0) {
    return;
  }

  links := findLinks(node);

  if len(links) == 0 {
    return;
  }

  for _, link := range links {
    link.Count *= TriggerInfluence;
    TriggerNode(link.To, depth - 1);
  }
}

func TriggerSentence(line string, depth int) {
  words := tokenize([]byte(line));

  for _, w := range words {
    node := findNode(string(w));
    TriggerNode(node, depth);
  }
}
