package mvola

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTransactionService_GetTransactionDetails(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/mvola/mm/transactions/type/merchantpay/1.0.0/636085941", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"amount": "10000.00",
			"currency": "Ar",
			"transactionReference": "636085941",
			"transactionStatus": "completed",
			"creationDate": "2022-05-03T18:09:10.391Z",
			"requestDate": "2022-05-10T08:15:39.513Z",
			"debitParty": [
				{
					"key": "msisdn",
					"value": "0343500003"
				}
			],
			"creditParty": [
				{
					"key": "msisdn",
					"value": "0343500004"
				}
			],
			"metadata": [
				{
					"key": "originalTransactionResult",
					"value": "0"
				},
				{
					"key": "originalTransactionResultDesc",
					"value": "0"
				}
			],
			"fees": [
				{
					"feeAmount": "150"
				}
			]
		}
		`)
	})

	client := NewClient(server.URL)
	client.base.Client(httpClient)

	SetAccessToken(client, accessToken)

	correlationID := uuid.NewString()
	opt := Options{
		Version:               "1.0",
		CorrelationID:         correlationID,
		UserAccountIdentifier: "msisdn;0343500003",
	}
	SetOptions(client, opt)

	response, err := client.Transaction.GetTransactionDetails("636085941")

	expected := &TransactionDetails{
		Amount:               "10000.00",
		Currency:             "Ar",
		TransactionReference: "636085941",
		TransactionStatus:    "completed",
		CreationDate:         "2022-05-03T18:09:10.391Z",
		RequestDate:          "2022-05-10T08:15:39.513Z",
		DebitParty: []KeyValue{
			{
				Key:   "msisdn",
				Value: "0343500003",
			},
		},
		CreditParty: []KeyValue{
			{
				Key:   "msisdn",
				Value: "0343500004",
			},
		},
		Metadata: []KeyValue{
			{
				Key:   "originalTransactionResult",
				Value: "0",
			},
			{
				Key:   "originalTransactionResultDesc",
				Value: "0",
			},
		},
		Fees: []Fee{
			{
				FeeAmount: "150",
			},
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, response)
}

func TestTransactionService_GetTransactionStatus(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/mvola/mm/transactions/type/merchantpay/1.0.0/status/ec2320b5-3167-4cf0-b2a1-d77e04902360", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"status": "completed",
			"serverCorrelationId": "2ba1d66a-25cf-4c12-8a6f-4cb01255148e",
			"notificationMethod": "polling",
			"objectReference": "636042511"
		}`)
	})

	client := NewClient(server.URL)
	client.base.Client(httpClient)

	SetAccessToken(client, accessToken)

	correlationID := uuid.NewString()
	opt := Options{
		Version:               "1.0",
		CorrelationID:         correlationID,
		UserLanguage:          "FR",
		UserAccountIdentifier: "msisdn;0343500003",
		PartnerName:           "TestMVola",
	}
	SetOptions(client, opt)

	response, err := client.Transaction.GetTransactionStatus("ec2320b5-3167-4cf0-b2a1-d77e04902360")

	expected := &TransactionStatus{
		Status:              "completed",
		ServerCorrelationID: "2ba1d66a-25cf-4c12-8a6f-4cb01255148e",
		NotificationMethod:  "polling",
		ObjectReference:     "636042511",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, response)
}

func TestTransactionService_SendPayment(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/mvola/mm/transactions/type/merchantpay/1.0.0/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"status": "pending",
			"serverCorrelationId": "a6b5569b-6181-4fc9-bee3-b9f928dd7ae3",
			"notificationMethod": "polling"
		}`)
	})

	client := NewClient(server.URL)
	client.base.Client(httpClient)

	SetAccessToken(client, accessToken)

	const ISO8601 = "2006-01-02T15:04:05.999Z"

	correlationID := uuid.NewString()
	transactionRef := uuid.NewString()

	opt := Options{
		Version:               "1.0",
		CorrelationID:         correlationID,
		UserLanguage:          "FR",
		UserAccountIdentifier: "msisdn;0343500003",
		PartnerName:           "TestMVola",
		CallbackURL:           nil,
	}
	SetOptions(client, opt)

	requestDate := time.Now().UTC().Format(ISO8601)

	tx := TransactionRequest{
		Amount:          "1000",
		Currency:        "Ar",
		DescriptionText: "test",
		RequestingOrganisationTransactionReference: transactionRef,
		RequestDate:                  requestDate,
		OriginalTransactionReference: transactionRef,
		DebitParty: []KeyValue{{
			Key:   "msisdn",
			Value: "0343500003",
		}},
		CreditParty: []KeyValue{{
			Key:   "msisdn",
			Value: "0343500004",
		}},
		Metadata: []KeyValue{{
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

	response, err := client.Transaction.SendPayment(&tx)

	expected := &TransactionResponse{
		Status:              "pending",
		ServerCorrelationID: "a6b5569b-6181-4fc9-bee3-b9f928dd7ae3",
		NotificationMethod:  "polling",
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, response)
}
