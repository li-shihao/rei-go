# sui-go
Go API for SUI
## Usage
```yml
sui := new(SUIClient) //Create instance of client
sui.Init("http://127.0.0.1:9000") //Initialise on local node

_ := sui.GetTotalTransactionNumber
_ = sui.GetTransactionsInRange(0, 10)
_ = sui.GetTransaction("Um5bXvoCztqZlhOy/xWslobwSTrXVxVt6QxDjYG+ep0=")
obj = sui.GetObject("0xde1e02902f1c591d6e71d68d41e663105a4e8f25")

_ = obj.GetOwner()
```
