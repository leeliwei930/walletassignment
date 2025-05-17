package models

import (
	"time"

	"github.com/google/uuid"
)

type WalletStatus struct {
	ID               uuid.UUID `json:"id,omitempty"`
	Balance          int       `json:"balance"`
	Currency         string    `json:"currency"`
	FormattedBalance string    `json:"formattedBalance"`
}

type WalletTransaction struct {
	ID        uuid.UUID `json:"id"`
	Amount    int       `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type,omitempty"`
}

type WalletDeposit struct {
	Wallet      WalletStatus      `json:"wallet"`
	Transaction WalletTransaction `json:"transaction"`
}

type WalletTransactions struct {
	Transactions []*WalletTransaction `json:"transactions"`
	Pagination   *Pagination          `json:"pagination"`
}

type Pagination struct {
	CurrentPage int  `json:"currentPage"`
	TotalPages  int  `json:"totalPages"`
	TotalItems  int  `json:"totalItems"`
	Limit       int  `json:"limit"`
	HasNext     bool `json:"hasNext"`
	HasPrev     bool `json:"hasPrev"`
}

type WalletWithdrawal struct {
	Wallet      WalletStatus      `json:"wallet"`
	Transaction WalletTransaction `json:"transaction"`
}

type WalletTransfer struct {
	Wallet      WalletStatus      `json:"wallet"`
	Transaction WalletTransaction `json:"transaction"`
}
