package wallet

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/leeliwei930/walletassignment/ent/wallet"
	"github.com/leeliwei930/walletassignment/internal/app/models"
	"github.com/leeliwei930/walletassignment/pkg/formatter"
)

func (s *walletService) Status(ctx context.Context, params models.WalletStatusParams) (*models.WalletStatus, error) {

	entClient := s.app.GetEnt()

	wallet, err := entClient.Wallet.Query().
		Where(wallet.UserID(params.UserID)).
		ForShare(
			sql.WithLockTables(wallet.Table),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	formmatedBalance := formatter.FormatCurrencyAmount(
		wallet.Balance,
		wallet.CurrencyCode,
		wallet.DecimalPlaces,
	)

	return &models.WalletStatus{
		ID:               wallet.ID,
		Balance:          wallet.Balance,
		Currency:         wallet.CurrencyCode,
		FormattedBalance: formmatedBalance,
	}, nil
}
