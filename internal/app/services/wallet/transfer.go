package wallet

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/leeliwei930/walletassignment/ent/wallet"
	pkgappcontext "github.com/leeliwei930/walletassignment/internal/app/context"
	"github.com/leeliwei930/walletassignment/internal/app/models"
	apperrors "github.com/leeliwei930/walletassignment/internal/errors"
	"github.com/leeliwei930/walletassignment/pkg/formatter"
)

const (
	TRX_TYPE_TRANSFER_OUT                        = "transfer_out"
	TRX_TYPE_TRANSFER_IN                         = "transfer_in"
	TRX_TYPE_TRANSFER_OUT_DESCRIPTION_LOCALE_KEY = "wallet::transfer::out::description"
	TRX_TYPE_TRANSFER_IN_DESCRIPTION_LOCALE_KEY  = "wallet::transfer::in::description"
)

func (s *walletService) Transfer(ctx context.Context, params models.WalletTransferParams) (*models.WalletTransfer, error) {

	entClient := s.app.GetEnt()
	userSvc := s.app.GetUserService()
	appCtx, err := pkgappcontext.GetApplicationContext(ctx)
	if err != nil {
		return nil, err
	}

	authUserID := appCtx.GetAuthUserID()

	if params.RecipientUserID == authUserID {
		return nil, apperrors.IdenticalSourceAndDestinationTransferErr
	}

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

	sourceWallet, err := tx.Wallet.
		Query().
		Where(wallet.UserID(authUserID)).
		ForUpdate(sql.WithLockTables(wallet.Table)).
		WithUser().
		First(ctx)
	if err != nil {
		return nil, err
	}

	if sourceWallet.Balance < params.Amount {
		return nil, apperrors.InsufficientBalanceTransferErr
	}

	destinationWallet, err := tx.Wallet.
		Query().
		Where(wallet.UserIDEQ(params.RecipientUserID)).
		ForUpdate(sql.WithLockTables(wallet.Table)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	sourceWalletBalance := sourceWallet.Balance - params.Amount
	destinationWalletBalance := destinationWallet.Balance + params.Amount

	sourceWallet, err = sourceWallet.Update().
		SetBalance(sourceWalletBalance).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	destinationWallet, err = destinationWallet.Update().
		SetBalance(destinationWalletBalance).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	locale := s.app.GetLocale()
	ut := locale.GetUT().GetFallback()

	recipient, err := userSvc.GetUserByID(ctx, params.RecipientUserID)
	if err != nil {
		return nil, err
	}

	recipientFullName := userSvc.GetFullName(ctx, recipient)

	transferOutDescription, _ := ut.T(TRX_TYPE_TRANSFER_OUT_DESCRIPTION_LOCALE_KEY, recipientFullName)
	// ledger for source wallet
	sourceLedger, err := tx.Ledger.Create().
		SetWalletID(sourceWallet.ID).
		SetAmount(params.Amount).
		SetDescription(transferOutDescription).
		SetRecipientReferenceNote(params.RecipientReferenceNote).
		SetTransactionType(TRX_TYPE_TRANSFER_OUT).
		SetCreatedAt(sourceWallet.UpdatedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	sender, err := userSvc.GetUserByID(ctx, authUserID)
	if err != nil {
		return nil, err
	}

	senderFullName := userSvc.GetFullName(ctx, sender)
	transferInDescription, _ := ut.T(TRX_TYPE_TRANSFER_IN_DESCRIPTION_LOCALE_KEY, senderFullName)

	// ledger for destination wallet
	_, err = tx.Ledger.Create().
		SetWalletID(destinationWallet.ID).
		SetAmount(params.Amount).
		SetDescription(transferInDescription).
		SetRecipientReferenceNote(params.RecipientReferenceNote).
		SetTransactionType(TRX_TYPE_TRANSFER_IN).
		SetCreatedAt(destinationWallet.UpdatedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	formattedSourceWalletBalance := formatter.FormatCurrencyAmount(
		sourceWallet.Balance,
		sourceWallet.CurrencyCode,
		sourceWallet.DecimalPlaces,
	)
	return &models.WalletTransfer{
		Wallet: models.WalletStatus{
			ID:               sourceWallet.ID,
			Balance:          sourceWallet.Balance,
			Currency:         sourceWallet.CurrencyCode,
			FormattedBalance: formattedSourceWalletBalance,
		},
		Transaction: models.WalletTransaction{
			ID:                     sourceLedger.ID,
			Amount:                 sourceLedger.Amount,
			Timestamp:              sourceLedger.CreatedAt,
			Type:                   sourceLedger.TransactionType,
			RecipientReferenceNote: sourceLedger.RecipientReferenceNote,
		},
	}, nil
}
