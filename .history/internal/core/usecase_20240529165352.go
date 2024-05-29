package core

import (
	"context"

	ltp "github.com/devpablocristo/dive-challenge/internal/core/ltp"
)

type UseCase struct {
	ltp ltp.RepositoryPort
	acl ltp.APIClientPort
}

func NewUseCase(r ltp.RepositoryPort, a ltp.APIClientPort) UseCasePort {
	return &UseCase{
		ltp: r,
		acl: a,
	}
}

func (u *UseCase) GetLTP(ctx context.Context, pairs []string) ([]ltp.LTP, error) {
	// pairs := []string{}
	// pairs = append(pairs, "BTC/CHF", "BTC/EUR", "BTC/USD")
	lp, err := u.acl.GetKrakenLTPs(ctx, pairs)
	if err != nil {
		return nil, err
	}

	return lp, nil
}
