# sui-go
Go API for SUI
## Usage
```Go
sui := new(SUIClient) //Create instance of client
sui.Init("http://127.0.0.1:9000") //Initialise on local node

max := sui.GetTotalTransactionNumber()

arr := sui.GetTransactionsInRange(0, 10)

tx := sui.GetTransaction("Um5bXvoCztqZlhOy/xWslobwSTrXVxVt6QxDjYG+ep0=")
ct := tx.GetContractDeploy()

obj = sui.GetObject("0xde1e02902f1c591d6e71d68d41e663105a4e8f25")
owner = obj.GetOwner()
```
For more guidance look [here](/internal/sui/types.go)
