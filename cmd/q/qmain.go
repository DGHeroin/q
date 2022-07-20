package main

import (
    "fmt"
    "github.com/DGHeroin/q"
    "io"
    "os"
)

func main() {
    s := q.NewStream(os.Stdin)
    for {
        obj, err := s.Decode()
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Println(err)
            os.Exit(-1)
            return
        }
        handleFilter(obj)
    }
}

func handleFilter(obj *q.Q) {
    obj.Select()
    fmt.Println(obj.ToJsonString())
}
