package wallet

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/leeliwei930/walletassignment/ent/user"
	"github.com/leeliwei930/walletassignment/ent/wallet"
	"github.com/leeliwei930/walletassignment/internal/app/models"
	"github.com/leeliwei930/walletassignment/internal/errors"
	"github.com/leeliwei930/walletassignment/pkg/formatter"
)

const (
	TRX_TYPE_DEPOSIT                = "deposit"
	TRX_TYPE_DESCRIPTION_LOCALE_KEY = "wallet::deposit::trx::description"
	// minimum deposit amount is 100 currency units
	MINIMUM_DEPOSIT_AMOUNT = 100
)

func (s *walletService) Deposit(ctx context.Context, params models.WalletDepositParams) (*models.WalletDeposit, error) {

	entClient := s.app.GetEnt()
	locale := s.app.GetLocale()
	ut := locale.GetUT().GetFallback()
	userSvc := s.app.GetUserService()

	tx, err := entClient.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	userRec, err := tx.User.Query().Where(user.ID(params.UserID)).First(ctx)
	if err != nil {
		return nil, err
	}

	userFullName := userSvc.GetFullName(ctx, userRec)
	trxDescription, _ := ut.T(TRX_TYPE_DESCRIPTION_LOCALE_KEY, userFullName)

	// retrieve current wallet
	wallet, err := tx.Wallet.
		Query().
		Where(wallet.UserID(params.UserID)).
		ForUpdate(sql.WithLockTables(wallet.Table)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	// validate deposit amount
	if params.Amount < MINIMUM_DEPOSIT_AMOUNT {
		formattedAmount := formatter.FormatCurrencyAmount(MINIMUM_DEPOSIT_AMOUNT, wallet.CurrencyCode, wallet.DecimalPlaces)
		err = errors.MinimumDepositAmountRequiredErr(ut, formattedAmount)
		return nil, err
	}

	// update wallet balance
	newBalance := wallet.Balance + params.Amount
	walletRec, err := wallet.Update().SetBalance(newBalance).Save(ctx)
	if err != nil {
		return nil, err
	}

	// create transactione in ledger
	currentTime := time.Now()
	ledgerRec, err := tx.Ledger.Create().SetWalletID(wallet.ID).
		SetAmount(params.Amount).
		SetDescription(trxDescription).
		SetTransactionType(TRX_TYPE_DEPOSIT).
		SetCreatedAt(currentTime).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	formattedBalance := formatter.FormatCurrencyAmount(walletRec.Balance, walletRec.CurrencyCode, walletRec.DecimalPlaces)
	return &models.WalletDeposit{
		Wallet: models.WalletStatus{
			ID:               walletRec.ID,
			Balance:          walletRec.Balance,
			Currency:         walletRec.CurrencyCode,
			FormattedBalance: formattedBalance,
		},
		Transaction: models.WalletTransaction{
			ID:        ledgerRec.ID,
			Amount:    ledgerRec.Amount,
			Timestamp: ledgerRec.CreatedAt,
			Type:      ledgerRec.TransactionType,
		},
	}, nil
}
