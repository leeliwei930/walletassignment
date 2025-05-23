// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/leeliwei930/walletassignment/ent/ledger"
	"github.com/leeliwei930/walletassignment/ent/schema"
	"github.com/leeliwei930/walletassignment/ent/user"
	"github.com/leeliwei930/walletassignment/ent/wallet"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	ledgerFields := schema.Ledger{}.Fields()
	_ = ledgerFields
	// ledgerDescAmount is the schema descriptor for amount field.
	ledgerDescAmount := ledgerFields[2].Descriptor()
	// ledger.AmountValidator is a validator for the "amount" field. It is called by the builders before save.
	ledger.AmountValidator = ledgerDescAmount.Validators[0].(func(int) error)
	// ledgerDescCreatedAt is the schema descriptor for created_at field.
	ledgerDescCreatedAt := ledgerFields[6].Descriptor()
	// ledger.DefaultCreatedAt holds the default value on creation for the created_at field.
	ledger.DefaultCreatedAt = ledgerDescCreatedAt.Default.(func() time.Time)
	// ledgerDescUpdatedAt is the schema descriptor for updated_at field.
	ledgerDescUpdatedAt := ledgerFields[7].Descriptor()
	// ledger.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	ledger.DefaultUpdatedAt = ledgerDescUpdatedAt.Default.(func() time.Time)
	// ledger.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	ledger.UpdateDefaultUpdatedAt = ledgerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// ledgerDescID is the schema descriptor for id field.
	ledgerDescID := ledgerFields[0].Descriptor()
	// ledger.DefaultID holds the default value on creation for the id field.
	ledger.DefaultID = ledgerDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescFirstName is the schema descriptor for first_name field.
	userDescFirstName := userFields[1].Descriptor()
	// user.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	user.FirstNameValidator = userDescFirstName.Validators[0].(func(string) error)
	// userDescLastName is the schema descriptor for last_name field.
	userDescLastName := userFields[2].Descriptor()
	// user.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	user.LastNameValidator = userDescLastName.Validators[0].(func(string) error)
	// userDescPhoneNumber is the schema descriptor for phone_number field.
	userDescPhoneNumber := userFields[3].Descriptor()
	// user.PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	user.PhoneNumberValidator = userDescPhoneNumber.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[5].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	walletFields := schema.Wallet{}.Fields()
	_ = walletFields
	// walletDescBalance is the schema descriptor for balance field.
	walletDescBalance := walletFields[2].Descriptor()
	// wallet.DefaultBalance holds the default value on creation for the balance field.
	wallet.DefaultBalance = walletDescBalance.Default.(int)
	// walletDescCurrencyCode is the schema descriptor for currency_code field.
	walletDescCurrencyCode := walletFields[3].Descriptor()
	// wallet.DefaultCurrencyCode holds the default value on creation for the currency_code field.
	wallet.DefaultCurrencyCode = walletDescCurrencyCode.Default.(string)
	// walletDescDecimalPlaces is the schema descriptor for decimal_places field.
	walletDescDecimalPlaces := walletFields[4].Descriptor()
	// wallet.DefaultDecimalPlaces holds the default value on creation for the decimal_places field.
	wallet.DefaultDecimalPlaces = walletDescDecimalPlaces.Default.(int)
	// walletDescCreatedAt is the schema descriptor for created_at field.
	walletDescCreatedAt := walletFields[5].Descriptor()
	// wallet.DefaultCreatedAt holds the default value on creation for the created_at field.
	wallet.DefaultCreatedAt = walletDescCreatedAt.Default.(func() time.Time)
	// walletDescUpdatedAt is the schema descriptor for updated_at field.
	walletDescUpdatedAt := walletFields[6].Descriptor()
	// wallet.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	wallet.DefaultUpdatedAt = walletDescUpdatedAt.Default.(func() time.Time)
	// wallet.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	wallet.UpdateDefaultUpdatedAt = walletDescUpdatedAt.UpdateDefault.(func() time.Time)
	// walletDescID is the schema descriptor for id field.
	walletDescID := walletFields[0].Descriptor()
	// wallet.DefaultID holds the default value on creation for the id field.
	wallet.DefaultID = walletDescID.Default.(func() uuid.UUID)
}
