package q

import (
    "errors"
    "fmt"
    "strconv"
    "strings"
)

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
