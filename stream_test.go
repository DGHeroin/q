package q

import (
    "bytes"
    "log"
    "testing"
)

func TestNewStream(t *testing.T) {
    r := bytes.NewBufferString(`
{"value":9}
{"value":10}
{"value":11}
{"value":12}
{"name":"mix"}
{"value":100}
{"value":200}
`)
    s := NewStream(r)
    s.Filter("value", ">=", 10)
    s.Filter("value", "<", 150)
    for {
        q, err := s.Decode()
        if err != nil {
            break
        }
        if q == nil {
            continue
        }
        log.Println(q.ToJsonString())
    }
}
