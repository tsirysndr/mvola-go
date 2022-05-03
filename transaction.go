package mvola

type TransactionService service

func (s *TransactionService) GetTransactionDetails(transactionID string) {

}

func (s *TransactionService) GetTransactionStatus(serverCorrelationID string) {

}

func (s *TransactionService) SendPayment(tx *TransactionRequest) (*TransactionResponse, error) {
	var err error
	res := new(TransactionResponse)
	endpoint := "/mvola/mm/transactions/type/merchantpay/1.0.0/"
	s.client.base.Post(endpoint).BodyJSON(tx).Receive(res, err)

	return res, err
}
