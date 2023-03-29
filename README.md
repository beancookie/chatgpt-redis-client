## 快速开始
```bash
go get github.com/beancookie/chatgpt-redis-service
```

```go
package main

import (
	"fmt"

	chatgptClient "github.com/beancookie/chatgpt-redis-client"
)

func main() {
	c := chatgptClient.NewClient(
		"", // RedisHost
		"", // RedisPassword
		0, // RedisPort
	)
	fmt.Println(c.Call("九牛一毛是什么意思"))
	fmt.Println(c.Call("九牛二毛是什么意思"))
	fmt.Println(c.Call("九牛三毛是什么意思"))
	fmt.Println(c.Call("九牛四毛是什么意思"))
	fmt.Println(c.Call("九牛五毛是什么意思"))
}

```
