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
