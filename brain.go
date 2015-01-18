package main

import (
  "fmt"
  )

type Node struct {
  Word string
}

type Link struct {
  From *Node
  To *Node
  Next *Link;
  Count int64;
}

var Nodes []*Node;
var Links []*Link;

func Forget() {
  Nodes = []*Node{};
  Links = []*Link{};
}

func LearnNode(word string) (*Node) {
  node := new(Node);
  node.Word = word;
  Nodes = append(Nodes, node);
  return node
}

func LearnLink(fromNode, toNode *Node, prev *Link) (*Link) {

  link := new(Link);
  link.From = fromNode;
  link.To = toNode;
  link.Count = 1;
  link.Next = nil;

  if prev != nil {
    prev.Next = link;
  }
  Links = append(Links, link);

  return link;
}

func findLink(from, to *Node) (*Link) {
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

func findNode(word string) (*Node){
  for _, node := range Nodes {
    if node.Word == word {
      return node;
    }
  }

  return LearnNode(word)
}

func findLinks(from *Node) ([]*Link) {
  out := []*Link{};

  for _, link := range Links {
    if link.From == from {
      out = append(out, link);
    }
  }
  return out;
}

func dumpBrain() {

  for _, node := range Nodes {
    links := findLinks(node);
    fmt.Println(node.Word);

    for _, link := range links {
      fmt.Println("\tTo: ", link.To.Word);
    }
  }
  fmt.Println("");

  currLink := Links[0];
  for currLink.Next != nil {
    fmt.Print(currLink.From.Word, " ")
    currLink = currLink.Next;
  }

  fmt.Println(currLink.To.Word);
}
