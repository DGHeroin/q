package q

import (
    "encoding/json"
    "fmt"
    "reflect"
    "strconv"
    "strings"
)

type (
    H  map[string]interface{}
    Hs []H
)

func (h H) GetFloat64(key string) (float64, bool) {
    v, ok := h[key]
    if !ok {
        return 0, false
    }
    str := fmt.Sprintf("%v", v)
    f64, err := strconv.ParseFloat(str, 10)
    if err != nil {
        return 0, false
    }
    return f64, true
}
func (h H) GetInt64(key string) (int64, bool) {
    v, ok := h[key]
    if !ok {
        return 0, false
    }
    str := fmt.Sprintf("%v", v)
    f64, err := strconv.ParseFloat(str, 10)
    if err != nil {
        return 0, false
    }
    return int64(f64), true
}
func (h H) GetInt(key string) (int, bool) {
    v, ok := h.GetInt64(key)
    return int(v), ok
}
func (h H) GetBool(key string) (bool, bool) {
    v, ok := h[key]
    if !ok {
        return false, false
    }
    str := fmt.Sprintf("%v", v)
    if str == "true" {
        return true, true
    } else if str == "false" {
        return false, true
    }
    return false, false
}
func (h H) GetString(key string) (string, bool) {
    if h == nil {
        return "", false
    }
    v, ok := h[key]
    if !ok {
        return "", false
    }
    return fmt.Sprintf("%v", v), true
}

func (h H) GetIntArray(key string) (result []int, ok bool) {
    str := ""
    obj := reflect.TypeOf(h[key])
    if obj == nil {
        return nil, false
    }
    switch obj.Kind() {
    case reflect.Float64:
        f := h[key].(float64)
        str = fmt.Sprintf("%d", int64(f))
    default:
        str = fmt.Sprint(h[key])
    }
    tokens := strings.Split(str, ",")

    for _, tok := range tokens {
        if tok != "" {
            val, _ := strconv.Atoi(tok)
            result = append(result, val)
        }
    }

    return result, true
}
func (h H) GetInt64Array(key string) (result []int64, ok bool) {
    str := ""
    obj := reflect.TypeOf(h[key])
    if obj == nil {
        return nil, false
    }
    switch obj.Kind() {
    case reflect.Float64:
        f := h[key].(float64)
        str = fmt.Sprintf("%d", int64(f))
    default:
        str = fmt.Sprint(h[key])
    }
    tokens := strings.Split(str, ",")

    for _, tok := range tokens {
        if tok != "" {
            val, _ := strconv.ParseInt(tok, 10, 64)
            result = append(result, val)
        }
    }

    return result, true
}
func (h H) GetSingleIntArray(key string) (result []int, ok bool) {
    str, _ := h.GetString(key)
    tokens := strings.Split(str, "|")

    for _, tok := range tokens {
        if tok != "" {
            val, _ := strconv.Atoi(tok)
            result = append(result, val)
        }
    }

    return result, true
}
func (h H) GetSingleInt64Array(key string) (result []int64, ok bool) {
    str, _ := h.GetString(key)
    tokens := strings.Split(str, "|")

    for _, tok := range tokens {
        if tok != "" {
            val, _ := strconv.ParseInt(tok, 10, 64)
            result = append(result, val)
        }
    }

    return result, true
}

func (h H) Float64(key string) float64 {
    v, ok := h[key]
    if !ok {
        return 0
    }
    str := fmt.Sprintf("%v", v)
    f64, err := strconv.ParseFloat(str, 10)
    if err != nil {
        return 0
    }
    return f64
}
func (h H) Int64(key string) int64 {
    v, ok := h[key]
    if !ok {
        return 0
    }
    str := fmt.Sprintf("%v", v)
    f64, err := strconv.ParseFloat(str, 10)
    if err != nil {
        return 0
    }
    return int64(f64)
}
func (h H) Int(key string) int {
    v, _ := h.GetInt64(key)
    return int(v)
}
func (h H) Bool(key string) bool {
    v, ok := h[key]
    if !ok {
        return false
    }
    str := fmt.Sprintf("%v", v)
    if str == "true" {
        return true
    } else if str == "false" {
        return false
    }
    return false
}
func (h H) String(key string) string {
    if h == nil {
        return ""
    }
    v, ok := h[key]
    if !ok {
        return ""
    }
    return fmt.Sprintf("%v", v)
}

