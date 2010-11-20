package dump_test

import (
  . "dump"
  "testing"
  "go/parser"
  "go/token"
  "fmt"
)


var emptyString = ""

type S struct {
  A int
  B int
}

type T struct {
  S
  C int
}

type Circular struct {
  c *Circular
}

func TestDump(t *testing.T) {
  file, e := parser.ParseFile("dump_test.go", nil, parser.ParseComments)
  if e != nil {
    fmt.Println("error", e)
  } else {
    //fmt.Printf("%#v\n", file);
    PrintDump(file)
    PrintDump(map[string]int{"satu": 1, "dua": 2})
    PrintDump([]int{1, 2, 3})
    PrintDump([3]int{1, 2, 3})
    PrintDump(&[][]int{[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}})
    PrintDump(&emptyString)
    PrintDump(T{S{1, 2}, 3})
    PrintDump(token.STRING)

    bulet := make([]Circular, 3)
    bulet[0].c = &bulet[1]
    bulet[1].c = &bulet[2]
    bulet[2].c = &bulet[0]

    PrintDump(struct{ a []Circular }{bulet})
  }
}
