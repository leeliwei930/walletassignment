package wallet

import (
	"context"

	"github.com/leeliwei930/walletassignment/internal/app/models"
)

func (s *walletService) Status(ctx context.Context, params models.WalletStatusParams) (*models.WalletStatus, error) {
	return &models.WalletStatus{}, nil
}
