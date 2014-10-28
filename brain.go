package main

import (
  "fmt"
  )

type Idea struct {
  Word string
}

type Link struct {
  From *Idea
  To *Idea
  Count int64
}

var Ideas []*Idea;
var Links []*Link;

func LearnIdea(word string) (*Idea) {
  idea := new(Idea);
  idea.Word = word;
  Ideas = append(Ideas, idea);
  return idea
}

func LearnLink(from, to string) (*Link) {
  fromIdea := findIdea(from);
  toIdea := findIdea(to);

  link := findLink(fromIdea, toIdea)

  link.Count += 1;
  return link;
}

func findLink(from, to *Idea) (*Link) {
  for _, link := range Links {
    if link.To == to && link.From == from {
      return link;
    }
  }

  link := new(Link);
  link.From = from;
  link.To = to;
  link.Count = 0;
  Links = append(Links, link);
  return link;
}

func findIdea(word string) (*Idea){
  for _, idea := range Ideas {
    if idea.Word == word {
      return idea;
    }
  }

  return LearnIdea(word)
}

func findLinks(from *Idea) ([]*Link) {
  out := []*Link{};

  for _, link := range Links {
    if link.From == from {
      out = append(out, link);
    }
  }
  return out;
}

func dumpBrain() {

  for _, idea := range Ideas {
    links := findLinks(idea);
    fmt.Println(idea.Word);

    for _, link := range links {
      fmt.Print("\t", link.To.Word, " ", link.Count, "\n");
    }
  }
}
