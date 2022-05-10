<h1>mvola-go</h1>
<p>
	<a href="https://app.travis-ci.com/github/tsirysndr/mvola-go" target="_blank">
    <img src="https://app.travis-ci.com/tsirysndr/mvola-go.svg?branch=master" />
  </a>
  <a href="https://codecov.io/gh/tsirysndr/mvola-go" target="_blank">
    <img src="https://codecov.io/gh/tsirysndr/mvola-go/branch/master/graph/badge.svg?token=" />
  </a>
  <a href="#" target="_blank">
		<a href="https://pkg.go.dev/github.com/tsirysndr/mvola-go" target="_blank">
			<img alt="Go Reference" src="https://pkg.go.dev/badge/github.com/tsirysndr/mvola-go" />
		</a>
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-green.svg" />
  </a>
</p>

[MVola](https://www.mvola.mg/devportal) Go client library.

## Install

```sh
  go get -u github.com/tsirysndr/mvola-go
```

## Usage

```go
import (
	"fmt"
	"log"
	"os"
	"time"

	uuid "github.com/google/uuid"
	mvola "github.com/tsirysndr/mvola-go"
)

const ISO8601 = "2006-01-02T15:04:05.999Z"

func main() {
	var (
		consumerKey    = os.Getenv("CONSUMER_KEY")
		consumerSecret = os.Getenv("CONSUMER_SECRET")
	)
	client := mvola.NewClient(mvola.SANDBOX_URL)
	res, err := client.Auth.GenerateToken(consumerKey, consumerSecret)
	if err != nil {
		log.Fatal(err)
	}

	mvola.SetAccessToken(client, res.AccessToken)

	correlationID := uuid.NewString()
	transactionRef := uuid.NewString()
	fmt.Println("correlationID:", correlationID)
	fmt.Println("transactionRef:", transactionRef)

	opt := mvola.Options{
		Version:               "1.0",
		CorrelationID:         correlationID,
		UserLanguage:          "FR",
		UserAccountIdentifier: "msisdn;0343500003",
		PartnerName:           "TestMVola",
		CallbackURL:           nil,
	}
	mvola.SetOptions(client, opt)

	requestDate := time.Now().UTC().Format(ISO8601)

	tx := mvola.TransactionRequest{
		Amount:          "1000",
		Currency:        "Ar",
		DescriptionText: "test",
		RequestingOrganisationTransactionReference: transactionRef,
		RequestDate:                  requestDate,
		OriginalTransactionReference: transactionRef,
		DebitParty: []mvola.KeyValue{{
			Key:   "msisdn",
			Value: "0343500003",
		}},
		CreditParty: []mvola.KeyValue{{
			Key:   "msisdn",
			Value: "0343500004",
		}},
		Metadata: []mvola.KeyValue{{
			Key:   "partnerName",
			Value: "TestMVola",
		},
			{
				Key:   "fc",
				Value: "USD",
			}, {
				Key:   "amountFc",
				Value: "1",
			}},
	}

	txres, err := client.Transaction.SendPayment(&tx)
	fmt.Println(txres)
	fmt.Println(err)
}

```

## Documentation

For details on all the functionality in this library, see the [Go
documentation](https://pkg.go.dev/github.com/tsirysndr/mvola-go).

## Author

👤 **Tsiry Sandratraina <tsiry.sndr@aol.com>**

* Twitter: [@tsiry_sndr](https://twitter.com/tsiry_sndr)
* Github: [@tsirysndr](https://github.com/tsirysndr)

## Show your support

Give a ⭐️ if this project helped you!
