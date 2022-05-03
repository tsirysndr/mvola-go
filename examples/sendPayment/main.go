package main

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
		Amount:          "100000",
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
