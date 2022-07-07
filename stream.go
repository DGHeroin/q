package q

import (
    "encoding/json"
    "io"
)

type (
    stream struct {
        dec *json.Decoder
    }
)

func NewStream(r io.Reader) *stream {
    s := &stream{
        dec: json.NewDecoder(r),
    }
    return s
}
func (s *stream) Decode() (*Q, error) {
    var m map[string]interface{}
    err := s.dec.Decode(&m)
    if err != nil {
        return nil, err
    }
    p := New()
    p.content = m
    return p, nil
}
