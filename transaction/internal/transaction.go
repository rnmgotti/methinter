package internal

import (
	"log"
	"time"
)

type Transaction interface {
	GetID() int
	GetStatus() string
	GetTime() int64
}
type TransactionStr struct {
	ID     int
	Status string
	Time   int64
}

func (t TransactionStr) GetID() int {
	return t.ID

}

func (t TransactionStr) GetStatus() string {
	return t.Status
}

func (t TransactionStr) GetTime() int64 {
	return int64(t.Time)
}

func CheckLastTransaction(lasttrans Transaction) {
	if lasttrans.GetStatus() == "failed" {
		log.Println("Последняя транзакции ошибочна, id:", lasttrans.GetID())
	} else {
		log.Println("Последняя транзакция успешна.")
	}
}

func MainTransaction() {
	InputTransaction := TransactionStr{1, "comp", (time.Now().Unix())}
	CheckLastTransaction(InputTransaction)
}
