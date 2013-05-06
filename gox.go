package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "github.com/moovweb/gokogiri"
)

func extract(file string, xpath string) {
  xml, err := ioutil.ReadFile(file)
  if err != nil {
    fmt.Fprintf(os.Stderr, "gox: Could not read %s\n", file)
  }

  doc, err := gokogiri.ParseXml(xml)
  defer doc.Free()

  if err != nil {
    fmt.Fprintf(os.Stderr, "gox: Could not parse %s\n", file)
  }

  nodes, _ := doc.Search(xpath)
  for n := range nodes { fmt.Println(nodes[n].String()) }
}

func main() {
  if len(os.Args) == 3 {
    file  := os.Args[1]
    xpath := os.Args[2]

    extract(file, xpath)
  } else {
    fmt.Fprintf(os.Stderr, "Usage: gox <file> <xpath>\n")
  }
}
