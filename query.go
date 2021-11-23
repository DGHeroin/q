package q

import (
    "fmt"
    "reflect"
)

const (
    operatorEq    = "="
    operatorNotEq = "!="
    operatorGt    = ">"
    operatorLt    = "<"
    operatorGtE   = ">="
    operatorLtE   = "<="
)

type (
    QueryFunc func(x, y interface{}) (bool, error)
)

func defaultQueries() map[string]QueryFunc {
    return map[string]QueryFunc{
        operatorEq:    eq,
        operatorNotEq: enq,
        operatorGt:    gt,
        operatorLt:    lq,
        operatorGtE:   gte,
        operatorLtE:   lte,
    }
}

func eq(x interface{}, y interface{}) (bool, error) {
    if v, ok := toFloat(y); ok {
        y = v
    }
    return reflect.DeepEqual(x, y), nil
}

func enq(x interface{}, y interface{}) (bool, error) {
    v, err := eq(x, y)
    return !v, err
}

func gt(x interface{}, y interface{}) (bool, error) {
    xv, ok := x.(float64)
    if !ok {
        return false, fmt.Errorf("%v is not numeric", x)
    }
    if fv, ok := toFloat(y); ok {
        return xv > fv, nil
    }
    return false, nil
}

func lq(x interface{}, y interface{}) (bool, error) {
    xv, ok := x.(float64)
    if !ok {
        return false, fmt.Errorf("%v is not numeric", x)
    }
    if fv, ok := toFloat(y); ok {
        return xv < fv, nil
    }
    return false, nil
}

func gte(x interface{}, y interface{}) (bool, error) {
    xv, ok := x.(float64)
    if !ok {
        return false, fmt.Errorf("%v is not numeric", x)
    }
    if fv, ok := toFloat(y); ok {
        return xv >= fv, nil
    }
    return false, nil
}

func lte(x interface{}, y interface{}) (bool, error) {
    xv, ok := x.(float64)
    if !ok {
        return false, fmt.Errorf("%v is not numeric", x)
    }
    if fv, ok := toFloat(y); ok {
        return xv <= fv, nil
    }
    return false, nil
}
