package q

import (
    "log"
    "testing"
)

func TestQSelect(t *testing.T) {
    query := NewWithString(`
{
"menu": {
  "id": "file",
  "value": "File",
  "count":100,
  "popup": {
    "name":"pop-name",
    "prices":["abc", 1, true, 2.2],
    "menuitem": [
      {"value": "New", "onclick": "CreateNewDoc()"},
      {"value": "Open", "onclick": "OpenDoc()"},
      {"value": "Close", "onclick": "CloseDoc()"}
    ]
  }
},
  "menu_bar": [
    {"name":"1", "type":1},
    {"name":"2", "type":2}
  ]
}
`)

    log.Println(query.Select(
        "menu.id",
        "menu.value",
        "menu.count",
        "menu.popup.prices",
        "menu.popup.menuitem.[1].onclick",
        "menu.popup.menuitem.[1].value",
        "menu_bar.[1]",
    ).ToJsonStringPretty())

    log.Println(query.Select(
        "menu.popup.prices.[0]",
        "menu.popup.menuitem.[1]",
        "menu.popup.menuitem.[1].value",
        "menu.popup.menuitem.[2].value",
        "menu_bar.[0]",
        "menu_bar.[1].type",
    ).ToJsonStringPretty())
}

func TestQSelectSlice(t *testing.T) {
    query := NewWithString(`
[
  {"value": "New", "onclick": "CreateNewDoc()"},
  {"value": "Open", "onclick": "OpenDoc()"},
  {"value": "Close", "onclick": "CloseDoc()"}
]
`)

    log.Println(query.Select(
        "[0]",
        "[1].value",
    ).ToJsonStringPretty())
}

func TestQMergeMap(t *testing.T) {
    a := NewWithString(`
{
  "name":"a",
  "a_ext":"a name"
}
`)

    b := NewWithString(`
{
  "name":"b",
  "b_ext":"b name",
  "price":100
}
`)
    c := a.Merge(b)
    t.Log(c.ToJsonStringPretty())
}

func TestQMergeSlice(t *testing.T) {
    a := NewWithString(`
[1,2,3]
`)

    b := NewWithString(`
[3,4,5]
`)
    c := a.Merge(b)
    t.Log(c.ToJsonStringPretty())
}

func TestQSliceUniq(t *testing.T) {
    a := NewWithString(`
[1,2,3,3,4,5]
`)
    t.Log(a.Uniq().ToJsonStringPretty())
}

func TestQSliceSum(t *testing.T) {
    a := NewWithString(`
[1,2,3,4,5]
`)
    t.Log(a.Sum())
}

func TestQWhere(t *testing.T) {
    a := NewWithString(`
[
  {"value": 10},
  {"value": 11},
  {"value": 12}
]
`)
    t.Log(a.Where("value", ">", 10).ToJsonStringPretty())
}
func TestQFilter(t *testing.T) {
    a := NewWithString(`{"value": 10}`)
    b := NewWithString(`{"value": 11}`)
    c := NewWithString(`{"value": 12}`)
    t.Log(a.Filter("value", ">", 10))
    t.Log(b.Filter("value", ">", 10))
    t.Log(c.Filter("value", ">", 10))

}
