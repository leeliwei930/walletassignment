package models

import "github.com/google/uuid"

type WalletStatusParams struct {
	UserID uuid.UUID
}

type WalletDepositParams struct {
	UserID uuid.UUID
	Amount int
}

type WalletWithdrawalParams struct {
	UserID uuid.UUID
	Amount int
}

type WalletTransferParams struct {
	RecipientUserID        uuid.UUID
	RecipientReferenceNote string
	Amount                 int
}

type WalletTransactionsParams struct {
	UserID uuid.UUID
	Page   int
	Limit  int
}
