package wallet

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/leeliwei930/walletassignment/ent/user"
	"github.com/leeliwei930/walletassignment/ent/wallet"
	pkgappctx "github.com/leeliwei930/walletassignment/internal/app/context"
	"github.com/leeliwei930/walletassignment/internal/app/models"
	"github.com/leeliwei930/walletassignment/pkg/formatter"
)

const TRX_TYPE_DEPOSIT = "deposit"
const TRX_TYPE_DESCRIPTION_LOCALE_KEY = "wallet::deposit::trx::description"

func (s *walletService) Deposit(ctx context.Context, params models.WalletDepositParams) (*models.WalletDeposit, error) {

	appCtx, err := pkgappctx.GetApplicationContext(ctx)
	if err != nil {
		return nil, err
	}
	userID := appCtx.GetAuthUserID()
	entClient := s.app.GetEnt()
	locale := s.app.GetLocale()
	ut := locale.GetUT().GetFallback()

	tx, err := entClient.Tx(ctx)
	if err != nil {
		return nil, err
	}

	userRec, err := tx.User.Query().Where(user.ID(userID)).First(ctx)
	if err != nil {
		return nil, err
	}

	trxDescription, _ := ut.T(TRX_TYPE_DESCRIPTION_LOCALE_KEY, fmt.Sprintf("%s %s", userRec.LastName, userRec.FirstName))

	// retrieve current wallet
	wallet, err := tx.Wallet.
		Query().
		Where(wallet.UserID(userID)).
		ForUpdate(sql.WithLockTables(wallet.Table)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	// update wallet balance
	newBalance := wallet.Balance + params.Amount
	walletRec, err := wallet.Update().SetBalance(newBalance).Save(ctx)
	if err != nil {
		return nil, tx.Rollback()
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
		return nil, tx.Rollback()
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
