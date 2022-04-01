### q

一个足够小足够好用的对象查询库.

```golang

import "github.com/DGHeroin/q"

func T()  {
    query := q.New()
    query.FromJson([]byte(`{"name":"q"}`))
    query.String("name")
}

```