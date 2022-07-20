package q

import (
    "encoding/json"
)

var (
    empty interface{}
)

type (
    Q struct {
        content interface{}
    }
)

func New() *Q {
    f := &Q{
        content: empty,
    }
    return f
}
func NewWithString(str string) *Q {
    ptr := New()
    if err := ptr.FromJsonString(str); err != nil {
        return nil
    }
    return ptr
}
func (q *Q) Get(nodes ...string) interface{} {
    if len(nodes) != 1 {
        return q.content
    }
    r, err := getNestedValue(q.content, nodes[0], ".")
    if err == nil {
        return r
    } else {
        return nil
    }
}
func (q *Q) Find(key string) *Q {
    result := New()
    if r, err := getNestedValue(q.content, key, "."); err == nil && r != nil {
        result.content = r
    }

    return result
}
func (q *Q) SetContent(p interface{}) {
    q.content = p
}
func (q *Q) FromJson(data []byte) error {
    return json.Unmarshal(data, &q.content)
}
func (q *Q) FromJsonString(str string) error {
    return q.FromJson([]byte(str))
}
func (q *Q) Count() int {
    arr, ok := q.content.([]interface{})
    if ok {
        return len(arr)
    }

    return 0
}
func (q *Q) Join(o *Q) {
    if q.content == empty {
        q.content = []interface{}{}
    }
    if arr1, ok1 := q.content.([]interface{}); ok1 {
        if arr2, ok2 := o.content.([]interface{}); ok2 {
            for _, val := range arr2 {
                arr1 = append(arr1, val)
            }
        }
        q.content = arr1
    }
}
func (q *Q) Set(key string, p interface{}) {
    if q.content == empty {
        q.content = map[string]interface{}{}
    }
    m, ok := q.content.(map[string]interface{})
    if !ok {
        return
    }
    m[key] = p
}
func (q *Q) Where(key string, cond string, i interface{}) *Q {
    fn := getQuery(cond)
    if fn == nil {
        return New()
    }
    arr, ok := q.content.([]interface{})
    if !ok {
        return New()
    }
    var result []interface{}
    for _, obj := range arr {
        if r, err := getNestedValue(obj, key, "."); err == nil && r != nil {
            if ok, err := fn(r, i); ok && err == nil {
                result = append(result, obj)
            }
        }
    }
    r := New()
    r.content = result
    return r
}
func (q *Q) Select(keys ...string) *Q {
    arr, ok := q.content.([]interface{})
    if !ok {
        return New()
    }
    var result []interface{}
    for _, obj := range arr {
        var m = map[string]interface{}{}
        for _, key := range keys {
            if r, err := getNestedValue(obj, key, "."); err == nil && r != nil {
                m[key] = r
            } else {
                break
            }
        }
        if len(m) == 0 {
            continue
        }
        result = append(result, m)
    }
    r := New()
    r.content = result
    return r
}

func (q *Q) Int(key string) int64 {
    val, _ := q.Get(key).(float64)
    return int64(val)
}
func (q *Q) Float(key string) float64 {
    val, _ := q.Get(key).(float64)
    return val
}

func (q *Q) String(key string) string {
    val, _ := q.Get(key).(string)
    return val
}

func (q *Q) Bool(key string) bool {
    val, _ := q.Get(key).(bool)
    return val
}

func (q *Q) ToJsonString() string {
    if data, err := json.Marshal(q.content); err == nil {
        return string(data)
    }
    return "null"
}
func (q *Q) ToJsonStringPretty(indent ...string) string {
    jsonIndent := "  "
    if len(indent) == 1 {
        jsonIndent = indent[0]
    }
    if data, err := json.MarshalIndent(q.content, "", jsonIndent); err == nil {
        return string(data)
    }
    return "null"
}
