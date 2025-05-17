package wallet

import (
	"context"

	"github.com/leeliwei930/walletassignment/internal/app/models"
)

type TransferRequest struct {
	Amount               int    `json:"amount" validate:"required,min=1,max=1000000" localeKey:"wallet::transfer::amount"`
	RecipientPhoneNumber string `json:"recipient_phone_number" validate:"required" localeKey:"wallet::transfer::recipient_phone_number"`
}

type TransferResponse struct {
	Wallet      models.WalletStatus      `json:"wallet"`
	Transaction models.WalletTransaction `json:"transaction"`
}

func (s *walletService) Transfer(ctx context.Context, params models.WalletTransferParams) (*models.WalletTransfer, error) {
	return &models.WalletTransfer{}, nil
}
