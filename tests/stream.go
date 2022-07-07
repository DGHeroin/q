package main

import (
    "github.com/DGHeroin/q"
    "log"
    "time"
)

func testStream() {
    r := q.NewBlockingBuffer()
    s := q.NewStream(r)
    r.WriteString(`{"menu": {
  "id": "file",
  "value": "File",
  "popup": {
    "menuitem": [
      {"value": "New", "onclick": "CreateNewDoc()"},
      {"value": "Open", "onclick": "OpenDoc()"},
      {"value": "Close", "onclick": "CloseDoc()"}
    ]
  }
}}`)
    r.WriteString(`{"menu": {
  "id": "file2",
  "value": "File",
  "popup": {
    "menuitem": [
      {"value": "New", "onclick": "CreateNewDoc()"},
      {"value": "Open", "onclick": "OpenDoc()"},
      {"value": "Close", "onclick": "CloseDoc()"}
    ]
  }
}}`)
    time.AfterFunc(time.Second, func() {
        r.Close()
    })
    for {
        query, err := s.Decode()
        if err != nil {
            log.Println(err)
            return
        }
        log.Println("找到!", query)
    }
}
func main() {
    testStream()
}
