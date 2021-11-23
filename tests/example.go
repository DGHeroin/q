package main

import (
    "github.com/DGHeroin/q"
    "log"
    "sync"
    "time"
)

var (
    json1 = `{
    "name":"computers",
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
}`
)
var (
    query = q.NewString(json1)
    wg    sync.WaitGroup
)

func doQuery() {
    name := query.String("name")
    description := query.String("description")
    prices := query.FloatSlice("prices")
    smallPrices := query.IntSlice("small_prices")
    names := query.StringSlice("names")
    productCode := query.Int64("product.code")
    productName := query.String("product.name")
    item0Name := query.String("items.[0].name")
    wg.Done()
    if true {
        return
    }
    log.Println("name:", name)
    log.Println("description:", description)
    log.Println("prices:", prices)
    log.Println("small_prices:", smallPrices)
    log.Println("names:", names)
    log.Println("produce.code:", productCode)
    log.Println("product name:", productName)
    log.Println("items.[0].name:", item0Name)
}
func main() {
    startTime := time.Now()

    max := 1000 * 1000*10
    wg.Add(max)
    for i := 0; i < max; i++ {
        go doQuery()
    }
    wg.Wait()
    dur := time.Since(startTime)
    qps := int(float64(max) / float64(dur) * float64(time.Second))
    log.Println("elapsed time", dur, qps)
}
