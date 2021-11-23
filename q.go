package q

import (
    "encoding/json"
    "errors"
    "fmt"
    "strconv"
    "strings"
    "sync"
)

var empty interface{}

const (
    cacheString      = "string"
    cacheFloat64     = "float64"
    cacheInt         = "int"
    cacheInt64       = "int64"
    cacheStringSlice = "stringSlice"
    cacheFloatSlice  = "floatSlice"
    cacheIntSlice    = "intSlice"
    cacheInt64Slice  = "int64Slice"
)

type Q struct {
    raw     json.RawMessage
    content interface{}
    cache   map[string]*sync.Map
}

func (q *Q) decode() {
    _ = json.Unmarshal(q.raw, &q.content)
}
func (q *Q) init() {
    q.cache = map[string]*sync.Map{
        cacheString:      {},
        cacheFloat64:     {},
        cacheInt:         {},
        cacheInt64:       {},
        cacheStringSlice: {},
        cacheFloatSlice:  {},
        cacheIntSlice:    {},
        cacheInt64Slice:  {},
    }
}
func (q *Q) From(node string) (interface{}, error) {
    if r, err := getNestedValue(q.content, node, "."); err == nil {
        return r, nil
    } else {
        return nil, err
    }
}
func (q *Q) GetCache(cType string) *sync.Map {
    return q.cache[cType]
}
func (q *Q) String(node string) string {
    cache := q.GetCache(cacheString)
    if v, ok := cache.Load(node); ok {
        return v.(string)
    }
    v, err := q.From(node)
    if err != nil {
        return ""
    }
    r, ok := v.(string)
    if ok {
        cache.Store(node, r)
    }
    return r
}
func (q *Q) Float(node string) float64 {
    cache := q.GetCache(cacheFloat64)
    if v, ok := cache.Load(node); ok {
        return v.(float64)
    }
    v, err := q.From(node)
    if err != nil {
        return 0
    }
    r, ok := v.(float64)
    if ok {
        cache.Store(node, r)
    }
    return r
}
func (q *Q) Int64(node string) int64 {
    return int64(q.Float(node))
}
func (q *Q) Int(node string) int {
    return int(q.Float(node))
}
func (q *Q) FloatSlice(node string) []float64 {
    cache := q.GetCache(cacheFloatSlice)
    if v, ok := cache.Load(node); ok {
        return v.([]float64)
    }
    var result []float64
    v, err := q.From(node)
    if err != nil {
        return result
    }
    if r, ok := v.([]interface{}); ok {
        for _, val := range r {
            if fv, ok := val.(float64); ok {
                result = append(result, fv)
            }
        }
        cache.Store(node, result)
    }
    return result
}
func (q *Q) Int64Slice(node string) []int64 {
    cache := q.GetCache(cacheInt64Slice)
    if v, ok := cache.Load(node); ok {
        return v.([]int64)
    }
    var result []int64
    v, err := q.From(node)
    if err != nil {
        return result
    }
    if r, ok := v.([]interface{}); ok {
        for _, val := range r {
            if fv, ok := val.(float64); ok {
                result = append(result, int64(fv))
            }
        }
        cache.Store(node, result)
    }
    return result
}
func (q *Q) IntSlice(node string) []int {
    cache := q.GetCache(cacheIntSlice)
    if v, ok := cache.Load(node); ok {
        return v.([]int)
    }
    var result []int
    v, err := q.From(node)
    if err != nil {
        return result
    }
    if r, ok := v.([]interface{}); ok {
        for _, val := range r {
            if fv, ok := val.(float64); ok {
                result = append(result, int(fv))
            }
        }
        cache.Store(node, result)
    }
    return result
}
func (q *Q) StringSlice(node string) []string {
    cache := q.GetCache(cacheStringSlice)
    if v, ok := cache.Load(node); ok {
        return v.([]string)
    }
    var result []string
    v, err := q.From(node)
    if err != nil {
        return result
    }
    if r, ok := v.([]interface{}); ok {
        for _, val := range r {
            if fv, ok := val.(string); ok {
                result = append(result, fv)
            }
        }
        cache.Store(node, result)
    }
    return result
}
func NewString(str string) *Q {
    q := &Q{}
    q.raw = []byte(str)
    q.init()
    q.decode()
    return q
}

func toFloat(v interface{}) (float64, bool) {
    var f float64
    flag := true
    switch u := v.(type) {
    case int:
        f = float64(u)
    case int8:
        f = float64(u)
    case int16:
        f = float64(u)
    case int32:
        f = float64(u)
    case int64:
        f = float64(u)
    case float32:
        f = float64(u)
    case float64:
        f = u
    default:
        flag = false
    }
    return f, flag
}
func getNestedValue(input interface{}, node, separator string) (interface{}, error) {
    toks := strings.Split(node, separator)
    for _, n := range toks {
        if isIndex(n) {
            // find slice/array
            if arr, ok := input.([]interface{}); ok {
                indx, err := getIndex(n)
                if err != nil {
                    return input, err
                }
                arrLen := len(arr)
                if arrLen == 0 ||
                    indx > arrLen-1 {
                    return empty, errors.New("empty array")
                }
                input = arr[indx]
            }
        } else {
            // find in map
            validNode := false
            if mp, ok := input.(map[string]interface{}); ok {
                input, ok = mp[n]
                validNode = ok
            }

            // find in group data
            if mp, ok := input.(map[string][]interface{}); ok {
                input, ok = mp[n]
                validNode = ok
            }

            if !validNode {
                return empty, fmt.Errorf("invalid node name %s", n)
            }
        }
    }

    return input, nil
}
func isIndex(in string) bool {
    return strings.HasPrefix(in, "[") && strings.HasSuffix(in, "]")
}
func getIndex(in string) (int, error) {
    if !isIndex(in) {
        return -1, fmt.Errorf("invalid index")
    }
    is := strings.TrimLeft(in, "[")
    is = strings.TrimRight(is, "]")
    oint, err := strconv.Atoi(is)
    if err != nil {
        return -1, err
    }
    return oint, nil
}
func toString(v interface{}) string {
    return fmt.Sprintf("%v", v)
}
