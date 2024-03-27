package svc

import "github.com/ramadhan1445sprint/sprint_segokuning/repo"

type BalanceSvc interface {
}

type balanceSvc struct {
	repo repo.BalanceRepo
}

func NewBalanceSvc(r repo.BalanceRepo) BalanceSvc {
	return &balanceSvc{
		repo: r,
	}
}
