package main

type Transaction struct {
	from   string
	to     string
	amount int
}

func NewTransaction(from string, to string, amount int) Transaction {
	return Transaction{
		from:   from,
		to:     to,
		amount: amount,
	}
}
