package wallet

import (
	"context"

	"github.com/leeliwei930/walletassignment/internal/app/models"
)

func (s *walletService) Withdraw(ctx context.Context, params models.WalletWithdrawalParams) (*models.WalletWithdrawal, error) {
	return &models.WalletWithdrawal{}, nil
}
