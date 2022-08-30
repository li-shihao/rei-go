package database

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
	"rei.io/rei/ent/accounts"
	"rei.io/rei/ent/arguments"
	"rei.io/rei/ent/events"
	"rei.io/rei/ent/nfts"
	"rei.io/rei/ent/objects"
	"rei.io/rei/ent/packages"
	"rei.io/rei/ent/transactions"
	"rei.io/rei/internal/sui"
)

func TestCreateTransaction(t *testing.T) {
	p := postgres.Preset(
		postgres.WithUser("gnomock", "gnomick"),
		postgres.WithDatabase("mydb"),
	)

	container, _ := gnomock.Start(p)
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
		container.Host, container.DefaultPort(),
		"gnomock", "gnomick", "mydb",
	)

	sc := new(sui.SUIClient)
	sc.Init("http://127.0.0.1:9000")

	db := new(EntClient)
	db.Init("postgres", connStr)

	testTX1, _ := sc.GetTransaction("AGcI2C8xd7H4Vs26Z6uMbj924S+JIAIPHSZttqQoJUk=")
	db.CreateTransaction(testTX1)

	if got1, _ := db.client.Transactions.Query().Where(
		transactions.SenderEQ("0xc4173a804406a365e69dfb297d4eaaf002546ebd")).Exist(context.Background()); !got1 {
		t.Errorf("Result was incorrect, got %t, want %t.", got1, true)
	}

	testTX2, _ := sc.GetTransaction("dNOK2AVQrXkE6aiFghswKAI9qDJOdOaI0NzU+jKtIJw=")
	db.CreateTransaction(testTX2)

	if got2, _ := db.client.Transactions.Query().Where(
		transactions.TransactionIDEQ("dNOK2AVQrXkE6aiFghswKAI9qDJOdOaI0NzU+jKtIJw=")).First(context.Background()); got2.Type != "Call" {
		t.Errorf("Result was incorrect, got %s, want %s.", got2, "Call")
	}
}

func TestCreateEvent(t *testing.T) {
	p := postgres.Preset(
		postgres.WithUser("gnomock", "gnomick"),
		postgres.WithDatabase("mydb"),
	)

	container, _ := gnomock.Start(p)
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
		container.Host, container.DefaultPort(),
		"gnomock", "gnomick", "mydb",
	)

	sc := new(sui.SUIClient)
	sc.Init("http://127.0.0.1:9000")

	db := new(EntClient)
	db.Init("postgres", connStr)

	testTX1, _ := sc.GetTransaction("dNOK2AVQrXkE6aiFghswKAI9qDJOdOaI0NzU+jKtIJw=")
	testEVT1 := *testTX1.Events
	db.CreateEvent(testEVT1[0])

	if got1, _ := db.client.Events.Query().Where(
		events.ObjectIDEQ("0x9e4784cd1990ad24a703e7805b62ed83f76cd43f")).Exist(context.Background()); !got1 {
		t.Errorf("Result was incorrect, got %t, want %t.", got1, true)
	}

	testTX2, _ := sc.GetTransaction("VCXm2TQzXosGJjeEJU19gBK7hVFYFLWRtyb4eR7gT6s=")
	testEVT2 := *testTX2.Events
	db.CreateEvent(testEVT2[0])

	if got2, _ := db.client.Events.Query().Where(
		events.SenderEQ("0x6b5af19a5686938f7ec7e72a660a4c2b9ccd18b1")).Exist(context.Background()); !got2 {
		t.Errorf("Result was incorrect, got %t, want %t.", got2, true)
	}
}

func TestCreateAccount(t *testing.T) {
	p := postgres.Preset(
		postgres.WithUser("gnomock", "gnomick"),
		postgres.WithDatabase("mydb"),
	)

	container, _ := gnomock.Start(p)
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
		container.Host, container.DefaultPort(),
		"gnomock", "gnomick", "mydb",
	)

	sc := new(sui.SUIClient)
	sc.Init("http://127.0.0.1:9000")

	db := new(EntClient)
	db.Init("postgres", connStr)

	testACC1, _ := sc.GetAccount("0x5f7b658e9efbdbd0580e676fc14b72bb86f7fcc6")
	db.CreateAccount(testACC1, 0)

	if got1, _ := db.client.Accounts.Query().Where(
		accounts.AccountID("0x5f7b658e9efbdbd0580e676fc14b72bb86f7fcc6")).Exist(context.Background()); !got1 {
		t.Errorf("Result was incorrect, got %t, want %t.", got1, true)
	}

	testACC2, _ := sc.GetAccount("0x1049ad810645c762786f8144d3ca8740c0f851cf")
	db.CreateAccount(testACC2, 0)

	if got2, _ := db.client.Accounts.Query().Where(
		accounts.BalanceEQ(250000)).Exist(context.Background()); !got2 {
		t.Errorf("Result was incorrect, got %t, want %t.", got2, true)
	}
}

