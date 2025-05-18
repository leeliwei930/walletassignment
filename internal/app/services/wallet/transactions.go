package wallet

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/leeliwei930/walletassignment/ent"
	"github.com/leeliwei930/walletassignment/ent/ledger"
	"github.com/leeliwei930/walletassignment/ent/user"
	"github.com/leeliwei930/walletassignment/internal/app/models"
	"github.com/leeliwei930/walletassignment/pkg/paginator"
)

func (s *walletService) Transactions(ctx context.Context, params models.WalletTransactionsParams) (*models.WalletTransactions, error) {

	app := s.app
	entClient := app.GetEnt()

	limitAndOffSet := paginator.GetLimitAndOffSet(
		paginator.WithLimit(ctx, params.Limit),
		paginator.WithCurrentPage(ctx, params.Page),
	)

	transactionQuery := entClient.Ledger.
		Query().
		WithWallet(
			func(wq *ent.WalletQuery) {
				wq.WithUser(
					func(uq *ent.UserQuery) {
						uq.Where(user.IDEQ(params.UserID))
					},
				)
			},
		)

	totalTransactionsCount, err := transactionQuery.Count(ctx)
	if err != nil {
		return nil, err
	}

	transactions, err := transactionQuery.
		Order(ledger.ByCreatedAt(sql.OrderDesc())).
		Limit(limitAndOffSet.Limit).
		Offset(limitAndOffSet.Offset).All(ctx)

	if err != nil {
		return nil, err
	}

	walletTransactions := make([]*models.WalletTransaction, 0, len(transactions))
	for _, transaction := range transactions {
		walletTransactions = append(walletTransactions, &models.WalletTransaction{
			ID:                     transaction.ID,
			Amount:                 transaction.Amount,
			Timestamp:              transaction.CreatedAt,
			Type:                   transaction.TransactionType,
			RecipientReferenceNote: transaction.RecipientReferenceNote,
		})
	}

	paginationInfo := paginator.GetPaginationInfo(paginator.PaginationInfoParams{
		TotalItems:   totalTransactionsCount,
		ItemsPerPage: len(walletTransactions),
		CurrentPage:  params.Page,
		Limit:        limitAndOffSet.Limit,
	})

	return &models.WalletTransactions{
		Transactions: walletTransactions,
		Pagination: &models.Pagination{
			CurrentPage:  paginationInfo.CurrentPage,
			TotalPages:   paginationInfo.TotalPages,
			TotalItems:   paginationInfo.TotalItems,
			ItemsPerPage: paginationInfo.Limit,
			HasNext:      paginationInfo.HasNext,
			HasPrev:      paginationInfo.HasPrev,
		},
	}, nil
}
