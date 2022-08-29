package sui

import (
	"reflect"
	"testing"
	"time"
)

func TestGetTotalTransactionNumber(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	if result := tc.GetTotalTransactionNumber(); result == 0 {
		t.Errorf("Error while getting total transaction number")
	}
}

func TestGetTransaction(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetTransaction("tsl8gn6Wbeq17lmh8MCRkacdHYCiuKDj7r80MD3OhpA=")
	got1 := uint32(result1.Result.Certificate.Data.Transactions[0]["TransferSui"].(map[string]interface{})["amount"].(float64))

	if got1 != 50000 {
		t.Errorf("Result was incorrect, got %d, want %d.", got1, 50000)
	}

	result2, _ := tc.GetTransaction("k1c9GX/kJT7DU0TFMwmrIOgXpkO0SHUYZBy4S/44eSA=")
	got2 := result2.Result.Certificate.Data.Transactions[0]["Call"].(map[string]interface{})["module"]

	if got2 != "suiname_nft" {
		t.Errorf("Result was incorrect, got %s, want %s.", got2, "suiname_nft")
	}

	result3, _ := tc.GetTransaction("IUV4c4wSB7dXvYiKTD3y99QeqK3cCUk034CoFm0NKv0=")
	got3 := result3.Result.Certificate.Data.Transactions[0]["TransferObject"].(map[string]interface{})["recipient"]

	if got3 != "0x661e76241de2413cb5d7f4ce380e283125941b1b" {
		t.Errorf("Result was incorrect, got %s, want %s.", got3, "0x661e76241de2413cb5d7f4ce380e283125941b1b")
	}

	result4, _ := tc.GetTransaction("ytTaRyKDc/eBVxlfijdOAPQZbvA42LlGQoK+s2XJhz0=")
	got4 := *result4.Arguments
	exp4 := []Arg{
		{ID: "ytTaRyKDc/eBVxlfijdOAPQZbvA42LlGQoK+s2XJhz0=", Type: "U8", Data: ""},
		{ID: "ytTaRyKDc/eBVxlfijdOAPQZbvA42LlGQoK+s2XJhz0=", Type: "U8", Data: float64(255)},
		{ID: "ytTaRyKDc/eBVxlfijdOAPQZbvA42LlGQoK+s2XJhz0=", Type: "U8", Data: ""},
	}

	if !reflect.DeepEqual(got4, exp4) {
		t.Errorf("Result was incorrect, got %v, want %v.", got4, exp4)
	}

	result5, _ := tc.GetTransaction("f+eG0euJ2s5X/md68P+gQwHwVCn6PEAtdUXC+ocOeQA=")
	got5 := *result5.Events
	exp5 := []Event{
		{TX: "f+eG0euJ2s5X/md68P+gQwHwVCn6PEAtdUXC+ocOeQA=", Type: "mint", Sender: "0x84278a52a92e9532ffcc06a73a3d9b9a79696936", Recipient: "0x84278a52a92e9532ffcc06a73a3d9b9a79696936", ObjectId: "0xc4b9a5e7dd8acd17697762b5068050de68ca4d61", Version: uint32(0)},
		{TX: "f+eG0euJ2s5X/md68P+gQwHwVCn6PEAtdUXC+ocOeQA=", Type: "transfer", Sender: "0x84278a52a92e9532ffcc06a73a3d9b9a79696936", Recipient: "0x84278a52a92e9532ffcc06a73a3d9b9a79696936", ObjectId: "0xde1e02902f1c591d6e71d68d41e663105a4e8f25", Version: uint32(1)},
		{TX: "f+eG0euJ2s5X/md68P+gQwHwVCn6PEAtdUXC+ocOeQA=", Type: "burn", Sender: "0x84278a52a92e9532ffcc06a73a3d9b9a79696936", ObjectId: "0x3c18ae6cde75b48702a737f7d6663359bd4865c1"},
	}

	if !reflect.DeepEqual(got5, exp5) {
		t.Errorf("Result was incorrect, got %v, want %v.", got5, exp5)
	}
}

func TestGetTransactionInRange(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetTransactionsInRange(0, 10)

	if len(result1) != 10 {
		t.Errorf("Result was incorrect, got %d, want %d.", len(result1), 10)
	}

	result2, _ := tc.GetTransactionsInRange(0, 100)

	if len(result2) != 100 {
		t.Errorf("Result was incorrect, got %d, want %d.", len(result2), 10)
	}
}

func TestGetTransactionsByObject(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetTransactionsByObject("0x42fe8e8893586133fe7b5e5e225b288baed58ec7")
	got1 := false
	exp1 := true

	for _, v := range result1 {
		if v == "EglTi7UMnl4MtRuwphX9+jM/1+qXVeH++Goxbp4FpMs=" {
			got1 = true
		}
	}

	if !reflect.DeepEqual(got1, exp1) {
		t.Errorf("Result was incorrect, got %t, want %t.", got1, exp1)
	}

	result2, _ := tc.GetTransactionsByObject("0x632ac9a2e3762fd220f7b66019ab5000584911dd")
	got2 := false
	exp2 := true

	for _, v := range result2 {
		if v == "m8RLiVx8+gGG9lVgj/aEcyv3ByjU+2GFjk3xPBaZNuo=" {
			got2 = true
		}
	}

	if !reflect.DeepEqual(got2, exp2) {
		t.Errorf("Result was incorrect, got %t, want %t.", got2, exp2)
	}
}

func TestGetType(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetTransaction("WCyfaBotBOKG1qBwakZ92wKtm79itrRpvJVj9zLaFIY=")
	got1 := result1.GetType()

	if got1 != "Call" {
		t.Errorf("Result was incorrect, got %s, want %s.", got1, "Call")
	}

	result2, _ := tc.GetTransaction("tsl8gn6Wbeq17lmh8MCRkacdHYCiuKDj7r80MD3OhpA=")
	got2 := result2.GetType()

	if got2 != "TransferSui" {
		t.Errorf("Result was incorrect, got %s, want %s.", got2, "TransferSui")
	}

	result3, _ := tc.GetTransaction("qurUKphegqg9QJjmqyoxFg2U5G+HAK49Wbk0YxDOI7c=")
	got3 := result3.GetType()

	if got3 != "TransferObject" {
		t.Errorf("Result was incorrect, got %s, want %s.", got3, "TransferObject")
	}
}

func TestGetTime(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetTransaction("WCyfaBotBOKG1qBwakZ92wKtm79itrRpvJVj9zLaFIY=")
	got1 := result1.GetTime()
	exp1, _ := time.Parse(time.RFC3339, "2022-08-25T14:05:55.835000038Z")

	if got1 != exp1 {
		t.Errorf("Result was incorrect, got %s, want %s.", got1, exp1)
	}

	result2, _ := tc.GetTransaction("tsl8gn6Wbeq17lmh8MCRkacdHYCiuKDj7r80MD3OhpA=")
	got2 := result2.GetTime()
	exp2, _ := time.Parse(time.RFC3339, "2022-08-25T09:40:38.845000028Z")

	if got2 != exp2 {
		t.Errorf("Result was incorrect, got %s, want %s.", got2, exp2)
	}

	result3, _ := tc.GetTransaction("qurUKphegqg9QJjmqyoxFg2U5G+HAK49Wbk0YxDOI7c=")
	got3 := result3.GetTime()
	exp3, _ := time.Parse(time.RFC3339, "2022-08-25T14:11:13.457999944Z")

	if got3 != exp3 {
		t.Errorf("Result was incorrect, got %s, want %s.", got3, exp3)
	}
}

func TestGetID(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetTransaction("WCyfaBotBOKG1qBwakZ92wKtm79itrRpvJVj9zLaFIY=")
	got1 := result1.GetID()

	if got1 != "WCyfaBotBOKG1qBwakZ92wKtm79itrRpvJVj9zLaFIY=" {
		t.Errorf("Result was incorrect, got %s, want %s.", got1, "WCyfaBotBOKG1qBwakZ92wKtm79itrRpvJVj9zLaFIY=")
	}

	result2, _ := tc.GetTransaction("tsl8gn6Wbeq17lmh8MCRkacdHYCiuKDj7r80MD3OhpA=")
	got2 := result2.GetID()

	if got2 != "tsl8gn6Wbeq17lmh8MCRkacdHYCiuKDj7r80MD3OhpA=" {
		t.Errorf("Result was incorrect, got %s, want %s.", got2, "tsl8gn6Wbeq17lmh8MCRkacdHYCiuKDj7r80MD3OhpA=")
	}

	result3, _ := tc.GetTransaction("qurUKphegqg9QJjmqyoxFg2U5G+HAK49Wbk0YxDOI7c=")
	got3 := result3.GetID()

	if got3 != "qurUKphegqg9QJjmqyoxFg2U5G+HAK49Wbk0YxDOI7c=" {
		t.Errorf("Result was incorrect, got %s, want %s.", got3, "qurUKphegqg9QJjmqyoxFg2U5G+HAK49Wbk0YxDOI7c=")
	}
}

func TestGetStatus(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetTransaction("WCyfaBotBOKG1qBwakZ92wKtm79itrRpvJVj9zLaFIY=")
	got1 := result1.GetStatus()

	if got1 != true {
		t.Errorf("Result was incorrect, got %t, want %t.", got1, true)
	}

	result2, _ := tc.GetTransaction("1MI3iWk0MBEe2BwpMDpsiX1UWgkTg5Oh6KJzG9DrCxM=")
	got2 := result2.GetStatus()

	if got2 != false {
		t.Errorf("Result was incorrect, got %t, want %t.", got2, false)
	}
}

func TestGetSender(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetTransaction("WCyfaBotBOKG1qBwakZ92wKtm79itrRpvJVj9zLaFIY=")
	got1 := result1.GetSender()

	if got1 != "0x823a57d9055e36a28db14abe8c73f23a6e765d3d" {
		t.Errorf("Result was incorrect, got %s, want %s.", got1, "0x823a57d9055e36a28db14abe8c73f23a6e765d3d")
	}

	result2, _ := tc.GetTransaction("tsl8gn6Wbeq17lmh8MCRkacdHYCiuKDj7r80MD3OhpA=")
	got2 := result2.GetSender()

	if got2 != "0xc4173a804406a365e69dfb297d4eaaf002546ebd" {
		t.Errorf("Result was incorrect, got %s, want %s.", got2, "0xc4173a804406a365e69dfb297d4eaaf002546ebd")
	}

	result3, _ := tc.GetTransaction("qurUKphegqg9QJjmqyoxFg2U5G+HAK49Wbk0YxDOI7c=")
	got3 := result3.GetSender()

	if got3 != "0xf3a838034e5d7b9cce68bbba683749fe588c0f92" {
		t.Errorf("Result was incorrect, got %s, want %s.", got3, "0xf3a838034e5d7b9cce68bbba683749fe588c0f92")
	}
}

func TestGetRecipient(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetTransaction("WCyfaBotBOKG1qBwakZ92wKtm79itrRpvJVj9zLaFIY=")
	got1 := result1.GetRecipient()

	if got1 != nil {
		t.Errorf("Result was incorrect, got %s, want %s.", *got1, "nil")
	}

	result2, _ := tc.GetTransaction("tsl8gn6Wbeq17lmh8MCRkacdHYCiuKDj7r80MD3OhpA=")
	got2 := result2.GetRecipient()

	if *got2 != "0xce1313e3b8e5326568e91460cfc449f2cb7ebfc0" {
		t.Errorf("Result was incorrect, got %s, want %s.", *got2, "0xce1313e3b8e5326568e91460cfc449f2cb7ebfc0")
	}

	result3, _ := tc.GetTransaction("qurUKphegqg9QJjmqyoxFg2U5G+HAK49Wbk0YxDOI7c=")
	got3 := result3.GetRecipient()

	if *got3 != "0x5f00c4a09b4f2cf91c7d6478be22a6a522d9299f" {
		t.Errorf("Result was incorrect, got %s, want %s.", *got3, "0x5f00c4a09b4f2cf91c7d6478be22a6a522d9299f")
	}
}

func TestGetTransferAmount(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetTransaction("WCyfaBotBOKG1qBwakZ92wKtm79itrRpvJVj9zLaFIY=")
	got1 := result1.GetTransferAmount()

	if got1 != nil {
		t.Errorf("Result was incorrect, got %f, want %f.", *got1, 0.0)
	}

	result2, _ := tc.GetTransaction("tsl8gn6Wbeq17lmh8MCRkacdHYCiuKDj7r80MD3OhpA=")
	got2 := result2.GetTransferAmount()

	if *got2 != 50000 {
		t.Errorf("Result was incorrect, got %f, want %f.", *got2, 50000.0)
	}

	result3, _ := tc.GetTransaction("qurUKphegqg9QJjmqyoxFg2U5G+HAK49Wbk0YxDOI7c=")
	got3 := result3.GetTransferAmount()

	if got3 != nil {
		t.Errorf("Result was incorrect, got %f, want %f.", *got3, 0.0)
	}
}

func TestGetGetContractPackage(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetTransaction("WCyfaBotBOKG1qBwakZ92wKtm79itrRpvJVj9zLaFIY=")
	got1 := result1.GetContractPackage()

	if *got1 != "0x0000000000000000000000000000000000000002" {
		t.Errorf("Result was incorrect, got %s, want %s.", *got1, "0x0000000000000000000000000000000000000002")
	}

	result2, _ := tc.GetTransaction("tsl8gn6Wbeq17lmh8MCRkacdHYCiuKDj7r80MD3OhpA=")
	got2 := result2.GetContractPackage()

	if got2 != nil {
		t.Errorf("Result was incorrect, got %s, want %s.", *got2, "nil")
	}

	result3, _ := tc.GetTransaction("qurUKphegqg9QJjmqyoxFg2U5G+HAK49Wbk0YxDOI7c=")
	got3 := result3.GetContractPackage()

	if got3 != nil {
		t.Errorf("Result was incorrect, got %s, want %s.", *got3, "nil")
	}

	result4, _ := tc.GetTransaction("1MI3iWk0MBEe2BwpMDpsiX1UWgkTg5Oh6KJzG9DrCxM=")
	got4 := result4.GetContractPackage()

	if *got4 != "0xa41ca8645474a1cacc08ce4559275ba07136c412" {
		t.Errorf("Result was incorrect, got %s, want %s.", *got4, "0xa41ca8645474a1cacc08ce4559275ba07136c412")
	}

	result5, _ := tc.GetTransaction("KEErop0BQhU0vM8sQNb2X0IygEVUar60XBn9eOFqlTQ=")
	got5 := result5.GetContractPackage()

	if *got5 != "0x0000000000000000000000000000000000000002" {
		t.Errorf("Result was incorrect, got %s, want %s.", *got5, "0x0000000000000000000000000000000000000002")
	}

	result6, _ := tc.GetTransaction("f+eG0euJ2s5X/md68P+gQwHwVCn6PEAtdUXC+ocOeQA=")
	got6 := result6.GetContractPackage()

	if *got6 != "0x6774c132abf5521bf9bd8e8c6948c03f0f0cb3f4" {
		t.Errorf("Result was incorrect, got %s, want %s.", *got6, "0x6774c132abf5521bf9bd8e8c6948c03f0f0cb3f4")
	}

	result7, _ := tc.GetTransaction("ytTaRyKDc/eBVxlfijdOAPQZbvA42LlGQoK+s2XJhz0=")
	got7 := result7.GetContractPackage()

	if *got7 != "0xb82c2e33acecc71f2b2b742c366d902bd640a1bf" {
		t.Errorf("Result was incorrect, got %s, want %s.", *got7, "0xb82c2e33acecc71f2b2b742c366d902bd640a1bf")
	}
}

func TestGetContractModule(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetTransaction("WCyfaBotBOKG1qBwakZ92wKtm79itrRpvJVj9zLaFIY=")
	got1 := result1.GetContractModule()

	if *got1 != "devnet_nft" {
		t.Errorf("Result was incorrect, got %s, want %s.", *got1, "devnet_nft")
	}

	result2, _ := tc.GetTransaction("tsl8gn6Wbeq17lmh8MCRkacdHYCiuKDj7r80MD3OhpA=")
	got2 := result2.GetContractModule()

	if got2 != nil {
		t.Errorf("Result was incorrect, got %s, want %s.", *got2, "nil")
	}

	result3, _ := tc.GetTransaction("qurUKphegqg9QJjmqyoxFg2U5G+HAK49Wbk0YxDOI7c=")
	got3 := result3.GetContractModule()

	if got3 != nil {
		t.Errorf("Result was incorrect, got %s, want %s.", *got3, "nil")
	}

	result4, _ := tc.GetTransaction("1MI3iWk0MBEe2BwpMDpsiX1UWgkTg5Oh6KJzG9DrCxM=")
	got4 := result4.GetContractModule()

	if *got4 != "suiname_nft" {
		t.Errorf("Result was incorrect, got %s, want %s.", *got4, "suiname_nft")
	}

	result5, _ := tc.GetTransaction("KEErop0BQhU0vM8sQNb2X0IygEVUar60XBn9eOFqlTQ=")
	got5 := result5.GetContractModule()

	if *got5 != "coin" {
		t.Errorf("Result was incorrect, got %s, want %s.", *got5, "coin")
	}

	result6, _ := tc.GetTransaction("f+eG0euJ2s5X/md68P+gQwHwVCn6PEAtdUXC+ocOeQA=")
	got6 := result6.GetContractModule()

	if *got6 != "market" {
		t.Errorf("Result was incorrect, got %s, want %s.", *got6, "market")
	}

	result7, _ := tc.GetTransaction("ytTaRyKDc/eBVxlfijdOAPQZbvA42LlGQoK+s2XJhz0=")
	got7 := result7.GetContractModule()

	if *got7 != "rgb" {
		t.Errorf("Result was incorrect, got %s, want %s.", *got7, "rgb")
	}
}

func TestGetContractFunction(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetTransaction("WCyfaBotBOKG1qBwakZ92wKtm79itrRpvJVj9zLaFIY=")
	got1 := result1.GetContractFunction()

	if *got1 != "mint" {
		t.Errorf("Result was incorrect, got %s, want %s.", *got1, "mint")
	}

	result2, _ := tc.GetTransaction("tsl8gn6Wbeq17lmh8MCRkacdHYCiuKDj7r80MD3OhpA=")
	got2 := result2.GetContractFunction()

	if got2 != nil {
		t.Errorf("Result was incorrect, got %s, want %s.", *got2, "nil")
	}

	result3, _ := tc.GetTransaction("qurUKphegqg9QJjmqyoxFg2U5G+HAK49Wbk0YxDOI7c=")
	got3 := result3.GetContractFunction()

	if got3 != nil {
		t.Errorf("Result was incorrect, got %s, want %s.", *got3, "nil")
	}

	result4, _ := tc.GetTransaction("1MI3iWk0MBEe2BwpMDpsiX1UWgkTg5Oh6KJzG9DrCxM=")
	got4 := result4.GetContractFunction()

	if *got4 != "mint" {
		t.Errorf("Result was incorrect, got %s, want %s.", *got4, "mint")
	}

	result5, _ := tc.GetTransaction("KEErop0BQhU0vM8sQNb2X0IygEVUar60XBn9eOFqlTQ=")
	got5 := result5.GetContractFunction()

	if *got5 != "join" {
		t.Errorf("Result was incorrect, got %s, want %s.", *got5, "join")
	}

	result6, _ := tc.GetTransaction("f+eG0euJ2s5X/md68P+gQwHwVCn6PEAtdUXC+ocOeQA=")
	got6 := result6.GetContractFunction()

	if *got6 != "new_order" {
		t.Errorf("Result was incorrect, got %s, want %s.", *got6, "new_order")
	}

	result7, _ := tc.GetTransaction("ytTaRyKDc/eBVxlfijdOAPQZbvA42LlGQoK+s2XJhz0=")
	got7 := result7.GetContractFunction()

	if *got7 != "create" {
		t.Errorf("Result was incorrect, got %s, want %s.", *got7, "create")
	}
}

func TestGetRawContractArguments(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetTransaction("WCyfaBotBOKG1qBwakZ92wKtm79itrRpvJVj9zLaFIY=")
	got1, _ := result1.GetRawContractArguments()
	exp1 := []interface{}{"Example NFT", "An NFT created by Sui Wallet", "ipfs://bafkreibngqhl3gaa7daob4i2vccziay2jjlp435cf66vhono7nrvww53ty"}

	if !reflect.DeepEqual(got1, exp1) {
		t.Errorf("Result was incorrect, got %s, want %s.", got1, exp1)
	}

	result2, _ := tc.GetTransaction("tsl8gn6Wbeq17lmh8MCRkacdHYCiuKDj7r80MD3OhpA=")
	got2, _ := result2.GetRawContractArguments()

	if got2 != nil {
		t.Errorf("Result was incorrect, got %s, want %s.", got2, "")
	}

	result3, _ := tc.GetTransaction("qurUKphegqg9QJjmqyoxFg2U5G+HAK49Wbk0YxDOI7c=")
	got3, _ := result3.GetRawContractArguments()

	if got3 != nil {
		t.Errorf("Result was incorrect, got %s, want %s.", got3, "")
	}

	result4, _ := tc.GetTransaction("1MI3iWk0MBEe2BwpMDpsiX1UWgkTg5Oh6KJzG9DrCxM=")
	got4, _ := result4.GetRawContractArguments()
	exp4 := []interface{}{"0x376084b3f088a16d4612269eef3f71b6953f4872",
		"0x357d18562aaad69251088bcd2dff9c7d54851f91",
		"olegopimah",
		"https://sui-name-service.herokuapp.com/olegopimah",
		"0x98c091a1bec98dc11c8b03390b0d85b817387e6"}

	if !reflect.DeepEqual(got4, exp4) {
		t.Errorf("Result was incorrect, got %s, want %s.", got4, exp4)
	}

	result5, _ := tc.GetTransaction("KEErop0BQhU0vM8sQNb2X0IygEVUar60XBn9eOFqlTQ=")
	got5, _ := result5.GetRawContractArguments()
	exp5 := []interface{}{"0x1428785d4a757154008beedf95598da4829d0202",
		"0x306c4bcacce95292bd573a1bc7a44f2889bd4be5"}

	if !reflect.DeepEqual(got5, exp5) {
		t.Errorf("Result was incorrect, got %s, want %s.", got5, exp5)
	}

	result6, _ := tc.GetTransaction("f+eG0euJ2s5X/md68P+gQwHwVCn6PEAtdUXC+ocOeQA=")
	got6, _ := result6.GetRawContractArguments()
	exp6 := []interface{}{"0xdcb4a80adab73a99e419241a0df5f22bd5df58eb",
		"0x5171d12b094e918abea348854d840a5b81cb8e06",
		"0x742cca31b7fb39070edf23f1eefed746f6e5cc30",
		"0x3c18ae6cde75b48702a737f7d6663359bd4865c1",
		"0xde1e02902f1c591d6e71d68d41e663105a4e8f25",
		[]interface{}{
			float64(1),
			float64(0),
			float64(0),
			float64(0),
			float64(0),
			float64(0),
			float64(0),
			float64(0),
			float64(0),
			float64(0),
			float64(0),
			float64(0),
			float64(0),
			float64(0),
			float64(0),
			float64(0),
		}, float64(12884901888),
		float64(11),
		float64(3245234524),
		float64(123),
		"",
		float64(1),
		float64(1),
		float64(4),
		float64(2)}

	if !reflect.DeepEqual(got6, exp6) {
		t.Errorf("Result was incorrect, got %s, want %s.", got6, exp6)
	}

	result7, _ := tc.GetTransaction("ytTaRyKDc/eBVxlfijdOAPQZbvA42LlGQoK+s2XJhz0=")
	got7, _ := result7.GetRawContractArguments()
	exp7 := []interface{}{"", float64(255), ""}

	if !reflect.DeepEqual(got7, exp7) {
		t.Errorf("Result was incorrect, got %s, want %s.", got7, exp7)
	}
}

func TestGetObject(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetObject("0xb12cb53ad48c6ace9231dec7495bc8bae86ffb66")
	got1 := result1.Result.Details.Data.DataType
	exp1 := "moveObject"

	if got1 != exp1 {
		t.Errorf("Result was incorrect, got %s, want %s.", got1, exp1)
	}

	result2, _ := tc.GetObject("0x56fe7fe3f0fc3b47a1e8eaeb41467294a354f517")
	got2 := result2.Result.Details.Data.Fields["id"].(map[string]interface{})["id"]
	exp2 := "0x56fe7fe3f0fc3b47a1e8eaeb41467294a354f517"

	if got2 != exp2 {
		t.Errorf("Result was incorrect, got %s, want %s.", got2, exp2)
	}

	result3, _ := tc.GetObject("0xdcb4a80adab73a99e419241a0df5f22bd5df58eb")
	got3 := result3.Result.Details.Data.Fields["asks"].(map[string]interface{})["type"]
	exp3 := "0x6774c132abf5521bf9bd8e8c6948c03f0f0cb3f4::critbit::CritBitTree<0x6774c132abf5521bf9bd8e8c6948c03f0f0cb3f4::market::Order>"

	if got3 != exp3 {
		t.Errorf("Result was incorrect, got %s, want %s.", got3, exp3)
	}

	result4, _ := tc.GetObject("0x5171d12b094e918abea348854d840a5b81cb8e06")
	got4 := result4.Result.Details.Data.Fields["events"].([]interface{})[0].(map[string]interface{})["fields"].(map[string]interface{})["account_id"]
	exp4 := "0x742cca31b7fb39070edf23f1eefed746f6e5cc30"

	if got4 != exp4 {
		t.Errorf("Result was incorrect, got %s, want %s.", got4, exp4)
	}
}

func TestGetDataType(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetObject("0xb12cb53ad48c6ace9231dec7495bc8bae86ffb66")
	got1 := result1.GetObjectDataType()
	exp1 := "moveObject"

	if got1 != exp1 {
		t.Errorf("Result was incorrect, got %s, want %s.", got1, exp1)
	}

	result2, _ := tc.GetObject("0x56fe7fe3f0fc3b47a1e8eaeb41467294a354f517")
	got2 := result2.GetObjectDataType()
	exp2 := "moveObject"

	if got2 != exp2 {
		t.Errorf("Result was incorrect, got %s, want %s.", got2, exp2)
	}

	result3, _ := tc.GetObject("0xdcb4a80adab73a99e419241a0df5f22bd5df58eb")
	got3 := result3.GetObjectDataType()
	exp3 := "moveObject"

	if got3 != exp3 {
		t.Errorf("Result was incorrect, got %s, want %s.", got3, exp3)
	}

	result4, _ := tc.GetObject("0x5171d12b094e918abea348854d840a5b81cb8e06")
	got4 := result4.GetObjectDataType()
	exp4 := "moveObject"

	if got4 != exp4 {
		t.Errorf("Result was incorrect, got %s, want %s.", got4, exp4)
	}
}

func TestGetObjectType(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetObject("0xb12cb53ad48c6ace9231dec7495bc8bae86ffb66")
	got1 := result1.GetObjectType()
	exp1 := "0x2::devnet_nft::DevNetNFT"

	if got1 != exp1 {
		t.Errorf("Result was incorrect, got %s, want %s.", got1, exp1)
	}

	result2, _ := tc.GetObject("0x56fe7fe3f0fc3b47a1e8eaeb41467294a354f517")
	got2 := result2.GetObjectType()
	exp2 := "0x2::coin::Coin<0x2::sui::SUI>"

	if got2 != exp2 {
		t.Errorf("Result was incorrect, got %s, want %s.", got2, exp2)
	}

	result3, _ := tc.GetObject("0xdcb4a80adab73a99e419241a0df5f22bd5df58eb")
	got3 := result3.GetObjectType()
	exp3 := "0x6774c132abf5521bf9bd8e8c6948c03f0f0cb3f4::market::Market<0x6774c132abf5521bf9bd8e8c6948c03f0f0cb3f4::base_coin::BASE_COIN, 0x6774c132abf5521bf9bd8e8c6948c03f0f0cb3f4::quote_coin::QUOTE_COIN>"

	if got3 != exp3 {
		t.Errorf("Result was incorrect, got %s, want %s.", got3, exp3)
	}

	result4, _ := tc.GetObject("0x5171d12b094e918abea348854d840a5b81cb8e06")
	got4 := result4.GetObjectType()
	exp4 := "0x6774c132abf5521bf9bd8e8c6948c03f0f0cb3f4::market::EventQueue"

	if got4 != exp4 {
		t.Errorf("Result was incorrect, got %s, want %s.", got4, exp4)
	}
}

func TestHasPublicTransfer(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetObject("0xb12cb53ad48c6ace9231dec7495bc8bae86ffb66")
	got1 := result1.HasPublicTransfer()
	exp1 := true

	if got1 != exp1 {
		t.Errorf("Result was incorrect, got %t, want %t.", got1, exp1)
	}

	result2, _ := tc.GetObject("0x56fe7fe3f0fc3b47a1e8eaeb41467294a354f517")
	got2 := result2.HasPublicTransfer()
	exp2 := true

	if got2 != exp2 {
		t.Errorf("Result was incorrect, got %t, want %t.", got2, exp2)
	}

	result3, _ := tc.GetObject("0xdcb4a80adab73a99e419241a0df5f22bd5df58eb")
	got3 := result3.HasPublicTransfer()
	exp3 := false

	if got3 != exp3 {
		t.Errorf("Result was incorrect, got %t, want %t.", got3, exp3)
	}

	result4, _ := tc.GetObject("0x5171d12b094e918abea348854d840a5b81cb8e06")
	got4 := result4.HasPublicTransfer()
	exp4 := false

	if got4 != exp4 {
		t.Errorf("Result was incorrect, got %t, want %t.", got4, exp4)
	}
}

func TestGetOwner(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetObject("0xb12cb53ad48c6ace9231dec7495bc8bae86ffb66")
	got1 := result1.GetOwner()
	exp1 := "0x661e76241de2413cb5d7f4ce380e283125941b1b"

	if got1 != exp1 {
		t.Errorf("Result was incorrect, got %s, want %s.", got1, exp1)
	}

	result2, _ := tc.GetObject("0x56fe7fe3f0fc3b47a1e8eaeb41467294a354f517")
	got2 := result2.GetOwner()
	exp2 := "0x7b91e88803a832a836037285bd708c9d43803c16"

	if got2 != exp2 {
		t.Errorf("Result was incorrect, got %s, want %s.", got2, exp2)
	}

	result3, _ := tc.GetObject("0xdcb4a80adab73a99e419241a0df5f22bd5df58eb")
	got3 := result3.GetOwner()
	exp3 := "Shared"

	if got3 != exp3 {
		t.Errorf("Result was incorrect, got %s, want %s.", got3, exp3)
	}

	result4, _ := tc.GetObject("0x5171d12b094e918abea348854d840a5b81cb8e06")
	got4 := result4.GetOwner()
	exp4 := "Shared"

	if got4 != exp4 {
		t.Errorf("Result was incorrect, got %s, want %s.", got4, exp4)
	}
}

func TestGetObjectID(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetObject("0xb12cb53ad48c6ace9231dec7495bc8bae86ffb66")
	got1 := result1.GetObjectID()
	exp1 := "0xb12cb53ad48c6ace9231dec7495bc8bae86ffb66"

	if got1 != exp1 {
		t.Errorf("Result was incorrect, got %s, want %s.", got1, exp1)
	}

	result2, _ := tc.GetObject("0x56fe7fe3f0fc3b47a1e8eaeb41467294a354f517")
	got2 := result2.GetObjectID()
	exp2 := "0x56fe7fe3f0fc3b47a1e8eaeb41467294a354f517"

	if got2 != exp2 {
		t.Errorf("Result was incorrect, got %s, want %s.", got2, exp2)
	}

	result3, _ := tc.GetObject("0xdcb4a80adab73a99e419241a0df5f22bd5df58eb")
	got3 := result3.GetObjectID()
	exp3 := "0xdcb4a80adab73a99e419241a0df5f22bd5df58eb"

	if got3 != exp3 {
		t.Errorf("Result was incorrect, got %s, want %s.", got3, exp3)
	}

	result4, _ := tc.GetObject("0x5171d12b094e918abea348854d840a5b81cb8e06")
	got4 := result4.GetObjectID()
	exp4 := "0x5171d12b094e918abea348854d840a5b81cb8e06"

	if got4 != exp4 {
		t.Errorf("Result was incorrect, got %s, want %s.", got4, exp4)
	}
}

func TestGetObjectPackage(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetObject("0xb12cb53ad48c6ace9231dec7495bc8bae86ffb66")
	got1 := result1.GetObjectPackage()
	exp1 := "0x2"

	if got1 != exp1 {
		t.Errorf("Result was incorrect, got %s, want %s.", got1, exp1)
	}

	result2, _ := tc.GetObject("0x56fe7fe3f0fc3b47a1e8eaeb41467294a354f517")
	got2 := result2.GetObjectPackage()
	exp2 := "0x2"

	if got2 != exp2 {
		t.Errorf("Result was incorrect, got %s, want %s.", got2, exp2)
	}

	result3, _ := tc.GetObject("0xdcb4a80adab73a99e419241a0df5f22bd5df58eb")
	got3 := result3.GetObjectPackage()
	exp3 := "0x6774c132abf5521bf9bd8e8c6948c03f0f0cb3f4"

	if got3 != exp3 {
		t.Errorf("Result was incorrect, got %s, want %s.", got3, exp3)
	}

	result4, _ := tc.GetObject("0x13db212f43507081be243dbd18d79df761a29953")
	got4 := result4.GetObjectPackage()
	exp4 := "0xb82c2e33acecc71f2b2b742c366d902bd640a1bf"

	if got4 != exp4 {
		t.Errorf("Result was incorrect, got %s, want %s.", got4, exp4)
	}
}

func TestGetObjectModule(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetObject("0xb12cb53ad48c6ace9231dec7495bc8bae86ffb66")
	got1 := result1.GetObjectModule()
	exp1 := "0x2::devnet_nft"

	if got1 != exp1 {
		t.Errorf("Result was incorrect, got %s, want %s.", got1, exp1)
	}

	result2, _ := tc.GetObject("0x56fe7fe3f0fc3b47a1e8eaeb41467294a354f517")
	got2 := result2.GetObjectModule()
	exp2 := "0x2::coin"

	if got2 != exp2 {
		t.Errorf("Result was incorrect, got %s, want %s.", got2, exp2)
	}

	result3, _ := tc.GetObject("0xdcb4a80adab73a99e419241a0df5f22bd5df58eb")
	got3 := result3.GetObjectModule()
	exp3 := "0x6774c132abf5521bf9bd8e8c6948c03f0f0cb3f4::market"

	if got3 != exp3 {
		t.Errorf("Result was incorrect, got %s, want %s.", got3, exp3)
	}

	result4, _ := tc.GetObject("0x13db212f43507081be243dbd18d79df761a29953")
	got4 := result4.GetObjectModule()
	exp4 := "0xb82c2e33acecc71f2b2b742c366d902bd640a1bf::rgb"

	if got4 != exp4 {
		t.Errorf("Result was incorrect, got %s, want %s.", got4, exp4)
	}
}