func TestCreateArgument(t *testing.T) {
	p := postgres.Preset(
		postgres.WithUser("gnomock", "gnomick"),
		postgres.WithDatabase("mydb"),
	)

	container, _ := gnomock.Start(p)
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
		container.Host, container.DefaultPort(),
		"gnomock", "gnomick", "mydb",
	)

	sc := new(sui.SUIClient)
	sc.Init("http://127.0.0.1:9000")

	db := new(EntClient)
	db.Init("postgres", connStr)

	testTX1, _ := sc.GetTransaction("LIJIF6jmAtp1XqnKstBrFXh4otlGdBnFnrIuc7YHQ/Y=")
	testArg1 := *testTX1.Arguments
	for _, k := range testArg1 {
		db.CreateArgument(k)
	}

	if got1, _ := db.client.Arguments.Query().Where(
		arguments.DataEQ("0xb59f06c24c8656693632df77724e326a8365763f")).Exist(context.Background()); !got1 {
		t.Errorf("Result was incorrect, got %t, want %t.", got1, true)
	}

	testTX2, _ := sc.GetTransaction("dNOK2AVQrXkE6aiFghswKAI9qDJOdOaI0NzU+jKtIJw=")
	testArg2 := *testTX2.Arguments
	for _, k := range testArg2 {
		db.CreateArgument(k)
	}

	if got2, _ := db.client.Arguments.Query().Where(
		arguments.NameEQ("Coin")).Exist(context.Background()); !got2 {
		t.Errorf("Result was incorrect, got %t, want %t.", got2, true)
	}
}

func TestCreateNFT(t *testing.T) {
	p := postgres.Preset(
		postgres.WithUser("gnomock", "gnomick"),
		postgres.WithDatabase("mydb"),
	)

	container, _ := gnomock.Start(p)
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
		container.Host, container.DefaultPort(),
		"gnomock", "gnomick", "mydb",
	)

	sc := new(sui.SUIClient)
	sc.Init("http://127.0.0.1:9000")

	db := new(EntClient)
	db.Init("postgres", connStr)

	testACC1, _ := sc.GetAccount("0x5f7b658e9efbdbd0580e676fc14b72bb86f7fcc6")
	testNFT1 := testACC1.GetAccountNFTs()
	for _, k := range testNFT1 {
		db.CreateNFT(k, 0)
	}

	if got1, _ := db.client.NFTs.Query().Where(
		nfts.ObjectIDEQ("0xeb1f1d1980f6f2d63bd0c994c82fc132e204954d")).Exist(context.Background()); !got1 {
		t.Errorf("Result was incorrect, got %t, want %t.", got1, true)
	}

	testACC2, _ := sc.GetAccount("0xcfd56e539d7bf7675e3c21215a1156bb23aab042")
	testNFT2 := testACC2.GetAccountNFTs()
	for _, k := range testNFT2 {
		db.CreateNFT(k, 0)
	}

	if got2, _ := db.client.NFTs.Query().Where(
		nfts.TypeEQ("0x2::devnet_nft::DevNetNFT")).Exist(context.Background()); !got2 {
		t.Errorf("Result was incorrect, got %t, want %t.", got2, false)
	}
}

func TestCreateObject(t *testing.T) {
	p := postgres.Preset(
		postgres.WithUser("gnomock", "gnomick"),
		postgres.WithDatabase("mydb"),
	)

	container, _ := gnomock.Start(p)
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
		container.Host, container.DefaultPort(),
		"gnomock", "gnomick", "mydb",
	)

	sc := new(sui.SUIClient)
	sc.Init("http://127.0.0.1:9000")

	db := new(EntClient)
	db.Init("postgres", connStr)

	testOBJ1, _ := sc.GetObject("0xfd4da6455ea26fde9f71596015ee50d35bb147e9")
	db.CreateObject(testOBJ1, 0)

	if got1, _ := db.client.Objects.Query().Where(
		objects.DataTypeEQ("moveObject")).Exist(context.Background()); !got1 {
		t.Errorf("Result was incorrect, got %t, want %t.", got1, true)
	}

	testOBJ2, _ := sc.GetObject("0xa1733e5df5241daf7d8ff32d29b6e2b77b6db90e")
	db.CreateObject(testOBJ2, 0)

	if got2, _ := db.client.Objects.Query().Where(
		objects.OwnerEQ("0xcfd56e539d7bf7675e3c21215a1156bb23aab042")).Exist(context.Background()); !got2 {
		t.Errorf("Result was incorrect, got %t, want %t.", got2, true)
	}
}

func TestCreatePackage(t *testing.T) {
	p := postgres.Preset(
		postgres.WithUser("gnomock", "gnomick"),
		postgres.WithDatabase("mydb"),
	)

	container, _ := gnomock.Start(p)
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
		container.Host, container.DefaultPort(),
		"gnomock", "gnomick", "mydb",
	)

	sc := new(sui.SUIClient)
	sc.Init("http://127.0.0.1:9000")

	db := new(EntClient)
	db.Init("postgres", connStr)

	testTX1, _ := sc.GetTransaction("t5bCqgkQxB2//+JVcawxpz3Z0WSLFlXFkBEQh66BNqc=")
	testPKG1, _ := testTX1.GetContractDeploy()
	db.CreatePackage(testPKG1)

	if got1, _ := db.client.Packages.Query().Where(
		packages.ObjectIDEQ("0x88c3ef1ede377d7a010ea10738ccb1b77766666b")).Exist(context.Background()); !got1 {
		t.Errorf("Result was incorrect, got %t, want %t.", got1, true)
	}

	testTX2, _ := sc.GetTransaction("Zx6fcsPsTs4wAF9IJQfje6ioSRBPUPcccu+CtwmC2ws=")
	testPKG2, _ := testTX2.GetContractDeploy()
	db.CreatePackage(testPKG2)

	if got2, _ := db.client.Packages.Query().Where(
		packages.ObjectIDEQ("0x0dd10e37cf1dd93d538b607adcbddd1ed20472d6")).Exist(context.Background()); !got2 {
		t.Errorf("Result was incorrect, got %t, want %t.", got2, true)
	}
}
