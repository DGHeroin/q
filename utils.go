package q

import (
    "errors"
    "fmt"
    "strconv"
    "strings"
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
