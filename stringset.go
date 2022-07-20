package q

type StringSet []string

func (s *StringSet) add(str string) StringSet {
    for _, val := range *s {
        if str == val {
            return *s
        }
    }
    return append(*s, str)
}
func (s *StringSet) Add(ss ...string) StringSet {
    for _, str := range ss {
        *s = s.add(str)
    }
    return *s
}
func (s *StringSet) Contains(str string) bool {
    for _, val := range *s {
        if str == val {
            return true
        }
    }
    return false
}
func (s *StringSet) Remove(str string) StringSet {
    var ks StringSet
    for _, val := range *s {
        if str == val {
            continue
        }
        ks = append(ks, val)
    }
    *s = ks
    return ks
}
func (s *StringSet) Count() int {
    return len(*s)
}
func (s *StringSet) Diff(o StringSet) StringSet {
    var ks StringSet
    for _, v1 := range *s {
        isMatch := false
        for _, v2 := range o {
            if v2 == v1 {
                isMatch = true
                break
            }
        }
        if !isMatch {
            ks = append(ks, v1)
        }
    }
    return ks
}
func (s *StringSet) Inter(o StringSet) StringSet {
    var ks StringSet
    for _, v1 := range *s {
        isMatch := false
        for _, v2 := range o {
            if v2 == v1 {
                isMatch = true
                break
            }
        }
        if isMatch {
            ks = append(ks, v1)
        }
    }
    return ks
}
func (s *StringSet) Merge(o StringSet) StringSet {
    var ks StringSet
    for _, v1 := range *s {
        ks = ks.Add(v1)
    }
    for _, v2 := range o {
        ks = ks.Add(v2)
    }
    return ks
}
