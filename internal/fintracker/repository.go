package fintracker

// UserRepository describes the ability to store users.
type UserRepository interface {
	AddUser(User) (string, error)
	FetchUser(string) (User, error)
}

// TransactionRepository describes the ability to store transactions.
type TransactionRepository interface {
	AddTransaction(Transaction) (string, error)
	FetchTransaction(string) (Transaction, error)
}
