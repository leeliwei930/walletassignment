package user

import (
	"context"

	"github.com/leeliwei930/walletassignment/ent"
)

func (s *userService) SetupUser(ctx context.Context, phoneNumber string, firstName string, lastName string) (*ent.User, error) {
	entClient := s.app.GetEnt()

	tx, err := entClient.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	user, err := tx.User.Create().
		SetPhoneNumber(phoneNumber).
		SetFirstName(firstName).
		SetLastName(lastName).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	_, err = tx.Wallet.Create().
		SetUserID(user.ID).
		SetBalance(0).
		SetCurrencyCode("USD").
		SetDecimalPlaces(2).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return user, nil
}
