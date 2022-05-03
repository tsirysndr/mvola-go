package mvola

type TransactionService service

func (s *TransactionService) GetTransactionDetails(transactionID string) (*TransactionDetails, error) {
	var err error
	res := new(TransactionDetails)
	endpoint := "/mvola/mm/transactions/type/merchantpay/1.0.0/" + transactionID
	s.client.base.Get(endpoint).Receive(res, err)
	return res, err
}

func (s *TransactionService) GetTransactionStatus(serverCorrelationID string) (*TransactionStatus, error) {
	var err error
	res := new(TransactionStatus)
	endpoint := "/mvola/mm/transactions/type/merchantpay/1.0.0/status/" + serverCorrelationID
	s.client.base.Get(endpoint).Receive(res, err)
	return res, err
}

func (s *TransactionService) SendPayment(tx *TransactionRequest) (*TransactionResponse, error) {
	var err error
	res := new(TransactionResponse)
	endpoint := "/mvola/mm/transactions/type/merchantpay/1.0.0/"
	s.client.base.Post(endpoint).BodyJSON(tx).Receive(res, err)
	return res, err
}
