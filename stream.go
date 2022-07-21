package q

import (
    "encoding/json"
    "io"
)

type (
    stream struct {
        dec     *json.Decoder
        filters []filter
    }
    filter struct {
        key   string
        cond  string
        value interface{}
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
    if len(s.filters) != 0 {
        for _, f := range s.filters {
            if !p.Filter(f.key, f.cond, f.value) {
                return nil, nil
            }
        }
    }
    return p, nil
}
func (s *stream) Filter(key string, cond string, val interface{}) {
    s.filters = append(s.filters, filter{
        key:   key,
        cond:  cond,
        value: val,
    })
}
