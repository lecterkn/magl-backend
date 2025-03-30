package port

import "context"

type TransactionProvider interface {
	Transact(func(context.Context) error) error
}
