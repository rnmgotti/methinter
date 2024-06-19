package internal

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