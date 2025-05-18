package wallet

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/leeliwei930/walletassignment/ent/wallet"
	"github.com/leeliwei930/walletassignment/internal/app/models"
	apperrors "github.com/leeliwei930/walletassignment/internal/errors"
	"github.com/leeliwei930/walletassignment/pkg/formatter"
)

const (
	TRX_TYPE_WITHDRAW_DESCRIPTION_LOCALE_KEY = "wallet::withdraw::trx::description"
	TRX_TYPE_WITHDRAW                        = "withdrawal"
	// minimum withdrawal amount is 100 currency units
	MINIMUM_WITHDRAWAL_AMOUNT = 100
)

func (s *walletService) Withdraw(ctx context.Context, params models.WalletWithdrawalParams) (*models.WalletWithdrawal, error) {
	entClient := s.app.GetEnt()
	locale := s.app.GetLocale()
	ut := locale.GetTranslatorFromContext(ctx)

	tx, err := entClient.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	// Ensure transaction is rolled back on error
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	walletRec, err := tx.Wallet.
		Query().
		Where(wallet.UserID(params.UserID)).
		ForUpdate(sql.WithLockTables(wallet.Table)).
		WithUser().
		First(ctx)
	if err != nil {
		return nil, err
	}

	if walletRec.Balance < params.Amount {
		return nil, apperrors.InsuficcientBalanceWithdrawalErr
	}

	if params.Amount < MINIMUM_WITHDRAWAL_AMOUNT {
		formattedAmount := formatter.FormatCurrencyAmount(MINIMUM_WITHDRAWAL_AMOUNT, walletRec.CurrencyCode, walletRec.DecimalPlaces)
		return nil, apperrors.MinimumWithdrawalAmountRequiredErr(ut, formattedAmount)
	}

	userRec := walletRec.Edges.User

	walletBalance := walletRec.Balance - params.Amount
	walletRec, err = walletRec.Update().
		SetBalance(walletBalance).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	systemUT := locale.GetUT().GetFallback()
	userSvc := s.app.GetUserService()
	userFullName := userSvc.GetFullName(ctx, userRec)

	description, _ := systemUT.T(TRX_TYPE_WITHDRAW_DESCRIPTION_LOCALE_KEY, userFullName)
	ledgerRec, err := tx.Ledger.Create().
		SetWalletID(walletRec.ID).
		SetAmount(params.Amount).
		SetDescription(description).
		SetTransactionType(TRX_TYPE_WITHDRAW).
		SetCreatedAt(walletRec.UpdatedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	formattedBalance := formatter.FormatCurrencyAmount(
		walletRec.Balance,
		walletRec.CurrencyCode,
		wallet.DefaultDecimalPlaces,
	)

	return &models.WalletWithdrawal{
		Wallet: models.WalletStatus{
			ID:               walletRec.ID,
			Currency:         walletRec.CurrencyCode,
			FormattedBalance: formattedBalance,
			Balance:          walletRec.Balance,
		},
		Transaction: models.WalletTransaction{
			ID:        ledgerRec.ID,
			Amount:    ledgerRec.Amount,
			Timestamp: ledgerRec.CreatedAt,
			Type:      ledgerRec.TransactionType,
		},
	}, nil
}
