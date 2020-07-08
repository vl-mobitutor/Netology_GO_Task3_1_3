package main

import (
	"fmt"
	"github.com/vl-mobitutor/Netology_GO_Task3_1_3/pkg/card"
	"time"
)

func main() {

	//ИНИЦИАЛИЗАЦИЯ ДАННЫХ КАРТЫ И ПЕРВЫХ ТРАНЗАКЦИЙ

	//Установка срока действия карты в формате UTC
	dateLayout:= "02/01/2006"
	cardExpiredDate, _ := time.Parse(dateLayout, "01/01/2022")

	//Установка даты и времени транзакции в формате UTC
	dateTimeLayout := "01/02/2006 3:04:05 PM"
	transactionDateTime, _ := time.Parse(dateTimeLayout, "06/05/2020 5:09:01 PM")

	myCard := &card.Card{
		Id:1,
		Number: "1111 2222 3333 4444",
		ExpiredDate: cardExpiredDate,
		CvCode: "001",
		PaymentSystem: "Master",
		BankIssuer: "SuperBank",
		CardholderName: "Ilon Mask",
		Balance: 100_000_00,
		Currency: "RUR",
		Transactions: []card.Transaction{	//начало массива транзакций
			{
				Id: 1001,
				Amount:  -735_55,
				Currency: "RUR",
				DateTime: transactionDateTime,
				Description: "Gipermarket",
				Status:  "Операция в обработке",
				MccCode: "5411",
			},

			{
				Id: 1002,
				Amount: 2000_00,
				Currency: "RUR",
				DateTime: transactionDateTime,
				Description: "Пополнение через Сбербанк",
				Status:  "Обработана",
				MccCode: "6539",
			},

			{
				Id: 1003,
				Amount: -1203_91,
				Currency: "RUR",
				DateTime: transactionDateTime,
				Description: "Gastrobar",
				Status:  "Обработана",
				MccCode: "5812",
			},

		},	//конец массива транзакций
	}

	fmt.Println(myCard)
	fmt.Println("----------------------------------------------------------------------")



	//ДОБАВЛЕНИЕ ТРАНЗАКЦИЙ ЧЕРЕЗ ФУНКЦИЮ

	transactionDateTime, _ = time.Parse(dateTimeLayout, "07/07/2020 2:02:08 PM")
	newTransaction := &card.Transaction{
		Id: 1004,
		Amount:  -1000_00,
		Currency: "RUR",
		DateTime: transactionDateTime,
		Description: "Gipermarket",
		Status:  "Операция в обработке",
		MccCode: "5411",
	}
	card.AddTransaction(myCard, newTransaction)

	newTransaction = &card.Transaction{
		Id: 1005,
		Amount:  -2000_00,
		Currency: "RUR",
		DateTime: transactionDateTime,
		Description: "Gipermarket",
		Status:  "Операция в обработке",
		MccCode: "5400",
	}
	card.AddTransaction(myCard, newTransaction)

	fmt.Println(myCard)
	fmt.Println("----------------------------------------------------------------------")

	//СУММИРОВАНИЕ ПО КОДУ МСС
	selectCodes := []string{"5411"}
	totalSum :=card.SumByMCC(myCard.Transactions, selectCodes)
	fmt.Printf("Общая сумма операций по кодам МСС: %s - составила %d копеек \n", selectCodes,totalSum)


	//Демо функции - определение категории торговой точки по MСС-коду транзакции
	transactionNumber := 4
	category := card.TranslateMCC(myCard.Transactions[transactionNumber].MccCode) //
	fmt.Printf("Транзакция № %d имеет категорию торговой точки: %s \n", transactionNumber, category)

}
