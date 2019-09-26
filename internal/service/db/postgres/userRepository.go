package postgres

type User struct {
	Id          int
	Login       string
	MoneyAmount int
	CardNumber  string
	Status      bool
}
