package wallet

import (
	"context"

	"github.com/leeliwei930/walletassignment/internal/app/models"
)

func (s *walletService) Transfer(ctx context.Context, params models.WalletTransferParams) (*models.WalletTransfer, error) {
	return &models.WalletTransfer{}, nil
}
