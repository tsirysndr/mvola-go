package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
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

	mvola.SetAccessToken(client, res.AccessToken)

	correlationID := uuid.NewString()
	opt := mvola.Options{
		Version:               "1.0",
		CorrelationID:         correlationID,
		UserLanguage:          "FR",
		UserAccountIdentifier: "msisdn;0343500003",
		PartnerName:           "TestMVola",
	}
	mvola.SetOptions(client, opt)

	txstatus, err := client.Transaction.GetTransactionStatus("2ba1d66a-25cf-4c12-8a6f-4cb01255148e")
	fmt.Println(txstatus)
	fmt.Println(err)
}
