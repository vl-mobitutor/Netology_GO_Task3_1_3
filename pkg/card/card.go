package card

type Transaction struct {
	Id int64
	Amount int64 //сумма операции
	Currency string //валюта операции
	DateTime uint64 //дата в формате UTC
	AuthorizationCode string //код авторизации  - уникальный код подтверждения совершения операции
	MccCode string //код точки продаж (POS-терминал) по каталогу Merchant Category Code
	Status string //текущий статус операции
}

type Card struct {
	Id int64
	Number string //номер карты
	ExpiredDate uint64 //срок действия в формате UTC
	CvCode string //CV-код
	PaymentSystem string //платежная система
	BankIssuer string //банк-эмитент карты
	CardholderName string //имя и фамилия держателя карты
	Balance int64 //доступный баланс карты
	Currency string //валюта карточного счета
	Transactions []Transaction
}

