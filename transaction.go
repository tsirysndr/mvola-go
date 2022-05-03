package mvola

import (
	"fmt"
	"log"
	"net/http/httputil"
)

type TransactionService service

func (s *TransactionService) GetTransactionDetails(transactionID string) {

}

func (s *TransactionService) GetTransactionStatus(serverCorrelationID string) {

}

func (s *TransactionService) SendPayment(tx *TransactionRequest) (*TransactionResponse, error) {
	var err error
	res := new(TransactionResponse)
	endpoint := "/mvola/mm/transactions/type/merchantpay/1.0.0/"
	hres, herr := s.client.base.Post(endpoint).BodyJSON(tx).Receive(res, err)
	fmt.Println(hres)
	fmt.Println(herr)

	respDump, err := httputil.DumpResponse(hres, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("RESPONSE:\n%s", string(respDump))

	return res, err
}
