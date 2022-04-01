package main

import (
    "fmt"
    "github.com/DGHeroin/q"
)

func main() {
    var (
        q1   = q.New()
        q2   = q.New()
        q3   = q.New()
        join = q.New()
    )
    q1.FromJson(a)
    q2.FromJson(b)
    q3.FromJson(c)
    join.Join(q2)
    join.Join(q3)

    fmt.Println("name:", q1.Get("name"))
    fmt.Println("items:", q1.Get("items.[0]"))
    fmt.Println("product.code:", q1.Get("product.code"))
    fmt.Println("prices:", q1.Get("prices"))

    fmt.Println("name:", q1.String("name"))
    fmt.Println("id:", q1.Int("id"))
    fmt.Println("weight:", q1.Float("weight"))
    fmt.Println("isSelling:", q1.Bool("isSelling"))
    fmt.Println("product.name:", q1.String("product.name"))

    fmt.Println("q2 count:", q2.Count())
    fmt.Println("q3 count:", q3.Count())
    fmt.Println("join:", join.Count())

    w1 := join.Where("id", ">", 1).Where("id", "<=", 3)
    fmt.Println("where:", w1.Get())

    w2 := join.Where("name", "=", "Fujitsu")
    fmt.Println("where:", w2.Get())

    x := join.Select("id", "price")
    fmt.Println("select:", x.Get())

    f := q1.Find("items")
    fmt.Println("find:", f.Get())

    q1.Set("provider", "best seller")
    fmt.Println("provider:", q1.String("provider"))
}

var (
    a = []byte(`{
    "name":"computers",
    "id":100,
    "weight":10.4,
    "isSelling":true,
    "description":"List of computer products",
    "prices":[2400, 2100, 1200, 400.87, 89.90, 150.10],
    "small_prices":[2400000, 2100000, 1200000, 400870, 89900, 150100],
    "names":["John Doe", "Jane Doe", "Tom", "Jerry", "Nicolas", "Abby"],
    "product":{
        "code":124,
        "name":"good"
    },
    "items":[
      {
         "id":1,
         "name":"MacBook Pro 13 inch retina",
         "price":1350
      },
      {
         "id":2,
         "name":"MacBook Pro 15 inch retina",
         "price":1700
      },
      {
         "id":3,
         "name":"Sony VAIO",
         "price":1200
      },
      {
         "id":4,
         "name":"Fujitsu",
         "price":850
      },
      {
         "id":null,
         "name":"HP core i3 SSD",
         "price":850
      }
   ]
}`)
    b = []byte(`
[
  {
    "id": 1,
    "name": "MacBook Pro 13 inch retina",
    "price": 1350
  },
  {
    "id": 2,
    "name": "MacBook Pro 15 inch retina",
    "price": 1700
  }
  
]
`)
    c = []byte(`
[
{
    "id": 3,
    "name": "Sony VAIO",
    "price": 1200
  },
  {
    "id": 4,
    "name": "Fujitsu",
    "price": 850
  },
  {
    "id": null,
    "name": "HP core i3 SSD",
    "price": 850
  }
]
`)
)
