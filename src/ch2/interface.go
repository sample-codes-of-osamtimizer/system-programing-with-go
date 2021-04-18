package main

import (
  "fmt"
)

// interface
type Talker interface {
  Talk()
}

// struct
type Greeter struct {
  name string
}

func (g Greeter) Talk() {
  fmt.Printf("Hello, mu name is %s\n", g.name)
}

func main() {
  var talker Talker
  talker = &Greeter{"wozozo"}
  talker.Talk()
}
