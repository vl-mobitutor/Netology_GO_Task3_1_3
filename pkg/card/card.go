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


//Функция добавления новой транзакции в историю операций по карте
func AddTransaction(card *Card, transaction *Transaction) {
	card.Transactions = append(card.Transactions, *transaction)
}


//Функция суммирования операций по набору кодов MCC
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


//Функция определения категории точки POS  по категории МСС
const mccFindError = "Категория не указана"

func TranslateMCC(code string) string {
	// Справочник MCC-кодов
	mcc := map[string]string{
		"5411": "Супермаркеты",
		"6539": "Фондирование операций",
		"5812": "Рестораны и точки питания",
		"5942": "Книжные магазины",
		"5993": "Магазины сигар",
	}
	value, ok := mcc[code]
	if ok {
		return value
	} else {
		return mccFindError
	}

}