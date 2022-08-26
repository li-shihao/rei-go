# sui-go
Go API for SUI
[](https://github.com/li-shihao/sui-go/actions/workflows/go.yml/badge.svg)
## Usage
```Go
import (
  "rei.io/rei/internal/helpers"
  "rei.io/rei/internal/sui"
)

sc := new(sui.SUIClient) //Create instance of client
sc.Init("http://127.0.0.1:9000") //Initialise on local node

max := sc.GetTotalTransactionNumber()

arr := sc.GetTransactionsInRange(0, 10)

tx := sc.GetTransaction("Um5bXvoCztqZlhOy/xWslobwSTrXVxVt6QxDjYG+ep0=")
ct := tx.GetContractDeploy()

obj = sc.GetObject("0xde1e02902f1c591d6e71d68d41e663105a4e8f25")
owner = obj.GetOwner()
```
For more guidance look [here](/internal/sui/types.go)
