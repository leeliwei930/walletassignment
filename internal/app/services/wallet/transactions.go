package wallet

import (
	"context"

	"github.com/leeliwei930/walletassignment/internal/app/models"
)

func (s *walletService) Transactions(ctx context.Context, params models.WalletTransactionsParams) (*models.WalletTransactions, error) {
	return &models.WalletTransactions{}, nil
}
