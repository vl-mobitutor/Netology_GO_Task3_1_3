package card

import (
	"time"
)

type Transaction struct {
	Id int64 // id-карты
	Amount int64 //сумма операции в копейках/центах
	Currency string //валюта операции
	DateTime time.Time //дата и время операции в формате UTC
	Description string //описание операции
	Status string //текущий статус операции
	MccCode string //код точки продаж по каталогу Merchant Category Code
	MccType string //тип точки продаж
}

type Card struct {
	Id int64 // id-транзакции
	Number string //номер карты
	ExpiredDate time.Time //срок действия в формате UTC
	CvCode string //CV-код
	PaymentSystem string //платежная система
	BankIssuer string //банк-эмитент карты
	CardholderName string //имя и фамилия держателя карты
	Balance int64 //текущий баланс карты в копейках/центах
	Currency string //валюта карточного счета
	Transactions []Transaction
}

func AddTransaction(card *Card, transaction *Transaction) {
	card.Transactions = append(card.Transactions, *transaction)
}

func SumByMCC(transactions []Transaction, mcc []string) int64 {
	var totalSum int64 = 0

	for _, myTransaction := range transactions { //Перебор элементов транзакций

		for _, myMcc := range mcc { //Перебор элементов набора MCC
			if myTransaction.MccCode == myMcc {
				totalSum += myTransaction.Amount
			}
		}
	}
	return totalSum
}


//Функция выборки последних N транзакций
func LastNTransactions(card *Card, n int) []Transaction {
	var (
		lastTransactions []Transaction
		sliceLength int
	)

	sliceLength = len(card.Transactions)
	if n > sliceLength {
		n = sliceLength
	}

	for i := 1; i <= n; i++ {
		lastTransactions = append(lastTransactions, card.Transactions[sliceLength - i])
	}
	return lastTransactions
}