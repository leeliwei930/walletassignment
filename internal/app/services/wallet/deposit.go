package wallet

import (
	"context"

	"github.com/leeliwei930/walletassignment/internal/app/models"
)

func (s *walletService) Deposit(ctx context.Context, params models.WalletDepositParams) (*models.WalletDeposit, error) {
	return &models.WalletDeposit{}, nil
}
