package interfaces

import (
	"context"

	"github.com/leeliwei930/walletassignment/internal/app/models"
)

type WalletService interface {
	Status(ctx context.Context, params models.WalletStatusParams) (*models.WalletStatus, error)
	Deposit(ctx context.Context, params models.WalletDepositParams) (*models.WalletDeposit, error)
	Withdraw(ctx context.Context, params models.WalletWithdrawalParams) (*models.WalletWithdrawal, error)
	Transfer(ctx context.Context, params models.WalletTransferParams) (*models.WalletTransfer, error)
	Transactions(ctx context.Context, params models.WalletTransactionsParams) (*models.WalletTransactions, error)
}