func TestGetObjectMetadata(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetObject("0xb12cb53ad48c6ace9231dec7495bc8bae86ffb66")
	got1 := result1.GetObjectMetadata()
	exp1 := map[string]interface{}{
		"description": "An NFT created by Sui Wallet",
		"id": map[string]interface{}{
			"id": "0xb12cb53ad48c6ace9231dec7495bc8bae86ffb66"},
		"name": "Example NFT",
		"url":  "ipfs://bafkreibngqhl3gaa7daob4i2vccziay2jjlp435cf66vhono7nrvww53ty",
	}

	if !reflect.DeepEqual(got1, exp1) {
		t.Errorf("Result was incorrect, got %s, want %s.", got1, exp1)
	}

	result2, _ := tc.GetObject("0x56fe7fe3f0fc3b47a1e8eaeb41467294a354f517")
	got2 := result2.GetObjectMetadata()
	exp2 := map[string]interface{}{
		"balance": float64(44923),
		"id": map[string]interface{}{
			"id": "0x56fe7fe3f0fc3b47a1e8eaeb41467294a354f517",
		},
	}

	if !reflect.DeepEqual(got2, exp2) {
		t.Errorf("Result was incorrect, got %s, want %s.", got2, exp2)
	}

	result3, _ := tc.GetObject("0x13db212f43507081be243dbd18d79df761a29953")
	got3 := result3.GetObjectMetadata()
	exp3 := map[string]interface{}{
		"blue":  float64(0),
		"green": float64(255),
		"id": map[string]interface{}{
			"id": "0x13db212f43507081be243dbd18d79df761a29953",
		},
		"red": float64(0),
	}

	if !reflect.DeepEqual(got3, exp3) {
		t.Errorf("Result was incorrect, got %s, want %s.", got3, exp3)
	}

	result4, _ := tc.GetObject("0x42fe8e8893586133fe7b5e5e225b288baed58ec7")
	got4 := result4.GetObjectMetadata()
	exp4 := map[string]interface{}{
		"description": "An NFT created by Sui Wallet",
		"id": map[string]interface{}{
			"id": "0x42fe8e8893586133fe7b5e5e225b288baed58ec7",
		},
		"name": "Example NFT",
		"url":  "ipfs://bafkreibngqhl3gaa7daob4i2vccziay2jjlp435cf66vhono7nrvww53ty",
	}

	if !reflect.DeepEqual(got4, exp4) {
		t.Errorf("Result was incorrect, got %s, want %s.", got4, exp4)
	}
}

func TestGetAccount(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetAccount("0x66f68a701b4cb5f2b3d3446c093eb91fae8af34f")
	got1 := result1.Balance
	exp1 := uint64(250000)

	if got1 != exp1 {
		t.Errorf("Result was incorrect, got %d, want %d.", got1, exp1)
	}

	result2, _ := tc.GetAccount("0x6a2c51a7621b7d6a967c9b630a2bbaea12b49f1d")
	got2 := false
	exp2 := true

	for _, v := range result2.Objects {
		if v.ObjectId == "0x1fa4bdef12c4b531e4d0d53dd4dfd55914af53fc" {
			got2 = true
		}
	}

	if got2 != exp2 {
		t.Errorf("Result was incorrect, got %t, want %t.", got2, exp2)
	}

	result3, _ := tc.GetAccount("0x018309d97970190a54e3a84f3459f976ebf0ac7e")
	got3 := false
	exp3 := true

	for _, v := range result3.Transactions {
		if v == "kQZs3QvSPN4ij7xKIJOkzK9S83Bj6EHq1nk2ncUjfXY=" {
			got3 = true
		}
	}

	if got3 != exp3 {
		t.Errorf("Result was incorrect, got %t, want %t.", got3, exp3)
	}

}

func TestGetAccountNFTs(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetAccount("0x66f68a701b4cb5f2b3d3446c093eb91fae8af34f")
	got1 := result1.GetAccountNFTs()
	exp1 := []AccObject{}

	if !reflect.DeepEqual(got1, exp1) {
		t.Errorf("Result was incorrect, got %v, want %v.", got1, exp1)
	}

	result2, _ := tc.GetAccount("0x1a6254d89ee1698ed62c03481d28eee88c685b94")
	got2 := result2.GetAccountNFTs()
	exp2 := []AccObject{{
		ObjectId: "0xa61de7bb233df7870bca7ed3459f1261f393ec7f",
		Type:     "0x95cd99feeeb3f49a52f2b4267743f551c828d5b2::rgb::ColorObject",
		Metadata: map[string]interface{}{
			"blue":  float64(0),
			"green": float64(255),
			"red":   float64(0),
		},
	}}

	if !reflect.DeepEqual(got2, exp2) {
		t.Errorf("Result was incorrect, got %+v, want %+v.", got2, exp2)
	}

}

