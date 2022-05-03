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
		UserAccountIdentifier: "msisdn;0343500003",
	}
	mvola.SetOptions(client, opt)

	txdetails, err := client.Transaction.GetTransactionDetails("636042511")
	fmt.Println(txdetails)
	fmt.Println(err)
}
