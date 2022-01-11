package transaction

type Domain struct {
	ID                  uint
	UserID              uint
	DetailTransactionID uint
	Total               int
	Link                string
	TransactionStatus   string
	FraudStatus         string
	PaymentType         string
	Provider            string
}

type DetailDomain struct {
	ID        uint
	ProductID uint
	Discount  int
	Subtotal  int
}

type CreateVA struct {
	ProductID uint
	Discount  int
	Tax       int
	Subtotal  int
	Bank      string
}

type Notification struct {
	TransactionStatus string
	OrderID           string
	PaymentType       string
	FraudStatus       string
}

type EmailDriver struct {
	Sender  string
	ToEmail string
	Subject string
}

type HistoryDomain struct {
	ID                uint
	ProductID         uint
	ProductName       string
	Discount          int
	Tax               int
	Total             int
	Link              string
	TransactionStatus string
	FraudStatus       string
	PaymentType       string
	Provider          string
}

type ITransactionUsecase interface {
	GetVirtualAccount(id int, inVA CreateVA) (string, error)
	ProcessNotification(input Notification) error
	GetAllTxUser(id int) ([]HistoryDomain, error)
	GetTxByID(id int) (HistoryDomain, error)
}

type ITransactionRepository interface {
	CreateDetail(det DetailDomain) (uint, error)
	CreateTx(tx Domain) error
	GetTxByID(id int) (Domain, error)
	UpdateTx(tx Domain) error
	GetUserEmail(id int) (string, string)
	GetUserTxByID(int) ([]HistoryDomain, error)
	GetNameNTax(id int) (string, int)
	GetTxHistoryByID(id int) (HistoryDomain, error)
}