func (h H) StringArray(key string) []string {
    var res []string
    if h == nil {
        return res
    }

    str, ok := h.GetString(key)
    if !ok {
        return res
    }

    tokens := strings.Split(str, ",")
    for _, token := range tokens {
        res = append(res, token)
    }

    return res
}
func (h H) ParseTimeN() (result [][2]int64) {
    if n, ok := h.GetInt("TimeN"); ok {
        for i := 1; i <= n; i++ {
            startTime, ok1 := h.GetInt64(fmt.Sprintf("StartTime%d", i))
            stopTime, ok2 := h.GetInt64(fmt.Sprintf("StopTime%d", i))
            if ok1 && ok2 {
                result = append(result, [2]int64{startTime, stopTime})
            }
        }
    }
    return
}
func (h H) IntArray(key string) (result []int) {
    str := ""
    obj := reflect.TypeOf(h[key])
    if obj == nil {
        return nil
    }
    switch obj.Kind() {
    case reflect.Float64:
        f := h[key].(float64)
        str = fmt.Sprintf("%d", int64(f))
    default:
        str = fmt.Sprint(h[key])
    }
    tokens := strings.Split(str, ",")

    for _, tok := range tokens {
        if tok != "" {
            val, _ := strconv.Atoi(tok)
            result = append(result, val)
        }
    }
    return result
}
func (h H) Int64Array(key string) (result []int64) {
    str := ""
    obj := reflect.TypeOf(h[key])
    if obj == nil {
        return nil
    }
    switch obj.Kind() {
    case reflect.Float64:
        f := h[key].(float64)
        str = fmt.Sprintf("%d", int64(f))
    default:
        str = fmt.Sprint(h[key])
    }
    tokens := strings.Split(str, ",")

    for _, tok := range tokens {
        if tok != "" {
            val, _ := strconv.ParseInt(tok, 10, 64)
            result = append(result, val)
        }
    }

    return result
}
func (h H) SingleIntArray(key string) (result []int) {
    str, _ := h.GetString(key)
    tokens := strings.Split(str, "|")

    for _, tok := range tokens {
        if tok != "" {
            val, _ := strconv.Atoi(tok)
            result = append(result, val)
        }
    }

    return result
}
func (h H) SingleInt64Array(key string) (result []int64) {
    str, _ := h.GetString(key)
    tokens := strings.Split(str, "|")

    for _, tok := range tokens {
        if tok != "" {
            val, _ := strconv.ParseInt(tok, 10, 64)
            result = append(result, val)
        }
    }

    return result
}

func (h H) IntSlice(key string) (result []int) {
    p := h.Get(key)
    if p == nil {
        return
    }
    if arr, ok := p.([]interface{}); ok {
        for _, i := range arr {
            if val, ok := toInt(i); ok {
                result = append(result, val)
            }
        }
    }
    return
}
func (h H) ToJson() string {
    if data, err := json.Marshal(&h); err == nil {
        return string(data)
    }
    return "{}"
}
func (h H) HasKey(key string) bool {
    if h == nil {
        return false
    }
    _, ok := h[key]
    return ok
}
func (h H) Get(s string) interface{} {
    if h == nil {
        return nil
    }
    return h[s]
}
func (h H) GetMap(s string) (H, bool) {
    if h == nil {
        return nil, false
    }
    obj, ok := h[s]
    if !ok {
        return nil, false
    }

    r, ok := obj.(map[string]interface{})
    return r, ok
}
func (h H) TryGet(s string) (interface{}, bool) {
    if h == nil {
        return nil, false
    }
    r, ok := h[s]
    return r, ok
}
func (h H) ContainsStringWithSplit(key string, s string, sep string) bool {
    if !h.HasKey(key) {
        return false
    }
    str, _ := h.GetString(key)
    tokens := strings.Split(str, sep)
    for _, token := range tokens {
        if s == token {
            return true
        }
    }
    return false
}
func (h H) ToQ() *Q {
    return NewWithString(h.ToJson())
}

func (hs Hs) GetInt64(line int, key string) (int64, bool) {
    h := hs[line]
    return h.GetInt64(key)
}
func (hs Hs) GetInt(line int, key string) (int, bool) {
    h := hs[line]
    return h.GetInt(key)
}
func (hs Hs) GetString(line int, key string) (string, bool) {
    h := hs[line]
    return h.GetString(key)
}
func (hs Hs) Len() int {
    return len(hs)
}
