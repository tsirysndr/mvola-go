package main

import (
	"fmt"
	"log"
	"os"
	"time"

	uuid "github.com/google/uuid"
	mvola "github.com/tsirysndr/mvola-go"
)

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
	fmt.Println(res)

	mvola.SetAccessToken(client, res.AccessToken)

	correlationID := uuid.NewString()
	transactionRef := uuid.NewString()
	fmt.Println("correlationID:", correlationID)
	fmt.Println("transactionRef:", transactionRef)

	opt := mvola.Options{
		Version:               "1.0",
		CorrelationID:         correlationID,
		UserLanguage:          "FR",
		UserAccountIdentifier: "msisdn;0340017983",
		PartnerName:           "mvola",
		CallbackURL:           nil,
	}
	mvola.SetOptions(client, opt)

	fmt.Println(time.Now().Format(time.RFC3339))

	tx := mvola.TransactionRequest{
		Amount:          "100000",
		Currency:        "Ar",
		DescriptionText: "test",
		RequestingOrganisationTransactionReference: "xxx",
		RequestDate:                  time.Now().Format(time.RFC3339),
		OriginalTransactionReference: transactionRef,
		DebitParty: []mvola.KeyValue{{
			Key:   "msisdn",
			Value: "0344289931",
		}},
		CreditParty: []mvola.KeyValue{{
			Key:   "msisdn",
			Value: "0340017983",
		}},
		Metadata: []mvola.KeyValue{{
			Key:   "partnerName",
			Value: "partner test",
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
