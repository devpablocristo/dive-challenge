package ltp

import "context"

type RepositoryPort interface {
	GetLTP(context.Context, string) error
}

type APIClientPort interface {
	GetKrakenLTPs(context.Context, []string) ([]LTP, error)
}
