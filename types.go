package mvola

type TransactionRequest struct {
	Amount                                     string     `json:"amount"`
	Currency                                   string     `json:"currency"`
	DescriptionText                            string     `json:"descriptionText"`
	RequestDate                                string     `json:"requestDate"`
	DebitParty                                 []KeyValue `json:"debitParty"`
	CreditParty                                []KeyValue `json:"creditParty"`
	Metadata                                   []KeyValue `json:"metadata"`
	RequestingOrganisationTransactionReference string     `json:"requestingOrganisationTransactionReference"`
	OriginalTransactionReference               string     `json:"originalTransactionReference"`
}

type TransactionResponse struct {
	Status              string `json:"status"`
	ServerCorrelationID string `json:"serverCorrelationId"`
	NotificationMethod  string `json:"notificationMethod"`
}

type TransactionDetails struct {
	Amount               string     `json:"amount"`
	Currency             string     `json:"currency"`
	TransactionReference string     `json:"transactionReference"`
	TransactionStatus    string     `json:"transactionStatus"`
	CreationDate         string     `json:"creationDate"`
	RequestDate          string     `json:"requestDate"`
	DebitParty           []KeyValue `json:"debitParty"`
	CreditParty          []KeyValue `json:"creditParty"`
	Metadata             []KeyValue `json:"metadata"`
	Fees                 []Fee      `json:"fees"`
}

type Fee struct {
	FeeAmount string `json:"feeAmount"`
}

type TransactionStatus struct {
	Status              string `json:"status"`
	ServerCorrelationID string `json:"serverCorrelationId"`
	NotificationMethod  string `json:"notificationMethod"`
	ObjectReference     string `json:"objectReference"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Options struct {
	Version               string
	CorrelationID         string
	UserLanguage          string
	UserAccountIdentifier string
	PartnerName           string
	CallbackURL           *string
}
