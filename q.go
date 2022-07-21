package q

import (
    "encoding/json"
    "log"
    "strings"
)

func init() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
}

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
    if q.content == nil {
        return New()
    }
    switch u := q.content.(type) {
    case []interface{}:
        newQ := New()
        newQ.content = buildMapValue(u, keys...)
        return newQ
    case map[string]interface{}:
        newQ := New()
        newQ.content = buildMapValue(u, keys...)
        return newQ
    }
    return nil
}
func buildMapValue(root interface{}, keys ...string) interface{} {
    var result = map[string]interface{}{}
    for _, key := range keys {
        makeNestedValue(root, result, strings.Split(key, "."))
    }
    // 过滤+修复
    return fixSlice(result)
}

func makeNestedValue(root interface{}, result map[string]interface{}, tokens []string) {
    if len(tokens) == 0 {
        return
    }
    for _, n := range tokens {
        if isIndex(n) { // 取数组元素
            // find slice/array
            if arr, ok := root.([]interface{}); ok {
                idx, err := getIndex(n)
                if err != nil {
                    return
                }
                arrLen := len(arr)
                if arrLen == 0 ||
                    idx > arrLen-1 {
                    return
                }
                nextRoot := arr[idx]
                if len(tokens[1:]) == 0 { // 达到最终节点
                    result[n] = nextRoot
                    return
                }

                var nextResult map[string]interface{}
                if u, ok := result[n]; !ok {
                    nextResult = map[string]interface{}{}
                    result[n] = nextResult
                } else {
                    nextResult = u.(map[string]interface{})
                }
                makeNestedValue(nextRoot, nextResult, tokens[1:])
                return
            }
        } else {
            // find in map
            validNode := false
            if mp, ok := root.(map[string]interface{}); ok {
                nextRoot, ok := mp[n]
                validNode = ok
                // 找到子节点
                if _, ok := nextRoot.(map[string]interface{}); ok {
                    if len(tokens[1:]) == 0 { // 达到最终节点

                        result[n] = nextRoot
                        return
                    }
                    var nextResult map[string]interface{}
                    if u, ok := result[n]; !ok {
                        nextResult = map[string]interface{}{}
                        result[n] = nextResult
                    } else {
                        nextResult = u.(map[string]interface{})
                    }
                    makeNestedValue(nextRoot, nextResult, tokens[1:])
                    return
                } else if _, ok := nextRoot.([]interface{}); ok {
                    if len(tokens[1:]) == 0 { // 达到最终节点
                        result[n] = nextRoot
                        return
                    }
                    var nextResult map[string]interface{}
                    if u, ok := result[n]; !ok {
                        nextResult = map[string]interface{}{}
                        result[n] = nextResult
                    } else {
                        nextResult = u.(map[string]interface{})
                    }
                    makeNestedValue(nextRoot, nextResult, tokens[1:])
                    return
                } else {
                    // 根节点
                    result[n] = nextRoot
                    return
                }
            }

            // find in group data
            if mp, ok := root.(map[string][]interface{}); ok {
                root, ok = mp[n]
                validNode = ok
            }

            if !validNode {
                log.Println("  没有找到", n, len(tokens))
                return
            }
        }
    }

    return
}
func fixSlice(r map[string]interface{}) interface{} {
    var (
        arr  []interface{}
        flag = false
    )

    for k, v := range r {
        if isIndex(k) {
            arr = append(arr, v)
            flag = true
        } else {
            if rr, ok := v.(map[string]interface{}); ok {
                r[k] = fixSlice(rr)
            }
        }
    }
    if flag {
        return arr
    }
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
    if q == nil || q.content == nil {
        return "null"
    }
    if data, err := json.Marshal(q.content); err == nil {
        return string(data)
    }
    return "null"
}
func (q *Q) ToJsonStringPretty(indent ...string) string {
    if q == nil || q.content == nil {
        return "null"
    }
    jsonIndent := "  "
    if len(indent) == 1 {
        jsonIndent = indent[0]
    }
    if data, err := json.MarshalIndent(q.content, "", jsonIndent); err == nil {
        return string(data)
    }
    return "null"
}