func TestGetGetContractDeploy(t *testing.T) {
	tc := new(SUIClient)
	tc.Init("http://158.140.129.74:9000")

	result1, _ := tc.GetTransaction("xVFZ7KO5uMrBSLtwLku4TMLZaIEMmhob0G/erUmTa6U=")
	got1, _ := result1.GetContractDeploy()
	suiname_nft := "// Move bytecode v5\nmodule 0.suiname_nft {\nstruct AdminCap has key {\n\tid: UID\n}\nstruct GroupsInfo has key {\n\tid: UID,\n\tdata: VecMap<u8, ID>\n}\nstruct NFTMinted has copy, drop {\n\tobject_id: ID,\n\tcreator: address,\n\tname: String\n}\nstruct NamesGroup has key {\n\tid: UID,\n\ttype: u8,\n\tnames: VecMap<String, ID>,\n\tbalance: Balance<SUI>\n}\nstruct SuiNameNFT has store, key {\n\tid: UID,\n\tname: String,\n\turl: Url,\n\tis_active: bool\n}\n\nentry public change_name_status(Arg0: &mut SuiNameNFT, Arg1: bool) {\nB0:\n\t0: MoveLoc[1](Arg1: bool)\n\t1: MoveLoc[0](Arg0: &mut SuiNameNFT)\n\t2: MutBorrowField[0](SuiNameNFT.is_active: bool)\n\t3: WriteRef\n\t4: Ret\n}\nentry public collect_payments(Arg0: &AdminCap, Arg1: &mut NamesGroup, Arg2: &mut TxContext) {\nB0:\n\t0: CopyLoc[1](Arg1: &mut NamesGroup)\n\t1: ImmBorrowField[1](NamesGroup.balance: Balance<SUI>)\n\t2: Call[0](value<SUI>(&Balance<SUI>): u64)\n\t3: StLoc[3](loc0: u64)\n\t4: MoveLoc[1](Arg1: &mut NamesGroup)\n\t5: MutBorrowField[1](NamesGroup.balance: Balance<SUI>)\n\t6: MoveLoc[3](loc0: u64)\n\t7: CopyLoc[2](Arg2: &mut TxContext)\n\t8: Call[1](take<SUI>(&mut Balance<SUI>, u64, &mut TxContext): Coin<SUI>)\n\t9: StLoc[4](loc1: Coin<SUI>)\n\t10: MoveLoc[4](loc1: Coin<SUI>)\n\t11: MoveLoc[2](Arg2: &mut TxContext)\n\t12: FreezeRef\n\t13: Call[12](sender(&TxContext): address)\n\t14: Call[2](transfer<Coin<SUI>>(Coin<SUI>, address))\n\t15: Ret\n}\npublic get_name(Arg0: &SuiNameNFT): String {\nB0:\n\t0: MoveLoc[0](Arg0: &SuiNameNFT)\n\t1: ImmBorrowField[2](SuiNameNFT.name: String)\n\t2: ReadRef\n\t3: Ret\n}\npublic get_price(Arg0: &vector<u8>): u64 {\nL0:\tloc1: u64\nB0:\n\t0: MoveLoc[0](Arg0: &vector<u8>)\n\t1: VecLen(7)\n\t2: StLoc[1](loc0: u64)\n\t3: CopyLoc[1](loc0: u64)\n\t4: LdU64(1)\n\t5: Eq\n\t6: BrTrue(8)\nB1:\n\t7: Branch(11)\nB2:\n\t8: LdU64(5000)\n\t9: StLoc[2](loc1: u64)\n\t10: Branch(29)\nB3:\n\t11: CopyLoc[1](loc0: u64)\n\t12: LdU64(2)\n\t13: Eq\n\t14: BrTrue(16)\nB4:\n\t15: Branch(19)\nB5:\n\t16: LdU64(3000)\n\t17: StLoc[2](loc1: u64)\n\t18: Branch(29)\nB6:\n\t19: MoveLoc[1](loc0: u64)\n\t20: LdU64(3)\n\t21: Eq\n\t22: BrTrue(24)\nB7:\n\t23: Branch(27)\nB8:\n\t24: LdU64(2000)\n\t25: StLoc[2](loc1: u64)\n\t26: Branch(29)\nB9:\n\t27: LdU64(1000)\n\t28: StLoc[2](loc1: u64)\nB10:\n\t29: MoveLoc[2](loc1: u64)\n\t30: Ret\n}\npublic get_type(Arg0: &vector<u8>): u8 {\nL0:\tloc1: u64\nL1:\tloc2: u64\nB0:\n\t0: LdU8(0)\n\t1: StLoc[1](loc0: u8)\n\t2: CopyLoc[0](Arg0: &vector<u8>)\n\t3: VecLen(7)\n\t4: StLoc[3](loc2: u64)\n\t5: LdU64(0)\n\t6: StLoc[2](loc1: u64)\nB1:\n\t7: CopyLoc[2](loc1: u64)\n\t8: CopyLoc[3](loc2: u64)\n\t9: Lt\n\t10: BrTrue(12)\nB2:\n\t11: Branch(26)\nB3:\n\t12: MoveLoc[1](loc0: u8)\n\t13: CopyLoc[0](Arg0: &vector<u8>)\n\t14: CopyLoc[2](loc1: u64)\n\t15: VecImmBorrow(7)\n\t16: ReadRef\n\t17: Add\n\t18: LdConst[7](U8: [64])\n\t19: Mod\n\t20: StLoc[1](loc0: u8)\n\t21: MoveLoc[2](loc1: u64)\n\t22: LdU64(1)\n\t23: Add\n\t24: StLoc[2](loc1: u64)\n\t25: Branch(7)\nB4:\n\t26: MoveLoc[0](Arg0: &vector<u8>)\n\t27: Pop\n\t28: MoveLoc[1](loc0: u8)\n\t29: LdU8(1)\n\t30: Add\n\t31: Ret\n}\ninit(Arg0: &mut TxContext) {\nL0:\tloc1: VecMap<u8, ID>\nL1:\tloc2: UID\nB0:\n\t0: CopyLoc[0](Arg0: &mut TxContext)\n\t1: Call[14](new(&mut TxContext): UID)\n\t2: Pack[0](AdminCap)\n\t3: CopyLoc[0](Arg0: &mut TxContext)\n\t4: FreezeRef\n\t5: Call[12](sender(&TxContext): address)\n\t6: Call[3](transfer<AdminCap>(AdminCap, address))\n\t7: Call[4](empty<u8, ID>(): VecMap<u8, ID>)\n\t8: StLoc[2](loc1: VecMap<u8, ID>)\n\t9: LdU8(1)\n\t10: StLoc[1](loc0: u8)\nB1:\n\t11: CopyLoc[1](loc0: u8)\n\t12: LdConst[7](U8: [64])\n\t13: Le\n\t14: BrTrue(16)\nB2:\n\t15: Branch(35)\nB3:\n\t16: CopyLoc[0](Arg0: &mut TxContext)\n\t17: Call[14](new(&mut TxContext): UID)\n\t18: StLoc[3](loc2: UID)\n\t19: MutBorrowLoc[2](loc1: VecMap<u8, ID>)\n\t20: CopyLoc[1](loc0: u8)\n\t21: ImmBorrowLoc[3](loc2: UID)\n\t22: Call[16](uid_to_inner(&UID): ID)\n\t23: Call[5](insert<u8, ID>(&mut VecMap<u8, ID>, u8, ID))\n\t24: MoveLoc[3](loc2: UID)\n\t25: CopyLoc[1](loc0: u8)\n\t26: Call[6](empty<String, ID>(): VecMap<String, ID>)\n\t27: Call[7](zero<SUI>(): Balance<SUI>)\n\t28: Pack[3](NamesGroup)\n\t29: Call[8](share_object<NamesGroup>(NamesGroup))\n\t30: MoveLoc[1](loc0: u8)\n\t31: LdU8(1)\n\t32: Add\n\t33: StLoc[1](loc0: u8)\n\t34: Branch(11)\nB4:\n\t35: MoveLoc[0](Arg0: &mut TxContext)\n\t36: Call[14](new(&mut TxContext): UID)\n\t37: MoveLoc[2](loc1: VecMap<u8, ID>)\n\t38: Pack[1](GroupsInfo)\n\t39: Call[9](share_object<GroupsInfo>(GroupsInfo))\n\t40: Ret\n}\npublic is_name_available(Arg0: &VecMap<String, ID>, Arg1: &String): bool {\nB0:\n\t0: MoveLoc[0](Arg0: &VecMap<String, ID>)\n\t1: MoveLoc[1](Arg1: &String)\n\t2: Call[10](contains<String, ID>(&VecMap<String, ID>, &String): bool)\n\t3: Not\n\t4: Ret\n}\npublic is_name_correct(Arg0: &vector<u8>): bool {\nL0:\tloc1: bool\nL1:\tloc2: bool\nL2:\tloc3: bool\nL3:\tloc4: u8\nL4:\tloc5: u64\nL5:\tloc6: u64\nB0:\n\t0: CopyLoc[0](Arg0: &vector<u8>)\n\t1: VecLen(7)\n\t2: StLoc[7](loc6: u64)\n\t3: CopyLoc[7](loc6: u64)\n\t4: LdU64(1)\n\t5: Lt\n\t6: BrTrue(8)\nB1:\n\t7: Branch(11)\nB2:\n\t8: LdTrue\n\t9: StLoc[1](loc0: bool)\n\t10: Branch(15)\nB3:\n\t11: CopyLoc[7](loc6: u64)\n\t12: LdU64(24)\n\t13: Gt\n\t14: StLoc[1](loc0: bool)\nB4:\n\t15: MoveLoc[1](loc0: bool)\n\t16: BrTrue(18)\nB5:\n\t17: Branch(22)\nB6:\n\t18: MoveLoc[0](Arg0: &vector<u8>)\n\t19: Pop\n\t20: LdFalse\n\t21: Ret\nB7:\n\t22: LdU64(0)\n\t23: StLoc[6](loc5: u64)\nB8:\n\t24: CopyLoc[6](loc5: u64)\n\t25: CopyLoc[7](loc6: u64)\n\t26: Lt\n\t27: BrTrue(29)\nB9:\n\t28: Branch(79)\nB10:\n\t29: CopyLoc[0](Arg0: &vector<u8>)\n\t30: CopyLoc[6](loc5: u64)\n\t31: VecImmBorrow(7)\n\t32: ReadRef\n\t33: StLoc[5](loc4: u8)\n\t34: LdU8(48)\n\t35: CopyLoc[5](loc4: u8)\n\t36: Le\n\t37: BrTrue(39)\nB11:\n\t38: Branch(44)\nB12:\n\t39: CopyLoc[5](loc4: u8)\n\t40: LdU8(57)\n\t41: Le\n\t42: StLoc[2](loc1: bool)\n\t43: Branch(46)\nB13:\n\t44: LdFalse\n\t45: StLoc[2](loc1: bool)\nB14:\n\t46: MoveLoc[2](loc1: bool)\n\t47: BrTrue(49)\nB15:\n\t48: Branch(52)\nB16:\n\t49: LdTrue\n\t50: StLoc[4](loc3: bool)\n\t51: Branch(66)\nB17:\n\t52: LdU8(97)\n\t53: CopyLoc[5](loc4: u8)\n\t54: Le\n\t55: BrTrue(57)\nB18:\n\t56: Branch(62)\nB19:\n\t57: MoveLoc[5](loc4: u8)\n\t58: LdU8(122)\n\t59: Le\n\t60: StLoc[3](loc2: bool)\n\t61: Branch(64)\nB20:\n\t62: LdFalse\n\t63: StLoc[3](loc2: bool)\nB21:\n\t64: MoveLoc[3](loc2: bool)\n\t65: StLoc[4](loc3: bool)\nB22:\n\t66: MoveLoc[4](loc3: bool)\n\t67: Not\n\t68: BrTrue(70)\nB23:\n\t69: Branch(74)\nB24:\n\t70: MoveLoc[0](Arg0: &vector<u8>)\n\t71: Pop\n\t72: LdFalse\n\t73: Ret\nB25:\n\t74: MoveLoc[6](loc5: u64)\n\t75: LdU64(1)\n\t76: Add\n\t77: StLoc[6](loc5: u64)\n\t78: Branch(24)\nB26:\n\t79: MoveLoc[0](Arg0: &vector<u8>)\n\t80: Pop\n\t81: LdTrue\n\t82: Ret\n}\nentry public mint(Arg0: &mut NamesGroup, Arg1: &GroupsInfo, Arg2: vector<u8>, Arg3: vector<u8>, Arg4: &mut Coin<SUI>, Arg5: &mut TxContext) {\nB0:\n\t0: CopyLoc[5](Arg5: &mut TxContext)\n\t1: FreezeRef\n\t2: Call[12](sender(&TxContext): address)\n\t3: StLoc[9](loc3: address)\n\t4: CopyLoc[9](loc3: address)\n\t5: LdConst[8](Address: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0])\n\t6: Neq\n\t7: BrTrue(18)\nB1:\n\t8: MoveLoc[4](Arg4: &mut Coin<SUI>)\n\t9: Pop\n\t10: MoveLoc[0](Arg0: &mut NamesGroup)\n\t11: Pop\n\t12: MoveLoc[1](Arg1: &GroupsInfo)\n\t13: Pop\n\t14: MoveLoc[5](Arg5: &mut TxContext)\n\t15: Pop\n\t16: LdConst[1](U64: [106, 0, 0, 0, 0, 0, 0, 0])\n\t17: Abort\nB2:\n\t18: ImmBorrowLoc[2](Arg2: vector<u8>)\n\t19: Call[7](is_name_correct(&vector<u8>): bool)\n\t20: BrTrue(31)\nB3:\n\t21: MoveLoc[4](Arg4: &mut Coin<SUI>)\n\t22: Pop\n\t23: MoveLoc[0](Arg0: &mut NamesGroup)\n\t24: Pop\n\t25: MoveLoc[1](Arg1: &GroupsInfo)\n\t26: Pop\n\t27: MoveLoc[5](Arg5: &mut TxContext)\n\t28: Pop\n\t29: LdConst[2](U64: [100, 0, 0, 0, 0, 0, 0, 0])\n\t30: Abort\nB4:\n\t31: ImmBorrowLoc[2](Arg2: vector<u8>)\n\t32: Call[4](get_type(&vector<u8>): u8)\n\t33: StLoc[11](loc5: u8)\n\t34: CopyLoc[11](loc5: u8)\n\t35: LdU8(0)\n\t36: Gt\n\t37: BrTrue(48)\nB5:\n\t38: MoveLoc[4](Arg4: &mut Coin<SUI>)\n\t39: Pop\n\t40: MoveLoc[0](Arg0: &mut NamesGroup)\n\t41: Pop\n\t42: MoveLoc[1](Arg1: &GroupsInfo)\n\t43: Pop\n\t44: MoveLoc[5](Arg5: &mut TxContext)\n\t45: Pop\n\t46: LdConst[6](U64: [104, 0, 0, 0, 0, 0, 0, 0])\n\t47: Abort\nB6:\n\t48: MoveLoc[1](Arg1: &GroupsInfo)\n\t49: ImmBorrowField[3](GroupsInfo.data: VecMap<u8, ID>)\n\t50: ImmBorrowLoc[11](loc5: u8)\n\t51: Call[11](get<u8, ID>(&VecMap<u8, ID>, &u8): &ID)\n\t52: CopyLoc[0](Arg0: &mut NamesGroup)\n\t53: ImmBorrowField[4](NamesGroup.id: UID)\n\t54: Call[22](uid_as_inner(&UID): &ID)\n\t55: Eq\n\t56: BrTrue(65)\nB7:\n\t57: MoveLoc[4](Arg4: &mut Coin<SUI>)\n\t58: Pop\n\t59: MoveLoc[0](Arg0: &mut NamesGroup)\n\t60: Pop\n\t61: MoveLoc[5](Arg5: &mut TxContext)\n\t62: Pop\n\t63: LdConst[5](U64: [105, 0, 0, 0, 0, 0, 0, 0])\n\t64: Abort\nB8:\n\t65: CopyLoc[2](Arg2: vector<u8>)\n\t66: Call[23](string_unsafe(vector<u8>): String)\n\t67: StLoc[10](loc4: String)\n\t68: CopyLoc[0](Arg0: &mut NamesGroup)\n\t69: ImmBorrowField[5](NamesGroup.names: VecMap<String, ID>)\n\t70: ImmBorrowLoc[10](loc4: String)\n\t71: Call[6](is_name_available(&VecMap<String, ID>, &String): bool)\n\t72: BrTrue(81)\nB9:\n\t73: MoveLoc[4](Arg4: &mut Coin<SUI>)\n\t74: Pop\n\t75: MoveLoc[0](Arg0: &mut NamesGroup)\n\t76: Pop\n\t77: MoveLoc[5](Arg5: &mut TxContext)\n\t78: Pop\n\t79: LdConst[3](U64: [101, 0, 0, 0, 0, 0, 0, 0])\n\t80: Abort\nB10:\n\t81: ImmBorrowLoc[2](Arg2: vector<u8>)\n\t82: Call[3](get_price(&vector<u8>): u64)\n\t83: StLoc[8](loc2: u64)\n\t84: CopyLoc[8](loc2: u64)\n\t85: CopyLoc[4](Arg4: &mut Coin<SUI>)\n\t86: FreezeRef\n\t87: Call[12](value<SUI>(&Coin<SUI>): u64)\n\t88: Le\n\t89: BrTrue(98)\nB11:\n\t90: MoveLoc[4](Arg4: &mut Coin<SUI>)\n\t91: Pop\n\t92: MoveLoc[0](Arg0: &mut NamesGroup)\n\t93: Pop\n\t94: MoveLoc[5](Arg5: &mut TxContext)\n\t95: Pop\n\t96: LdConst[0](U64: [103, 0, 0, 0, 0, 0, 0, 0])\n\t97: Abort\nB12:\n\t98: MoveLoc[4](Arg4: &mut Coin<SUI>)\n\t99: Call[13](balance_mut<SUI>(&mut Coin<SUI>): &mut Balance<SUI>)\n\t100: StLoc[6](loc0: &mut Balance<SUI>)\n\t101: CopyLoc[0](Arg0: &mut NamesGroup)\n\t102: MutBorrowField[1](NamesGroup.balance: Balance<SUI>)\n\t103: MoveLoc[6](loc0: &mut Balance<SUI>)\n\t104: MoveLoc[8](loc2: u64)\n\t105: Call[14](split<SUI>(&mut Balance<SUI>, u64): Balance<SUI>)\n\t106: Call[15](join<SUI>(&mut Balance<SUI>, Balance<SUI>): u64)\n\t107: Pop\n\t108: MoveLoc[5](Arg5: &mut TxContext)\n\t109: Call[14](new(&mut TxContext): UID)\n\t110: CopyLoc[10](loc4: String)\n\t111: MoveLoc[3](Arg3: vector<u8>)\n\t112: Call[28](new_unsafe_from_bytes(vector<u8>): Url)\n\t113: LdFalse\n\t114: Pack[4](SuiNameNFT)\n\t115: StLoc[7](loc1: SuiNameNFT)\n\t116: ImmBorrowLoc[7](loc1: SuiNameNFT)\n\t117: ImmBorrowField[6](SuiNameNFT.id: UID)\n\t118: Call[16](uid_to_inner(&UID): ID)\n\t119: CopyLoc[9](loc3: address)\n\t120: ImmBorrowLoc[7](loc1: SuiNameNFT)\n\t121: ImmBorrowField[2](SuiNameNFT.name: String)\n\t122: ReadRef\n\t123: Pack[2](NFTMinted)\n\t124: Call[16](emit<NFTMinted>(NFTMinted))\n\t125: MoveLoc[0](Arg0: &mut NamesGroup)\n\t126: MutBorrowField[5](NamesGroup.names: VecMap<String, ID>)\n\t127: MoveLoc[10](loc4: String)\n\t128: ImmBorrowLoc[7](loc1: SuiNameNFT)\n\t129: ImmBorrowField[6](SuiNameNFT.id: UID)\n\t130: Call[16](uid_to_inner(&UID): ID)\n\t131: Call[17](insert<String, ID>(&mut VecMap<String, ID>, String, ID))\n\t132: MoveLoc[7](loc1: SuiNameNFT)\n\t133: MoveLoc[9](loc3: address)\n\t134: Call[18](transfer<SuiNameNFT>(SuiNameNFT, address))\n\t135: Ret\n}\nentry public transfer(Arg0: SuiNameNFT, Arg1: address) {\nB0:\n\t0: MoveLoc[0](Arg0: SuiNameNFT)\n\t1: MoveLoc[1](Arg1: address)\n\t2: Call[18](transfer<SuiNameNFT>(SuiNameNFT, address))\n\t3: Ret\n}\n}"
	exp1 := Package{
		DeployTX: "xVFZ7KO5uMrBSLtwLku4TMLZaIEMmhob0G/erUmTa6U=",
		ID:       "0xa41ca8645474a1cacc08ce4559275ba07136c412",
		Bytecode: map[string]interface{}{
			"suiname_nft": suiname_nft,
		},
	}

	if !reflect.DeepEqual(got1, exp1) {
		t.Errorf("Result was incorrect, got %s, want %s.", got1, exp1)
	}

	result2, _ := tc.GetTransaction("AXVQbp46p/THvGhkEoP4/fIFdnnFyNHxYccBJqACzs8=")
	got2, _ := result2.GetContractDeploy()
	base_coin := "// Move bytecode v5\nmodule 0.base_coin {\nstruct BASE_COIN has drop {\n\tdummy_field: bool\n}\n\ninit(Arg0: BASE_COIN, Arg1: &mut TxContext) {\nB0:\n\t0: MoveLoc[0](Arg0: BASE_COIN)\n\t1: MoveLoc[1](Arg1: &mut TxContext)\n\t2: Call[0](create_currency<BASE_COIN>(BASE_COIN, &mut TxContext): TreasuryCap<BASE_COIN>)\n\t3: StLoc[2](loc0: TreasuryCap<BASE_COIN>)\n\t4: MoveLoc[2](loc0: TreasuryCap<BASE_COIN>)\n\t5: Call[1](share_object<TreasuryCap<BASE_COIN>>(TreasuryCap<BASE_COIN>))\n\t6: Ret\n}\nentry public mint(Arg0: &mut TreasuryCap<BASE_COIN>, Arg1: u64, Arg2: &mut TxContext) {\nB0:\n\t0: MoveLoc[0](Arg0: &mut TreasuryCap<BASE_COIN>)\n\t1: MoveLoc[1](Arg1: u64)\n\t2: CopyLoc[2](Arg2: &mut TxContext)\n\t3: Call[2](mint<BASE_COIN>(&mut TreasuryCap<BASE_COIN>, u64, &mut TxContext): Coin<BASE_COIN>)\n\t4: StLoc[3](loc0: Coin<BASE_COIN>)\n\t5: MoveLoc[3](loc0: Coin<BASE_COIN>)\n\t6: MoveLoc[2](Arg2: &mut TxContext)\n\t7: FreezeRef\n\t8: Call[6](sender(&TxContext): address)\n\t9: Call[3](transfer<BASE_COIN>(Coin<BASE_COIN>, address))\n\t10: Ret\n}\nentry public mint_zero(Arg0: &mut TxContext) {\nB0:\n\t0: CopyLoc[0](Arg0: &mut TxContext)\n\t1: Call[4](zero<BASE_COIN>(&mut TxContext): Coin<BASE_COIN>)\n\t2: StLoc[1](loc0: Coin<BASE_COIN>)\n\t3: MoveLoc[1](loc0: Coin<BASE_COIN>)\n\t4: MoveLoc[0](Arg0: &mut TxContext)\n\t5: FreezeRef\n\t6: Call[6](sender(&TxContext): address)\n\t7: Call[3](transfer<BASE_COIN>(Coin<BASE_COIN>, address))\n\t8: Ret\n}\n}"
	critbit := "// Move bytecode v5\nmodule 0.critbit {\nstruct CritBitTree<Ty0> has store {\n\troot: u64,\n\tinner_nodes: vector<InnerNode>,\n\touter_nodes: vector<OuterNode<Ty0>>\n}\nstruct InnerNode has store {\n\tcritical_bit: u8,\n\tparent_index: u64,\n\tleft_child_index: u64,\n\tright_child_index: u64\n}\nstruct OuterNode<Ty0> has store {\n\tkey: u128,\n\tvalue: Ty0,\n\tparent_index: u64\n}\n\npublic borrow<Ty0>(Arg0: &CritBitTree<Ty0>, Arg1: u128): &Ty0 {\nB0:\n\t0: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t1: Call[0](is_empty<Ty0>(&CritBitTree<Ty0>): bool)\n\t2: Not\n\t3: BrTrue(8)\nB1:\n\t4: MoveLoc[0](Arg0: &CritBitTree<Ty0>)\n\t5: Pop\n\t6: LdConst[1](U64: [3, 0, 0, 0, 0, 0, 0, 0])\n\t7: Abort\nB2:\n\t8: MoveLoc[0](Arg0: &CritBitTree<Ty0>)\n\t9: CopyLoc[1](Arg1: u128)\n\t10: Call[1](borrow_closest_outer_node<Ty0>(&CritBitTree<Ty0>, u128): &OuterNode<Ty0>)\n\t11: StLoc[2](loc0: &OuterNode<Ty0>)\n\t12: CopyLoc[2](loc0: &OuterNode<Ty0>)\n\t13: ImmBorrowFieldGeneric[0](OuterNode.key: u128)\n\t14: ReadRef\n\t15: MoveLoc[1](Arg1: u128)\n\t16: Eq\n\t17: BrTrue(22)\nB3:\n\t18: MoveLoc[2](loc0: &OuterNode<Ty0>)\n\t19: Pop\n\t20: LdConst[6](U64: [4, 0, 0, 0, 0, 0, 0, 0])\n\t21: Abort\nB4:\n\t22: MoveLoc[2](loc0: &OuterNode<Ty0>)\n\t23: ImmBorrowFieldGeneric[1](OuterNode.value: Ty0)\n\t24: Ret\n}\nborrow_closest_outer_node<Ty0>(Arg0: &CritBitTree<Ty0>, Arg1: u128): &OuterNode<Ty0> {\nL0:\tloc2: &InnerNode\nB0:\n\t0: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t1: ImmBorrowFieldGeneric[2](CritBitTree.root: u64)\n\t2: ReadRef\n\t3: Call[18](is_outer_node(u64): bool)\n\t4: BrTrue(6)\nB1:\n\t5: Branch(14)\nB2:\n\t6: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t7: ImmBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t8: MoveLoc[0](Arg0: &CritBitTree<Ty0>)\n\t9: ImmBorrowFieldGeneric[2](CritBitTree.root: u64)\n\t10: ReadRef\n\t11: Call[26](outer_node_vector_index(u64): u64)\n\t12: VecImmBorrow(37)\n\t13: Ret\nB3:\n\t14: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t15: ImmBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t16: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t17: ImmBorrowFieldGeneric[2](CritBitTree.root: u64)\n\t18: ReadRef\n\t19: VecImmBorrow(38)\n\t20: StLoc[4](loc2: &InnerNode)\nB4:\n\t21: CopyLoc[1](Arg1: u128)\n\t22: CopyLoc[4](loc2: &InnerNode)\n\t23: ImmBorrowField[5](InnerNode.critical_bit: u8)\n\t24: ReadRef\n\t25: Call[19](is_set(u128, u8): bool)\n\t26: BrTrue(28)\nB5:\n\t27: Branch(33)\nB6:\n\t28: MoveLoc[4](loc2: &InnerNode)\n\t29: ImmBorrowField[6](InnerNode.right_child_index: u64)\n\t30: ReadRef\n\t31: StLoc[2](loc0: u64)\n\t32: Branch(37)\nB7:\n\t33: MoveLoc[4](loc2: &InnerNode)\n\t34: ImmBorrowField[7](InnerNode.left_child_index: u64)\n\t35: ReadRef\n\t36: StLoc[2](loc0: u64)\nB8:\n\t37: MoveLoc[2](loc0: u64)\n\t38: StLoc[3](loc1: u64)\n\t39: CopyLoc[3](loc1: u64)\n\t40: Call[18](is_outer_node(u64): bool)\n\t41: BrTrue(43)\nB9:\n\t42: Branch(49)\nB10:\n\t43: MoveLoc[0](Arg0: &CritBitTree<Ty0>)\n\t44: ImmBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t45: MoveLoc[3](loc1: u64)\n\t46: Call[26](outer_node_vector_index(u64): u64)\n\t47: VecImmBorrow(37)\n\t48: Ret\nB11:\n\t49: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t50: ImmBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t51: MoveLoc[3](loc1: u64)\n\t52: VecImmBorrow(38)\n\t53: StLoc[4](loc2: &InnerNode)\n\t54: Branch(21)\n}\nborrow_closest_outer_node_mut<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128): &mut OuterNode<Ty0> {\nL0:\tloc2: &InnerNode\nB0:\n\t0: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: ImmBorrowFieldGeneric[2](CritBitTree.root: u64)\n\t2: ReadRef\n\t3: Call[18](is_outer_node(u64): bool)\n\t4: BrTrue(6)\nB1:\n\t5: Branch(14)\nB2:\n\t6: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t7: MutBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t8: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t9: ImmBorrowFieldGeneric[2](CritBitTree.root: u64)\n\t10: ReadRef\n\t11: Call[26](outer_node_vector_index(u64): u64)\n\t12: VecMutBorrow(37)\n\t13: Ret\nB3:\n\t14: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t15: ImmBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t16: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t17: ImmBorrowFieldGeneric[2](CritBitTree.root: u64)\n\t18: ReadRef\n\t19: VecImmBorrow(38)\n\t20: StLoc[4](loc2: &InnerNode)\nB4:\n\t21: CopyLoc[1](Arg1: u128)\n\t22: CopyLoc[4](loc2: &InnerNode)\n\t23: ImmBorrowField[5](InnerNode.critical_bit: u8)\n\t24: ReadRef\n\t25: Call[19](is_set(u128, u8): bool)\n\t26: BrTrue(28)\nB5:\n\t27: Branch(33)\nB6:\n\t28: MoveLoc[4](loc2: &InnerNode)\n\t29: ImmBorrowField[6](InnerNode.right_child_index: u64)\n\t30: ReadRef\n\t31: StLoc[2](loc0: u64)\n\t32: Branch(37)\nB7:\n\t33: MoveLoc[4](loc2: &InnerNode)\n\t34: ImmBorrowField[7](InnerNode.left_child_index: u64)\n\t35: ReadRef\n\t36: StLoc[2](loc0: u64)\nB8:\n\t37: MoveLoc[2](loc0: u64)\n\t38: StLoc[3](loc1: u64)\n\t39: CopyLoc[3](loc1: u64)\n\t40: Call[18](is_outer_node(u64): bool)\n\t41: BrTrue(43)\nB9:\n\t42: Branch(49)\nB10:\n\t43: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t44: MutBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t45: MoveLoc[3](loc1: u64)\n\t46: Call[26](outer_node_vector_index(u64): u64)\n\t47: VecMutBorrow(37)\n\t48: Ret\nB11:\n\t49: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t50: ImmBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t51: MoveLoc[3](loc1: u64)\n\t52: VecImmBorrow(38)\n\t53: StLoc[4](loc2: &InnerNode)\n\t54: Branch(21)\n}\npublic borrow_mut<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128): &mut Ty0 {\nB0:\n\t0: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: FreezeRef\n\t2: Call[0](is_empty<Ty0>(&CritBitTree<Ty0>): bool)\n\t3: Not\n\t4: BrTrue(9)\nB1:\n\t5: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t6: Pop\n\t7: LdConst[1](U64: [3, 0, 0, 0, 0, 0, 0, 0])\n\t8: Abort\nB2:\n\t9: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t10: CopyLoc[1](Arg1: u128)\n\t11: Call[2](borrow_closest_outer_node_mut<Ty0>(&mut CritBitTree<Ty0>, u128): &mut OuterNode<Ty0>)\n\t12: StLoc[2](loc0: &mut OuterNode<Ty0>)\n\t13: CopyLoc[2](loc0: &mut OuterNode<Ty0>)\n\t14: ImmBorrowFieldGeneric[0](OuterNode.key: u128)\n\t15: ReadRef\n\t16: MoveLoc[1](Arg1: u128)\n\t17: Eq\n\t18: BrTrue(23)\nB3:\n\t19: MoveLoc[2](loc0: &mut OuterNode<Ty0>)\n\t20: Pop\n\t21: LdConst[6](U64: [4, 0, 0, 0, 0, 0, 0, 0])\n\t22: Abort\nB4:\n\t23: MoveLoc[2](loc0: &mut OuterNode<Ty0>)\n\t24: MutBorrowFieldGeneric[1](OuterNode.value: Ty0)\n\t25: Ret\n}\ncheck_length(Arg0: u64) {\nB0:\n\t0: MoveLoc[0](Arg0: u64)\n\t1: LdConst[9](U64: [255, 255, 255, 255, 255, 255, 255, 255])\n\t2: LdConst[2](U64: [1, 0, 0, 0, 0, 0, 0, 0])\n\t3: LdConst[12](U8: [63])\n\t4: Shl\n\t5: Xor\n\t6: Lt\n\t7: BrTrue(10)\nB1:\n\t8: LdConst[4](U64: [5, 0, 0, 0, 0, 0, 0, 0])\n\t9: Abort\nB2:\n\t10: Ret\n}\ncrit_bit(Arg0: u128, Arg1: u128): u8 {\nL0:\tloc2: u128\nL1:\tloc3: u8\nL2:\tloc4: u128\nB0:\n\t0: MoveLoc[0](Arg0: u128)\n\t1: MoveLoc[1](Arg1: u128)\n\t2: Xor\n\t3: StLoc[6](loc4: u128)\n\t4: LdU8(0)\n\t5: StLoc[2](loc0: u8)\n\t6: LdConst[11](U8: [127])\n\t7: StLoc[5](loc3: u8)\nB1:\n\t8: CopyLoc[2](loc0: u8)\n\t9: CopyLoc[5](loc3: u8)\n\t10: Add\n\t11: LdU8(2)\n\t12: Div\n\t13: StLoc[3](loc1: u8)\n\t14: CopyLoc[6](loc4: u128)\n\t15: CopyLoc[3](loc1: u8)\n\t16: Shr\n\t17: StLoc[4](loc2: u128)\n\t18: CopyLoc[4](loc2: u128)\n\t19: LdU128(1)\n\t20: Eq\n\t21: BrTrue(23)\nB2:\n\t22: Branch(25)\nB3:\n\t23: MoveLoc[3](loc1: u8)\n\t24: Ret\nB4:\n\t25: MoveLoc[4](loc2: u128)\n\t26: LdU128(1)\n\t27: Gt\n\t28: BrTrue(30)\nB5:\n\t29: Branch(35)\nB6:\n\t30: MoveLoc[3](loc1: u8)\n\t31: LdU8(1)\n\t32: Add\n\t33: StLoc[2](loc0: u8)\n\t34: Branch(39)\nB7:\n\t35: MoveLoc[3](loc1: u8)\n\t36: LdU8(1)\n\t37: Sub\n\t38: StLoc[5](loc3: u8)\nB8:\n\t39: Branch(8)\n}\npublic destroy_empty<Ty0>(Arg0: CritBitTree<Ty0>) {\nL0:\tloc1: vector<OuterNode<Ty0>>\nB0:\n\t0: ImmBorrowLoc[0](Arg0: CritBitTree<Ty0>)\n\t1: Call[0](is_empty<Ty0>(&CritBitTree<Ty0>): bool)\n\t2: BrTrue(5)\nB1:\n\t3: LdConst[2](U64: [1, 0, 0, 0, 0, 0, 0, 0])\n\t4: Abort\nB2:\n\t5: MoveLoc[0](Arg0: CritBitTree<Ty0>)\n\t6: UnpackGeneric[0](CritBitTree<Ty0>)\n\t7: StLoc[2](loc1: vector<OuterNode<Ty0>>)\n\t8: StLoc[1](loc0: vector<InnerNode>)\n\t9: Pop\n\t10: MoveLoc[1](loc0: vector<InnerNode>)\n\t11: VecUnpack(38, 0)\n\t12: MoveLoc[2](loc1: vector<OuterNode<Ty0>>)\n\t13: VecUnpack(37, 0)\n\t14: Ret\n}\npublic empty<Ty0>(): CritBitTree<Ty0> {\nB0:\n\t0: LdU64(0)\n\t1: VecPack(38, 0)\n\t2: VecPack(37, 0)\n\t3: PackGeneric[0](CritBitTree<Ty0>)\n\t4: Ret\n}\npublic has_key<Ty0>(Arg0: &CritBitTree<Ty0>, Arg1: u128): bool {\nB0:\n\t0: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t1: Call[0](is_empty<Ty0>(&CritBitTree<Ty0>): bool)\n\t2: BrTrue(4)\nB1:\n\t3: Branch(8)\nB2:\n\t4: MoveLoc[0](Arg0: &CritBitTree<Ty0>)\n\t5: Pop\n\t6: LdFalse\n\t7: Ret\nB3:\n\t8: MoveLoc[0](Arg0: &CritBitTree<Ty0>)\n\t9: CopyLoc[1](Arg1: u128)\n\t10: Call[1](borrow_closest_outer_node<Ty0>(&CritBitTree<Ty0>, u128): &OuterNode<Ty0>)\n\t11: ImmBorrowFieldGeneric[0](OuterNode.key: u128)\n\t12: ReadRef\n\t13: MoveLoc[1](Arg1: u128)\n\t14: Eq\n\t15: Ret\n}\npublic insert<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128, Arg2: Ty0) {\nB0:\n\t0: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: FreezeRef\n\t2: Call[3](length<Ty0>(&CritBitTree<Ty0>): u64)\n\t3: StLoc[3](loc0: u64)\n\t4: CopyLoc[3](loc0: u64)\n\t5: Call[4](check_length(u64))\n\t6: CopyLoc[3](loc0: u64)\n\t7: LdU64(0)\n\t8: Eq\n\t9: BrTrue(11)\nB1:\n\t10: Branch(16)\nB2:\n\t11: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t12: MoveLoc[1](Arg1: u128)\n\t13: MoveLoc[2](Arg2: Ty0)\n\t14: Call[4](insert_empty<Ty0>(&mut CritBitTree<Ty0>, u128, Ty0))\n\t15: Branch(31)\nB3:\n\t16: CopyLoc[3](loc0: u64)\n\t17: LdU64(1)\n\t18: Eq\n\t19: BrTrue(21)\nB4:\n\t20: Branch(26)\nB5:\n\t21: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t22: MoveLoc[1](Arg1: u128)\n\t23: MoveLoc[2](Arg2: Ty0)\n\t24: Call[5](insert_singleton<Ty0>(&mut CritBitTree<Ty0>, u128, Ty0))\n\t25: Branch(31)\nB6:\n\t26: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t27: MoveLoc[1](Arg1: u128)\n\t28: MoveLoc[2](Arg2: Ty0)\n\t29: MoveLoc[3](loc0: u64)\n\t30: Call[6](insert_general<Ty0>(&mut CritBitTree<Ty0>, u128, Ty0, u64))\nB7:\n\t31: Ret\n}\ninsert_above<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128, Arg2: Ty0, Arg3: u64, Arg4: u64, Arg5: u64, Arg6: u8) {\nB0:\n\t0: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: ImmBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t2: MoveLoc[5](Arg5: u64)\n\t3: VecImmBorrow(38)\n\t4: ImmBorrowField[8](InnerNode.parent_index: u64)\n\t5: ReadRef\n\t6: StLoc[8](loc1: u64)\nB1:\n\t7: CopyLoc[8](loc1: u64)\n\t8: LdConst[9](U64: [255, 255, 255, 255, 255, 255, 255, 255])\n\t9: Eq\n\t10: BrTrue(12)\nB2:\n\t11: Branch(20)\nB3:\n\t12: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t13: MoveLoc[1](Arg1: u128)\n\t14: MoveLoc[2](Arg2: Ty0)\n\t15: MoveLoc[3](Arg3: u64)\n\t16: MoveLoc[4](Arg4: u64)\n\t17: MoveLoc[6](Arg6: u8)\n\t18: Call[7](insert_above_root<Ty0>(&mut CritBitTree<Ty0>, u128, Ty0, u64, u64, u8))\n\t19: Ret\nB4:\n\t20: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t21: MutBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t22: CopyLoc[8](loc1: u64)\n\t23: VecMutBorrow(38)\n\t24: StLoc[7](loc0: &mut InnerNode)\n\t25: CopyLoc[6](Arg6: u8)\n\t26: CopyLoc[7](loc0: &mut InnerNode)\n\t27: ImmBorrowField[5](InnerNode.critical_bit: u8)\n\t28: ReadRef\n\t29: Lt\n\t30: BrTrue(32)\nB5:\n\t31: Branch(43)\nB6:\n\t32: MoveLoc[7](loc0: &mut InnerNode)\n\t33: Pop\n\t34: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t35: MoveLoc[1](Arg1: u128)\n\t36: MoveLoc[2](Arg2: Ty0)\n\t37: MoveLoc[3](Arg3: u64)\n\t38: MoveLoc[4](Arg4: u64)\n\t39: MoveLoc[8](loc1: u64)\n\t40: MoveLoc[6](Arg6: u8)\n\t41: Call[8](insert_below_walk<Ty0>(&mut CritBitTree<Ty0>, u128, Ty0, u64, u64, u64, u8))\n\t42: Ret\nB7:\n\t43: MoveLoc[7](loc0: &mut InnerNode)\n\t44: ImmBorrowField[8](InnerNode.parent_index: u64)\n\t45: ReadRef\n\t46: StLoc[8](loc1: u64)\n\t47: Branch(7)\n}\ninsert_above_root<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128, Arg2: Ty0, Arg3: u64, Arg4: u64, Arg5: u8) {\nB0:\n\t0: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: ImmBorrowFieldGeneric[2](CritBitTree.root: u64)\n\t2: ReadRef\n\t3: StLoc[6](loc0: u64)\n\t4: CopyLoc[4](Arg4: u64)\n\t5: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t6: MutBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t7: CopyLoc[6](loc0: u64)\n\t8: VecMutBorrow(38)\n\t9: MutBorrowField[8](InnerNode.parent_index: u64)\n\t10: WriteRef\n\t11: CopyLoc[4](Arg4: u64)\n\t12: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t13: MutBorrowFieldGeneric[2](CritBitTree.root: u64)\n\t14: WriteRef\n\t15: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t16: CopyLoc[1](Arg1: u128)\n\t17: MoveLoc[2](Arg2: Ty0)\n\t18: MoveLoc[4](Arg4: u64)\n\t19: CopyLoc[5](Arg5: u8)\n\t20: LdConst[9](U64: [255, 255, 255, 255, 255, 255, 255, 255])\n\t21: MoveLoc[1](Arg1: u128)\n\t22: MoveLoc[5](Arg5: u8)\n\t23: Call[19](is_set(u128, u8): bool)\n\t24: MoveLoc[6](loc0: u64)\n\t25: MoveLoc[3](Arg3: u64)\n\t26: Call[25](outer_node_child_index(u64): u64)\n\t27: Call[9](push_back_insert_nodes<Ty0>(&mut CritBitTree<Ty0>, u128, Ty0, u64, u8, u64, bool, u64, u64))\n\t28: Ret\n}\ninsert_below<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128, Arg2: Ty0, Arg3: u64, Arg4: u64, Arg5: u64, Arg6: bool, Arg7: u128, Arg8: u64, Arg9: u8) {\nB0:\n\t0: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: MutBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t2: CopyLoc[8](Arg8: u64)\n\t3: VecMutBorrow(38)\n\t4: StLoc[10](loc0: &mut InnerNode)\n\t5: MoveLoc[6](Arg6: bool)\n\t6: LdConst[10](Bool: [1])\n\t7: Eq\n\t8: BrTrue(10)\nB1:\n\t9: Branch(15)\nB2:\n\t10: CopyLoc[4](Arg4: u64)\n\t11: MoveLoc[10](loc0: &mut InnerNode)\n\t12: MutBorrowField[7](InnerNode.left_child_index: u64)\n\t13: WriteRef\n\t14: Branch(19)\nB3:\n\t15: CopyLoc[4](Arg4: u64)\n\t16: MoveLoc[10](loc0: &mut InnerNode)\n\t17: MutBorrowField[6](InnerNode.right_child_index: u64)\n\t18: WriteRef\nB4:\n\t19: CopyLoc[4](Arg4: u64)\n\t20: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t21: MutBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t22: CopyLoc[5](Arg5: u64)\n\t23: Call[26](outer_node_vector_index(u64): u64)\n\t24: VecMutBorrow(37)\n\t25: MutBorrowFieldGeneric[5](OuterNode.parent_index: u64)\n\t26: WriteRef\n\t27: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t28: CopyLoc[1](Arg1: u128)\n\t29: MoveLoc[2](Arg2: Ty0)\n\t30: MoveLoc[4](Arg4: u64)\n\t31: MoveLoc[9](Arg9: u8)\n\t32: MoveLoc[8](Arg8: u64)\n\t33: MoveLoc[1](Arg1: u128)\n\t34: MoveLoc[7](Arg7: u128)\n\t35: Lt\n\t36: MoveLoc[3](Arg3: u64)\n\t37: Call[25](outer_node_child_index(u64): u64)\n\t38: MoveLoc[5](Arg5: u64)\n\t39: Call[9](push_back_insert_nodes<Ty0>(&mut CritBitTree<Ty0>, u128, Ty0, u64, u8, u64, bool, u64, u64))\n\t40: Ret\n}\ninsert_below_walk<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128, Arg2: Ty0, Arg3: u64, Arg4: u64, Arg5: u64, Arg6: u8) {\nB0:\n\t0: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: MutBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t2: CopyLoc[5](Arg5: u64)\n\t3: VecMutBorrow(38)\n\t4: StLoc[9](loc2: &mut InnerNode)\n\t5: CopyLoc[1](Arg1: u128)\n\t6: CopyLoc[9](loc2: &mut InnerNode)\n\t7: ImmBorrowField[5](InnerNode.critical_bit: u8)\n\t8: ReadRef\n\t9: Call[19](is_set(u128, u8): bool)\n\t10: BrTrue(12)\nB1:\n\t11: Branch(19)\nB2:\n\t12: LdConst[13](Bool: [0])\n\t13: CopyLoc[9](loc2: &mut InnerNode)\n\t14: ImmBorrowField[6](InnerNode.right_child_index: u64)\n\t15: ReadRef\n\t16: StLoc[8](loc1: u64)\n\t17: StLoc[7](loc0: bool)\n\t18: Branch(25)\nB3:\n\t19: LdConst[10](Bool: [1])\n\t20: CopyLoc[9](loc2: &mut InnerNode)\n\t21: ImmBorrowField[7](InnerNode.left_child_index: u64)\n\t22: ReadRef\n\t23: StLoc[8](loc1: u64)\n\t24: StLoc[7](loc0: bool)\nB4:\n\t25: MoveLoc[7](loc0: bool)\n\t26: MoveLoc[8](loc1: u64)\n\t27: StLoc[10](loc3: u64)\n\t28: StLoc[11](loc4: bool)\n\t29: MoveLoc[11](loc4: bool)\n\t30: LdConst[10](Bool: [1])\n\t31: Eq\n\t32: BrTrue(34)\nB5:\n\t33: Branch(39)\nB6:\n\t34: CopyLoc[4](Arg4: u64)\n\t35: MoveLoc[9](loc2: &mut InnerNode)\n\t36: MutBorrowField[7](InnerNode.left_child_index: u64)\n\t37: WriteRef\n\t38: Branch(43)\nB7:\n\t39: CopyLoc[4](Arg4: u64)\n\t40: MoveLoc[9](loc2: &mut InnerNode)\n\t41: MutBorrowField[6](InnerNode.right_child_index: u64)\n\t42: WriteRef\nB8:\n\t43: CopyLoc[4](Arg4: u64)\n\t44: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t45: MutBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t46: CopyLoc[10](loc3: u64)\n\t47: VecMutBorrow(38)\n\t48: MutBorrowField[8](InnerNode.parent_index: u64)\n\t49: WriteRef\n\t50: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t51: CopyLoc[1](Arg1: u128)\n\t52: MoveLoc[2](Arg2: Ty0)\n\t53: MoveLoc[4](Arg4: u64)\n\t54: CopyLoc[6](Arg6: u8)\n\t55: MoveLoc[5](Arg5: u64)\n\t56: MoveLoc[1](Arg1: u128)\n\t57: MoveLoc[6](Arg6: u8)\n\t58: Call[19](is_set(u128, u8): bool)\n\t59: MoveLoc[10](loc3: u64)\n\t60: MoveLoc[3](Arg3: u64)\n\t61: Call[25](outer_node_child_index(u64): u64)\n\t62: Call[9](push_back_insert_nodes<Ty0>(&mut CritBitTree<Ty0>, u128, Ty0, u64, u8, u64, bool, u64, u64))\n\t63: Ret\n}\ninsert_empty<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128, Arg2: Ty0) {\nB0:\n\t0: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: MutBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t2: MoveLoc[1](Arg1: u128)\n\t3: MoveLoc[2](Arg2: Ty0)\n\t4: LdConst[9](U64: [255, 255, 255, 255, 255, 255, 255, 255])\n\t5: PackGeneric[1](OuterNode<Ty0>)\n\t6: VecPushBack(37)\n\t7: LdConst[2](U64: [1, 0, 0, 0, 0, 0, 0, 0])\n\t8: LdConst[12](U8: [63])\n\t9: Shl\n\t10: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t11: MutBorrowFieldGeneric[2](CritBitTree.root: u64)\n\t12: WriteRef\n\t13: Ret\n}\ninsert_general<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128, Arg2: Ty0, Arg3: u64) {\nL0:\tloc4: bool\nL1:\tloc5: u64\nL2:\tloc6: u128\nL3:\tloc7: u8\nL4:\tloc8: u64\nB0:\n\t0: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: ImmBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t2: VecLen(38)\n\t3: StLoc[7](loc3: u64)\n\t4: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t5: CopyLoc[1](Arg1: u128)\n\t6: StLoc[5](loc1: u128)\n\t7: StLoc[4](loc0: &mut CritBitTree<Ty0>)\n\t8: MoveLoc[4](loc0: &mut CritBitTree<Ty0>)\n\t9: FreezeRef\n\t10: MoveLoc[5](loc1: u128)\n\t11: Call[10](search_outer<Ty0>(&CritBitTree<Ty0>, u128): u64 * bool * u128 * u64 * u8)\n\t12: StLoc[11](loc7: u8)\n\t13: StLoc[12](loc8: u64)\n\t14: StLoc[10](loc6: u128)\n\t15: StLoc[8](loc4: bool)\n\t16: StLoc[9](loc5: u64)\n\t17: CopyLoc[10](loc6: u128)\n\t18: CopyLoc[1](Arg1: u128)\n\t19: Neq\n\t20: BrTrue(25)\nB1:\n\t21: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t22: Pop\n\t23: LdConst[3](U64: [2, 0, 0, 0, 0, 0, 0, 0])\n\t24: Abort\nB2:\n\t25: CopyLoc[10](loc6: u128)\n\t26: CopyLoc[1](Arg1: u128)\n\t27: Call[5](crit_bit(u128, u128): u8)\n\t28: StLoc[6](loc2: u8)\n\t29: CopyLoc[6](loc2: u8)\n\t30: MoveLoc[11](loc7: u8)\n\t31: Lt\n\t32: BrTrue(34)\nB3:\n\t33: Branch(46)\nB4:\n\t34: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t35: MoveLoc[1](Arg1: u128)\n\t36: MoveLoc[2](Arg2: Ty0)\n\t37: MoveLoc[3](Arg3: u64)\n\t38: MoveLoc[7](loc3: u64)\n\t39: MoveLoc[9](loc5: u64)\n\t40: MoveLoc[8](loc4: bool)\n\t41: MoveLoc[10](loc6: u128)\n\t42: MoveLoc[12](loc8: u64)\n\t43: MoveLoc[6](loc2: u8)\n\t44: Call[11](insert_below<Ty0>(&mut CritBitTree<Ty0>, u128, Ty0, u64, u64, u64, bool, u128, u64, u8))\n\t45: Branch(54)\nB5:\n\t46: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t47: MoveLoc[1](Arg1: u128)\n\t48: MoveLoc[2](Arg2: Ty0)\n\t49: MoveLoc[3](Arg3: u64)\n\t50: MoveLoc[7](loc3: u64)\n\t51: MoveLoc[12](loc8: u64)\n\t52: MoveLoc[6](loc2: u8)\n\t53: Call[12](insert_above<Ty0>(&mut CritBitTree<Ty0>, u128, Ty0, u64, u64, u64, u8))\nB6:\n\t54: Ret\n}\ninsert_singleton<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128, Arg2: Ty0) {\nB0:\n\t0: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: ImmBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t2: LdU64(0)\n\t3: VecImmBorrow(37)\n\t4: StLoc[4](loc1: &OuterNode<Ty0>)\n\t5: CopyLoc[1](Arg1: u128)\n\t6: CopyLoc[4](loc1: &OuterNode<Ty0>)\n\t7: ImmBorrowFieldGeneric[0](OuterNode.key: u128)\n\t8: ReadRef\n\t9: Neq\n\t10: BrTrue(17)\nB1:\n\t11: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t12: Pop\n\t13: MoveLoc[4](loc1: &OuterNode<Ty0>)\n\t14: Pop\n\t15: LdConst[3](U64: [2, 0, 0, 0, 0, 0, 0, 0])\n\t16: Abort\nB2:\n\t17: CopyLoc[4](loc1: &OuterNode<Ty0>)\n\t18: ImmBorrowFieldGeneric[0](OuterNode.key: u128)\n\t19: ReadRef\n\t20: CopyLoc[1](Arg1: u128)\n\t21: Call[5](crit_bit(u128, u128): u8)\n\t22: StLoc[3](loc0: u8)\n\t23: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t24: CopyLoc[1](Arg1: u128)\n\t25: MoveLoc[2](Arg2: Ty0)\n\t26: LdU64(0)\n\t27: MoveLoc[3](loc0: u8)\n\t28: LdConst[9](U64: [255, 255, 255, 255, 255, 255, 255, 255])\n\t29: MoveLoc[1](Arg1: u128)\n\t30: MoveLoc[4](loc1: &OuterNode<Ty0>)\n\t31: ImmBorrowFieldGeneric[0](OuterNode.key: u128)\n\t32: ReadRef\n\t33: Gt\n\t34: LdU64(0)\n\t35: Call[25](outer_node_child_index(u64): u64)\n\t36: LdU64(1)\n\t37: Call[25](outer_node_child_index(u64): u64)\n\t38: Call[9](push_back_insert_nodes<Ty0>(&mut CritBitTree<Ty0>, u128, Ty0, u64, u8, u64, bool, u64, u64))\n\t39: LdU64(0)\n\t40: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t41: MutBorrowFieldGeneric[2](CritBitTree.root: u64)\n\t42: WriteRef\n\t43: LdU64(0)\n\t44: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t45: MutBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t46: LdU64(0)\n\t47: VecMutBorrow(37)\n\t48: MutBorrowFieldGeneric[5](OuterNode.parent_index: u64)\n\t49: WriteRef\n\t50: Ret\n}\npublic is_empty<Ty0>(Arg0: &CritBitTree<Ty0>): bool {\nB0:\n\t0: MoveLoc[0](Arg0: &CritBitTree<Ty0>)\n\t1: ImmBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t2: Call[13](is_empty<OuterNode<Ty0>>(&vector<OuterNode<Ty0>>): bool)\n\t3: Ret\n}\nis_outer_node(Arg0: u64): bool {\nB0:\n\t0: MoveLoc[0](Arg0: u64)\n\t1: LdConst[12](U8: [63])\n\t2: Shr\n\t3: LdConst[2](U64: [1, 0, 0, 0, 0, 0, 0, 0])\n\t4: BitAnd\n\t5: LdConst[2](U64: [1, 0, 0, 0, 0, 0, 0, 0])\n\t6: Eq\n\t7: Ret\n}\nis_set(Arg0: u128, Arg1: u8): bool {\nB0:\n\t0: MoveLoc[0](Arg0: u128)\n\t1: MoveLoc[1](Arg1: u8)\n\t2: Shr\n\t3: LdU128(1)\n\t4: BitAnd\n\t5: LdU128(1)\n\t6: Eq\n\t7: Ret\n}\npublic length<Ty0>(Arg0: &CritBitTree<Ty0>): u64 {\nB0:\n\t0: MoveLoc[0](Arg0: &CritBitTree<Ty0>)\n\t1: ImmBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t2: VecLen(37)\n\t3: Ret\n}\npublic max_key<Ty0>(Arg0: &CritBitTree<Ty0>): u128 {\nB0:\n\t0: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t1: Call[0](is_empty<Ty0>(&CritBitTree<Ty0>): bool)\n\t2: Not\n\t3: BrTrue(8)\nB1:\n\t4: MoveLoc[0](Arg0: &CritBitTree<Ty0>)\n\t5: Pop\n\t6: LdConst[5](U64: [7, 0, 0, 0, 0, 0, 0, 0])\n\t7: Abort\nB2:\n\t8: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t9: ImmBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t10: MoveLoc[0](Arg0: &CritBitTree<Ty0>)\n\t11: Call[14](max_node_child_index<Ty0>(&CritBitTree<Ty0>): u64)\n\t12: Call[26](outer_node_vector_index(u64): u64)\n\t13: VecImmBorrow(37)\n\t14: ImmBorrowFieldGeneric[0](OuterNode.key: u128)\n\t15: ReadRef\n\t16: Ret\n}\nmax_node_child_index<Ty0>(Arg0: &CritBitTree<Ty0>): u64 {\nB0:\n\t0: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t1: ImmBorrowFieldGeneric[2](CritBitTree.root: u64)\n\t2: ReadRef\n\t3: StLoc[1](loc0: u64)\nB1:\n\t4: CopyLoc[1](loc0: u64)\n\t5: Call[18](is_outer_node(u64): bool)\n\t6: BrTrue(8)\nB2:\n\t7: Branch(12)\nB3:\n\t8: MoveLoc[0](Arg0: &CritBitTree<Ty0>)\n\t9: Pop\n\t10: MoveLoc[1](loc0: u64)\n\t11: Ret\nB4:\n\t12: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t13: ImmBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t14: MoveLoc[1](loc0: u64)\n\t15: VecImmBorrow(38)\n\t16: ImmBorrowField[6](InnerNode.right_child_index: u64)\n\t17: ReadRef\n\t18: StLoc[1](loc0: u64)\n\t19: Branch(4)\n}\npublic min_key<Ty0>(Arg0: &CritBitTree<Ty0>): u128 {\nB0:\n\t0: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t1: Call[0](is_empty<Ty0>(&CritBitTree<Ty0>): bool)\n\t2: Not\n\t3: BrTrue(8)\nB1:\n\t4: MoveLoc[0](Arg0: &CritBitTree<Ty0>)\n\t5: Pop\n\t6: LdConst[5](U64: [7, 0, 0, 0, 0, 0, 0, 0])\n\t7: Abort\nB2:\n\t8: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t9: ImmBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t10: MoveLoc[0](Arg0: &CritBitTree<Ty0>)\n\t11: Call[15](min_node_child_index<Ty0>(&CritBitTree<Ty0>): u64)\n\t12: Call[26](outer_node_vector_index(u64): u64)\n\t13: VecImmBorrow(37)\n\t14: ImmBorrowFieldGeneric[0](OuterNode.key: u128)\n\t15: ReadRef\n\t16: Ret\n}\nmin_node_child_index<Ty0>(Arg0: &CritBitTree<Ty0>): u64 {\nB0:\n\t0: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t1: ImmBorrowFieldGeneric[2](CritBitTree.root: u64)\n\t2: ReadRef\n\t3: StLoc[1](loc0: u64)\nB1:\n\t4: CopyLoc[1](loc0: u64)\n\t5: Call[18](is_outer_node(u64): bool)\n\t6: BrTrue(8)\nB2:\n\t7: Branch(12)\nB3:\n\t8: MoveLoc[0](Arg0: &CritBitTree<Ty0>)\n\t9: Pop\n\t10: MoveLoc[1](loc0: u64)\n\t11: Ret\nB4:\n\t12: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t13: ImmBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t14: MoveLoc[1](loc0: u64)\n\t15: VecImmBorrow(38)\n\t16: ImmBorrowField[7](InnerNode.left_child_index: u64)\n\t17: ReadRef\n\t18: StLoc[1](loc0: u64)\n\t19: Branch(4)\n}\nouter_node_child_index(Arg0: u64): u64 {\nB0:\n\t0: MoveLoc[0](Arg0: u64)\n\t1: LdConst[2](U64: [1, 0, 0, 0, 0, 0, 0, 0])\n\t2: LdConst[12](U8: [63])\n\t3: Shl\n\t4: BitOr\n\t5: Ret\n}\nouter_node_vector_index(Arg0: u64): u64 {\nB0:\n\t0: MoveLoc[0](Arg0: u64)\n\t1: LdConst[9](U64: [255, 255, 255, 255, 255, 255, 255, 255])\n\t2: BitAnd\n\t3: LdConst[2](U64: [1, 0, 0, 0, 0, 0, 0, 0])\n\t4: LdConst[12](U8: [63])\n\t5: Shl\n\t6: Xor\n\t7: Ret\n}\npublic pop<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128): Ty0 {\nB0:\n\t0: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: FreezeRef\n\t2: Call[0](is_empty<Ty0>(&CritBitTree<Ty0>): bool)\n\t3: Not\n\t4: BrTrue(9)\nB1:\n\t5: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t6: Pop\n\t7: LdConst[7](U64: [6, 0, 0, 0, 0, 0, 0, 0])\n\t8: Abort\nB2:\n\t9: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t10: FreezeRef\n\t11: Call[3](length<Ty0>(&CritBitTree<Ty0>): u64)\n\t12: StLoc[3](loc1: u64)\n\t13: CopyLoc[3](loc1: u64)\n\t14: LdU64(1)\n\t15: Eq\n\t16: BrTrue(18)\nB3:\n\t17: Branch(23)\nB4:\n\t18: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t19: MoveLoc[1](Arg1: u128)\n\t20: Call[16](pop_singleton<Ty0>(&mut CritBitTree<Ty0>, u128): Ty0)\n\t21: StLoc[2](loc0: Ty0)\n\t22: Branch(28)\nB5:\n\t23: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t24: MoveLoc[1](Arg1: u128)\n\t25: MoveLoc[3](loc1: u64)\n\t26: Call[17](pop_general<Ty0>(&mut CritBitTree<Ty0>, u128, u64): Ty0)\n\t27: StLoc[2](loc0: Ty0)\nB6:\n\t28: MoveLoc[2](loc0: Ty0)\n\t29: Ret\n}\npop_destroy_nodes<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u64, Arg2: u64, Arg3: u64): Ty0 {\nB0:\n\t0: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: ImmBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t2: VecLen(38)\n\t3: StLoc[4](loc0: u64)\n\t4: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t5: MutBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t6: CopyLoc[1](Arg1: u64)\n\t7: Call[18](swap_remove<InnerNode>(&mut vector<InnerNode>, u64): InnerNode)\n\t8: Unpack[1](InnerNode)\n\t9: Pop\n\t10: Pop\n\t11: Pop\n\t12: Pop\n\t13: CopyLoc[1](Arg1: u64)\n\t14: CopyLoc[4](loc0: u64)\n\t15: LdU64(1)\n\t16: Sub\n\t17: Lt\n\t18: BrTrue(20)\nB1:\n\t19: Branch(24)\nB2:\n\t20: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t21: MoveLoc[1](Arg1: u64)\n\t22: MoveLoc[4](loc0: u64)\n\t23: Call[19](stitch_swap_remove<Ty0>(&mut CritBitTree<Ty0>, u64, u64))\nB3:\n\t24: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t25: MutBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t26: CopyLoc[2](Arg2: u64)\n\t27: Call[26](outer_node_vector_index(u64): u64)\n\t28: Call[20](swap_remove<OuterNode<Ty0>>(&mut vector<OuterNode<Ty0>>, u64): OuterNode<Ty0>)\n\t29: UnpackGeneric[1](OuterNode<Ty0>)\n\t30: Pop\n\t31: StLoc[5](loc1: Ty0)\n\t32: Pop\n\t33: CopyLoc[2](Arg2: u64)\n\t34: Call[26](outer_node_vector_index(u64): u64)\n\t35: CopyLoc[3](Arg3: u64)\n\t36: LdU64(1)\n\t37: Sub\n\t38: Lt\n\t39: BrTrue(41)\nB4:\n\t40: Branch(46)\nB5:\n\t41: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t42: MoveLoc[2](Arg2: u64)\n\t43: MoveLoc[3](Arg3: u64)\n\t44: Call[19](stitch_swap_remove<Ty0>(&mut CritBitTree<Ty0>, u64, u64))\n\t45: Branch(48)\nB6:\n\t46: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t47: Pop\nB7:\n\t48: MoveLoc[5](loc1: Ty0)\n\t49: Ret\n}\npop_general<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128, Arg2: u64): Ty0 {\nL0:\tloc3: u64\nL1:\tloc4: u128\nL2:\tloc5: u64\nB0:\n\t0: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: CopyLoc[1](Arg1: u128)\n\t2: StLoc[4](loc1: u128)\n\t3: StLoc[3](loc0: &mut CritBitTree<Ty0>)\n\t4: MoveLoc[3](loc0: &mut CritBitTree<Ty0>)\n\t5: FreezeRef\n\t6: MoveLoc[4](loc1: u128)\n\t7: Call[10](search_outer<Ty0>(&CritBitTree<Ty0>, u128): u64 * bool * u128 * u64 * u8)\n\t8: Pop\n\t9: StLoc[8](loc5: u64)\n\t10: StLoc[7](loc4: u128)\n\t11: StLoc[5](loc2: bool)\n\t12: StLoc[6](loc3: u64)\n\t13: MoveLoc[7](loc4: u128)\n\t14: MoveLoc[1](Arg1: u128)\n\t15: Eq\n\t16: BrTrue(21)\nB1:\n\t17: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t18: Pop\n\t19: LdConst[6](U64: [4, 0, 0, 0, 0, 0, 0, 0])\n\t20: Abort\nB2:\n\t21: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t22: MoveLoc[5](loc2: bool)\n\t23: CopyLoc[8](loc5: u64)\n\t24: Call[21](pop_update_relationships<Ty0>(&mut CritBitTree<Ty0>, bool, u64))\n\t25: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t26: MoveLoc[8](loc5: u64)\n\t27: MoveLoc[6](loc3: u64)\n\t28: MoveLoc[2](Arg2: u64)\n\t29: Call[22](pop_destroy_nodes<Ty0>(&mut CritBitTree<Ty0>, u64, u64, u64): Ty0)\n\t30: Ret\n}\npop_singleton<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128): Ty0 {\nB0:\n\t0: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: ImmBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t2: LdU64(0)\n\t3: VecImmBorrow(37)\n\t4: ImmBorrowFieldGeneric[0](OuterNode.key: u128)\n\t5: ReadRef\n\t6: MoveLoc[1](Arg1: u128)\n\t7: Eq\n\t8: BrTrue(13)\nB1:\n\t9: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t10: Pop\n\t11: LdConst[6](U64: [4, 0, 0, 0, 0, 0, 0, 0])\n\t12: Abort\nB2:\n\t13: LdU64(0)\n\t14: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t15: MutBorrowFieldGeneric[2](CritBitTree.root: u64)\n\t16: WriteRef\n\t17: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t18: MutBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t19: VecPopBack(37)\n\t20: UnpackGeneric[1](OuterNode<Ty0>)\n\t21: Pop\n\t22: StLoc[2](loc0: Ty0)\n\t23: Pop\n\t24: MoveLoc[2](loc0: Ty0)\n\t25: Ret\n}\npop_update_relationships<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: bool, Arg2: u64) {\nL0:\tloc3: &InnerNode\nL1:\tloc4: u64\nB0:\n\t0: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: ImmBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t2: CopyLoc[2](Arg2: u64)\n\t3: VecImmBorrow(38)\n\t4: StLoc[6](loc3: &InnerNode)\n\t5: MoveLoc[1](Arg1: bool)\n\t6: LdConst[10](Bool: [1])\n\t7: Eq\n\t8: BrTrue(10)\nB1:\n\t9: Branch(15)\nB2:\n\t10: CopyLoc[6](loc3: &InnerNode)\n\t11: ImmBorrowField[6](InnerNode.right_child_index: u64)\n\t12: ReadRef\n\t13: StLoc[3](loc0: u64)\n\t14: Branch(19)\nB3:\n\t15: CopyLoc[6](loc3: &InnerNode)\n\t16: ImmBorrowField[7](InnerNode.left_child_index: u64)\n\t17: ReadRef\n\t18: StLoc[3](loc0: u64)\nB4:\n\t19: MoveLoc[3](loc0: u64)\n\t20: StLoc[7](loc4: u64)\n\t21: MoveLoc[6](loc3: &InnerNode)\n\t22: ImmBorrowField[8](InnerNode.parent_index: u64)\n\t23: ReadRef\n\t24: StLoc[5](loc2: u64)\n\t25: CopyLoc[7](loc4: u64)\n\t26: Call[18](is_outer_node(u64): bool)\n\t27: BrTrue(29)\nB5:\n\t28: Branch(38)\nB6:\n\t29: CopyLoc[5](loc2: u64)\n\t30: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t31: MutBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t32: CopyLoc[7](loc4: u64)\n\t33: Call[26](outer_node_vector_index(u64): u64)\n\t34: VecMutBorrow(37)\n\t35: MutBorrowFieldGeneric[5](OuterNode.parent_index: u64)\n\t36: WriteRef\n\t37: Branch(45)\nB7:\n\t38: CopyLoc[5](loc2: u64)\n\t39: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t40: MutBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t41: CopyLoc[7](loc4: u64)\n\t42: VecMutBorrow(38)\n\t43: MutBorrowField[8](InnerNode.parent_index: u64)\n\t44: WriteRef\nB8:\n\t45: CopyLoc[5](loc2: u64)\n\t46: LdConst[9](U64: [255, 255, 255, 255, 255, 255, 255, 255])\n\t47: Eq\n\t48: BrTrue(50)\nB9:\n\t49: Branch(55)\nB10:\n\t50: MoveLoc[7](loc4: u64)\n\t51: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t52: MutBorrowFieldGeneric[2](CritBitTree.root: u64)\n\t53: WriteRef\n\t54: Branch(76)\nB11:\n\t55: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t56: MutBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t57: MoveLoc[5](loc2: u64)\n\t58: VecMutBorrow(38)\n\t59: StLoc[4](loc1: &mut InnerNode)\n\t60: CopyLoc[4](loc1: &mut InnerNode)\n\t61: ImmBorrowField[7](InnerNode.left_child_index: u64)\n\t62: ReadRef\n\t63: MoveLoc[2](Arg2: u64)\n\t64: Eq\n\t65: BrTrue(67)\nB12:\n\t66: Branch(72)\nB13:\n\t67: MoveLoc[7](loc4: u64)\n\t68: MoveLoc[4](loc1: &mut InnerNode)\n\t69: MutBorrowField[7](InnerNode.left_child_index: u64)\n\t70: WriteRef\n\t71: Branch(76)\nB14:\n\t72: MoveLoc[7](loc4: u64)\n\t73: MoveLoc[4](loc1: &mut InnerNode)\n\t74: MutBorrowField[6](InnerNode.right_child_index: u64)\n\t75: WriteRef\nB15:\n\t76: Ret\n}\npush_back_insert_nodes<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128, Arg2: Ty0, Arg3: u64, Arg4: u8, Arg5: u64, Arg6: bool, Arg7: u64, Arg8: u64) {\nB0:\n\t0: MoveLoc[6](Arg6: bool)\n\t1: BrTrue(3)\nB1:\n\t2: Branch(8)\nB2:\n\t3: MoveLoc[7](Arg7: u64)\n\t4: MoveLoc[8](Arg8: u64)\n\t5: StLoc[10](loc1: u64)\n\t6: StLoc[9](loc0: u64)\n\t7: Branch(12)\nB3:\n\t8: MoveLoc[8](Arg8: u64)\n\t9: MoveLoc[7](Arg7: u64)\n\t10: StLoc[10](loc1: u64)\n\t11: StLoc[9](loc0: u64)\nB4:\n\t12: MoveLoc[9](loc0: u64)\n\t13: MoveLoc[10](loc1: u64)\n\t14: StLoc[12](loc3: u64)\n\t15: StLoc[11](loc2: u64)\n\t16: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t17: MutBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t18: MoveLoc[1](Arg1: u128)\n\t19: MoveLoc[2](Arg2: Ty0)\n\t20: MoveLoc[3](Arg3: u64)\n\t21: PackGeneric[1](OuterNode<Ty0>)\n\t22: VecPushBack(37)\n\t23: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t24: MutBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t25: MoveLoc[4](Arg4: u8)\n\t26: MoveLoc[5](Arg5: u64)\n\t27: MoveLoc[11](loc2: u64)\n\t28: MoveLoc[12](loc3: u64)\n\t29: Pack[1](InnerNode)\n\t30: VecPushBack(38)\n\t31: Ret\n}\nsearch_outer<Ty0>(Arg0: &CritBitTree<Ty0>, Arg1: u128): u64 * bool * u128 * u64 * u8 {\nL0:\tloc2: u64\nL1:\tloc3: &OuterNode<Ty0>\nL2:\tloc4: &InnerNode\nL3:\tloc5: bool\nB0:\n\t0: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t1: ImmBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t2: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t3: ImmBorrowFieldGeneric[2](CritBitTree.root: u64)\n\t4: ReadRef\n\t5: VecImmBorrow(38)\n\t6: StLoc[6](loc4: &InnerNode)\nB1:\n\t7: CopyLoc[1](Arg1: u128)\n\t8: CopyLoc[6](loc4: &InnerNode)\n\t9: ImmBorrowField[5](InnerNode.critical_bit: u8)\n\t10: ReadRef\n\t11: Call[19](is_set(u128, u8): bool)\n\t12: BrTrue(14)\nB2:\n\t13: Branch(21)\nB3:\n\t14: CopyLoc[6](loc4: &InnerNode)\n\t15: ImmBorrowField[6](InnerNode.right_child_index: u64)\n\t16: ReadRef\n\t17: LdConst[13](Bool: [0])\n\t18: StLoc[3](loc1: bool)\n\t19: StLoc[2](loc0: u64)\n\t20: Branch(27)\nB4:\n\t21: CopyLoc[6](loc4: &InnerNode)\n\t22: ImmBorrowField[7](InnerNode.left_child_index: u64)\n\t23: ReadRef\n\t24: LdConst[10](Bool: [1])\n\t25: StLoc[3](loc1: bool)\n\t26: StLoc[2](loc0: u64)\nB5:\n\t27: MoveLoc[2](loc0: u64)\n\t28: MoveLoc[3](loc1: bool)\n\t29: StLoc[7](loc5: bool)\n\t30: StLoc[4](loc2: u64)\n\t31: CopyLoc[4](loc2: u64)\n\t32: Call[18](is_outer_node(u64): bool)\n\t33: BrTrue(35)\nB6:\n\t34: Branch(53)\nB7:\n\t35: MoveLoc[0](Arg0: &CritBitTree<Ty0>)\n\t36: ImmBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t37: CopyLoc[4](loc2: u64)\n\t38: Call[26](outer_node_vector_index(u64): u64)\n\t39: VecImmBorrow(37)\n\t40: StLoc[5](loc3: &OuterNode<Ty0>)\n\t41: MoveLoc[4](loc2: u64)\n\t42: MoveLoc[7](loc5: bool)\n\t43: CopyLoc[5](loc3: &OuterNode<Ty0>)\n\t44: ImmBorrowFieldGeneric[0](OuterNode.key: u128)\n\t45: ReadRef\n\t46: MoveLoc[5](loc3: &OuterNode<Ty0>)\n\t47: ImmBorrowFieldGeneric[5](OuterNode.parent_index: u64)\n\t48: ReadRef\n\t49: MoveLoc[6](loc4: &InnerNode)\n\t50: ImmBorrowField[5](InnerNode.critical_bit: u8)\n\t51: ReadRef\n\t52: Ret\nB8:\n\t53: MoveLoc[6](loc4: &InnerNode)\n\t54: Pop\n\t55: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t56: ImmBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t57: MoveLoc[4](loc2: u64)\n\t58: VecImmBorrow(38)\n\t59: StLoc[6](loc4: &InnerNode)\n\t60: Branch(7)\n}\npublic singleton<Ty0>(Arg0: u128, Arg1: Ty0): CritBitTree<Ty0> {\nB0:\n\t0: LdU64(0)\n\t1: VecPack(38, 0)\n\t2: VecPack(37, 0)\n\t3: PackGeneric[0](CritBitTree<Ty0>)\n\t4: StLoc[2](loc0: CritBitTree<Ty0>)\n\t5: MutBorrowLoc[2](loc0: CritBitTree<Ty0>)\n\t6: MoveLoc[0](Arg0: u128)\n\t7: MoveLoc[1](Arg1: Ty0)\n\t8: Call[4](insert_empty<Ty0>(&mut CritBitTree<Ty0>, u128, Ty0))\n\t9: MoveLoc[2](loc0: CritBitTree<Ty0>)\n\t10: Ret\n}\nstitch_child_of_parent<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u64, Arg2: u64, Arg3: u64) {\nB0:\n\t0: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: MutBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t2: MoveLoc[2](Arg2: u64)\n\t3: VecMutBorrow(38)\n\t4: StLoc[4](loc0: &mut InnerNode)\n\t5: CopyLoc[4](loc0: &mut InnerNode)\n\t6: ImmBorrowField[7](InnerNode.left_child_index: u64)\n\t7: ReadRef\n\t8: MoveLoc[3](Arg3: u64)\n\t9: Eq\n\t10: BrTrue(12)\nB1:\n\t11: Branch(17)\nB2:\n\t12: MoveLoc[1](Arg1: u64)\n\t13: MoveLoc[4](loc0: &mut InnerNode)\n\t14: MutBorrowField[7](InnerNode.left_child_index: u64)\n\t15: WriteRef\n\t16: Branch(21)\nB3:\n\t17: MoveLoc[1](Arg1: u64)\n\t18: MoveLoc[4](loc0: &mut InnerNode)\n\t19: MutBorrowField[6](InnerNode.right_child_index: u64)\n\t20: WriteRef\nB4:\n\t21: Ret\n}\nstitch_parent_of_child<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u64, Arg2: u64) {\nB0:\n\t0: CopyLoc[2](Arg2: u64)\n\t1: Call[18](is_outer_node(u64): bool)\n\t2: BrTrue(4)\nB1:\n\t3: Branch(13)\nB2:\n\t4: MoveLoc[1](Arg1: u64)\n\t5: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t6: MutBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t7: MoveLoc[2](Arg2: u64)\n\t8: Call[26](outer_node_vector_index(u64): u64)\n\t9: VecMutBorrow(37)\n\t10: MutBorrowFieldGeneric[5](OuterNode.parent_index: u64)\n\t11: WriteRef\n\t12: Branch(20)\nB3:\n\t13: MoveLoc[1](Arg1: u64)\n\t14: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t15: MutBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t16: MoveLoc[2](Arg2: u64)\n\t17: VecMutBorrow(38)\n\t18: MutBorrowField[8](InnerNode.parent_index: u64)\n\t19: WriteRef\nB4:\n\t20: Ret\n}\nstitch_swap_remove<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u64, Arg2: u64) {\nL0:\tloc3: u64\nL1:\tloc4: u64\nB0:\n\t0: CopyLoc[1](Arg1: u64)\n\t1: Call[18](is_outer_node(u64): bool)\n\t2: BrTrue(4)\nB1:\n\t3: Branch(31)\nB2:\n\t4: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t5: ImmBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t6: CopyLoc[1](Arg1: u64)\n\t7: Call[26](outer_node_vector_index(u64): u64)\n\t8: VecImmBorrow(37)\n\t9: ImmBorrowFieldGeneric[5](OuterNode.parent_index: u64)\n\t10: ReadRef\n\t11: StLoc[5](loc2: u64)\n\t12: CopyLoc[5](loc2: u64)\n\t13: LdConst[9](U64: [255, 255, 255, 255, 255, 255, 255, 255])\n\t14: Eq\n\t15: BrTrue(17)\nB3:\n\t16: Branch(22)\nB4:\n\t17: MoveLoc[1](Arg1: u64)\n\t18: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t19: MutBorrowFieldGeneric[2](CritBitTree.root: u64)\n\t20: WriteRef\n\t21: Ret\nB5:\n\t22: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t23: MoveLoc[1](Arg1: u64)\n\t24: MoveLoc[5](loc2: u64)\n\t25: MoveLoc[2](Arg2: u64)\n\t26: LdU64(1)\n\t27: Sub\n\t28: Call[25](outer_node_child_index(u64): u64)\n\t29: Call[23](stitch_child_of_parent<Ty0>(&mut CritBitTree<Ty0>, u64, u64, u64))\n\t30: Branch(73)\nB6:\n\t31: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t32: ImmBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t33: CopyLoc[1](Arg1: u64)\n\t34: VecImmBorrow(38)\n\t35: StLoc[4](loc1: &InnerNode)\n\t36: CopyLoc[4](loc1: &InnerNode)\n\t37: ImmBorrowField[8](InnerNode.parent_index: u64)\n\t38: ReadRef\n\t39: CopyLoc[4](loc1: &InnerNode)\n\t40: ImmBorrowField[7](InnerNode.left_child_index: u64)\n\t41: ReadRef\n\t42: MoveLoc[4](loc1: &InnerNode)\n\t43: ImmBorrowField[6](InnerNode.right_child_index: u64)\n\t44: ReadRef\n\t45: StLoc[7](loc4: u64)\n\t46: StLoc[3](loc0: u64)\n\t47: StLoc[6](loc3: u64)\n\t48: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t49: CopyLoc[1](Arg1: u64)\n\t50: MoveLoc[3](loc0: u64)\n\t51: Call[24](stitch_parent_of_child<Ty0>(&mut CritBitTree<Ty0>, u64, u64))\n\t52: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t53: CopyLoc[1](Arg1: u64)\n\t54: MoveLoc[7](loc4: u64)\n\t55: Call[24](stitch_parent_of_child<Ty0>(&mut CritBitTree<Ty0>, u64, u64))\n\t56: CopyLoc[6](loc3: u64)\n\t57: LdConst[9](U64: [255, 255, 255, 255, 255, 255, 255, 255])\n\t58: Eq\n\t59: BrTrue(61)\nB7:\n\t60: Branch(66)\nB8:\n\t61: MoveLoc[1](Arg1: u64)\n\t62: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t63: MutBorrowFieldGeneric[2](CritBitTree.root: u64)\n\t64: WriteRef\n\t65: Ret\nB9:\n\t66: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t67: MoveLoc[1](Arg1: u64)\n\t68: MoveLoc[6](loc3: u64)\n\t69: MoveLoc[2](Arg2: u64)\n\t70: LdU64(1)\n\t71: Sub\n\t72: Call[23](stitch_child_of_parent<Ty0>(&mut CritBitTree<Ty0>, u64, u64, u64))\nB10:\n\t73: Ret\n}\npublic traverse_end_pop<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u64, Arg2: u64, Arg3: u64): Ty0 {\nB0:\n\t0: CopyLoc[3](Arg3: u64)\n\t1: LdU64(1)\n\t2: Eq\n\t3: BrTrue(5)\nB1:\n\t4: Branch(19)\nB2:\n\t5: LdU64(0)\n\t6: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t7: MutBorrowFieldGeneric[2](CritBitTree.root: u64)\n\t8: WriteRef\n\t9: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t10: MutBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t11: VecPopBack(37)\n\t12: UnpackGeneric[1](OuterNode<Ty0>)\n\t13: Pop\n\t14: StLoc[6](loc2: Ty0)\n\t15: Pop\n\t16: MoveLoc[6](loc2: Ty0)\n\t17: StLoc[4](loc0: Ty0)\n\t18: Branch(38)\nB3:\n\t19: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t20: ImmBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t21: CopyLoc[1](Arg1: u64)\n\t22: VecImmBorrow(38)\n\t23: ImmBorrowField[7](InnerNode.left_child_index: u64)\n\t24: ReadRef\n\t25: CopyLoc[2](Arg2: u64)\n\t26: Eq\n\t27: StLoc[5](loc1: bool)\n\t28: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t29: MoveLoc[5](loc1: bool)\n\t30: CopyLoc[1](Arg1: u64)\n\t31: Call[21](pop_update_relationships<Ty0>(&mut CritBitTree<Ty0>, bool, u64))\n\t32: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t33: MoveLoc[1](Arg1: u64)\n\t34: MoveLoc[2](Arg2: u64)\n\t35: MoveLoc[3](Arg3: u64)\n\t36: Call[22](pop_destroy_nodes<Ty0>(&mut CritBitTree<Ty0>, u64, u64, u64): Ty0)\n\t37: StLoc[4](loc0: Ty0)\nB4:\n\t38: MoveLoc[4](loc0: Ty0)\n\t39: Ret\n}\npublic traverse_init_mut<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: bool): u128 * &mut Ty0 * u64 * u64 {\nL0:\tloc2: &mut OuterNode<Ty0>\nB0:\n\t0: MoveLoc[1](Arg1: bool)\n\t1: LdConst[10](Bool: [1])\n\t2: Eq\n\t3: BrTrue(5)\nB1:\n\t4: Branch(10)\nB2:\n\t5: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t6: FreezeRef\n\t7: Call[14](max_node_child_index<Ty0>(&CritBitTree<Ty0>): u64)\n\t8: StLoc[2](loc0: u64)\n\t9: Branch(14)\nB3:\n\t10: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t11: FreezeRef\n\t12: Call[15](min_node_child_index<Ty0>(&CritBitTree<Ty0>): u64)\n\t13: StLoc[2](loc0: u64)\nB4:\n\t14: MoveLoc[2](loc0: u64)\n\t15: StLoc[3](loc1: u64)\n\t16: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t17: MutBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t18: CopyLoc[3](loc1: u64)\n\t19: Call[26](outer_node_vector_index(u64): u64)\n\t20: VecMutBorrow(37)\n\t21: StLoc[4](loc2: &mut OuterNode<Ty0>)\n\t22: CopyLoc[4](loc2: &mut OuterNode<Ty0>)\n\t23: ImmBorrowFieldGeneric[0](OuterNode.key: u128)\n\t24: ReadRef\n\t25: CopyLoc[4](loc2: &mut OuterNode<Ty0>)\n\t26: MutBorrowFieldGeneric[1](OuterNode.value: Ty0)\n\t27: MoveLoc[4](loc2: &mut OuterNode<Ty0>)\n\t28: ImmBorrowFieldGeneric[5](OuterNode.parent_index: u64)\n\t29: ReadRef\n\t30: MoveLoc[3](loc1: u64)\n\t31: Ret\n}\npublic traverse_mut<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128, Arg2: u64, Arg3: bool): u128 * &mut Ty0 * u64 * u64 {\nL0:\tloc4: &mut OuterNode<Ty0>\nL1:\tloc5: u64\nB0:\n\t0: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: MoveLoc[1](Arg1: u128)\n\t2: MoveLoc[2](Arg2: u64)\n\t3: MoveLoc[3](Arg3: bool)\n\t4: StLoc[7](loc3: bool)\n\t5: StLoc[6](loc2: u64)\n\t6: StLoc[5](loc1: u128)\n\t7: StLoc[4](loc0: &mut CritBitTree<Ty0>)\n\t8: MoveLoc[4](loc0: &mut CritBitTree<Ty0>)\n\t9: FreezeRef\n\t10: MoveLoc[5](loc1: u128)\n\t11: MoveLoc[6](loc2: u64)\n\t12: MoveLoc[7](loc3: bool)\n\t13: Call[25](traverse_target_child_index<Ty0>(&CritBitTree<Ty0>, u128, u64, bool): u64)\n\t14: StLoc[9](loc5: u64)\n\t15: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t16: MutBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t17: CopyLoc[9](loc5: u64)\n\t18: Call[26](outer_node_vector_index(u64): u64)\n\t19: VecMutBorrow(37)\n\t20: StLoc[8](loc4: &mut OuterNode<Ty0>)\n\t21: CopyLoc[8](loc4: &mut OuterNode<Ty0>)\n\t22: ImmBorrowFieldGeneric[0](OuterNode.key: u128)\n\t23: ReadRef\n\t24: CopyLoc[8](loc4: &mut OuterNode<Ty0>)\n\t25: MutBorrowFieldGeneric[1](OuterNode.value: Ty0)\n\t26: MoveLoc[8](loc4: &mut OuterNode<Ty0>)\n\t27: ImmBorrowFieldGeneric[5](OuterNode.parent_index: u64)\n\t28: ReadRef\n\t29: MoveLoc[9](loc5: u64)\n\t30: Ret\n}\npublic traverse_pop_mut<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128, Arg2: u64, Arg3: u64, Arg4: u64, Arg5: bool): u128 * &mut Ty0 * u64 * u64 * Ty0 {\nL0:\tloc6: u64\nL1:\tloc7: &mut OuterNode<Ty0>\nB0:\n\t0: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: ImmBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t2: CopyLoc[2](Arg2: u64)\n\t3: VecImmBorrow(38)\n\t4: ImmBorrowField[7](InnerNode.left_child_index: u64)\n\t5: ReadRef\n\t6: CopyLoc[3](Arg3: u64)\n\t7: Eq\n\t8: StLoc[10](loc4: bool)\n\t9: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t10: MoveLoc[1](Arg1: u128)\n\t11: CopyLoc[2](Arg2: u64)\n\t12: MoveLoc[5](Arg5: bool)\n\t13: StLoc[9](loc3: bool)\n\t14: StLoc[8](loc2: u64)\n\t15: StLoc[7](loc1: u128)\n\t16: StLoc[6](loc0: &mut CritBitTree<Ty0>)\n\t17: MoveLoc[6](loc0: &mut CritBitTree<Ty0>)\n\t18: FreezeRef\n\t19: MoveLoc[7](loc1: u128)\n\t20: MoveLoc[8](loc2: u64)\n\t21: MoveLoc[9](loc3: bool)\n\t22: Call[25](traverse_target_child_index<Ty0>(&CritBitTree<Ty0>, u128, u64, bool): u64)\n\t23: StLoc[12](loc6: u64)\n\t24: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t25: MoveLoc[10](loc4: bool)\n\t26: CopyLoc[2](Arg2: u64)\n\t27: Call[21](pop_update_relationships<Ty0>(&mut CritBitTree<Ty0>, bool, u64))\n\t28: CopyLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t29: MoveLoc[2](Arg2: u64)\n\t30: CopyLoc[3](Arg3: u64)\n\t31: CopyLoc[4](Arg4: u64)\n\t32: Call[22](pop_destroy_nodes<Ty0>(&mut CritBitTree<Ty0>, u64, u64, u64): Ty0)\n\t33: StLoc[11](loc5: Ty0)\n\t34: CopyLoc[12](loc6: u64)\n\t35: Call[26](outer_node_vector_index(u64): u64)\n\t36: MoveLoc[4](Arg4: u64)\n\t37: LdU64(1)\n\t38: Sub\n\t39: Eq\n\t40: BrTrue(42)\nB1:\n\t41: Branch(44)\nB2:\n\t42: MoveLoc[3](Arg3: u64)\n\t43: StLoc[12](loc6: u64)\nB3:\n\t44: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t45: MutBorrowFieldGeneric[3](CritBitTree.outer_nodes: vector<OuterNode<Ty0>>)\n\t46: CopyLoc[12](loc6: u64)\n\t47: Call[26](outer_node_vector_index(u64): u64)\n\t48: VecMutBorrow(37)\n\t49: StLoc[13](loc7: &mut OuterNode<Ty0>)\n\t50: CopyLoc[13](loc7: &mut OuterNode<Ty0>)\n\t51: ImmBorrowFieldGeneric[0](OuterNode.key: u128)\n\t52: ReadRef\n\t53: CopyLoc[13](loc7: &mut OuterNode<Ty0>)\n\t54: MutBorrowFieldGeneric[1](OuterNode.value: Ty0)\n\t55: MoveLoc[13](loc7: &mut OuterNode<Ty0>)\n\t56: ImmBorrowFieldGeneric[5](OuterNode.parent_index: u64)\n\t57: ReadRef\n\t58: MoveLoc[12](loc6: u64)\n\t59: MoveLoc[11](loc5: Ty0)\n\t60: Ret\n}\npublic traverse_predecessor_init_mut<Ty0>(Arg0: &mut CritBitTree<Ty0>): u128 * &mut Ty0 * u64 * u64 {\nB0:\n\t0: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: LdConst[10](Bool: [1])\n\t2: Call[26](traverse_init_mut<Ty0>(&mut CritBitTree<Ty0>, bool): u128 * &mut Ty0 * u64 * u64)\n\t3: Ret\n}\npublic traverse_predecessor_mut<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128, Arg2: u64): u128 * &mut Ty0 * u64 * u64 {\nB0:\n\t0: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: MoveLoc[1](Arg1: u128)\n\t2: MoveLoc[2](Arg2: u64)\n\t3: LdConst[10](Bool: [1])\n\t4: Call[27](traverse_mut<Ty0>(&mut CritBitTree<Ty0>, u128, u64, bool): u128 * &mut Ty0 * u64 * u64)\n\t5: Ret\n}\npublic traverse_predecessor_pop_mut<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128, Arg2: u64, Arg3: u64, Arg4: u64): u128 * &mut Ty0 * u64 * u64 * Ty0 {\nB0:\n\t0: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: MoveLoc[1](Arg1: u128)\n\t2: MoveLoc[2](Arg2: u64)\n\t3: MoveLoc[3](Arg3: u64)\n\t4: MoveLoc[4](Arg4: u64)\n\t5: LdConst[10](Bool: [1])\n\t6: Call[28](traverse_pop_mut<Ty0>(&mut CritBitTree<Ty0>, u128, u64, u64, u64, bool): u128 * &mut Ty0 * u64 * u64 * Ty0)\n\t7: Ret\n}\npublic traverse_successor_init_mut<Ty0>(Arg0: &mut CritBitTree<Ty0>): u128 * &mut Ty0 * u64 * u64 {\nB0:\n\t0: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: LdConst[13](Bool: [0])\n\t2: Call[26](traverse_init_mut<Ty0>(&mut CritBitTree<Ty0>, bool): u128 * &mut Ty0 * u64 * u64)\n\t3: Ret\n}\npublic traverse_successor_mut<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128, Arg2: u64): u128 * &mut Ty0 * u64 * u64 {\nB0:\n\t0: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: MoveLoc[1](Arg1: u128)\n\t2: MoveLoc[2](Arg2: u64)\n\t3: LdConst[13](Bool: [0])\n\t4: Call[27](traverse_mut<Ty0>(&mut CritBitTree<Ty0>, u128, u64, bool): u128 * &mut Ty0 * u64 * u64)\n\t5: Ret\n}\npublic traverse_successor_pop_mut<Ty0>(Arg0: &mut CritBitTree<Ty0>, Arg1: u128, Arg2: u64, Arg3: u64, Arg4: u64): u128 * &mut Ty0 * u64 * u64 * Ty0 {\nB0:\n\t0: MoveLoc[0](Arg0: &mut CritBitTree<Ty0>)\n\t1: MoveLoc[1](Arg1: u128)\n\t2: MoveLoc[2](Arg2: u64)\n\t3: MoveLoc[3](Arg3: u64)\n\t4: MoveLoc[4](Arg4: u64)\n\t5: LdConst[13](Bool: [0])\n\t6: Call[28](traverse_pop_mut<Ty0>(&mut CritBitTree<Ty0>, u128, u64, u64, u64, bool): u128 * &mut Ty0 * u64 * u64 * Ty0)\n\t7: Ret\n}\ntraverse_target_child_index<Ty0>(Arg0: &CritBitTree<Ty0>, Arg1: u128, Arg2: u64, Arg3: bool): u64 {\nB0:\n\t0: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t1: ImmBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t2: MoveLoc[2](Arg2: u64)\n\t3: VecImmBorrow(38)\n\t4: StLoc[7](loc3: &InnerNode)\nB1:\n\t5: CopyLoc[3](Arg3: bool)\n\t6: CopyLoc[1](Arg1: u128)\n\t7: CopyLoc[7](loc3: &InnerNode)\n\t8: ImmBorrowField[5](InnerNode.critical_bit: u8)\n\t9: ReadRef\n\t10: Call[19](is_set(u128, u8): bool)\n\t11: Neq\n\t12: BrTrue(14)\nB2:\n\t13: Branch(22)\nB3:\n\t14: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t15: ImmBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t16: MoveLoc[7](loc3: &InnerNode)\n\t17: ImmBorrowField[8](InnerNode.parent_index: u64)\n\t18: ReadRef\n\t19: VecImmBorrow(38)\n\t20: StLoc[7](loc3: &InnerNode)\n\t21: Branch(5)\nB4:\n\t22: CopyLoc[3](Arg3: bool)\n\t23: LdConst[10](Bool: [1])\n\t24: Eq\n\t25: BrTrue(27)\nB5:\n\t26: Branch(32)\nB6:\n\t27: MoveLoc[7](loc3: &InnerNode)\n\t28: ImmBorrowField[7](InnerNode.left_child_index: u64)\n\t29: ReadRef\n\t30: StLoc[4](loc0: u64)\n\t31: Branch(36)\nB7:\n\t32: MoveLoc[7](loc3: &InnerNode)\n\t33: ImmBorrowField[6](InnerNode.right_child_index: u64)\n\t34: ReadRef\n\t35: StLoc[4](loc0: u64)\nB8:\n\t36: MoveLoc[4](loc0: u64)\n\t37: StLoc[6](loc2: u64)\nB9:\n\t38: CopyLoc[6](loc2: u64)\n\t39: Call[18](is_outer_node(u64): bool)\n\t40: Not\n\t41: BrTrue(43)\nB10:\n\t42: Branch(66)\nB11:\n\t43: CopyLoc[3](Arg3: bool)\n\t44: LdConst[10](Bool: [1])\n\t45: Eq\n\t46: BrTrue(48)\nB12:\n\t47: Branch(56)\nB13:\n\t48: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t49: ImmBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t50: MoveLoc[6](loc2: u64)\n\t51: VecImmBorrow(38)\n\t52: ImmBorrowField[6](InnerNode.right_child_index: u64)\n\t53: ReadRef\n\t54: StLoc[5](loc1: u64)\n\t55: Branch(63)\nB14:\n\t56: CopyLoc[0](Arg0: &CritBitTree<Ty0>)\n\t57: ImmBorrowFieldGeneric[4](CritBitTree.inner_nodes: vector<InnerNode>)\n\t58: MoveLoc[6](loc2: u64)\n\t59: VecImmBorrow(38)\n\t60: ImmBorrowField[7](InnerNode.left_child_index: u64)\n\t61: ReadRef\n\t62: StLoc[5](loc1: u64)\nB15:\n\t63: MoveLoc[5](loc1: u64)\n\t64: StLoc[6](loc2: u64)\n\t65: Branch(38)\nB16:\n\t66: MoveLoc[0](Arg0: &CritBitTree<Ty0>)\n\t67: Pop\n\t68: MoveLoc[6](loc2: u64)\n\t69: Ret\n}\n}"
	fp_math := "// Move bytecode v5\nmodule 0.fp_math {\n\n\nfp32_ceil_util(Arg0: u128): u128 {\nL0:\tloc1: u64\nL1:\tloc2: u64\nB0:\n\t0: CopyLoc[0](Arg0: u128)\n\t1: LdU8(96)\n\t2: Shr\n\t3: CastU64\n\t4: StLoc[3](loc2: u64)\n\t5: LdConst[3](U64: [255, 255, 255, 255, 0, 0, 0, 0])\n\t6: MoveLoc[3](loc2: u64)\n\t7: Sub\n\t8: StLoc[2](loc1: u64)\n\t9: MoveLoc[2](loc1: u64)\n\t10: LdU64(1)\n\t11: Call[4](u32_wrapping_add(u64, u64): u64)\n\t12: StLoc[1](loc0: u64)\n\t13: MoveLoc[0](Arg0: u128)\n\t14: MoveLoc[1](loc0: u64)\n\t15: CastU128\n\t16: Add\n\t17: Ret\n}\npublic fp32_div(Arg0: u64, Arg1: u64): u64 {\nB0:\n\t0: MoveLoc[1](Arg1: u64)\n\t1: Call[5](create_from_raw_value(u64): FixedPoint32)\n\t2: StLoc[2](loc0: FixedPoint32)\n\t3: MoveLoc[0](Arg0: u64)\n\t4: MoveLoc[2](loc0: FixedPoint32)\n\t5: Call[6](divide_u64(u64, FixedPoint32): u64)\n\t6: Ret\n}\npublic fp32_mul_ceil(Arg0: u64, Arg1: u64): u64 {\nL0:\tloc2: u128\nB0:\n\t0: MoveLoc[0](Arg0: u64)\n\t1: CastU128\n\t2: MoveLoc[1](Arg1: u64)\n\t3: CastU128\n\t4: Mul\n\t5: StLoc[3](loc1: u128)\n\t6: MoveLoc[3](loc1: u128)\n\t7: Call[0](fp32_ceil_util(u128): u128)\n\t8: StLoc[4](loc2: u128)\n\t9: MoveLoc[4](loc2: u128)\n\t10: LdU8(32)\n\t11: Shr\n\t12: StLoc[2](loc0: u128)\n\t13: CopyLoc[2](loc0: u128)\n\t14: LdConst[2](U128: [255, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0])\n\t15: Le\n\t16: BrTrue(20)\nB1:\n\t17: LdConst[0](U64: [2, 0, 0, 0, 0, 0, 0, 0])\n\t18: Call[7](limit_exceeded(u64): u64)\n\t19: Abort\nB2:\n\t20: MoveLoc[2](loc0: u128)\n\t21: CastU64\n\t22: Ret\n}\npublic fp32_mul_floor(Arg0: u64, Arg1: u64): u64 {\nB0:\n\t0: MoveLoc[1](Arg1: u64)\n\t1: Call[5](create_from_raw_value(u64): FixedPoint32)\n\t2: StLoc[2](loc0: FixedPoint32)\n\t3: MoveLoc[0](Arg0: u64)\n\t4: MoveLoc[2](loc0: FixedPoint32)\n\t5: Call[8](multiply_u64(u64, FixedPoint32): u64)\n\t6: Ret\n}\nu32_wrapping_add(Arg0: u64, Arg1: u64): u64 {\nB0:\n\t0: CopyLoc[0](Arg0: u64)\n\t1: LdConst[3](U64: [255, 255, 255, 255, 0, 0, 0, 0])\n\t2: Le\n\t3: BrTrue(6)\nB1:\n\t4: LdU64(1)\n\t5: Abort\nB2:\n\t6: CopyLoc[1](Arg1: u64)\n\t7: LdConst[3](U64: [255, 255, 255, 255, 0, 0, 0, 0])\n\t8: Le\n\t9: BrTrue(12)\nB3:\n\t10: LdU64(2)\n\t11: Abort\nB4:\n\t12: MoveLoc[0](Arg0: u64)\n\t13: MoveLoc[1](Arg1: u64)\n\t14: Add\n\t15: StLoc[2](loc0: u64)\n\t16: CopyLoc[2](loc0: u64)\n\t17: LdConst[3](U64: [255, 255, 255, 255, 0, 0, 0, 0])\n\t18: Gt\n\t19: BrTrue(21)\nB5:\n\t20: Branch(25)\nB6:\n\t21: MoveLoc[2](loc0: u64)\n\t22: LdConst[3](U64: [255, 255, 255, 255, 0, 0, 0, 0])\n\t23: BitAnd\n\t24: Ret\nB7:\n\t25: MoveLoc[2](loc0: u64)\n\t26: Ret\n}\n}"
	market := "// Move bytecode v5\nmodule 0.market {\nstruct Account has key {\n\tid: UID,\n\tmarket_id: ID,\n\ttag: u64,\n\towner: address,\n\tbase_token_free: u64,\n\tbase_token_locked: u64,\n\tquote_token_free: u64,\n\tquote_token_locked: u64,\n\taccumulated_rebates: u64,\n\taccumulated_maker_quote_volume: u64,\n\taccumulated_maker_base_volume: u64,\n\taccumulated_taker_quote_volume: u64,\n\taccumulated_taker_base_volume: u64,\n\tnumber_of_orders: u64,\n\torders: vector<AccountOrder>\n}\nstruct AccountOrder has drop, store {\n\tid: u128,\n\tclient_id: u128\n}\nstruct Event has copy, drop, store {\n\ttag: u8,\n\tside: u8,\n\tquote_size: u64,\n\tbase_size: u64,\n\tmaker_order_id: u128,\n\taccount_id: ID\n}\nstruct EventQueue has key {\n\tid: UID,\n\tmarket_id: ID,\n\tseq_num: u64,\n\tevents: vector<Event>\n}\nstruct LPCoin<phantom Ty0, phantom Ty1> has drop {\n\tdummy_field: bool\n}\nstruct Market<phantom Ty0, phantom Ty1> has key {\n\tid: UID,\n\tasks: CritBitTree<Order>,\n\tbids: CritBitTree<Order>,\n\tbase: Balance<Ty0>,\n\tquote: Balance<Ty1>,\n\tadmin: address,\n\tbase_volume: u64,\n\tquote_volume: u64,\n\taccumulated_fees: u64,\n\tmin_base_order_size: u64,\n\taccumulated_royalties: u64,\n\tbase_currency_multiplier: u64,\n\tquote_currency_multiplier: u64,\n\ttick_size: u64,\n\tfee_type: u8\n}\nstruct MarketRegistry has key {\n\tid: UID,\n\tmarket_ids: vector<ID>,\n\tevent_queue_ids: vector<ID>,\n\taccount_ids: vector<ID>\n}\nstruct Order has copy, drop, store {\n\tkey: u128,\n\tbase_parcels: u64,\n\taccount_id: ID\n}\n\npublic account_data(Arg0: &Account): ID * address * u64 * u64 * u64 * u64 * u64 * u64 * u64 * u64 * u64 {\nB0:\n\t0: CopyLoc[0](Arg0: &Account)\n\t1: ImmBorrowField[0](Account.market_id: ID)\n\t2: ReadRef\n\t3: CopyLoc[0](Arg0: &Account)\n\t4: ImmBorrowField[1](Account.owner: address)\n\t5: ReadRef\n\t6: CopyLoc[0](Arg0: &Account)\n\t7: ImmBorrowField[2](Account.base_token_free: u64)\n\t8: ReadRef\n\t9: CopyLoc[0](Arg0: &Account)\n\t10: ImmBorrowField[3](Account.base_token_locked: u64)\n\t11: ReadRef\n\t12: CopyLoc[0](Arg0: &Account)\n\t13: ImmBorrowField[4](Account.quote_token_free: u64)\n\t14: ReadRef\n\t15: CopyLoc[0](Arg0: &Account)\n\t16: ImmBorrowField[5](Account.quote_token_locked: u64)\n\t17: ReadRef\n\t18: CopyLoc[0](Arg0: &Account)\n\t19: ImmBorrowField[6](Account.accumulated_rebates: u64)\n\t20: ReadRef\n\t21: CopyLoc[0](Arg0: &Account)\n\t22: ImmBorrowField[7](Account.accumulated_maker_quote_volume: u64)\n\t23: ReadRef\n\t24: CopyLoc[0](Arg0: &Account)\n\t25: ImmBorrowField[8](Account.accumulated_maker_base_volume: u64)\n\t26: ReadRef\n\t27: CopyLoc[0](Arg0: &Account)\n\t28: ImmBorrowField[9](Account.accumulated_taker_quote_volume: u64)\n\t29: ReadRef\n\t30: MoveLoc[0](Arg0: &Account)\n\t31: ImmBorrowField[10](Account.accumulated_taker_base_volume: u64)\n\t32: ReadRef\n\t33: Ret\n}\nentry public cancel_order<Ty0, Ty1>(Arg0: &mut Market<Ty0, Ty1>, Arg1: &mut Account, Arg2: u128, Arg3: u8, Arg4: bool, Arg5: &mut TxContext) {\nL0:\tloc6: u64\nL1:\tloc7: u64\nL2:\tloc8: &mut CritBitTree<Order>\nB0:\n\t0: CopyLoc[1](Arg1: &mut Account)\n\t1: ImmBorrowField[1](Account.owner: address)\n\t2: ReadRef\n\t3: MoveLoc[5](Arg5: &mut TxContext)\n\t4: FreezeRef\n\t5: Call[25](sender(&TxContext): address)\n\t6: Eq\n\t7: BrTrue(14)\nB1:\n\t8: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t9: Pop\n\t10: MoveLoc[1](Arg1: &mut Account)\n\t11: Pop\n\t12: LdConst[9](U64: [5, 0, 0, 0, 0, 0, 0, 0])\n\t13: Abort\nB2:\n\t14: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t15: CopyLoc[3](Arg3: u8)\n\t16: Call[0](get_tree_mut<Ty0, Ty1>(&mut Market<Ty0, Ty1>, u8): &mut CritBitTree<Order>)\n\t17: StLoc[14](loc8: &mut CritBitTree<Order>)\n\t18: MoveLoc[14](loc8: &mut CritBitTree<Order>)\n\t19: CopyLoc[2](Arg2: u128)\n\t20: Call[1](pop<Order>(&mut CritBitTree<Order>, u128): Order)\n\t21: StLoc[9](loc3: Order)\n\t22: ImmBorrowLoc[9](loc3: Order)\n\t23: ImmBorrowField[11](Order.account_id: ID)\n\t24: ReadRef\n\t25: CopyLoc[1](Arg1: &mut Account)\n\t26: FreezeRef\n\t27: Call[2](id<Account>(&Account): ID)\n\t28: Eq\n\t29: BrTrue(34)\nB3:\n\t30: MoveLoc[1](Arg1: &mut Account)\n\t31: Pop\n\t32: LdConst[14](U64: [9, 0, 0, 0, 0, 0, 0, 0])\n\t33: Abort\nB4:\n\t34: CopyLoc[2](Arg2: u128)\n\t35: Call[28](price(u128): u64)\n\t36: StLoc[11](loc5: u64)\n\t37: ImmBorrowLoc[9](loc3: Order)\n\t38: ImmBorrowField[12](Order.base_parcels: u64)\n\t39: ReadRef\n\t40: StLoc[12](loc6: u64)\n\t41: CopyLoc[12](loc6: u64)\n\t42: MoveLoc[11](loc5: u64)\n\t43: Call[29](fp32_mul_floor(u64, u64): u64)\n\t44: StLoc[13](loc7: u64)\n\t45: MoveLoc[3](Arg3: u8)\n\t46: LdConst[1](U8: [0])\n\t47: Eq\n\t48: BrTrue(50)\nB5:\n\t49: Branch(67)\nB6:\n\t50: CopyLoc[1](Arg1: &mut Account)\n\t51: ImmBorrowField[4](Account.quote_token_free: u64)\n\t52: ReadRef\n\t53: CopyLoc[13](loc7: u64)\n\t54: Add\n\t55: CopyLoc[1](Arg1: &mut Account)\n\t56: MutBorrowField[4](Account.quote_token_free: u64)\n\t57: WriteRef\n\t58: CopyLoc[1](Arg1: &mut Account)\n\t59: ImmBorrowField[5](Account.quote_token_locked: u64)\n\t60: ReadRef\n\t61: MoveLoc[13](loc7: u64)\n\t62: Sub\n\t63: CopyLoc[1](Arg1: &mut Account)\n\t64: MutBorrowField[5](Account.quote_token_locked: u64)\n\t65: WriteRef\n\t66: Branch(83)\nB7:\n\t67: CopyLoc[1](Arg1: &mut Account)\n\t68: ImmBorrowField[2](Account.base_token_free: u64)\n\t69: ReadRef\n\t70: CopyLoc[12](loc6: u64)\n\t71: Add\n\t72: CopyLoc[1](Arg1: &mut Account)\n\t73: MutBorrowField[2](Account.base_token_free: u64)\n\t74: WriteRef\n\t75: CopyLoc[1](Arg1: &mut Account)\n\t76: ImmBorrowField[3](Account.base_token_locked: u64)\n\t77: ReadRef\n\t78: MoveLoc[12](loc6: u64)\n\t79: Sub\n\t80: CopyLoc[1](Arg1: &mut Account)\n\t81: MutBorrowField[3](Account.base_token_locked: u64)\n\t82: WriteRef\nB8:\n\t83: CopyLoc[1](Arg1: &mut Account)\n\t84: MoveLoc[2](Arg2: u128)\n\t85: MoveLoc[4](Arg4: bool)\n\t86: StLoc[8](loc2: bool)\n\t87: StLoc[7](loc1: u128)\n\t88: StLoc[6](loc0: &mut Account)\n\t89: MoveLoc[6](loc0: &mut Account)\n\t90: FreezeRef\n\t91: MoveLoc[7](loc1: u128)\n\t92: MoveLoc[8](loc2: bool)\n\t93: Call[9](get_account_order_index(&Account, u128, bool): u64)\n\t94: StLoc[10](loc4: u64)\n\t95: MoveLoc[1](Arg1: &mut Account)\n\t96: MutBorrowField[13](Account.orders: vector<AccountOrder>)\n\t97: MoveLoc[10](loc4: u64)\n\t98: Call[3](remove<AccountOrder>(&mut vector<AccountOrder>, u64): AccountOrder)\n\t99: Pop\n\t100: Ret\n}\nentry public consume_event<Ty0, Ty1>(Arg0: &mut Market<Ty0, Ty1>, Arg1: &mut EventQueue, Arg2: &mut Account) {\nL0:\tloc3: u128\nL1:\tloc4: bool\nL2:\tloc5: u8\nL3:\tloc6: ID\nL4:\tloc7: u64\nL5:\tloc8: &Event\nL6:\tloc9: &mut vector<Event>\nL7:\tloc10: u128\nL8:\tloc11: u64\nL9:\tloc12: u64\nL10:\tloc13: u64\nL11:\tloc14: u64\nL12:\tloc15: u64\nL13:\tloc16: u64\nL14:\tloc17: u8\nL15:\tloc18: u64\nL16:\tloc19: u64\nB0:\n\t0: MoveLoc[1](Arg1: &mut EventQueue)\n\t1: MutBorrowField[14](EventQueue.events: vector<Event>)\n\t2: StLoc[12](loc9: &mut vector<Event>)\n\t3: CopyLoc[12](loc9: &mut vector<Event>)\n\t4: StLoc[3](loc0: &mut vector<Event>)\n\t5: MoveLoc[3](loc0: &mut vector<Event>)\n\t6: FreezeRef\n\t7: LdU64(0)\n\t8: VecImmBorrow(23)\n\t9: StLoc[11](loc8: &Event)\n\t10: CopyLoc[11](loc8: &Event)\n\t11: ReadRef\n\t12: Unpack[2](Event)\n\t13: StLoc[9](loc6: ID)\n\t14: StLoc[13](loc10: u128)\n\t15: StLoc[10](loc7: u64)\n\t16: StLoc[18](loc15: u64)\n\t17: StLoc[20](loc17: u8)\n\t18: Pop\n\t19: MoveLoc[9](loc6: ID)\n\t20: CopyLoc[2](Arg2: &mut Account)\n\t21: FreezeRef\n\t22: Call[2](id<Account>(&Account): ID)\n\t23: Eq\n\t24: BrTrue(35)\nB1:\n\t25: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t26: Pop\n\t27: MoveLoc[12](loc9: &mut vector<Event>)\n\t28: Pop\n\t29: MoveLoc[11](loc8: &Event)\n\t30: Pop\n\t31: MoveLoc[2](Arg2: &mut Account)\n\t32: Pop\n\t33: LdConst[10](U64: [10, 0, 0, 0, 0, 0, 0, 0])\n\t34: Abort\nB2:\n\t35: MoveLoc[18](loc15: u64)\n\t36: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t37: ImmBorrowFieldGeneric[0](Market.quote_currency_multiplier: u64)\n\t38: ReadRef\n\t39: Mul\n\t40: StLoc[18](loc15: u64)\n\t41: MoveLoc[10](loc7: u64)\n\t42: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t43: ImmBorrowFieldGeneric[1](Market.base_currency_multiplier: u64)\n\t44: ReadRef\n\t45: Mul\n\t46: StLoc[10](loc7: u64)\n\t47: CopyLoc[11](loc8: &Event)\n\t48: ImmBorrowField[17](Event.tag: u8)\n\t49: ReadRef\n\t50: LdConst[1](U8: [0])\n\t51: Eq\n\t52: BrTrue(54)\nB3:\n\t53: Branch(177)\nB4:\n\t54: MoveLoc[11](loc8: &Event)\n\t55: Pop\n\t56: LdU64(0)\n\t57: StLoc[14](loc11: u64)\n\t58: LdU64(0)\n\t59: CopyLoc[14](loc11: u64)\n\t60: Sub\n\t61: LdU64(0)\n\t62: Sub\n\t63: StLoc[22](loc19: u64)\n\t64: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t65: ImmBorrowFieldGeneric[2](Market.accumulated_fees: u64)\n\t66: ReadRef\n\t67: MoveLoc[22](loc19: u64)\n\t68: Add\n\t69: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t70: MutBorrowFieldGeneric[2](Market.accumulated_fees: u64)\n\t71: WriteRef\n\t72: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t73: ImmBorrowFieldGeneric[3](Market.accumulated_royalties: u64)\n\t74: ReadRef\n\t75: LdU64(0)\n\t76: Add\n\t77: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t78: MutBorrowFieldGeneric[3](Market.accumulated_royalties: u64)\n\t79: WriteRef\n\t80: MoveLoc[20](loc17: u8)\n\t81: LdConst[1](U8: [0])\n\t82: Eq\n\t83: BrTrue(85)\nB5:\n\t84: Branch(112)\nB6:\n\t85: CopyLoc[2](Arg2: &mut Account)\n\t86: ImmBorrowField[4](Account.quote_token_free: u64)\n\t87: ReadRef\n\t88: CopyLoc[18](loc15: u64)\n\t89: Add\n\t90: CopyLoc[14](loc11: u64)\n\t91: Add\n\t92: CopyLoc[2](Arg2: &mut Account)\n\t93: MutBorrowField[4](Account.quote_token_free: u64)\n\t94: WriteRef\n\t95: CopyLoc[2](Arg2: &mut Account)\n\t96: ImmBorrowField[3](Account.base_token_locked: u64)\n\t97: ReadRef\n\t98: CopyLoc[10](loc7: u64)\n\t99: Sub\n\t100: CopyLoc[2](Arg2: &mut Account)\n\t101: MutBorrowField[3](Account.base_token_locked: u64)\n\t102: WriteRef\n\t103: CopyLoc[2](Arg2: &mut Account)\n\t104: ImmBorrowField[6](Account.accumulated_rebates: u64)\n\t105: ReadRef\n\t106: MoveLoc[14](loc11: u64)\n\t107: Add\n\t108: CopyLoc[2](Arg2: &mut Account)\n\t109: MutBorrowField[6](Account.accumulated_rebates: u64)\n\t110: WriteRef\n\t111: Branch(144)\nB7:\n\t112: CopyLoc[2](Arg2: &mut Account)\n\t113: ImmBorrowField[2](Account.base_token_free: u64)\n\t114: ReadRef\n\t115: CopyLoc[10](loc7: u64)\n\t116: Add\n\t117: CopyLoc[2](Arg2: &mut Account)\n\t118: MutBorrowField[2](Account.base_token_free: u64)\n\t119: WriteRef\n\t120: CopyLoc[2](Arg2: &mut Account)\n\t121: ImmBorrowField[5](Account.quote_token_locked: u64)\n\t122: ReadRef\n\t123: CopyLoc[18](loc15: u64)\n\t124: Sub\n\t125: CopyLoc[2](Arg2: &mut Account)\n\t126: MutBorrowField[5](Account.quote_token_locked: u64)\n\t127: WriteRef\n\t128: CopyLoc[2](Arg2: &mut Account)\n\t129: ImmBorrowField[4](Account.quote_token_free: u64)\n\t130: ReadRef\n\t131: CopyLoc[14](loc11: u64)\n\t132: Add\n\t133: CopyLoc[2](Arg2: &mut Account)\n\t134: MutBorrowField[4](Account.quote_token_free: u64)\n\t135: WriteRef\n\t136: CopyLoc[2](Arg2: &mut Account)\n\t137: ImmBorrowField[6](Account.accumulated_rebates: u64)\n\t138: ReadRef\n\t139: MoveLoc[14](loc11: u64)\n\t140: Add\n\t141: CopyLoc[2](Arg2: &mut Account)\n\t142: MutBorrowField[6](Account.accumulated_rebates: u64)\n\t143: WriteRef\nB8:\n\t144: CopyLoc[2](Arg2: &mut Account)\n\t145: ImmBorrowField[7](Account.accumulated_maker_quote_volume: u64)\n\t146: ReadRef\n\t147: CopyLoc[18](loc15: u64)\n\t148: Add\n\t149: CopyLoc[2](Arg2: &mut Account)\n\t150: MutBorrowField[7](Account.accumulated_maker_quote_volume: u64)\n\t151: WriteRef\n\t152: CopyLoc[2](Arg2: &mut Account)\n\t153: ImmBorrowField[8](Account.accumulated_maker_base_volume: u64)\n\t154: ReadRef\n\t155: CopyLoc[10](loc7: u64)\n\t156: Add\n\t157: MoveLoc[2](Arg2: &mut Account)\n\t158: MutBorrowField[8](Account.accumulated_maker_base_volume: u64)\n\t159: WriteRef\n\t160: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t161: ImmBorrowFieldGeneric[4](Market.quote_volume: u64)\n\t162: ReadRef\n\t163: MoveLoc[18](loc15: u64)\n\t164: Add\n\t165: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t166: MutBorrowFieldGeneric[4](Market.quote_volume: u64)\n\t167: WriteRef\n\t168: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t169: ImmBorrowFieldGeneric[5](Market.base_volume: u64)\n\t170: ReadRef\n\t171: MoveLoc[10](loc7: u64)\n\t172: Add\n\t173: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t174: MutBorrowFieldGeneric[5](Market.base_volume: u64)\n\t175: WriteRef\n\t176: Branch(254)\nB9:\n\t177: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t178: Pop\n\t179: MoveLoc[11](loc8: &Event)\n\t180: ImmBorrowField[17](Event.tag: u8)\n\t181: ReadRef\n\t182: LdConst[0](U8: [1])\n\t183: Eq\n\t184: BrTrue(186)\nB10:\n\t185: Branch(252)\nB11:\n\t186: CopyLoc[10](loc7: u64)\n\t187: LdU64(0)\n\t188: Neq\n\t189: BrTrue(191)\nB12:\n\t190: Branch(236)\nB13:\n\t191: MoveLoc[20](loc17: u8)\n\t192: LdConst[1](U8: [0])\n\t193: Eq\n\t194: BrTrue(196)\nB14:\n\t195: Branch(213)\nB15:\n\t196: CopyLoc[2](Arg2: &mut Account)\n\t197: ImmBorrowField[2](Account.base_token_free: u64)\n\t198: ReadRef\n\t199: CopyLoc[10](loc7: u64)\n\t200: Add\n\t201: CopyLoc[2](Arg2: &mut Account)\n\t202: MutBorrowField[2](Account.base_token_free: u64)\n\t203: WriteRef\n\t204: CopyLoc[2](Arg2: &mut Account)\n\t205: ImmBorrowField[3](Account.base_token_locked: u64)\n\t206: ReadRef\n\t207: MoveLoc[10](loc7: u64)\n\t208: Sub\n\t209: CopyLoc[2](Arg2: &mut Account)\n\t210: MutBorrowField[3](Account.base_token_locked: u64)\n\t211: WriteRef\n\t212: Branch(236)\nB16:\n\t213: CopyLoc[13](loc10: u128)\n\t214: Call[28](price(u128): u64)\n\t215: StLoc[16](loc13: u64)\n\t216: MoveLoc[10](loc7: u64)\n\t217: MoveLoc[16](loc13: u64)\n\t218: Call[29](fp32_mul_floor(u64, u64): u64)\n\t219: StLoc[17](loc14: u64)\n\t220: CopyLoc[2](Arg2: &mut Account)\n\t221: ImmBorrowField[4](Account.quote_token_free: u64)\n\t222: ReadRef\n\t223: CopyLoc[17](loc14: u64)\n\t224: Add\n\t225: CopyLoc[2](Arg2: &mut Account)\n\t226: MutBorrowField[4](Account.quote_token_free: u64)\n\t227: WriteRef\n\t228: CopyLoc[2](Arg2: &mut Account)\n\t229: ImmBorrowField[5](Account.quote_token_locked: u64)\n\t230: ReadRef\n\t231: MoveLoc[17](loc14: u64)\n\t232: Sub\n\t233: CopyLoc[2](Arg2: &mut Account)\n\t234: MutBorrowField[5](Account.quote_token_locked: u64)\n\t235: WriteRef\nB17:\n\t236: CopyLoc[2](Arg2: &mut Account)\n\t237: MoveLoc[13](loc10: u128)\n\t238: StLoc[6](loc3: u128)\n\t239: StLoc[5](loc2: &mut Account)\n\t240: MoveLoc[5](loc2: &mut Account)\n\t241: FreezeRef\n\t242: MoveLoc[6](loc3: u128)\n\t243: LdFalse\n\t244: Call[9](get_account_order_index(&Account, u128, bool): u64)\n\t245: StLoc[15](loc12: u64)\n\t246: MoveLoc[2](Arg2: &mut Account)\n\t247: MutBorrowField[13](Account.orders: vector<AccountOrder>)\n\t248: MoveLoc[15](loc12: u64)\n\t249: Call[3](remove<AccountOrder>(&mut vector<AccountOrder>, u64): AccountOrder)\n\t250: Pop\n\t251: Branch(254)\nB18:\n\t252: MoveLoc[2](Arg2: &mut Account)\n\t253: Pop\nB19:\n\t254: MoveLoc[12](loc9: &mut vector<Event>)\n\t255: LdU64(0)\n\t256: Call[4](remove<Event>(&mut vector<Event>, u64): Event)\n\t257: Pop\n\t258: Ret\n}\npublic create_account<Ty0, Ty1>(Arg0: &mut MarketRegistry, Arg1: &Market<Ty0, Ty1>, Arg2: &mut TxContext): ID {\nB0:\n\t0: CopyLoc[2](Arg2: &mut TxContext)\n\t1: Call[31](new(&mut TxContext): UID)\n\t2: StLoc[5](loc2: UID)\n\t3: MoveLoc[5](loc2: UID)\n\t4: MoveLoc[1](Arg1: &Market<Ty0, Ty1>)\n\t5: Call[5](id<Market<Ty0, Ty1>>(&Market<Ty0, Ty1>): ID)\n\t6: LdU64(0)\n\t7: MoveLoc[2](Arg2: &mut TxContext)\n\t8: FreezeRef\n\t9: Call[25](sender(&TxContext): address)\n\t10: LdU64(0)\n\t11: LdU64(0)\n\t12: LdU64(0)\n\t13: LdU64(0)\n\t14: LdU64(0)\n\t15: LdU64(0)\n\t16: LdU64(0)\n\t17: LdU64(0)\n\t18: LdU64(0)\n\t19: LdU64(0)\n\t20: VecPack(42, 0)\n\t21: Pack[0](Account)\n\t22: StLoc[3](loc0: Account)\n\t23: ImmBorrowLoc[3](loc0: Account)\n\t24: Call[2](id<Account>(&Account): ID)\n\t25: StLoc[4](loc1: ID)\n\t26: MoveLoc[3](loc0: Account)\n\t27: Call[6](share_object<Account>(Account))\n\t28: MoveLoc[0](Arg0: &mut MarketRegistry)\n\t29: MutBorrowField[22](MarketRegistry.account_ids: vector<ID>)\n\t30: CopyLoc[4](loc1: ID)\n\t31: VecPushBack(6)\n\t32: MoveLoc[4](loc1: ID)\n\t33: Ret\n}\nentry public create_account_<Ty0, Ty1>(Arg0: &mut MarketRegistry, Arg1: &Market<Ty0, Ty1>, Arg2: &mut TxContext) {\nB0:\n\t0: MoveLoc[0](Arg0: &mut MarketRegistry)\n\t1: MoveLoc[1](Arg1: &Market<Ty0, Ty1>)\n\t2: MoveLoc[2](Arg2: &mut TxContext)\n\t3: Call[7](create_account<Ty0, Ty1>(&mut MarketRegistry, &Market<Ty0, Ty1>, &mut TxContext): ID)\n\t4: Pop\n\t5: Ret\n}\nentry public create_market<Ty0, Ty1>(Arg0: &mut MarketRegistry, Arg1: Coin<Ty0>, Arg2: Coin<Ty1>, Arg3: u64, Arg4: u64, Arg5: u64, Arg6: u64, Arg7: &mut TxContext) {\nL0:\tloc8: CritBitTree<Order>\nL1:\tloc9: Balance<Ty0>\nL2:\tloc10: Balance<Ty1>\nL3:\tloc11: address\nL4:\tloc12: u64\nL5:\tloc13: u64\nL6:\tloc14: u64\nL7:\tloc15: Balance<Ty0>\nL8:\tloc16: EventQueue\nL9:\tloc17: ID\nL10:\tloc18: Market<Ty0, Ty1>\nL11:\tloc19: ID\nL12:\tloc20: UID\nL13:\tloc21: Balance<Ty1>\nB0:\n\t0: MoveLoc[1](Arg1: Coin<Ty0>)\n\t1: Call[8](into_balance<Ty0>(Coin<Ty0>): Balance<Ty0>)\n\t2: StLoc[23](loc15: Balance<Ty0>)\n\t3: MoveLoc[2](Arg2: Coin<Ty1>)\n\t4: Call[9](into_balance<Ty1>(Coin<Ty1>): Balance<Ty1>)\n\t5: StLoc[29](loc21: Balance<Ty1>)\n\t6: ImmBorrowLoc[23](loc15: Balance<Ty0>)\n\t7: Call[10](value<Ty0>(&Balance<Ty0>): u64)\n\t8: LdU64(0)\n\t9: Eq\n\t10: BrTrue(17)\nB1:\n\t11: MoveLoc[0](Arg0: &mut MarketRegistry)\n\t12: Pop\n\t13: MoveLoc[7](Arg7: &mut TxContext)\n\t14: Pop\n\t15: LdConst[4](U64: [3, 0, 0, 0, 0, 0, 0, 0])\n\t16: Abort\nB2:\n\t17: ImmBorrowLoc[29](loc21: Balance<Ty1>)\n\t18: Call[11](value<Ty1>(&Balance<Ty1>): u64)\n\t19: LdU64(0)\n\t20: Eq\n\t21: BrTrue(28)\nB3:\n\t22: MoveLoc[0](Arg0: &mut MarketRegistry)\n\t23: Pop\n\t24: MoveLoc[7](Arg7: &mut TxContext)\n\t25: Pop\n\t26: LdConst[7](U64: [4, 0, 0, 0, 0, 0, 0, 0])\n\t27: Abort\nB4:\n\t28: CopyLoc[7](Arg7: &mut TxContext)\n\t29: Call[31](new(&mut TxContext): UID)\n\t30: StLoc[28](loc20: UID)\n\t31: MoveLoc[28](loc20: UID)\n\t32: StLoc[8](loc0: UID)\n\t33: Call[12](empty<Order>(): CritBitTree<Order>)\n\t34: StLoc[15](loc7: CritBitTree<Order>)\n\t35: Call[12](empty<Order>(): CritBitTree<Order>)\n\t36: StLoc[16](loc8: CritBitTree<Order>)\n\t37: MoveLoc[23](loc15: Balance<Ty0>)\n\t38: StLoc[17](loc9: Balance<Ty0>)\n\t39: MoveLoc[29](loc21: Balance<Ty1>)\n\t40: StLoc[18](loc10: Balance<Ty1>)\n\t41: CopyLoc[7](Arg7: &mut TxContext)\n\t42: FreezeRef\n\t43: Call[25](sender(&TxContext): address)\n\t44: StLoc[19](loc11: address)\n\t45: MoveLoc[3](Arg3: u64)\n\t46: StLoc[10](loc2: u64)\n\t47: MoveLoc[4](Arg4: u64)\n\t48: StLoc[11](loc3: u64)\n\t49: MoveLoc[5](Arg5: u64)\n\t50: StLoc[12](loc4: u64)\n\t51: MoveLoc[6](Arg6: u64)\n\t52: StLoc[13](loc5: u64)\n\t53: MoveLoc[8](loc0: UID)\n\t54: MoveLoc[15](loc7: CritBitTree<Order>)\n\t55: MoveLoc[16](loc8: CritBitTree<Order>)\n\t56: MoveLoc[17](loc9: Balance<Ty0>)\n\t57: MoveLoc[18](loc10: Balance<Ty1>)\n\t58: MoveLoc[19](loc11: address)\n\t59: LdU64(0)\n\t60: LdU64(0)\n\t61: LdU64(0)\n\t62: MoveLoc[10](loc2: u64)\n\t63: LdU64(0)\n\t64: MoveLoc[12](loc4: u64)\n\t65: MoveLoc[13](loc5: u64)\n\t66: MoveLoc[11](loc3: u64)\n\t67: LdU8(0)\n\t68: PackGeneric[0](Market<Ty0, Ty1>)\n\t69: StLoc[26](loc18: Market<Ty0, Ty1>)\n\t70: ImmBorrowLoc[26](loc18: Market<Ty0, Ty1>)\n\t71: Call[5](id<Market<Ty0, Ty1>>(&Market<Ty0, Ty1>): ID)\n\t72: StLoc[27](loc19: ID)\n\t73: MoveLoc[7](Arg7: &mut TxContext)\n\t74: Call[31](new(&mut TxContext): UID)\n\t75: CopyLoc[27](loc19: ID)\n\t76: LdU64(0)\n\t77: VecPack(23, 0)\n\t78: Pack[3](EventQueue)\n\t79: StLoc[24](loc16: EventQueue)\n\t80: ImmBorrowLoc[24](loc16: EventQueue)\n\t81: Call[13](id<EventQueue>(&EventQueue): ID)\n\t82: StLoc[25](loc17: ID)\n\t83: MoveLoc[26](loc18: Market<Ty0, Ty1>)\n\t84: Call[14](share_object<Market<Ty0, Ty1>>(Market<Ty0, Ty1>))\n\t85: MoveLoc[24](loc16: EventQueue)\n\t86: Call[15](share_object<EventQueue>(EventQueue))\n\t87: CopyLoc[0](Arg0: &mut MarketRegistry)\n\t88: MutBorrowField[23](MarketRegistry.market_ids: vector<ID>)\n\t89: MoveLoc[27](loc19: ID)\n\t90: VecPushBack(6)\n\t91: MoveLoc[0](Arg0: &mut MarketRegistry)\n\t92: MutBorrowField[24](MarketRegistry.event_queue_ids: vector<ID>)\n\t93: MoveLoc[25](loc17: ID)\n\t94: VecPushBack(6)\n\t95: Ret\n}\ncreate_market_registry(Arg0: &mut TxContext): MarketRegistry {\nB0:\n\t0: MoveLoc[0](Arg0: &mut TxContext)\n\t1: Call[31](new(&mut TxContext): UID)\n\t2: VecPack(6, 0)\n\t3: VecPack(6, 0)\n\t4: VecPack(6, 0)\n\t5: Pack[6](MarketRegistry)\n\t6: Ret\n}\npublic gen_order_id(Arg0: &mut EventQueue, Arg1: u64, Arg2: u8): u128 {\nL0:\tloc3: u128\nB0:\n\t0: MoveLoc[0](Arg0: &mut EventQueue)\n\t1: Call[8](gen_seq_num(&mut EventQueue): u64)\n\t2: StLoc[5](loc2: u64)\n\t3: MoveLoc[1](Arg1: u64)\n\t4: CastU128\n\t5: LdU8(64)\n\t6: Shl\n\t7: StLoc[6](loc3: u128)\n\t8: MoveLoc[2](Arg2: u8)\n\t9: LdConst[1](U8: [0])\n\t10: Eq\n\t11: BrTrue(13)\nB1:\n\t12: Branch(18)\nB2:\n\t13: LdConst[25](U64: [255, 255, 255, 255, 255, 255, 255, 255])\n\t14: MoveLoc[5](loc2: u64)\n\t15: Sub\n\t16: StLoc[3](loc0: u64)\n\t17: Branch(20)\nB3:\n\t18: MoveLoc[5](loc2: u64)\n\t19: StLoc[3](loc0: u64)\nB4:\n\t20: MoveLoc[3](loc0: u64)\n\t21: StLoc[4](loc1: u64)\n\t22: MoveLoc[6](loc3: u128)\n\t23: MoveLoc[4](loc1: u64)\n\t24: CastU128\n\t25: BitOr\n\t26: Ret\n}\ngen_seq_num(Arg0: &mut EventQueue): u64 {\nB0:\n\t0: CopyLoc[0](Arg0: &mut EventQueue)\n\t1: ImmBorrowField[25](EventQueue.seq_num: u64)\n\t2: ReadRef\n\t3: StLoc[1](loc0: u64)\n\t4: CopyLoc[0](Arg0: &mut EventQueue)\n\t5: ImmBorrowField[25](EventQueue.seq_num: u64)\n\t6: ReadRef\n\t7: LdU64(1)\n\t8: Add\n\t9: MoveLoc[0](Arg0: &mut EventQueue)\n\t10: MutBorrowField[25](EventQueue.seq_num: u64)\n\t11: WriteRef\n\t12: MoveLoc[1](loc0: u64)\n\t13: Ret\n}\npublic get_account_order_index(Arg0: &Account, Arg1: u128, Arg2: bool): u64 {\nL0:\tloc3: u64\nB0:\n\t0: LdU64(0)\n\t1: StLoc[6](loc3: u64)\nB1:\n\t2: CopyLoc[0](Arg0: &Account)\n\t3: ImmBorrowField[13](Account.orders: vector<AccountOrder>)\n\t4: CopyLoc[6](loc3: u64)\n\t5: VecImmBorrow(42)\n\t6: StLoc[5](loc2: &AccountOrder)\n\t7: CopyLoc[2](Arg2: bool)\n\t8: Not\n\t9: BrTrue(11)\nB2:\n\t10: Branch(18)\nB3:\n\t11: CopyLoc[5](loc2: &AccountOrder)\n\t12: ImmBorrowField[26](AccountOrder.id: u128)\n\t13: ReadRef\n\t14: CopyLoc[1](Arg1: u128)\n\t15: Eq\n\t16: StLoc[3](loc0: bool)\n\t17: Branch(20)\nB4:\n\t18: LdFalse\n\t19: StLoc[3](loc0: bool)\nB5:\n\t20: MoveLoc[3](loc0: bool)\n\t21: BrTrue(23)\nB6:\n\t22: Branch(28)\nB7:\n\t23: MoveLoc[5](loc2: &AccountOrder)\n\t24: Pop\n\t25: MoveLoc[0](Arg0: &Account)\n\t26: Pop\n\t27: Branch(53)\nB8:\n\t28: CopyLoc[2](Arg2: bool)\n\t29: BrTrue(31)\nB9:\n\t30: Branch(38)\nB10:\n\t31: MoveLoc[5](loc2: &AccountOrder)\n\t32: ImmBorrowField[27](AccountOrder.client_id: u128)\n\t33: ReadRef\n\t34: CopyLoc[1](Arg1: u128)\n\t35: Eq\n\t36: StLoc[4](loc1: bool)\n\t37: Branch(42)\nB11:\n\t38: MoveLoc[5](loc2: &AccountOrder)\n\t39: Pop\n\t40: LdFalse\n\t41: StLoc[4](loc1: bool)\nB12:\n\t42: MoveLoc[4](loc1: bool)\n\t43: BrTrue(45)\nB13:\n\t44: Branch(48)\nB14:\n\t45: MoveLoc[0](Arg0: &Account)\n\t46: Pop\n\t47: Branch(53)\nB15:\n\t48: MoveLoc[6](loc3: u64)\n\t49: LdU64(1)\n\t50: Add\n\t51: StLoc[6](loc3: u64)\n\t52: Branch(2)\nB16:\n\t53: MoveLoc[6](loc3: u64)\n\t54: Ret\n}\npublic get_opposite_side(Arg0: u8): u8 {\nB0:\n\t0: MoveLoc[0](Arg0: u8)\n\t1: LdU8(0)\n\t2: Eq\n\t3: BrTrue(5)\nB1:\n\t4: Branch(8)\nB2:\n\t5: LdU8(1)\n\t6: StLoc[1](loc0: u8)\n\t7: Branch(10)\nB3:\n\t8: LdU8(0)\n\t9: StLoc[1](loc0: u8)\nB4:\n\t10: MoveLoc[1](loc0: u8)\n\t11: Ret\n}\npublic get_quote_from_base<Ty0, Ty1>(Arg0: &Market<Ty0, Ty1>, Arg1: u64, Arg2: u64): u64 {\nB0:\n\t0: MoveLoc[1](Arg1: u64)\n\t1: MoveLoc[2](Arg2: u64)\n\t2: Call[29](fp32_mul_floor(u64, u64): u64)\n\t3: CastU128\n\t4: CopyLoc[0](Arg0: &Market<Ty0, Ty1>)\n\t5: ImmBorrowFieldGeneric[0](Market.quote_currency_multiplier: u64)\n\t6: ReadRef\n\t7: CastU128\n\t8: Mul\n\t9: MoveLoc[0](Arg0: &Market<Ty0, Ty1>)\n\t10: ImmBorrowFieldGeneric[1](Market.base_currency_multiplier: u64)\n\t11: ReadRef\n\t12: CastU128\n\t13: Div\n\t14: CastU64\n\t15: Ret\n}\npublic get_tree<Ty0, Ty1>(Arg0: &mut Market<Ty0, Ty1>, Arg1: u8): &CritBitTree<Order> {\nB0:\n\t0: MoveLoc[1](Arg1: u8)\n\t1: LdConst[1](U8: [0])\n\t2: Eq\n\t3: BrTrue(5)\nB1:\n\t4: Branch(9)\nB2:\n\t5: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t6: ImmBorrowFieldGeneric[6](Market.bids: CritBitTree<Order>)\n\t7: StLoc[2](loc0: &CritBitTree<Order>)\n\t8: Branch(12)\nB3:\n\t9: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t10: ImmBorrowFieldGeneric[7](Market.asks: CritBitTree<Order>)\n\t11: StLoc[2](loc0: &CritBitTree<Order>)\nB4:\n\t12: MoveLoc[2](loc0: &CritBitTree<Order>)\n\t13: Ret\n}\npublic get_tree_mut<Ty0, Ty1>(Arg0: &mut Market<Ty0, Ty1>, Arg1: u8): &mut CritBitTree<Order> {\nB0:\n\t0: MoveLoc[1](Arg1: u8)\n\t1: LdConst[1](U8: [0])\n\t2: Eq\n\t3: BrTrue(5)\nB1:\n\t4: Branch(9)\nB2:\n\t5: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t6: MutBorrowFieldGeneric[6](Market.bids: CritBitTree<Order>)\n\t7: StLoc[2](loc0: &mut CritBitTree<Order>)\n\t8: Branch(12)\nB3:\n\t9: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t10: MutBorrowFieldGeneric[7](Market.asks: CritBitTree<Order>)\n\t11: StLoc[2](loc0: &mut CritBitTree<Order>)\nB4:\n\t12: MoveLoc[2](loc0: &mut CritBitTree<Order>)\n\t13: Ret\n}\ninit(Arg0: &mut TxContext) {\nB0:\n\t0: MoveLoc[0](Arg0: &mut TxContext)\n\t1: Call[6](create_market_registry(&mut TxContext): MarketRegistry)\n\t2: StLoc[1](loc0: MarketRegistry)\n\t3: MoveLoc[1](loc0: MarketRegistry)\n\t4: Call[16](share_object<MarketRegistry>(MarketRegistry))\n\t5: Ret\n}\npublic market_data<Ty0, Ty1>(Arg0: &Market<Ty0, Ty1>): address * u64 * u64 * u64 * u64 * u64 * u64 * u64 * u64 {\nB0:\n\t0: CopyLoc[0](Arg0: &Market<Ty0, Ty1>)\n\t1: ImmBorrowFieldGeneric[8](Market.admin: address)\n\t2: ReadRef\n\t3: CopyLoc[0](Arg0: &Market<Ty0, Ty1>)\n\t4: ImmBorrowFieldGeneric[5](Market.base_volume: u64)\n\t5: ReadRef\n\t6: CopyLoc[0](Arg0: &Market<Ty0, Ty1>)\n\t7: ImmBorrowFieldGeneric[4](Market.quote_volume: u64)\n\t8: ReadRef\n\t9: CopyLoc[0](Arg0: &Market<Ty0, Ty1>)\n\t10: ImmBorrowFieldGeneric[2](Market.accumulated_fees: u64)\n\t11: ReadRef\n\t12: CopyLoc[0](Arg0: &Market<Ty0, Ty1>)\n\t13: ImmBorrowFieldGeneric[9](Market.min_base_order_size: u64)\n\t14: ReadRef\n\t15: CopyLoc[0](Arg0: &Market<Ty0, Ty1>)\n\t16: ImmBorrowFieldGeneric[3](Market.accumulated_royalties: u64)\n\t17: ReadRef\n\t18: CopyLoc[0](Arg0: &Market<Ty0, Ty1>)\n\t19: ImmBorrowFieldGeneric[1](Market.base_currency_multiplier: u64)\n\t20: ReadRef\n\t21: CopyLoc[0](Arg0: &Market<Ty0, Ty1>)\n\t22: ImmBorrowFieldGeneric[0](Market.quote_currency_multiplier: u64)\n\t23: ReadRef\n\t24: MoveLoc[0](Arg0: &Market<Ty0, Ty1>)\n\t25: ImmBorrowFieldGeneric[10](Market.tick_size: u64)\n\t26: ReadRef\n\t27: Ret\n}\npublic new_event(Arg0: u8, Arg1: u8, Arg2: u64, Arg3: u64, Arg4: u128, Arg5: ID): Event {\nB0:\n\t0: MoveLoc[0](Arg0: u8)\n\t1: MoveLoc[1](Arg1: u8)\n\t2: MoveLoc[2](Arg2: u64)\n\t3: MoveLoc[3](Arg3: u64)\n\t4: MoveLoc[4](Arg4: u128)\n\t5: MoveLoc[5](Arg5: ID)\n\t6: Pack[2](Event)\n\t7: Ret\n}\nentry public new_order<Ty0, Ty1>(Arg0: &mut Market<Ty0, Ty1>, Arg1: &mut EventQueue, Arg2: &mut Account, Arg3: Coin<Ty0>, Arg4: Coin<Ty1>, Arg5: u128, Arg6: u64, Arg7: u64, Arg8: u64, Arg9: u64, Arg10: bool, Arg11: bool, Arg12: u8, Arg13: u8, Arg14: u8, Arg15: &mut TxContext) {\nL0:\tloc16: u64\nL1:\tloc17: u64\nL2:\tloc18: u64\nL3:\tloc19: u64\nL4:\tloc20: u64\nL5:\tloc21: u64\nL6:\tloc22: u64\nL7:\tloc23: u64\nL8:\tloc24: u64\nL9:\tloc25: u64\nL10:\tloc26: u64\nL11:\tloc27: u64\nL12:\tloc28: u64\nB0:\n\t0: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t1: MoveLoc[1](Arg1: &mut EventQueue)\n\t2: CopyLoc[2](Arg2: &mut Account)\n\t3: MoveLoc[5](Arg5: u128)\n\t4: CopyLoc[6](Arg6: u64)\n\t5: CopyLoc[7](Arg7: u64)\n\t6: CopyLoc[8](Arg8: u64)\n\t7: MoveLoc[9](Arg9: u64)\n\t8: MoveLoc[10](Arg10: bool)\n\t9: MoveLoc[11](Arg11: bool)\n\t10: CopyLoc[12](Arg12: u8)\n\t11: MoveLoc[14](Arg14: u8)\n\t12: CopyLoc[15](Arg15: &mut TxContext)\n\t13: Call[17](process_new_order<Ty0, Ty1>(&mut Market<Ty0, Ty1>, &mut EventQueue, &mut Account, u128, u64, u64, u64, u64, bool, bool, u8, u8, &mut TxContext): Option<u128> * u64 * u64 * u64)\n\t14: StLoc[43](loc27: u64)\n\t15: StLoc[44](loc28: u64)\n\t16: StLoc[42](loc26: u64)\n\t17: StLoc[31](loc15: Option<u128>)\n\t18: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t19: CopyLoc[43](loc27: u64)\n\t20: MoveLoc[6](Arg6: u64)\n\t21: StLoc[21](loc5: u64)\n\t22: StLoc[20](loc4: u64)\n\t23: StLoc[16](loc0: &mut Market<Ty0, Ty1>)\n\t24: MoveLoc[16](loc0: &mut Market<Ty0, Ty1>)\n\t25: FreezeRef\n\t26: MoveLoc[20](loc4: u64)\n\t27: MoveLoc[21](loc5: u64)\n\t28: Call[18](get_quote_from_base<Ty0, Ty1>(&Market<Ty0, Ty1>, u64, u64): u64)\n\t29: StLoc[32](loc16: u64)\n\t30: CopyLoc[12](Arg12: u8)\n\t31: LdConst[1](U8: [0])\n\t32: Eq\n\t33: BrTrue(35)\nB1:\n\t34: Branch(82)\nB2:\n\t35: CopyLoc[44](loc28: u64)\n\t36: CopyLoc[32](loc16: u64)\n\t37: Sub\n\t38: Pop\n\t39: MoveLoc[44](loc28: u64)\n\t40: LdU64(0)\n\t41: Add\n\t42: LdU64(0)\n\t43: Add\n\t44: StLoc[44](loc28: u64)\n\t45: LdU64(0)\n\t46: Pop\n\t47: CopyLoc[44](loc28: u64)\n\t48: CopyLoc[2](Arg2: &mut Account)\n\t49: ImmBorrowField[4](Account.quote_token_free: u64)\n\t50: ReadRef\n\t51: Call[36](u64_saturating_sub(u64, u64): u64)\n\t52: StLoc[33](loc17: u64)\n\t53: CopyLoc[2](Arg2: &mut Account)\n\t54: ImmBorrowField[4](Account.quote_token_free: u64)\n\t55: ReadRef\n\t56: CopyLoc[44](loc28: u64)\n\t57: Call[36](u64_saturating_sub(u64, u64): u64)\n\t58: CopyLoc[2](Arg2: &mut Account)\n\t59: MutBorrowField[4](Account.quote_token_free: u64)\n\t60: WriteRef\n\t61: CopyLoc[2](Arg2: &mut Account)\n\t62: ImmBorrowField[5](Account.quote_token_locked: u64)\n\t63: ReadRef\n\t64: CopyLoc[32](loc16: u64)\n\t65: Add\n\t66: CopyLoc[2](Arg2: &mut Account)\n\t67: MutBorrowField[5](Account.quote_token_locked: u64)\n\t68: WriteRef\n\t69: CopyLoc[42](loc26: u64)\n\t70: CopyLoc[43](loc27: u64)\n\t71: Sub\n\t72: CopyLoc[2](Arg2: &mut Account)\n\t73: ImmBorrowField[2](Account.base_token_free: u64)\n\t74: ReadRef\n\t75: Add\n\t76: CopyLoc[2](Arg2: &mut Account)\n\t77: MutBorrowField[2](Account.base_token_free: u64)\n\t78: WriteRef\n\t79: MoveLoc[33](loc17: u64)\n\t80: StLoc[22](loc6: u64)\n\t81: Branch(122)\nB3:\n\t82: CopyLoc[42](loc26: u64)\n\t83: CopyLoc[2](Arg2: &mut Account)\n\t84: ImmBorrowField[2](Account.base_token_free: u64)\n\t85: ReadRef\n\t86: Call[36](u64_saturating_sub(u64, u64): u64)\n\t87: StLoc[34](loc18: u64)\n\t88: CopyLoc[2](Arg2: &mut Account)\n\t89: ImmBorrowField[2](Account.base_token_free: u64)\n\t90: ReadRef\n\t91: CopyLoc[42](loc26: u64)\n\t92: Call[36](u64_saturating_sub(u64, u64): u64)\n\t93: CopyLoc[2](Arg2: &mut Account)\n\t94: MutBorrowField[2](Account.base_token_free: u64)\n\t95: WriteRef\n\t96: CopyLoc[2](Arg2: &mut Account)\n\t97: ImmBorrowField[3](Account.base_token_locked: u64)\n\t98: ReadRef\n\t99: CopyLoc[43](loc27: u64)\n\t100: Add\n\t101: CopyLoc[2](Arg2: &mut Account)\n\t102: MutBorrowField[3](Account.base_token_locked: u64)\n\t103: WriteRef\n\t104: CopyLoc[44](loc28: u64)\n\t105: CopyLoc[32](loc16: u64)\n\t106: Sub\n\t107: StLoc[39](loc23: u64)\n\t108: LdU64(0)\n\t109: Pop\n\t110: MoveLoc[39](loc23: u64)\n\t111: LdU64(0)\n\t112: Sub\n\t113: CopyLoc[2](Arg2: &mut Account)\n\t114: ImmBorrowField[4](Account.quote_token_free: u64)\n\t115: ReadRef\n\t116: Add\n\t117: CopyLoc[2](Arg2: &mut Account)\n\t118: MutBorrowField[4](Account.quote_token_free: u64)\n\t119: WriteRef\n\t120: MoveLoc[34](loc18: u64)\n\t121: StLoc[22](loc6: u64)\nB4:\n\t122: MoveLoc[22](loc6: u64)\n\t123: StLoc[35](loc19: u64)\n\t124: CopyLoc[13](Arg13: u8)\n\t125: LdConst[0](U8: [1])\n\t126: Eq\n\t127: BrTrue(129)\nB5:\n\t128: Branch(134)\nB6:\n\t129: CopyLoc[44](loc28: u64)\n\t130: LdU64(0)\n\t131: Eq\n\t132: StLoc[19](loc3: bool)\n\t133: Branch(172)\nB7:\n\t134: CopyLoc[13](Arg13: u8)\n\t135: LdConst[23](U8: [2])\n\t136: Eq\n\t137: BrTrue(139)\nB8:\n\t138: Branch(156)\nB9:\n\t139: CopyLoc[12](Arg12: u8)\n\t140: LdConst[1](U8: [0])\n\t141: Eq\n\t142: BrTrue(144)\nB10:\n\t143: Branch(149)\nB11:\n\t144: CopyLoc[44](loc28: u64)\n\t145: MoveLoc[8](Arg8: u64)\n\t146: Lt\n\t147: StLoc[23](loc7: bool)\n\t148: Branch(153)\nB12:\n\t149: CopyLoc[42](loc26: u64)\n\t150: MoveLoc[7](Arg7: u64)\n\t151: Lt\n\t152: StLoc[23](loc7: bool)\nB13:\n\t153: MoveLoc[23](loc7: bool)\n\t154: StLoc[18](loc2: bool)\n\t155: Branch(170)\nB14:\n\t156: MoveLoc[13](Arg13: u8)\n\t157: LdConst[24](U8: [3])\n\t158: Eq\n\t159: BrTrue(161)\nB15:\n\t160: Branch(166)\nB16:\n\t161: ImmBorrowLoc[31](loc15: Option<u128>)\n\t162: Call[19](is_some<u128>(&Option<u128>): bool)\n\t163: Not\n\t164: StLoc[17](loc1: bool)\n\t165: Branch(168)\nB17:\n\t166: LdFalse\n\t167: StLoc[17](loc1: bool)\nB18:\n\t168: MoveLoc[17](loc1: bool)\n\t169: StLoc[18](loc2: bool)\nB19:\n\t170: MoveLoc[18](loc2: bool)\n\t171: StLoc[19](loc3: bool)\nB20:\n\t172: MoveLoc[19](loc3: bool)\n\t173: StLoc[27](loc11: bool)\n\t174: MoveLoc[27](loc11: bool)\n\t175: Not\n\t176: BrTrue(185)\nB21:\n\t177: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t178: Pop\n\t179: MoveLoc[15](Arg15: &mut TxContext)\n\t180: Pop\n\t181: MoveLoc[2](Arg2: &mut Account)\n\t182: Pop\n\t183: LdConst[3](U64: [0, 0, 0, 0, 0, 0, 0, 0])\n\t184: Abort\nB22:\n\t185: MoveLoc[12](Arg12: u8)\n\t186: LdConst[1](U8: [0])\n\t187: Eq\n\t188: BrTrue(190)\nB23:\n\t189: Branch(213)\nB24:\n\t190: ImmBorrowLoc[4](Arg4: Coin<Ty1>)\n\t191: Call[20](value<Ty1>(&Coin<Ty1>): u64)\n\t192: StLoc[36](loc20: u64)\n\t193: MoveLoc[36](loc20: u64)\n\t194: MoveLoc[35](loc19: u64)\n\t195: Sub\n\t196: StLoc[29](loc13: u64)\n\t197: MutBorrowLoc[4](Arg4: Coin<Ty1>)\n\t198: MoveLoc[29](loc13: u64)\n\t199: CopyLoc[15](Arg15: &mut TxContext)\n\t200: Call[21](split<Ty1>(&mut Coin<Ty1>, u64, &mut TxContext))\n\t201: MoveLoc[3](Arg3: Coin<Ty0>)\n\t202: MoveLoc[15](Arg15: &mut TxContext)\n\t203: FreezeRef\n\t204: Call[25](sender(&TxContext): address)\n\t205: Call[22](transfer<Coin<Ty0>>(Coin<Ty0>, address))\n\t206: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t207: MutBorrowFieldGeneric[11](Market.quote: Balance<Ty1>)\n\t208: MoveLoc[4](Arg4: Coin<Ty1>)\n\t209: Call[9](into_balance<Ty1>(Coin<Ty1>): Balance<Ty1>)\n\t210: Call[23](join<Ty1>(&mut Balance<Ty1>, Balance<Ty1>): u64)\n\t211: Pop\n\t212: Branch(235)\nB25:\n\t213: ImmBorrowLoc[3](Arg3: Coin<Ty0>)\n\t214: Call[24](value<Ty0>(&Coin<Ty0>): u64)\n\t215: StLoc[28](loc12: u64)\n\t216: MoveLoc[28](loc12: u64)\n\t217: MoveLoc[35](loc19: u64)\n\t218: Sub\n\t219: StLoc[30](loc14: u64)\n\t220: MutBorrowLoc[3](Arg3: Coin<Ty0>)\n\t221: MoveLoc[30](loc14: u64)\n\t222: CopyLoc[15](Arg15: &mut TxContext)\n\t223: Call[25](split<Ty0>(&mut Coin<Ty0>, u64, &mut TxContext))\n\t224: MoveLoc[4](Arg4: Coin<Ty1>)\n\t225: MoveLoc[15](Arg15: &mut TxContext)\n\t226: FreezeRef\n\t227: Call[25](sender(&TxContext): address)\n\t228: Call[26](transfer<Coin<Ty1>>(Coin<Ty1>, address))\n\t229: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t230: MutBorrowFieldGeneric[12](Market.base: Balance<Ty0>)\n\t231: MoveLoc[3](Arg3: Coin<Ty0>)\n\t232: Call[8](into_balance<Ty0>(Coin<Ty0>): Balance<Ty0>)\n\t233: Call[27](join<Ty0>(&mut Balance<Ty0>, Balance<Ty0>): u64)\n\t234: Pop\nB26:\n\t235: CopyLoc[2](Arg2: &mut Account)\n\t236: ImmBorrowField[10](Account.accumulated_taker_base_volume: u64)\n\t237: ReadRef\n\t238: MoveLoc[42](loc26: u64)\n\t239: MoveLoc[43](loc27: u64)\n\t240: Call[36](u64_saturating_sub(u64, u64): u64)\n\t241: Add\n\t242: CopyLoc[2](Arg2: &mut Account)\n\t243: MutBorrowField[10](Account.accumulated_taker_base_volume: u64)\n\t244: WriteRef\n\t245: CopyLoc[2](Arg2: &mut Account)\n\t246: ImmBorrowField[9](Account.accumulated_taker_quote_volume: u64)\n\t247: ReadRef\n\t248: MoveLoc[44](loc28: u64)\n\t249: MoveLoc[32](loc16: u64)\n\t250: Call[36](u64_saturating_sub(u64, u64): u64)\n\t251: Add\n\t252: MoveLoc[2](Arg2: &mut Account)\n\t253: MutBorrowField[9](Account.accumulated_taker_quote_volume: u64)\n\t254: WriteRef\n\t255: Ret\n}\nprocess_new_order<Ty0, Ty1>(Arg0: &mut Market<Ty0, Ty1>, Arg1: &mut EventQueue, Arg2: &mut Account, Arg3: u128, Arg4: u64, Arg5: u64, Arg6: u64, Arg7: u64, Arg8: bool, Arg9: bool, Arg10: u8, Arg11: u8, Arg12: &mut TxContext): Option<u128> * u64 * u64 * u64 {\nL0:\tloc13: AccountOrder\nL1:\tloc14: u64\nL2:\tloc15: u64\nL3:\tloc16: u64\nL4:\tloc17: Order\nL5:\tloc18: &mut Order\nL6:\tloc19: Order\nL7:\tloc20: u128\nL8:\tloc21: u64\nL9:\tloc22: bool\nL10:\tloc23: Event\nL11:\tloc24: u64\nL12:\tloc25: Order\nL13:\tloc26: u128\nL14:\tloc27: u64\nL15:\tloc28: u8\nL16:\tloc29: &CritBitTree<Order>\nL17:\tloc30: &mut CritBitTree<Order>\nL18:\tloc31: &mut CritBitTree<Order>\nL19:\tloc32: &mut CritBitTree<Order>\nL20:\tloc33: bool\nL21:\tloc34: Event\nL22:\tloc35: Event\nL23:\tloc36: u64\nL24:\tloc37: u64\nL25:\tloc38: u64\nL26:\tloc39: u64\nL27:\tloc40: u64\nL28:\tloc41: u64\nL29:\tloc42: &mut CritBitTree<Order>\nB0:\n\t0: CopyLoc[1](Arg1: &mut EventQueue)\n\t1: ImmBorrowField[35](EventQueue.market_id: ID)\n\t2: ReadRef\n\t3: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t4: FreezeRef\n\t5: Call[5](id<Market<Ty0, Ty1>>(&Market<Ty0, Ty1>): ID)\n\t6: Eq\n\t7: BrTrue(18)\nB1:\n\t8: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t9: Pop\n\t10: MoveLoc[1](Arg1: &mut EventQueue)\n\t11: Pop\n\t12: MoveLoc[12](Arg12: &mut TxContext)\n\t13: Pop\n\t14: MoveLoc[2](Arg2: &mut Account)\n\t15: Pop\n\t16: LdConst[13](U64: [8, 0, 0, 0, 0, 0, 0, 0])\n\t17: Abort\nB2:\n\t18: CopyLoc[2](Arg2: &mut Account)\n\t19: ImmBorrowField[1](Account.owner: address)\n\t20: ReadRef\n\t21: MoveLoc[12](Arg12: &mut TxContext)\n\t22: FreezeRef\n\t23: Call[25](sender(&TxContext): address)\n\t24: Eq\n\t25: BrTrue(34)\nB3:\n\t26: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t27: Pop\n\t28: MoveLoc[1](Arg1: &mut EventQueue)\n\t29: Pop\n\t30: MoveLoc[2](Arg2: &mut Account)\n\t31: Pop\n\t32: LdConst[9](U64: [5, 0, 0, 0, 0, 0, 0, 0])\n\t33: Abort\nB4:\n\t34: CopyLoc[2](Arg2: &mut Account)\n\t35: FreezeRef\n\t36: Call[2](id<Account>(&Account): ID)\n\t37: StLoc[25](loc12: ID)\n\t38: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t39: ImmBorrowFieldGeneric[9](Market.min_base_order_size: u64)\n\t40: ReadRef\n\t41: StLoc[37](loc24: u64)\n\t42: CopyLoc[5](Arg5: u64)\n\t43: StLoc[27](loc14: u64)\n\t44: CopyLoc[6](Arg6: u64)\n\t45: StLoc[50](loc37: u64)\n\t46: LdTrue\n\t47: StLoc[35](loc22: bool)\nB5:\n\t48: CopyLoc[7](Arg7: u64)\n\t49: LdU64(0)\n\t50: Eq\n\t51: BrTrue(53)\nB6:\n\t52: Branch(54)\nB7:\n\t53: Branch(302)\nB8:\n\t54: CopyLoc[10](Arg10: u8)\n\t55: Call[10](get_opposite_side(u8): u8)\n\t56: StLoc[41](loc28: u8)\n\t57: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t58: CopyLoc[41](loc28: u8)\n\t59: Call[28](get_tree<Ty0, Ty1>(&mut Market<Ty0, Ty1>, u8): &CritBitTree<Order>)\n\t60: StLoc[42](loc29: &CritBitTree<Order>)\n\t61: CopyLoc[42](loc29: &CritBitTree<Order>)\n\t62: Call[29](is_empty<Order>(&CritBitTree<Order>): bool)\n\t63: BrTrue(65)\nB9:\n\t64: Branch(70)\nB10:\n\t65: MoveLoc[42](loc29: &CritBitTree<Order>)\n\t66: Pop\n\t67: LdFalse\n\t68: StLoc[35](loc22: bool)\n\t69: Branch(302)\nB11:\n\t70: CopyLoc[41](loc28: u8)\n\t71: LdConst[1](U8: [0])\n\t72: Eq\n\t73: BrTrue(75)\nB12:\n\t74: Branch(79)\nB13:\n\t75: CopyLoc[42](loc29: &CritBitTree<Order>)\n\t76: Call[30](max_key<Order>(&CritBitTree<Order>): u128)\n\t77: StLoc[13](loc0: u128)\n\t78: Branch(82)\nB14:\n\t79: CopyLoc[42](loc29: &CritBitTree<Order>)\n\t80: Call[31](min_key<Order>(&CritBitTree<Order>): u128)\n\t81: StLoc[13](loc0: u128)\nB15:\n\t82: MoveLoc[13](loc0: u128)\n\t83: StLoc[33](loc20: u128)\n\t84: MoveLoc[42](loc29: &CritBitTree<Order>)\n\t85: CopyLoc[33](loc20: u128)\n\t86: Call[32](borrow<Order>(&CritBitTree<Order>, u128): &Order)\n\t87: ReadRef\n\t88: StLoc[30](loc17: Order)\n\t89: CopyLoc[33](loc20: u128)\n\t90: Call[28](price(u128): u64)\n\t91: StLoc[34](loc21: u64)\n\t92: CopyLoc[10](Arg10: u8)\n\t93: LdConst[1](U8: [0])\n\t94: Eq\n\t95: BrTrue(97)\nB16:\n\t96: Branch(102)\nB17:\n\t97: CopyLoc[4](Arg4: u64)\n\t98: CopyLoc[34](loc21: u64)\n\t99: Ge\n\t100: StLoc[21](loc8: bool)\n\t101: Branch(106)\nB18:\n\t102: CopyLoc[4](Arg4: u64)\n\t103: CopyLoc[34](loc21: u64)\n\t104: Le\n\t105: StLoc[21](loc8: bool)\nB19:\n\t106: MoveLoc[21](loc8: bool)\n\t107: StLoc[35](loc22: bool)\n\t108: CopyLoc[8](Arg8: bool)\n\t109: BrTrue(111)\nB20:\n\t110: Branch(114)\nB21:\n\t111: LdTrue\n\t112: StLoc[22](loc9: bool)\n\t113: Branch(117)\nB22:\n\t114: CopyLoc[35](loc22: bool)\n\t115: Not\n\t116: StLoc[22](loc9: bool)\nB23:\n\t117: MoveLoc[22](loc9: bool)\n\t118: BrTrue(120)\nB24:\n\t119: Branch(121)\nB25:\n\t120: Branch(302)\nB26:\n\t121: ImmBorrowLoc[30](loc17: Order)\n\t122: ImmBorrowField[12](Order.base_parcels: u64)\n\t123: ReadRef\n\t124: StLoc[40](loc27: u64)\n\t125: MoveLoc[40](loc27: u64)\n\t126: CopyLoc[27](loc14: u64)\n\t127: Call[46](min(u64, u64): u64)\n\t128: CopyLoc[50](loc37: u64)\n\t129: CopyLoc[34](loc21: u64)\n\t130: Call[47](fp32_div(u64, u64): u64)\n\t131: Call[46](min(u64, u64): u64)\n\t132: StLoc[29](loc16: u64)\n\t133: CopyLoc[29](loc16: u64)\n\t134: LdU64(0)\n\t135: Eq\n\t136: BrTrue(138)\nB27:\n\t137: Branch(139)\nB28:\n\t138: Branch(302)\nB29:\n\t139: CopyLoc[50](loc37: u64)\n\t140: StLoc[24](loc11: u64)\n\t141: CopyLoc[10](Arg10: u8)\n\t142: LdConst[1](U8: [0])\n\t143: Eq\n\t144: BrTrue(146)\nB30:\n\t145: Branch(151)\nB31:\n\t146: CopyLoc[29](loc16: u64)\n\t147: MoveLoc[34](loc21: u64)\n\t148: Call[48](fp32_mul_ceil(u64, u64): u64)\n\t149: StLoc[23](loc10: u64)\n\t150: Branch(155)\nB32:\n\t151: CopyLoc[29](loc16: u64)\n\t152: MoveLoc[34](loc21: u64)\n\t153: Call[29](fp32_mul_floor(u64, u64): u64)\n\t154: StLoc[23](loc10: u64)\nB33:\n\t155: MoveLoc[24](loc11: u64)\n\t156: MoveLoc[23](loc10: u64)\n\t157: Call[46](min(u64, u64): u64)\n\t158: StLoc[49](loc36: u64)\n\t159: CopyLoc[49](loc36: u64)\n\t160: LdU64(0)\n\t161: Eq\n\t162: BrTrue(164)\nB34:\n\t163: Branch(165)\nB35:\n\t164: Branch(302)\nB36:\n\t165: CopyLoc[11](Arg11: u8)\n\t166: LdConst[1](U8: [0])\n\t167: Neq\n\t168: BrTrue(170)\nB37:\n\t169: Branch(229)\nB38:\n\t170: ImmBorrowLoc[30](loc17: Order)\n\t171: ImmBorrowField[11](Order.account_id: ID)\n\t172: ReadRef\n\t173: CopyLoc[2](Arg2: &mut Account)\n\t174: FreezeRef\n\t175: Call[2](id<Account>(&Account): ID)\n\t176: Eq\n\t177: StLoc[46](loc33: bool)\n\t178: MoveLoc[46](loc33: bool)\n\t179: BrTrue(181)\nB39:\n\t180: Branch(229)\nB40:\n\t181: CopyLoc[11](Arg11: u8)\n\t182: LdConst[23](U8: [2])\n\t183: Neq\n\t184: BrTrue(193)\nB41:\n\t185: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t186: Pop\n\t187: MoveLoc[1](Arg1: &mut EventQueue)\n\t188: Pop\n\t189: MoveLoc[2](Arg2: &mut Account)\n\t190: Pop\n\t191: LdConst[2](U64: [1, 0, 0, 0, 0, 0, 0, 0])\n\t192: Abort\nB42:\n\t193: CopyLoc[11](Arg11: u8)\n\t194: LdConst[0](U8: [1])\n\t195: Eq\n\t196: BrTrue(205)\nB43:\n\t197: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t198: Pop\n\t199: MoveLoc[1](Arg1: &mut EventQueue)\n\t200: Pop\n\t201: MoveLoc[2](Arg2: &mut Account)\n\t202: Pop\n\t203: LdConst[15](U64: [2, 0, 0, 0, 0, 0, 0, 0])\n\t204: Abort\nB44:\n\t205: LdConst[0](U8: [1])\n\t206: CopyLoc[41](loc28: u8)\n\t207: LdU64(0)\n\t208: LdU64(0)\n\t209: CopyLoc[33](loc20: u128)\n\t210: CopyLoc[25](loc12: ID)\n\t211: Call[16](new_event(u8, u8, u64, u64, u128, ID): Event)\n\t212: StLoc[48](loc35: Event)\n\t213: CopyLoc[1](Arg1: &mut EventQueue)\n\t214: MoveLoc[48](loc35: Event)\n\t215: Call[19](push_back_event(&mut EventQueue, Event))\n\t216: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t217: MoveLoc[41](loc28: u8)\n\t218: Call[0](get_tree_mut<Ty0, Ty1>(&mut Market<Ty0, Ty1>, u8): &mut CritBitTree<Order>)\n\t219: StLoc[43](loc30: &mut CritBitTree<Order>)\n\t220: MoveLoc[43](loc30: &mut CritBitTree<Order>)\n\t221: MoveLoc[33](loc20: u128)\n\t222: Call[1](pop<Order>(&mut CritBitTree<Order>, u128): Order)\n\t223: Pop\n\t224: MoveLoc[7](Arg7: u64)\n\t225: LdU64(1)\n\t226: Sub\n\t227: StLoc[7](Arg7: u64)\n\t228: Branch(48)\nB45:\n\t229: LdConst[1](U8: [0])\n\t230: CopyLoc[10](Arg10: u8)\n\t231: CopyLoc[49](loc36: u64)\n\t232: CopyLoc[29](loc16: u64)\n\t233: CopyLoc[33](loc20: u128)\n\t234: ImmBorrowLoc[30](loc17: Order)\n\t235: ImmBorrowField[11](Order.account_id: ID)\n\t236: ReadRef\n\t237: Call[16](new_event(u8, u8, u64, u64, u128, ID): Event)\n\t238: StLoc[36](loc23: Event)\n\t239: CopyLoc[1](Arg1: &mut EventQueue)\n\t240: MoveLoc[36](loc23: Event)\n\t241: Call[19](push_back_event(&mut EventQueue, Event))\n\t242: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t243: CopyLoc[41](loc28: u8)\n\t244: Call[0](get_tree_mut<Ty0, Ty1>(&mut Market<Ty0, Ty1>, u8): &mut CritBitTree<Order>)\n\t245: StLoc[44](loc31: &mut CritBitTree<Order>)\n\t246: MoveLoc[44](loc31: &mut CritBitTree<Order>)\n\t247: CopyLoc[33](loc20: u128)\n\t248: Call[33](borrow_mut<Order>(&mut CritBitTree<Order>, u128): &mut Order)\n\t249: StLoc[31](loc18: &mut Order)\n\t250: CopyLoc[31](loc18: &mut Order)\n\t251: ImmBorrowField[12](Order.base_parcels: u64)\n\t252: ReadRef\n\t253: CopyLoc[29](loc16: u64)\n\t254: Sub\n\t255: CopyLoc[31](loc18: &mut Order)\n\t256: MutBorrowField[12](Order.base_parcels: u64)\n\t257: WriteRef\n\t258: MoveLoc[31](loc18: &mut Order)\n\t259: ReadRef\n\t260: StLoc[32](loc19: Order)\n\t261: MoveLoc[27](loc14: u64)\n\t262: MoveLoc[29](loc16: u64)\n\t263: Sub\n\t264: StLoc[27](loc14: u64)\n\t265: MoveLoc[50](loc37: u64)\n\t266: MoveLoc[49](loc36: u64)\n\t267: Sub\n\t268: StLoc[50](loc37: u64)\n\t269: ImmBorrowLoc[32](loc19: Order)\n\t270: ImmBorrowField[12](Order.base_parcels: u64)\n\t271: ReadRef\n\t272: CopyLoc[37](loc24: u64)\n\t273: Lt\n\t274: BrTrue(276)\nB46:\n\t275: Branch(297)\nB47:\n\t276: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t277: CopyLoc[41](loc28: u8)\n\t278: Call[0](get_tree_mut<Ty0, Ty1>(&mut Market<Ty0, Ty1>, u8): &mut CritBitTree<Order>)\n\t279: StLoc[45](loc32: &mut CritBitTree<Order>)\n\t280: MoveLoc[45](loc32: &mut CritBitTree<Order>)\n\t281: CopyLoc[33](loc20: u128)\n\t282: Call[1](pop<Order>(&mut CritBitTree<Order>, u128): Order)\n\t283: Pop\n\t284: LdConst[0](U8: [1])\n\t285: MoveLoc[41](loc28: u8)\n\t286: LdU64(0)\n\t287: ImmBorrowLoc[32](loc19: Order)\n\t288: ImmBorrowField[12](Order.base_parcels: u64)\n\t289: ReadRef\n\t290: MoveLoc[33](loc20: u128)\n\t291: CopyLoc[25](loc12: ID)\n\t292: Call[16](new_event(u8, u8, u64, u64, u128, ID): Event)\n\t293: StLoc[47](loc34: Event)\n\t294: CopyLoc[1](Arg1: &mut EventQueue)\n\t295: MoveLoc[47](loc34: Event)\n\t296: Call[19](push_back_event(&mut EventQueue, Event))\nB48:\n\t297: MoveLoc[7](Arg7: u64)\n\t298: LdU64(1)\n\t299: Sub\n\t300: StLoc[7](Arg7: u64)\n\t301: Branch(48)\nB49:\n\t302: CopyLoc[4](Arg4: u64)\n\t303: LdU64(0)\n\t304: Eq\n\t305: BrTrue(307)\nB50:\n\t306: Branch(310)\nB51:\n\t307: LdConst[25](U64: [255, 255, 255, 255, 255, 255, 255, 255])\n\t308: StLoc[14](loc1: u64)\n\t309: Branch(314)\nB52:\n\t310: CopyLoc[50](loc37: u64)\n\t311: CopyLoc[4](Arg4: u64)\n\t312: Call[47](fp32_div(u64, u64): u64)\n\t313: StLoc[14](loc1: u64)\nB53:\n\t314: MoveLoc[14](loc1: u64)\n\t315: CopyLoc[27](loc14: u64)\n\t316: Call[46](min(u64, u64): u64)\n\t317: StLoc[28](loc15: u64)\n\t318: MoveLoc[35](loc22: bool)\n\t319: BrTrue(321)\nB54:\n\t320: Branch(324)\nB55:\n\t321: LdTrue\n\t322: StLoc[15](loc2: bool)\n\t323: Branch(327)\nB56:\n\t324: MoveLoc[9](Arg9: bool)\n\t325: Not\n\t326: StLoc[15](loc2: bool)\nB57:\n\t327: MoveLoc[15](loc2: bool)\n\t328: BrTrue(330)\nB58:\n\t329: Branch(333)\nB59:\n\t330: LdTrue\n\t331: StLoc[16](loc3: bool)\n\t332: Branch(337)\nB60:\n\t333: CopyLoc[28](loc15: u64)\n\t334: MoveLoc[37](loc24: u64)\n\t335: Lt\n\t336: StLoc[16](loc3: bool)\nB61:\n\t337: MoveLoc[16](loc3: bool)\n\t338: BrTrue(340)\nB62:\n\t339: Branch(359)\nB63:\n\t340: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t341: Pop\n\t342: MoveLoc[1](Arg1: &mut EventQueue)\n\t343: Pop\n\t344: MoveLoc[2](Arg2: &mut Account)\n\t345: Pop\n\t346: MoveLoc[5](Arg5: u64)\n\t347: MoveLoc[27](loc14: u64)\n\t348: Sub\n\t349: StLoc[51](loc38: u64)\n\t350: MoveLoc[6](Arg6: u64)\n\t351: MoveLoc[50](loc37: u64)\n\t352: Sub\n\t353: StLoc[53](loc40: u64)\n\t354: Call[34](none<u128>(): Option<u128>)\n\t355: MoveLoc[51](loc38: u64)\n\t356: MoveLoc[53](loc40: u64)\n\t357: LdU64(0)\n\t358: Ret\nB64:\n\t359: MoveLoc[1](Arg1: &mut EventQueue)\n\t360: CopyLoc[4](Arg4: u64)\n\t361: CopyLoc[10](Arg10: u8)\n\t362: Call[7](gen_order_id(&mut EventQueue, u64, u8): u128)\n\t363: StLoc[39](loc26: u128)\n\t364: CopyLoc[28](loc15: u64)\n\t365: StLoc[17](loc4: u64)\n\t366: CopyLoc[39](loc26: u128)\n\t367: StLoc[18](loc5: u128)\n\t368: CopyLoc[2](Arg2: &mut Account)\n\t369: FreezeRef\n\t370: Call[2](id<Account>(&Account): ID)\n\t371: StLoc[19](loc6: ID)\n\t372: MoveLoc[18](loc5: u128)\n\t373: MoveLoc[17](loc4: u64)\n\t374: MoveLoc[19](loc6: ID)\n\t375: Pack[7](Order)\n\t376: StLoc[38](loc25: Order)\n\t377: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t378: CopyLoc[10](Arg10: u8)\n\t379: Call[0](get_tree_mut<Ty0, Ty1>(&mut Market<Ty0, Ty1>, u8): &mut CritBitTree<Order>)\n\t380: StLoc[55](loc42: &mut CritBitTree<Order>)\n\t381: MoveLoc[55](loc42: &mut CritBitTree<Order>)\n\t382: CopyLoc[39](loc26: u128)\n\t383: MoveLoc[38](loc25: Order)\n\t384: Call[35](insert<Order>(&mut CritBitTree<Order>, u128, Order))\n\t385: CopyLoc[39](loc26: u128)\n\t386: MoveLoc[3](Arg3: u128)\n\t387: Pack[1](AccountOrder)\n\t388: StLoc[26](loc13: AccountOrder)\n\t389: MoveLoc[2](Arg2: &mut Account)\n\t390: MutBorrowField[13](Account.orders: vector<AccountOrder>)\n\t391: MoveLoc[26](loc13: AccountOrder)\n\t392: VecPushBack(42)\n\t393: MoveLoc[27](loc14: u64)\n\t394: CopyLoc[28](loc15: u64)\n\t395: Sub\n\t396: StLoc[27](loc14: u64)\n\t397: MoveLoc[10](Arg10: u8)\n\t398: LdConst[1](U8: [0])\n\t399: Eq\n\t400: BrTrue(402)\nB65:\n\t401: Branch(407)\nB66:\n\t402: CopyLoc[28](loc15: u64)\n\t403: MoveLoc[4](Arg4: u64)\n\t404: Call[48](fp32_mul_ceil(u64, u64): u64)\n\t405: StLoc[20](loc7: u64)\n\t406: Branch(411)\nB67:\n\t407: CopyLoc[28](loc15: u64)\n\t408: MoveLoc[4](Arg4: u64)\n\t409: Call[29](fp32_mul_floor(u64, u64): u64)\n\t410: StLoc[20](loc7: u64)\nB68:\n\t411: MoveLoc[50](loc37: u64)\n\t412: MoveLoc[20](loc7: u64)\n\t413: Sub\n\t414: StLoc[50](loc37: u64)\n\t415: MoveLoc[5](Arg5: u64)\n\t416: MoveLoc[27](loc14: u64)\n\t417: Sub\n\t418: StLoc[52](loc39: u64)\n\t419: MoveLoc[6](Arg6: u64)\n\t420: MoveLoc[50](loc37: u64)\n\t421: Sub\n\t422: StLoc[54](loc41: u64)\n\t423: MoveLoc[39](loc26: u128)\n\t424: Call[36](some<u128>(u128): Option<u128>)\n\t425: MoveLoc[52](loc39: u64)\n\t426: MoveLoc[54](loc41: u64)\n\t427: MoveLoc[28](loc15: u64)\n\t428: Ret\n}\npublic push_back_event(Arg0: &mut EventQueue, Arg1: Event) {\nB0:\n\t0: MoveLoc[0](Arg0: &mut EventQueue)\n\t1: MutBorrowField[14](EventQueue.events: vector<Event>)\n\t2: MoveLoc[1](Arg1: Event)\n\t3: VecPushBack(23)\n\t4: Ret\n}\npublic scale_base_amount<Ty0, Ty1>(Arg0: &Market<Ty0, Ty1>, Arg1: u64): u64 {\nB0:\n\t0: MoveLoc[1](Arg1: u64)\n\t1: MoveLoc[0](Arg0: &Market<Ty0, Ty1>)\n\t2: ImmBorrowFieldGeneric[1](Market.base_currency_multiplier: u64)\n\t3: ReadRef\n\t4: Div\n\t5: Ret\n}\npublic scale_quote_amount<Ty0, Ty1>(Arg0: &Market<Ty0, Ty1>, Arg1: u64): u64 {\nB0:\n\t0: MoveLoc[1](Arg1: u64)\n\t1: MoveLoc[0](Arg0: &Market<Ty0, Ty1>)\n\t2: ImmBorrowFieldGeneric[0](Market.quote_currency_multiplier: u64)\n\t3: ReadRef\n\t4: Div\n\t5: Ret\n}\nentry public settle<Ty0, Ty1>(Arg0: &mut Market<Ty0, Ty1>, Arg1: &mut Account, Arg2: &mut TxContext) {\nB0:\n\t0: CopyLoc[1](Arg1: &mut Account)\n\t1: ImmBorrowField[0](Account.market_id: ID)\n\t2: ReadRef\n\t3: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t4: FreezeRef\n\t5: Call[5](id<Market<Ty0, Ty1>>(&Market<Ty0, Ty1>): ID)\n\t6: Eq\n\t7: BrTrue(16)\nB1:\n\t8: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t9: Pop\n\t10: MoveLoc[2](Arg2: &mut TxContext)\n\t11: Pop\n\t12: MoveLoc[1](Arg1: &mut Account)\n\t13: Pop\n\t14: LdConst[11](U64: [7, 0, 0, 0, 0, 0, 0, 0])\n\t15: Abort\nB2:\n\t16: CopyLoc[1](Arg1: &mut Account)\n\t17: ImmBorrowField[1](Account.owner: address)\n\t18: ReadRef\n\t19: CopyLoc[2](Arg2: &mut TxContext)\n\t20: FreezeRef\n\t21: Call[25](sender(&TxContext): address)\n\t22: Eq\n\t23: BrTrue(32)\nB3:\n\t24: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t25: Pop\n\t26: MoveLoc[2](Arg2: &mut TxContext)\n\t27: Pop\n\t28: MoveLoc[1](Arg1: &mut Account)\n\t29: Pop\n\t30: LdConst[9](U64: [5, 0, 0, 0, 0, 0, 0, 0])\n\t31: Abort\nB4:\n\t32: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t33: MutBorrowFieldGeneric[11](Market.quote: Balance<Ty1>)\n\t34: CopyLoc[1](Arg1: &mut Account)\n\t35: ImmBorrowField[4](Account.quote_token_free: u64)\n\t36: ReadRef\n\t37: CopyLoc[2](Arg2: &mut TxContext)\n\t38: Call[37](take<Ty1>(&mut Balance<Ty1>, u64, &mut TxContext): Coin<Ty1>)\n\t39: CopyLoc[2](Arg2: &mut TxContext)\n\t40: FreezeRef\n\t41: Call[25](sender(&TxContext): address)\n\t42: Call[26](transfer<Coin<Ty1>>(Coin<Ty1>, address))\n\t43: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t44: MutBorrowFieldGeneric[12](Market.base: Balance<Ty0>)\n\t45: CopyLoc[1](Arg1: &mut Account)\n\t46: ImmBorrowField[2](Account.base_token_free: u64)\n\t47: ReadRef\n\t48: CopyLoc[2](Arg2: &mut TxContext)\n\t49: Call[38](take<Ty0>(&mut Balance<Ty0>, u64, &mut TxContext): Coin<Ty0>)\n\t50: MoveLoc[2](Arg2: &mut TxContext)\n\t51: FreezeRef\n\t52: Call[25](sender(&TxContext): address)\n\t53: Call[22](transfer<Coin<Ty0>>(Coin<Ty0>, address))\n\t54: LdU64(0)\n\t55: CopyLoc[1](Arg1: &mut Account)\n\t56: MutBorrowField[4](Account.quote_token_free: u64)\n\t57: WriteRef\n\t58: LdU64(0)\n\t59: MoveLoc[1](Arg1: &mut Account)\n\t60: MutBorrowField[2](Account.base_token_free: u64)\n\t61: WriteRef\n\t62: Ret\n}\nentry public swap<Ty0, Ty1>(Arg0: &mut Market<Ty0, Ty1>, Arg1: &mut EventQueue, Arg2: &mut Account, Arg3: Coin<Ty0>, Arg4: Coin<Ty1>, Arg5: u64, Arg6: u64, Arg7: u64, Arg8: u8, Arg9: bool, Arg10: &mut TxContext) {\nL0:\tloc11: Option<u128>\nL1:\tloc12: u64\nL2:\tloc13: u64\nL3:\tloc14: u64\nL4:\tloc15: u64\nL5:\tloc16: u128\nL6:\tloc17: u64\nL7:\tloc18: u64\nL8:\tloc19: bool\nL9:\tloc20: bool\nL10:\tloc21: bool\nL11:\tloc22: u64\nL12:\tloc23: u64\nL13:\tloc24: u64\nL14:\tloc25: Coin<Ty0>\nL15:\tloc26: Coin<Ty1>\nL16:\tloc27: bool\nL17:\tloc28: bool\nL18:\tloc29: u64\nL19:\tloc30: u64\nL20:\tloc31: u64\nL21:\tloc32: u8\nL22:\tloc33: u64\nL23:\tloc34: u64\nL24:\tloc35: u64\nL25:\tloc36: u64\nL26:\tloc37: u64\nB0:\n\t0: CopyLoc[5](Arg5: u64)\n\t1: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t2: ImmBorrowFieldGeneric[9](Market.min_base_order_size: u64)\n\t3: ReadRef\n\t4: Lt\n\t5: BrTrue(16)\nB1:\n\t6: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t7: Pop\n\t8: MoveLoc[1](Arg1: &mut EventQueue)\n\t9: Pop\n\t10: MoveLoc[10](Arg10: &mut TxContext)\n\t11: Pop\n\t12: MoveLoc[2](Arg2: &mut Account)\n\t13: Pop\n\t14: LdConst[19](U64: [11, 0, 0, 0, 0, 0, 0, 0])\n\t15: Abort\nB2:\n\t16: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t17: ImmBorrowFieldGeneric[10](Market.tick_size: u64)\n\t18: ReadRef\n\t19: StLoc[46](loc35: u64)\n\t20: CopyLoc[8](Arg8: u8)\n\t21: LdConst[1](U8: [0])\n\t22: Eq\n\t23: BrTrue(25)\nB3:\n\t24: Branch(43)\nB4:\n\t25: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t26: CopyLoc[6](Arg6: u64)\n\t27: StLoc[15](loc4: u64)\n\t28: StLoc[11](loc0: &mut Market<Ty0, Ty1>)\n\t29: LdConst[25](U64: [255, 255, 255, 255, 255, 255, 255, 255])\n\t30: MoveLoc[11](loc0: &mut Market<Ty0, Ty1>)\n\t31: FreezeRef\n\t32: MoveLoc[15](loc4: u64)\n\t33: Call[39](scale_quote_amount<Ty0, Ty1>(&Market<Ty0, Ty1>, u64): u64)\n\t34: LdConst[25](U64: [255, 255, 255, 255, 255, 255, 255, 255])\n\t35: LdConst[25](U64: [255, 255, 255, 255, 255, 255, 255, 255])\n\t36: MoveLoc[46](loc35: u64)\n\t37: Mod\n\t38: Sub\n\t39: StLoc[21](loc10: u64)\n\t40: StLoc[20](loc9: u64)\n\t41: StLoc[19](loc8: u64)\n\t42: Branch(56)\nB5:\n\t43: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t44: CopyLoc[5](Arg5: u64)\n\t45: StLoc[18](loc7: u64)\n\t46: StLoc[17](loc6: &mut Market<Ty0, Ty1>)\n\t47: MoveLoc[17](loc6: &mut Market<Ty0, Ty1>)\n\t48: FreezeRef\n\t49: MoveLoc[18](loc7: u64)\n\t50: Call[40](scale_base_amount<Ty0, Ty1>(&Market<Ty0, Ty1>, u64): u64)\n\t51: LdConst[25](U64: [255, 255, 255, 255, 255, 255, 255, 255])\n\t52: LdU64(0)\n\t53: StLoc[21](loc10: u64)\n\t54: StLoc[20](loc9: u64)\n\t55: StLoc[19](loc8: u64)\nB6:\n\t56: MoveLoc[19](loc8: u64)\n\t57: MoveLoc[20](loc9: u64)\n\t58: MoveLoc[21](loc10: u64)\n\t59: StLoc[33](loc22: u64)\n\t60: StLoc[35](loc24: u64)\n\t61: StLoc[34](loc23: u64)\n\t62: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t63: MoveLoc[1](Arg1: &mut EventQueue)\n\t64: MoveLoc[2](Arg2: &mut Account)\n\t65: LdU128(0)\n\t66: MoveLoc[33](loc22: u64)\n\t67: MoveLoc[34](loc23: u64)\n\t68: MoveLoc[35](loc24: u64)\n\t69: MoveLoc[7](Arg7: u64)\n\t70: LdFalse\n\t71: LdFalse\n\t72: CopyLoc[8](Arg8: u8)\n\t73: LdConst[1](U8: [0])\n\t74: CopyLoc[10](Arg10: &mut TxContext)\n\t75: Call[17](process_new_order<Ty0, Ty1>(&mut Market<Ty0, Ty1>, &mut EventQueue, &mut Account, u128, u64, u64, u64, u64, bool, bool, u8, u8, &mut TxContext): Option<u128> * u64 * u64 * u64)\n\t76: Pop\n\t77: StLoc[48](loc37: u64)\n\t78: StLoc[47](loc36: u64)\n\t79: Pop\n\t80: LdU64(0)\n\t81: Pop\n\t82: LdU64(0)\n\t83: StLoc[42](loc31: u64)\n\t84: CopyLoc[8](Arg8: u8)\n\t85: LdConst[1](U8: [0])\n\t86: Eq\n\t87: BrTrue(89)\nB7:\n\t88: Branch(106)\nB8:\n\t89: MoveLoc[48](loc37: u64)\n\t90: LdU64(0)\n\t91: Add\n\t92: MoveLoc[42](loc31: u64)\n\t93: Add\n\t94: StLoc[48](loc37: u64)\n\t95: CopyLoc[47](loc36: u64)\n\t96: MoveLoc[5](Arg5: u64)\n\t97: Ge\n\t98: StLoc[30](loc19: bool)\n\t99: MoveLoc[30](loc19: bool)\n\t100: MoveLoc[47](loc36: u64)\n\t101: MoveLoc[48](loc37: u64)\n\t102: StLoc[14](loc3: u64)\n\t103: StLoc[13](loc2: u64)\n\t104: StLoc[12](loc1: bool)\n\t105: Branch(120)\nB9:\n\t106: CopyLoc[48](loc37: u64)\n\t107: MoveLoc[6](Arg6: u64)\n\t108: Ge\n\t109: StLoc[31](loc20: bool)\n\t110: MoveLoc[31](loc20: bool)\n\t111: MoveLoc[47](loc36: u64)\n\t112: MoveLoc[48](loc37: u64)\n\t113: LdU64(0)\n\t114: MoveLoc[42](loc31: u64)\n\t115: Add\n\t116: Sub\n\t117: StLoc[14](loc3: u64)\n\t118: StLoc[13](loc2: u64)\n\t119: StLoc[12](loc1: bool)\nB10:\n\t120: MoveLoc[12](loc1: bool)\n\t121: MoveLoc[13](loc2: u64)\n\t122: MoveLoc[14](loc3: u64)\n\t123: StLoc[40](loc29: u64)\n\t124: StLoc[25](loc14: u64)\n\t125: StLoc[32](loc21: bool)\n\t126: MoveLoc[32](loc21: bool)\n\t127: BrTrue(134)\nB11:\n\t128: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t129: Pop\n\t130: MoveLoc[10](Arg10: &mut TxContext)\n\t131: Pop\n\t132: LdConst[8](U64: [12, 0, 0, 0, 0, 0, 0, 0])\n\t133: Abort\nB12:\n\t134: MoveLoc[8](Arg8: u8)\n\t135: LdConst[1](U8: [0])\n\t136: Eq\n\t137: BrTrue(139)\nB13:\n\t138: Branch(173)\nB14:\n\t139: ImmBorrowLoc[4](Arg4: Coin<Ty1>)\n\t140: Call[20](value<Ty1>(&Coin<Ty1>): u64)\n\t141: StLoc[41](loc30: u64)\n\t142: MoveLoc[41](loc30: u64)\n\t143: MoveLoc[40](loc29: u64)\n\t144: Sub\n\t145: StLoc[28](loc17: u64)\n\t146: MutBorrowLoc[4](Arg4: Coin<Ty1>)\n\t147: MoveLoc[28](loc17: u64)\n\t148: CopyLoc[10](Arg10: &mut TxContext)\n\t149: Call[21](split<Ty1>(&mut Coin<Ty1>, u64, &mut TxContext))\n\t150: MoveLoc[3](Arg3: Coin<Ty0>)\n\t151: CopyLoc[10](Arg10: &mut TxContext)\n\t152: FreezeRef\n\t153: Call[25](sender(&TxContext): address)\n\t154: Call[22](transfer<Coin<Ty0>>(Coin<Ty0>, address))\n\t155: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t156: MutBorrowFieldGeneric[11](Market.quote: Balance<Ty1>)\n\t157: MoveLoc[4](Arg4: Coin<Ty1>)\n\t158: Call[9](into_balance<Ty1>(Coin<Ty1>): Balance<Ty1>)\n\t159: Call[23](join<Ty1>(&mut Balance<Ty1>, Balance<Ty1>): u64)\n\t160: Pop\n\t161: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t162: MutBorrowFieldGeneric[12](Market.base: Balance<Ty0>)\n\t163: MoveLoc[25](loc14: u64)\n\t164: CopyLoc[10](Arg10: &mut TxContext)\n\t165: Call[38](take<Ty0>(&mut Balance<Ty0>, u64, &mut TxContext): Coin<Ty0>)\n\t166: StLoc[36](loc25: Coin<Ty0>)\n\t167: MoveLoc[36](loc25: Coin<Ty0>)\n\t168: MoveLoc[10](Arg10: &mut TxContext)\n\t169: FreezeRef\n\t170: Call[25](sender(&TxContext): address)\n\t171: Call[22](transfer<Coin<Ty0>>(Coin<Ty0>, address))\n\t172: Branch(206)\nB15:\n\t173: ImmBorrowLoc[3](Arg3: Coin<Ty0>)\n\t174: Call[24](value<Ty0>(&Coin<Ty0>): u64)\n\t175: StLoc[26](loc15: u64)\n\t176: MoveLoc[26](loc15: u64)\n\t177: MoveLoc[25](loc14: u64)\n\t178: Sub\n\t179: StLoc[29](loc18: u64)\n\t180: MutBorrowLoc[3](Arg3: Coin<Ty0>)\n\t181: MoveLoc[29](loc18: u64)\n\t182: CopyLoc[10](Arg10: &mut TxContext)\n\t183: Call[25](split<Ty0>(&mut Coin<Ty0>, u64, &mut TxContext))\n\t184: MoveLoc[4](Arg4: Coin<Ty1>)\n\t185: CopyLoc[10](Arg10: &mut TxContext)\n\t186: FreezeRef\n\t187: Call[25](sender(&TxContext): address)\n\t188: Call[26](transfer<Coin<Ty1>>(Coin<Ty1>, address))\n\t189: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t190: MutBorrowFieldGeneric[12](Market.base: Balance<Ty0>)\n\t191: MoveLoc[3](Arg3: Coin<Ty0>)\n\t192: Call[8](into_balance<Ty0>(Coin<Ty0>): Balance<Ty0>)\n\t193: Call[27](join<Ty0>(&mut Balance<Ty0>, Balance<Ty0>): u64)\n\t194: Pop\n\t195: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t196: MutBorrowFieldGeneric[11](Market.quote: Balance<Ty1>)\n\t197: MoveLoc[40](loc29: u64)\n\t198: CopyLoc[10](Arg10: &mut TxContext)\n\t199: Call[37](take<Ty1>(&mut Balance<Ty1>, u64, &mut TxContext): Coin<Ty1>)\n\t200: StLoc[37](loc26: Coin<Ty1>)\n\t201: MoveLoc[37](loc26: Coin<Ty1>)\n\t202: MoveLoc[10](Arg10: &mut TxContext)\n\t203: FreezeRef\n\t204: Call[25](sender(&TxContext): address)\n\t205: Call[26](transfer<Coin<Ty1>>(Coin<Ty1>, address))\nB16:\n\t206: Ret\n}\nentry public sweep_fees<Ty0, Ty1>(Arg0: &mut Market<Ty0, Ty1>, Arg1: &mut TxContext) {\nB0:\n\t0: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t1: ImmBorrowFieldGeneric[8](Market.admin: address)\n\t2: ReadRef\n\t3: CopyLoc[1](Arg1: &mut TxContext)\n\t4: FreezeRef\n\t5: Call[25](sender(&TxContext): address)\n\t6: Eq\n\t7: BrTrue(14)\nB1:\n\t8: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t9: Pop\n\t10: MoveLoc[1](Arg1: &mut TxContext)\n\t11: Pop\n\t12: LdConst[12](U64: [6, 0, 0, 0, 0, 0, 0, 0])\n\t13: Abort\nB2:\n\t14: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t15: ImmBorrowFieldGeneric[2](Market.accumulated_fees: u64)\n\t16: ReadRef\n\t17: StLoc[2](loc0: u64)\n\t18: CopyLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t19: MutBorrowFieldGeneric[11](Market.quote: Balance<Ty1>)\n\t20: MoveLoc[2](loc0: u64)\n\t21: CopyLoc[1](Arg1: &mut TxContext)\n\t22: Call[37](take<Ty1>(&mut Balance<Ty1>, u64, &mut TxContext): Coin<Ty1>)\n\t23: MoveLoc[1](Arg1: &mut TxContext)\n\t24: FreezeRef\n\t25: Call[25](sender(&TxContext): address)\n\t26: Call[26](transfer<Coin<Ty1>>(Coin<Ty1>, address))\n\t27: LdU64(0)\n\t28: MoveLoc[0](Arg0: &mut Market<Ty0, Ty1>)\n\t29: MutBorrowFieldGeneric[2](Market.accumulated_fees: u64)\n\t30: WriteRef\n\t31: Ret\n}\n}"
	order_id := "// Move bytecode v5\nmodule 0.order_id {\n\n\npublic order_id(Arg0: u64, Arg1: u64, Arg2: bool): u128 {\nB0:\n\t0: MoveLoc[2](Arg2: bool)\n\t1: LdConst[0](Bool: [1])\n\t2: Eq\n\t3: BrTrue(5)\nB1:\n\t4: Branch(10)\nB2:\n\t5: MoveLoc[0](Arg0: u64)\n\t6: MoveLoc[1](Arg1: u64)\n\t7: Call[1](order_id_ask(u64, u64): u128)\n\t8: StLoc[3](loc0: u128)\n\t9: Branch(14)\nB3:\n\t10: MoveLoc[0](Arg0: u64)\n\t11: MoveLoc[1](Arg1: u64)\n\t12: Call[2](order_id_bid(u64, u64): u128)\n\t13: StLoc[3](loc0: u128)\nB4:\n\t14: MoveLoc[3](loc0: u128)\n\t15: Ret\n}\npublic order_id_ask(Arg0: u64, Arg1: u64): u128 {\nB0:\n\t0: MoveLoc[0](Arg0: u64)\n\t1: CastU128\n\t2: LdConst[2](U8: [64])\n\t3: Shl\n\t4: MoveLoc[1](Arg1: u64)\n\t5: CastU128\n\t6: BitOr\n\t7: Ret\n}\npublic order_id_bid(Arg0: u64, Arg1: u64): u128 {\nB0:\n\t0: MoveLoc[0](Arg0: u64)\n\t1: CastU128\n\t2: LdConst[2](U8: [64])\n\t3: Shl\n\t4: MoveLoc[1](Arg1: u64)\n\t5: LdConst[3](U64: [255, 255, 255, 255, 255, 255, 255, 255])\n\t6: Xor\n\t7: CastU128\n\t8: BitOr\n\t9: Ret\n}\npublic price(Arg0: u128): u64 {\nB0:\n\t0: MoveLoc[0](Arg0: u128)\n\t1: LdConst[2](U8: [64])\n\t2: Shr\n\t3: CastU64\n\t4: Ret\n}\npublic serial_id_ask(Arg0: u128): u64 {\nB0:\n\t0: MoveLoc[0](Arg0: u128)\n\t1: LdConst[3](U64: [255, 255, 255, 255, 255, 255, 255, 255])\n\t2: CastU128\n\t3: BitAnd\n\t4: CastU64\n\t5: Ret\n}\npublic serial_id_bid(Arg0: u128): u64 {\nB0:\n\t0: MoveLoc[0](Arg0: u128)\n\t1: LdConst[3](U64: [255, 255, 255, 255, 255, 255, 255, 255])\n\t2: CastU128\n\t3: BitAnd\n\t4: CastU64\n\t5: LdConst[3](U64: [255, 255, 255, 255, 255, 255, 255, 255])\n\t6: Xor\n\t7: Ret\n}\n}"
	pool := "// Move bytecode v5\nmodule 0.pool {\nstruct LSP<phantom Ty0, phantom Ty1> has drop {\n\tdummy_field: bool\n}\nstruct Pool<phantom Ty0, phantom Ty1> has key {\n\tid: UID,\n\tbase: Balance<Ty0>,\n\tquote: Balance<Ty1>,\n\tlsp_supply: Supply<LSP<Ty0, Ty1>>,\n\tfee_percent: u64\n}\n\npublic add_liquidity<Ty0, Ty1>(Arg0: &mut Pool<Ty0, Ty1>, Arg1: Coin<Ty0>, Arg2: Coin<Ty1>, Arg3: &mut TxContext): Coin<LSP<Ty0, Ty1>> {\nL0:\tloc4: u64\nL1:\tloc5: u64\nL2:\tloc6: u64\nL3:\tloc7: u64\nL4:\tloc8: u64\nL5:\tloc9: Balance<Ty1>\nL6:\tloc10: u64\nB0:\n\t0: ImmBorrowLoc[1](Arg1: Coin<Ty0>)\n\t1: Call[0](value<Ty0>(&Coin<Ty0>): u64)\n\t2: LdU64(0)\n\t3: Gt\n\t4: BrTrue(11)\nB1:\n\t5: MoveLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t6: Pop\n\t7: MoveLoc[3](Arg3: &mut TxContext)\n\t8: Pop\n\t9: LdConst[4](U64: [0, 0, 0, 0, 0, 0, 0, 0])\n\t10: Abort\nB2:\n\t11: ImmBorrowLoc[2](Arg2: Coin<Ty1>)\n\t12: Call[1](value<Ty1>(&Coin<Ty1>): u64)\n\t13: LdU64(0)\n\t14: Gt\n\t15: BrTrue(22)\nB3:\n\t16: MoveLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t17: Pop\n\t18: MoveLoc[3](Arg3: &mut TxContext)\n\t19: Pop\n\t20: LdConst[4](U64: [0, 0, 0, 0, 0, 0, 0, 0])\n\t21: Abort\nB4:\n\t22: MoveLoc[1](Arg1: Coin<Ty0>)\n\t23: Call[2](into_balance<Ty0>(Coin<Ty0>): Balance<Ty0>)\n\t24: StLoc[7](loc3: Balance<Ty0>)\n\t25: MoveLoc[2](Arg2: Coin<Ty1>)\n\t26: Call[3](into_balance<Ty1>(Coin<Ty1>): Balance<Ty1>)\n\t27: StLoc[13](loc9: Balance<Ty1>)\n\t28: CopyLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t29: FreezeRef\n\t30: Call[4](get_amounts<Ty0, Ty1>(&Pool<Ty0, Ty1>): u64 * u64 * u64)\n\t31: StLoc[10](loc6: u64)\n\t32: StLoc[12](loc8: u64)\n\t33: StLoc[6](loc2: u64)\n\t34: ImmBorrowLoc[7](loc3: Balance<Ty0>)\n\t35: Call[5](value<Ty0>(&Balance<Ty0>): u64)\n\t36: StLoc[5](loc1: u64)\n\t37: ImmBorrowLoc[13](loc9: Balance<Ty1>)\n\t38: Call[6](value<Ty1>(&Balance<Ty1>): u64)\n\t39: StLoc[11](loc7: u64)\n\t40: MoveLoc[5](loc1: u64)\n\t41: CopyLoc[10](loc6: u64)\n\t42: Mul\n\t43: MoveLoc[6](loc2: u64)\n\t44: Div\n\t45: MoveLoc[11](loc7: u64)\n\t46: MoveLoc[10](loc6: u64)\n\t47: Mul\n\t48: MoveLoc[12](loc8: u64)\n\t49: Div\n\t50: Call[18](min(u64, u64): u64)\n\t51: StLoc[14](loc10: u64)\n\t52: CopyLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t53: MutBorrowFieldGeneric[0](Pool.base: Balance<Ty0>)\n\t54: MoveLoc[7](loc3: Balance<Ty0>)\n\t55: Call[7](join<Ty0>(&mut Balance<Ty0>, Balance<Ty0>): u64)\n\t56: StLoc[8](loc4: u64)\n\t57: CopyLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t58: MutBorrowFieldGeneric[1](Pool.quote: Balance<Ty1>)\n\t59: MoveLoc[13](loc9: Balance<Ty1>)\n\t60: Call[8](join<Ty1>(&mut Balance<Ty1>, Balance<Ty1>): u64)\n\t61: StLoc[9](loc5: u64)\n\t62: MoveLoc[8](loc4: u64)\n\t63: LdConst[6](U64: [203, 16, 199, 186, 184, 141, 6, 0])\n\t64: Lt\n\t65: BrTrue(72)\nB5:\n\t66: MoveLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t67: Pop\n\t68: MoveLoc[3](Arg3: &mut TxContext)\n\t69: Pop\n\t70: LdConst[0](U64: [4, 0, 0, 0, 0, 0, 0, 0])\n\t71: Abort\nB6:\n\t72: MoveLoc[9](loc5: u64)\n\t73: LdConst[6](U64: [203, 16, 199, 186, 184, 141, 6, 0])\n\t74: Lt\n\t75: BrTrue(82)\nB7:\n\t76: MoveLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t77: Pop\n\t78: MoveLoc[3](Arg3: &mut TxContext)\n\t79: Pop\n\t80: LdConst[0](U64: [4, 0, 0, 0, 0, 0, 0, 0])\n\t81: Abort\nB8:\n\t82: MoveLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t83: MutBorrowFieldGeneric[2](Pool.lsp_supply: Supply<LSP<Ty0, Ty1>>)\n\t84: MoveLoc[14](loc10: u64)\n\t85: Call[9](increase_supply<LSP<Ty0, Ty1>>(&mut Supply<LSP<Ty0, Ty1>>, u64): Balance<LSP<Ty0, Ty1>>)\n\t86: StLoc[4](loc0: Balance<LSP<Ty0, Ty1>>)\n\t87: MoveLoc[4](loc0: Balance<LSP<Ty0, Ty1>>)\n\t88: MoveLoc[3](Arg3: &mut TxContext)\n\t89: Call[10](from_balance<LSP<Ty0, Ty1>>(Balance<LSP<Ty0, Ty1>>, &mut TxContext): Coin<LSP<Ty0, Ty1>>)\n\t90: Ret\n}\nentry add_liquidity_<Ty0, Ty1>(Arg0: &mut Pool<Ty0, Ty1>, Arg1: Coin<Ty0>, Arg2: Coin<Ty1>, Arg3: &mut TxContext) {\nB0:\n\t0: MoveLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t1: MoveLoc[1](Arg1: Coin<Ty0>)\n\t2: MoveLoc[2](Arg2: Coin<Ty1>)\n\t3: CopyLoc[3](Arg3: &mut TxContext)\n\t4: Call[11](add_liquidity<Ty0, Ty1>(&mut Pool<Ty0, Ty1>, Coin<Ty0>, Coin<Ty1>, &mut TxContext): Coin<LSP<Ty0, Ty1>>)\n\t5: MoveLoc[3](Arg3: &mut TxContext)\n\t6: FreezeRef\n\t7: Call[22](sender(&TxContext): address)\n\t8: Call[12](transfer<Coin<LSP<Ty0, Ty1>>>(Coin<LSP<Ty0, Ty1>>, address))\n\t9: Ret\n}\npublic base_price<Ty0, Ty1>(Arg0: &Pool<Ty0, Ty1>, Arg1: u64): u64 {\nB0:\n\t0: CopyLoc[0](Arg0: &Pool<Ty0, Ty1>)\n\t1: Call[4](get_amounts<Ty0, Ty1>(&Pool<Ty0, Ty1>): u64 * u64 * u64)\n\t2: Pop\n\t3: StLoc[3](loc1: u64)\n\t4: StLoc[2](loc0: u64)\n\t5: MoveLoc[1](Arg1: u64)\n\t6: MoveLoc[3](loc1: u64)\n\t7: MoveLoc[2](loc0: u64)\n\t8: MoveLoc[0](Arg0: &Pool<Ty0, Ty1>)\n\t9: ImmBorrowFieldGeneric[3](Pool.fee_percent: u64)\n\t10: ReadRef\n\t11: Call[8](get_input_price(u64, u64, u64, u64): u64)\n\t12: Ret\n}\npublic buy<Ty0, Ty1>(Arg0: &mut Pool<Ty0, Ty1>, Arg1: Coin<Ty1>, Arg2: &mut TxContext): Coin<Ty0> {\nL0:\tloc3: Balance<Ty1>\nL1:\tloc4: u64\nB0:\n\t0: ImmBorrowLoc[1](Arg1: Coin<Ty1>)\n\t1: Call[1](value<Ty1>(&Coin<Ty1>): u64)\n\t2: LdU64(0)\n\t3: Gt\n\t4: BrTrue(11)\nB1:\n\t5: MoveLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t6: Pop\n\t7: MoveLoc[2](Arg2: &mut TxContext)\n\t8: Pop\n\t9: LdConst[4](U64: [0, 0, 0, 0, 0, 0, 0, 0])\n\t10: Abort\nB2:\n\t11: MoveLoc[1](Arg1: Coin<Ty1>)\n\t12: Call[3](into_balance<Ty1>(Coin<Ty1>): Balance<Ty1>)\n\t13: StLoc[6](loc3: Balance<Ty1>)\n\t14: CopyLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t15: FreezeRef\n\t16: Call[4](get_amounts<Ty0, Ty1>(&Pool<Ty0, Ty1>): u64 * u64 * u64)\n\t17: Pop\n\t18: StLoc[7](loc4: u64)\n\t19: StLoc[4](loc1: u64)\n\t20: CopyLoc[4](loc1: u64)\n\t21: LdU64(0)\n\t22: Gt\n\t23: BrTrue(25)\nB3:\n\t24: Branch(30)\nB4:\n\t25: CopyLoc[7](loc4: u64)\n\t26: LdU64(0)\n\t27: Gt\n\t28: StLoc[3](loc0: bool)\n\t29: Branch(32)\nB5:\n\t30: LdFalse\n\t31: StLoc[3](loc0: bool)\nB6:\n\t32: MoveLoc[3](loc0: bool)\n\t33: BrTrue(40)\nB7:\n\t34: MoveLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t35: Pop\n\t36: MoveLoc[2](Arg2: &mut TxContext)\n\t37: Pop\n\t38: LdConst[1](U64: [2, 0, 0, 0, 0, 0, 0, 0])\n\t39: Abort\nB8:\n\t40: ImmBorrowLoc[6](loc3: Balance<Ty1>)\n\t41: Call[6](value<Ty1>(&Balance<Ty1>): u64)\n\t42: MoveLoc[7](loc4: u64)\n\t43: MoveLoc[4](loc1: u64)\n\t44: CopyLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t45: ImmBorrowFieldGeneric[3](Pool.fee_percent: u64)\n\t46: ReadRef\n\t47: Call[8](get_input_price(u64, u64, u64, u64): u64)\n\t48: StLoc[5](loc2: u64)\n\t49: CopyLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t50: MutBorrowFieldGeneric[1](Pool.quote: Balance<Ty1>)\n\t51: MoveLoc[6](loc3: Balance<Ty1>)\n\t52: Call[8](join<Ty1>(&mut Balance<Ty1>, Balance<Ty1>): u64)\n\t53: Pop\n\t54: MoveLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t55: MutBorrowFieldGeneric[0](Pool.base: Balance<Ty0>)\n\t56: MoveLoc[5](loc2: u64)\n\t57: MoveLoc[2](Arg2: &mut TxContext)\n\t58: Call[13](take<Ty0>(&mut Balance<Ty0>, u64, &mut TxContext): Coin<Ty0>)\n\t59: Ret\n}\nentry buy_<Ty0, Ty1>(Arg0: &mut Pool<Ty0, Ty1>, Arg1: Coin<Ty1>, Arg2: &mut TxContext) {\nB0:\n\t0: MoveLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t1: MoveLoc[1](Arg1: Coin<Ty1>)\n\t2: CopyLoc[2](Arg2: &mut TxContext)\n\t3: Call[14](buy<Ty0, Ty1>(&mut Pool<Ty0, Ty1>, Coin<Ty1>, &mut TxContext): Coin<Ty0>)\n\t4: MoveLoc[2](Arg2: &mut TxContext)\n\t5: FreezeRef\n\t6: Call[22](sender(&TxContext): address)\n\t7: Call[15](transfer<Coin<Ty0>>(Coin<Ty0>, address))\n\t8: Ret\n}\npublic create_pool<Ty0, Ty1>(Arg0: Coin<Ty0>, Arg1: Coin<Ty1>, Arg2: u64, Arg3: &mut TxContext): Coin<LSP<Ty0, Ty1>> {\nL0:\tloc4: Balance<LSP<Ty0, Ty1>>\nL1:\tloc5: Supply<LSP<Ty0, Ty1>>\nL2:\tloc6: u64\nL3:\tloc7: u64\nB0:\n\t0: ImmBorrowLoc[1](Arg1: Coin<Ty1>)\n\t1: Call[1](value<Ty1>(&Coin<Ty1>): u64)\n\t2: StLoc[10](loc6: u64)\n\t3: ImmBorrowLoc[0](Arg0: Coin<Ty0>)\n\t4: Call[0](value<Ty0>(&Coin<Ty0>): u64)\n\t5: StLoc[7](loc3: u64)\n\t6: CopyLoc[7](loc3: u64)\n\t7: LdU64(0)\n\t8: Gt\n\t9: BrTrue(11)\nB1:\n\t10: Branch(16)\nB2:\n\t11: CopyLoc[10](loc6: u64)\n\t12: LdU64(0)\n\t13: Gt\n\t14: StLoc[4](loc0: bool)\n\t15: Branch(18)\nB3:\n\t16: LdFalse\n\t17: StLoc[4](loc0: bool)\nB4:\n\t18: MoveLoc[4](loc0: bool)\n\t19: BrTrue(24)\nB5:\n\t20: MoveLoc[3](Arg3: &mut TxContext)\n\t21: Pop\n\t22: LdConst[4](U64: [0, 0, 0, 0, 0, 0, 0, 0])\n\t23: Abort\nB6:\n\t24: CopyLoc[7](loc3: u64)\n\t25: LdConst[6](U64: [203, 16, 199, 186, 184, 141, 6, 0])\n\t26: Lt\n\t27: BrTrue(29)\nB7:\n\t28: Branch(34)\nB8:\n\t29: CopyLoc[10](loc6: u64)\n\t30: LdConst[6](U64: [203, 16, 199, 186, 184, 141, 6, 0])\n\t31: Lt\n\t32: StLoc[5](loc1: bool)\n\t33: Branch(36)\nB9:\n\t34: LdFalse\n\t35: StLoc[5](loc1: bool)\nB10:\n\t36: MoveLoc[5](loc1: bool)\n\t37: BrTrue(42)\nB11:\n\t38: MoveLoc[3](Arg3: &mut TxContext)\n\t39: Pop\n\t40: LdConst[0](U64: [4, 0, 0, 0, 0, 0, 0, 0])\n\t41: Abort\nB12:\n\t42: CopyLoc[2](Arg2: u64)\n\t43: LdU64(0)\n\t44: Ge\n\t45: BrTrue(47)\nB13:\n\t46: Branch(52)\nB14:\n\t47: CopyLoc[2](Arg2: u64)\n\t48: LdU64(10000)\n\t49: Lt\n\t50: StLoc[6](loc2: bool)\n\t51: Branch(54)\nB15:\n\t52: LdFalse\n\t53: StLoc[6](loc2: bool)\nB16:\n\t54: MoveLoc[6](loc2: bool)\n\t55: BrTrue(60)\nB17:\n\t56: MoveLoc[3](Arg3: &mut TxContext)\n\t57: Pop\n\t58: LdConst[3](U64: [1, 0, 0, 0, 0, 0, 0, 0])\n\t59: Abort\nB18:\n\t60: MoveLoc[10](loc6: u64)\n\t61: Call[25](sqrt(u64): u64)\n\t62: MoveLoc[7](loc3: u64)\n\t63: Call[25](sqrt(u64): u64)\n\t64: Mul\n\t65: StLoc[11](loc7: u64)\n\t66: LdFalse\n\t67: PackGeneric[0](LSP<Ty0, Ty1>)\n\t68: Call[16](create_supply<LSP<Ty0, Ty1>>(LSP<Ty0, Ty1>): Supply<LSP<Ty0, Ty1>>)\n\t69: StLoc[9](loc5: Supply<LSP<Ty0, Ty1>>)\n\t70: MutBorrowLoc[9](loc5: Supply<LSP<Ty0, Ty1>>)\n\t71: MoveLoc[11](loc7: u64)\n\t72: Call[9](increase_supply<LSP<Ty0, Ty1>>(&mut Supply<LSP<Ty0, Ty1>>, u64): Balance<LSP<Ty0, Ty1>>)\n\t73: StLoc[8](loc4: Balance<LSP<Ty0, Ty1>>)\n\t74: CopyLoc[3](Arg3: &mut TxContext)\n\t75: Call[27](new(&mut TxContext): UID)\n\t76: MoveLoc[0](Arg0: Coin<Ty0>)\n\t77: Call[2](into_balance<Ty0>(Coin<Ty0>): Balance<Ty0>)\n\t78: MoveLoc[1](Arg1: Coin<Ty1>)\n\t79: Call[3](into_balance<Ty1>(Coin<Ty1>): Balance<Ty1>)\n\t80: MoveLoc[9](loc5: Supply<LSP<Ty0, Ty1>>)\n\t81: MoveLoc[2](Arg2: u64)\n\t82: PackGeneric[1](Pool<Ty0, Ty1>)\n\t83: Call[17](share_object<Pool<Ty0, Ty1>>(Pool<Ty0, Ty1>))\n\t84: MoveLoc[8](loc4: Balance<LSP<Ty0, Ty1>>)\n\t85: MoveLoc[3](Arg3: &mut TxContext)\n\t86: Call[10](from_balance<LSP<Ty0, Ty1>>(Balance<LSP<Ty0, Ty1>>, &mut TxContext): Coin<LSP<Ty0, Ty1>>)\n\t87: Ret\n}\nentry create_pool_<Ty0, Ty1>(Arg0: Coin<Ty0>, Arg1: Coin<Ty1>, Arg2: u64, Arg3: &mut TxContext) {\nB0:\n\t0: MoveLoc[0](Arg0: Coin<Ty0>)\n\t1: MoveLoc[1](Arg1: Coin<Ty1>)\n\t2: MoveLoc[2](Arg2: u64)\n\t3: CopyLoc[3](Arg3: &mut TxContext)\n\t4: Call[18](create_pool<Ty0, Ty1>(Coin<Ty0>, Coin<Ty1>, u64, &mut TxContext): Coin<LSP<Ty0, Ty1>>)\n\t5: MoveLoc[3](Arg3: &mut TxContext)\n\t6: FreezeRef\n\t7: Call[22](sender(&TxContext): address)\n\t8: Call[12](transfer<Coin<LSP<Ty0, Ty1>>>(Coin<LSP<Ty0, Ty1>>, address))\n\t9: Ret\n}\npublic get_amounts<Ty0, Ty1>(Arg0: &Pool<Ty0, Ty1>): u64 * u64 * u64 {\nB0:\n\t0: CopyLoc[0](Arg0: &Pool<Ty0, Ty1>)\n\t1: ImmBorrowFieldGeneric[0](Pool.base: Balance<Ty0>)\n\t2: Call[5](value<Ty0>(&Balance<Ty0>): u64)\n\t3: CopyLoc[0](Arg0: &Pool<Ty0, Ty1>)\n\t4: ImmBorrowFieldGeneric[1](Pool.quote: Balance<Ty1>)\n\t5: Call[6](value<Ty1>(&Balance<Ty1>): u64)\n\t6: MoveLoc[0](Arg0: &Pool<Ty0, Ty1>)\n\t7: ImmBorrowFieldGeneric[2](Pool.lsp_supply: Supply<LSP<Ty0, Ty1>>)\n\t8: Call[19](supply_value<LSP<Ty0, Ty1>>(&Supply<LSP<Ty0, Ty1>>): u64)\n\t9: Ret\n}\npublic get_input_price(Arg0: u64, Arg1: u64, Arg2: u64, Arg3: u64): u64 {\nL0:\tloc4: u128\nL1:\tloc5: u128\nL2:\tloc6: u128\nB0:\n\t0: MoveLoc[0](Arg0: u64)\n\t1: CastU128\n\t2: MoveLoc[1](Arg1: u64)\n\t3: CastU128\n\t4: MoveLoc[2](Arg2: u64)\n\t5: CastU128\n\t6: MoveLoc[3](Arg3: u64)\n\t7: CastU128\n\t8: StLoc[5](loc1: u128)\n\t9: StLoc[10](loc6: u128)\n\t10: StLoc[8](loc4: u128)\n\t11: StLoc[6](loc2: u128)\n\t12: MoveLoc[6](loc2: u128)\n\t13: LdConst[5](U128: [16, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0])\n\t14: MoveLoc[5](loc1: u128)\n\t15: Sub\n\t16: Mul\n\t17: StLoc[7](loc3: u128)\n\t18: CopyLoc[7](loc3: u128)\n\t19: MoveLoc[10](loc6: u128)\n\t20: Mul\n\t21: StLoc[9](loc5: u128)\n\t22: MoveLoc[8](loc4: u128)\n\t23: LdConst[5](U128: [16, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0])\n\t24: Mul\n\t25: MoveLoc[7](loc3: u128)\n\t26: Add\n\t27: StLoc[4](loc0: u128)\n\t28: MoveLoc[9](loc5: u128)\n\t29: MoveLoc[4](loc0: u128)\n\t30: Div\n\t31: CastU64\n\t32: Ret\n}\ninit(Arg0: &mut TxContext) {\nB0:\n\t0: Ret\n}\npublic remove_liquidity<Ty0, Ty1>(Arg0: &mut Pool<Ty0, Ty1>, Arg1: Coin<LSP<Ty0, Ty1>>, Arg2: &mut TxContext): Coin<Ty0> * Coin<Ty1> {\nL0:\tloc3: u64\nL1:\tloc4: u64\nL2:\tloc5: u64\nB0:\n\t0: ImmBorrowLoc[1](Arg1: Coin<LSP<Ty0, Ty1>>)\n\t1: Call[20](value<LSP<Ty0, Ty1>>(&Coin<LSP<Ty0, Ty1>>): u64)\n\t2: StLoc[5](loc2: u64)\n\t3: CopyLoc[5](loc2: u64)\n\t4: LdU64(0)\n\t5: Gt\n\t6: BrTrue(13)\nB1:\n\t7: MoveLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t8: Pop\n\t9: MoveLoc[2](Arg2: &mut TxContext)\n\t10: Pop\n\t11: LdConst[4](U64: [0, 0, 0, 0, 0, 0, 0, 0])\n\t12: Abort\nB2:\n\t13: CopyLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t14: FreezeRef\n\t15: Call[4](get_amounts<Ty0, Ty1>(&Pool<Ty0, Ty1>): u64 * u64 * u64)\n\t16: StLoc[6](loc3: u64)\n\t17: StLoc[7](loc4: u64)\n\t18: StLoc[3](loc0: u64)\n\t19: MoveLoc[3](loc0: u64)\n\t20: CopyLoc[5](loc2: u64)\n\t21: Mul\n\t22: CopyLoc[6](loc3: u64)\n\t23: Div\n\t24: StLoc[4](loc1: u64)\n\t25: MoveLoc[7](loc4: u64)\n\t26: MoveLoc[5](loc2: u64)\n\t27: Mul\n\t28: MoveLoc[6](loc3: u64)\n\t29: Div\n\t30: StLoc[8](loc5: u64)\n\t31: CopyLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t32: MutBorrowFieldGeneric[2](Pool.lsp_supply: Supply<LSP<Ty0, Ty1>>)\n\t33: MoveLoc[1](Arg1: Coin<LSP<Ty0, Ty1>>)\n\t34: Call[21](into_balance<LSP<Ty0, Ty1>>(Coin<LSP<Ty0, Ty1>>): Balance<LSP<Ty0, Ty1>>)\n\t35: Call[22](decrease_supply<LSP<Ty0, Ty1>>(&mut Supply<LSP<Ty0, Ty1>>, Balance<LSP<Ty0, Ty1>>): u64)\n\t36: Pop\n\t37: CopyLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t38: MutBorrowFieldGeneric[0](Pool.base: Balance<Ty0>)\n\t39: MoveLoc[4](loc1: u64)\n\t40: CopyLoc[2](Arg2: &mut TxContext)\n\t41: Call[13](take<Ty0>(&mut Balance<Ty0>, u64, &mut TxContext): Coin<Ty0>)\n\t42: MoveLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t43: MutBorrowFieldGeneric[1](Pool.quote: Balance<Ty1>)\n\t44: MoveLoc[8](loc5: u64)\n\t45: MoveLoc[2](Arg2: &mut TxContext)\n\t46: Call[23](take<Ty1>(&mut Balance<Ty1>, u64, &mut TxContext): Coin<Ty1>)\n\t47: Ret\n}\nentry remove_liquidity_<Ty0, Ty1>(Arg0: &mut Pool<Ty0, Ty1>, Arg1: Coin<LSP<Ty0, Ty1>>, Arg2: &mut TxContext) {\nB0:\n\t0: MoveLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t1: MoveLoc[1](Arg1: Coin<LSP<Ty0, Ty1>>)\n\t2: CopyLoc[2](Arg2: &mut TxContext)\n\t3: Call[24](remove_liquidity<Ty0, Ty1>(&mut Pool<Ty0, Ty1>, Coin<LSP<Ty0, Ty1>>, &mut TxContext): Coin<Ty0> * Coin<Ty1>)\n\t4: StLoc[4](loc1: Coin<Ty1>)\n\t5: StLoc[3](loc0: Coin<Ty0>)\n\t6: MoveLoc[2](Arg2: &mut TxContext)\n\t7: FreezeRef\n\t8: Call[22](sender(&TxContext): address)\n\t9: StLoc[5](loc2: address)\n\t10: MoveLoc[3](loc0: Coin<Ty0>)\n\t11: CopyLoc[5](loc2: address)\n\t12: Call[15](transfer<Coin<Ty0>>(Coin<Ty0>, address))\n\t13: MoveLoc[4](loc1: Coin<Ty1>)\n\t14: MoveLoc[5](loc2: address)\n\t15: Call[25](transfer<Coin<Ty1>>(Coin<Ty1>, address))\n\t16: Ret\n}\npublic sell<Ty0, Ty1>(Arg0: &mut Pool<Ty0, Ty1>, Arg1: Coin<Ty0>, Arg2: &mut TxContext): Coin<Ty1> {\nL0:\tloc3: u64\nL1:\tloc4: u64\nB0:\n\t0: ImmBorrowLoc[1](Arg1: Coin<Ty0>)\n\t1: Call[0](value<Ty0>(&Coin<Ty0>): u64)\n\t2: LdU64(0)\n\t3: Gt\n\t4: BrTrue(11)\nB1:\n\t5: MoveLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t6: Pop\n\t7: MoveLoc[2](Arg2: &mut TxContext)\n\t8: Pop\n\t9: LdConst[4](U64: [0, 0, 0, 0, 0, 0, 0, 0])\n\t10: Abort\nB2:\n\t11: MoveLoc[1](Arg1: Coin<Ty0>)\n\t12: Call[2](into_balance<Ty0>(Coin<Ty0>): Balance<Ty0>)\n\t13: StLoc[4](loc1: Balance<Ty0>)\n\t14: CopyLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t15: FreezeRef\n\t16: Call[4](get_amounts<Ty0, Ty1>(&Pool<Ty0, Ty1>): u64 * u64 * u64)\n\t17: Pop\n\t18: StLoc[7](loc4: u64)\n\t19: StLoc[5](loc2: u64)\n\t20: CopyLoc[5](loc2: u64)\n\t21: LdU64(0)\n\t22: Gt\n\t23: BrTrue(25)\nB3:\n\t24: Branch(30)\nB4:\n\t25: CopyLoc[7](loc4: u64)\n\t26: LdU64(0)\n\t27: Gt\n\t28: StLoc[3](loc0: bool)\n\t29: Branch(32)\nB5:\n\t30: LdFalse\n\t31: StLoc[3](loc0: bool)\nB6:\n\t32: MoveLoc[3](loc0: bool)\n\t33: BrTrue(40)\nB7:\n\t34: MoveLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t35: Pop\n\t36: MoveLoc[2](Arg2: &mut TxContext)\n\t37: Pop\n\t38: LdConst[1](U64: [2, 0, 0, 0, 0, 0, 0, 0])\n\t39: Abort\nB8:\n\t40: ImmBorrowLoc[4](loc1: Balance<Ty0>)\n\t41: Call[5](value<Ty0>(&Balance<Ty0>): u64)\n\t42: MoveLoc[5](loc2: u64)\n\t43: MoveLoc[7](loc4: u64)\n\t44: CopyLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t45: ImmBorrowFieldGeneric[3](Pool.fee_percent: u64)\n\t46: ReadRef\n\t47: Call[8](get_input_price(u64, u64, u64, u64): u64)\n\t48: StLoc[6](loc3: u64)\n\t49: CopyLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t50: MutBorrowFieldGeneric[0](Pool.base: Balance<Ty0>)\n\t51: MoveLoc[4](loc1: Balance<Ty0>)\n\t52: Call[7](join<Ty0>(&mut Balance<Ty0>, Balance<Ty0>): u64)\n\t53: Pop\n\t54: MoveLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t55: MutBorrowFieldGeneric[1](Pool.quote: Balance<Ty1>)\n\t56: MoveLoc[6](loc3: u64)\n\t57: MoveLoc[2](Arg2: &mut TxContext)\n\t58: Call[23](take<Ty1>(&mut Balance<Ty1>, u64, &mut TxContext): Coin<Ty1>)\n\t59: Ret\n}\nentry sell_<Ty0, Ty1>(Arg0: &mut Pool<Ty0, Ty1>, Arg1: Coin<Ty0>, Arg2: &mut TxContext) {\nB0:\n\t0: MoveLoc[0](Arg0: &mut Pool<Ty0, Ty1>)\n\t1: MoveLoc[1](Arg1: Coin<Ty0>)\n\t2: CopyLoc[2](Arg2: &mut TxContext)\n\t3: Call[26](sell<Ty0, Ty1>(&mut Pool<Ty0, Ty1>, Coin<Ty0>, &mut TxContext): Coin<Ty1>)\n\t4: MoveLoc[2](Arg2: &mut TxContext)\n\t5: FreezeRef\n\t6: Call[22](sender(&TxContext): address)\n\t7: Call[25](transfer<Coin<Ty1>>(Coin<Ty1>, address))\n\t8: Ret\n}\npublic token_price<Ty0, Ty1>(Arg0: &Pool<Ty0, Ty1>, Arg1: u64): u64 {\nB0:\n\t0: CopyLoc[0](Arg0: &Pool<Ty0, Ty1>)\n\t1: Call[4](get_amounts<Ty0, Ty1>(&Pool<Ty0, Ty1>): u64 * u64 * u64)\n\t2: Pop\n\t3: StLoc[3](loc1: u64)\n\t4: StLoc[2](loc0: u64)\n\t5: MoveLoc[1](Arg1: u64)\n\t6: MoveLoc[2](loc0: u64)\n\t7: MoveLoc[3](loc1: u64)\n\t8: MoveLoc[0](Arg0: &Pool<Ty0, Ty1>)\n\t9: ImmBorrowFieldGeneric[3](Pool.fee_percent: u64)\n\t10: ReadRef\n\t11: Call[8](get_input_price(u64, u64, u64, u64): u64)\n\t12: Ret\n}\n}"
	quote_coin := "// Move bytecode v5\nmodule 0.quote_coin {\nstruct QUOTE_COIN has drop {\n\tdummy_field: bool\n}\n\ninit(Arg0: QUOTE_COIN, Arg1: &mut TxContext) {\nB0:\n\t0: MoveLoc[0](Arg0: QUOTE_COIN)\n\t1: MoveLoc[1](Arg1: &mut TxContext)\n\t2: Call[0](create_currency<QUOTE_COIN>(QUOTE_COIN, &mut TxContext): TreasuryCap<QUOTE_COIN>)\n\t3: StLoc[2](loc0: TreasuryCap<QUOTE_COIN>)\n\t4: MoveLoc[2](loc0: TreasuryCap<QUOTE_COIN>)\n\t5: Call[1](share_object<TreasuryCap<QUOTE_COIN>>(TreasuryCap<QUOTE_COIN>))\n\t6: Ret\n}\nentry public mint(Arg0: &mut TreasuryCap<QUOTE_COIN>, Arg1: u64, Arg2: &mut TxContext) {\nB0:\n\t0: MoveLoc[0](Arg0: &mut TreasuryCap<QUOTE_COIN>)\n\t1: MoveLoc[1](Arg1: u64)\n\t2: CopyLoc[2](Arg2: &mut TxContext)\n\t3: Call[2](mint<QUOTE_COIN>(&mut TreasuryCap<QUOTE_COIN>, u64, &mut TxContext): Coin<QUOTE_COIN>)\n\t4: StLoc[3](loc0: Coin<QUOTE_COIN>)\n\t5: MoveLoc[3](loc0: Coin<QUOTE_COIN>)\n\t6: MoveLoc[2](Arg2: &mut TxContext)\n\t7: FreezeRef\n\t8: Call[6](sender(&TxContext): address)\n\t9: Call[3](transfer<QUOTE_COIN>(Coin<QUOTE_COIN>, address))\n\t10: Ret\n}\nentry public mint_zero(Arg0: &mut TxContext) {\nB0:\n\t0: CopyLoc[0](Arg0: &mut TxContext)\n\t1: Call[4](zero<QUOTE_COIN>(&mut TxContext): Coin<QUOTE_COIN>)\n\t2: StLoc[1](loc0: Coin<QUOTE_COIN>)\n\t3: MoveLoc[1](loc0: Coin<QUOTE_COIN>)\n\t4: MoveLoc[0](Arg0: &mut TxContext)\n\t5: FreezeRef\n\t6: Call[6](sender(&TxContext): address)\n\t7: Call[3](transfer<QUOTE_COIN>(Coin<QUOTE_COIN>, address))\n\t8: Ret\n}\n}"
	util := "// Move bytecode v5\nmodule 0.util {\n\n\npublic u64_saturating_sub(Arg0: u64, Arg1: u64): u64 {\nB0:\n\t0: CopyLoc[0](Arg0: u64)\n\t1: CopyLoc[1](Arg1: u64)\n\t2: Ge\n\t3: BrTrue(5)\nB1:\n\t4: Branch(10)\nB2:\n\t5: MoveLoc[0](Arg0: u64)\n\t6: MoveLoc[1](Arg1: u64)\n\t7: Sub\n\t8: StLoc[2](loc0: u64)\n\t9: Branch(12)\nB3:\n\t10: LdU64(0)\n\t11: StLoc[2](loc0: u64)\nB4:\n\t12: MoveLoc[2](loc0: u64)\n\t13: Ret\n}\n}"
	exp2 := Package{
		DeployTX: "AXVQbp46p/THvGhkEoP4/fIFdnnFyNHxYccBJqACzs8=",
		ID:       "0x6774c132abf5521bf9bd8e8c6948c03f0f0cb3f4",
		Bytecode: map[string]interface{}{
			"base_coin":  base_coin,
			"critbit":    critbit,
			"fp_math":    fp_math,
			"market":     market,
			"order_id":   order_id,
			"pool":       pool,
			"quote_coin": quote_coin,
			"util":       util,
		},
	}

	if !reflect.DeepEqual(got2, exp2) {
		t.Errorf("Result was incorrect, got %s, want %s.", got2, exp2)
	}

	result3, _ := tc.GetTransaction("NKfaECvNnkH/hx3OvJeWYAhIgaRJwA2U1zobtkpiAoY=")
	got3, _ := result3.GetContractDeploy()
	rgb := "// Move bytecode v5\nmodule 0.rgb {\nstruct ColorObject has key {\n\tid: UID,\n\tred: u8,\n\tgreen: u8,\n\tblue: u8\n}\n\nentry public copy_into(Arg0: &ColorObject, Arg1: &mut ColorObject) {\nB0:\n\t0: CopyLoc[0](Arg0: &ColorObject)\n\t1: ImmBorrowField[0](ColorObject.red: u8)\n\t2: ReadRef\n\t3: CopyLoc[1](Arg1: &mut ColorObject)\n\t4: MutBorrowField[0](ColorObject.red: u8)\n\t5: WriteRef\n\t6: CopyLoc[0](Arg0: &ColorObject)\n\t7: ImmBorrowField[1](ColorObject.green: u8)\n\t8: ReadRef\n\t9: CopyLoc[1](Arg1: &mut ColorObject)\n\t10: MutBorrowField[1](ColorObject.green: u8)\n\t11: WriteRef\n\t12: MoveLoc[0](Arg0: &ColorObject)\n\t13: ImmBorrowField[2](ColorObject.blue: u8)\n\t14: ReadRef\n\t15: MoveLoc[1](Arg1: &mut ColorObject)\n\t16: MutBorrowField[2](ColorObject.blue: u8)\n\t17: WriteRef\n\t18: Ret\n}\nentry public create(Arg0: u8, Arg1: u8, Arg2: u8, Arg3: &mut TxContext) {\nB0:\n\t0: MoveLoc[0](Arg0: u8)\n\t1: MoveLoc[1](Arg1: u8)\n\t2: MoveLoc[2](Arg2: u8)\n\t3: CopyLoc[3](Arg3: &mut TxContext)\n\t4: Call[4](new(u8, u8, u8, &mut TxContext): ColorObject)\n\t5: StLoc[4](loc0: ColorObject)\n\t6: MoveLoc[4](loc0: ColorObject)\n\t7: MoveLoc[3](Arg3: &mut TxContext)\n\t8: FreezeRef\n\t9: Call[6](sender(&TxContext): address)\n\t10: Call[0](transfer<ColorObject>(ColorObject, address))\n\t11: Ret\n}\nentry public delete(Arg0: ColorObject) {\nB0:\n\t0: MoveLoc[0](Arg0: ColorObject)\n\t1: Unpack[0](ColorObject)\n\t2: Pop\n\t3: Pop\n\t4: Pop\n\t5: StLoc[1](loc0: UID)\n\t6: MoveLoc[1](loc0: UID)\n\t7: Call[8](delete(UID))\n\t8: Ret\n}\npublic get_color(Arg0: &ColorObject): u8 * u8 * u8 {\nB0:\n\t0: CopyLoc[0](Arg0: &ColorObject)\n\t1: ImmBorrowField[0](ColorObject.red: u8)\n\t2: ReadRef\n\t3: CopyLoc[0](Arg0: &ColorObject)\n\t4: ImmBorrowField[1](ColorObject.green: u8)\n\t5: ReadRef\n\t6: MoveLoc[0](Arg0: &ColorObject)\n\t7: ImmBorrowField[2](ColorObject.blue: u8)\n\t8: ReadRef\n\t9: Ret\n}\nnew(Arg0: u8, Arg1: u8, Arg2: u8, Arg3: &mut TxContext): ColorObject {\nB0:\n\t0: MoveLoc[3](Arg3: &mut TxContext)\n\t1: Call[9](new(&mut TxContext): UID)\n\t2: MoveLoc[0](Arg0: u8)\n\t3: MoveLoc[1](Arg1: u8)\n\t4: MoveLoc[2](Arg2: u8)\n\t5: Pack[0](ColorObject)\n\t6: Ret\n}\nentry public transfer(Arg0: ColorObject, Arg1: address, Arg2: &mut TxContext) {\nB0:\n\t0: MoveLoc[0](Arg0: ColorObject)\n\t1: MoveLoc[1](Arg1: address)\n\t2: Call[0](transfer<ColorObject>(ColorObject, address))\n\t3: Ret\n}\n}"
	exp3 := Package{
		DeployTX: "NKfaECvNnkH/hx3OvJeWYAhIgaRJwA2U1zobtkpiAoY=",
		ID:       "0xb82c2e33acecc71f2b2b742c366d902bd640a1bf",
		Bytecode: map[string]interface{}{
			"rgb": rgb,
		},
	}

	if !reflect.DeepEqual(got3, exp3) {
		t.Errorf("Result was incorrect, got %s, want %s.", got3, exp3)
	}
}
