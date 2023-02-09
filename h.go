package q

import (
    "encoding/json"
    "errors"
    "fmt"
    "strconv"
    "strings"
    "time"
)

type H map[string]interface{}
type Hs []H

func (h H) ToJson() string {
    if h == nil {
        return ""
    }
    data, _ := json.Marshal(&h)
    return string(data)
}
func (h H) ToJsonPretty() string {
    if h == nil {
        return ""
    }
    data, _ := json.MarshalIndent(&h, "", "  ")
    return string(data)
}

func (h Hs) ToJson() string {
    if h == nil {
        return ""
    }
    data, _ := json.Marshal(&h)
    return string(data)
}
func (h Hs) ToJsonPretty() string {
    if h == nil {
        return ""
    }
    data, _ := json.MarshalIndent(&h, "", "  ")
    return string(data)
}

func (h H) String(key string) string {
    var p map[string]interface{} = h
    v, err := getNestedValue(p, key, ".")
    if err != nil {
        return ""
    }
    if str, ok := v.(string); ok {
        return str
    }
    return ""
}
func (h H) Int(key string) int64 {
    var p map[string]interface{} = h
    v, err := getNestedValue(p, key, ".")
    if err != nil {
        return 0
    }
    var val int64
    switch vv := v.(type) {
    case float64:
        val = int64(vv)
    default:
        val, _ = strconv.ParseInt(fmt.Sprint(v), 10, 64)
    }
    return val
}
func (h H) Float(key string) float64 {
    var p map[string]interface{} = h
    v, err := getNestedValue(p, key, ".")
    if err != nil {
        return 0
    }
    val, _ := strconv.ParseFloat(fmt.Sprint(v), 64)
    return val
}
func (h H) Bool(key string) bool {
    var p map[string]interface{} = h
    v, err := getNestedValue(p, key, ".")
    if err != nil {
        return false
    }
    str := strings.ToLower(fmt.Sprint(v))
    if str == "true" {
        return true
    }
    if str == "1" {
        return true
    }
    return false
}
func (h H) UnixTime(key string) time.Time {
    sec := h.Int(key)
    return time.Unix(sec, 0)
}
func (h H) UnixTimeString(key string) string {
    sec := h.Int(key)
    return time.Unix(sec, 0).Format("2006-01-02 15:04:05.000 Z07:00")
}
func (h H) ByteSize(key string) string {
    v := h.Float(key)
    units := "B"
    if v > 1024 {
        v = v / 1024
        units = "K"
    }
    if v > 1024 {
        v = v / 1024
        units = "M"
    }
    if v > 1024 {
        v = v / 1024
        units = "G"
    }
    if v > 1024 {
        v = v / 1024
        units = "T"
    }
    v, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", v), 64)
    return fmt.Sprintf("%g%s", v, units)
}
func (h H) H(key string) H {
    var p map[string]interface{} = h
    v, err := getNestedValue(p, key, ".")
    if err != nil {
        return nil
    }
    tmp, ok := v.(map[string]interface{})
    if ok {
        return tmp
    }
    return nil
}
func (h H) Hs(key string) Hs {
    var p map[string]interface{} = h
    v, err := getNestedValue(p, key, ".")
    if err != nil {
        return nil
    }
    var result Hs
    if tmp, ok := v.([]interface{}); ok {
        for _, v := range tmp {
            if h, ok := v.(map[string]interface{}); ok {
                result = append(result, h)
            }
        }
    }
    return result
}

var (
// empty interface{}
)

func toFloat(v interface{}) (float64, bool) {
    switch u := v.(type) {
    case int:
        return float64(u), true
    case int8:
        return float64(u), true
    case int16:
        return float64(u), true
    case int32:
        return float64(u), true
    case int64:
        return float64(u), true
    case uint:
        return float64(u), true
    case uint8:
        return float64(u), true
    case uint16:
        return float64(u), true
    case uint32:
        return float64(u), true
    case uint64:
        return float64(u), true
    case uintptr:
        return float64(u), true
    case float32:
        return float64(u), true
    case float64:
        return u, true
    default:
        return 0, false
    }
}
func getNestedValue(input interface{}, node, separator string) (interface{}, error) {
    tokens := strings.Split(node, separator)
    for _, n := range tokens {
        if isIndex(n) {
            // find slice/array
            if arr, ok := input.([]interface{}); ok {
                idx, err := getIndex(n)
                if err != nil {
                    return input, err
                }
                arrLen := len(arr)
                if arrLen == 0 ||
                    idx > arrLen-1 {
                    return empty, errors.New("empty array")
                }
                input = arr[idx]
            }
        } else {
            // find in map
            validNode := false
            if mp, ok := input.(map[string]interface{}); ok {
                input, ok = mp[n]
                validNode = ok
            } else {
                if mp, ok := input.(H); ok {
                    input, ok = mp[n]
                    validNode = ok
                }
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
    val, err := strconv.Atoi(is)
    if err != nil {
        return -1, err
    }
    return val, nil
}
func toString(v interface{}) string {
    return fmt.Sprintf("%v", v)
}
func toInt(v interface{}) (int, bool) {
    if val, ok := toInt64(v); ok {
        return int(val), true
    }
    return 0, false
}
func toInt64(v interface{}) (int64, bool) {
    switch u := v.(type) {
    case int:
        return int64(u), true
    case int8:
        return int64(u), true
    case int16:
        return int64(u), true
    case int32:
        return int64(u), true
    case int64:
        return u, true
    case uint:
        return int64(u), true
    case uint8:
        return int64(u), true
    case uint16:
        return int64(u), true
    case uint32:
        return int64(u), true
    case uint64:
        return int64(u), true
    case uintptr:
        return int64(u), true
    case float32:
        return int64(u), true
    case float64:
        return int64(u), true
    default:
        return 0, false
    }
}
func isNumber(v interface{}) bool {
    switch v.(type) {
    case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
        return true
    case float32, float64:
        return true
    default:
        return false
    }
}
