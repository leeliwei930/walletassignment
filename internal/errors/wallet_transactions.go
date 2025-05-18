package errors

import ut "github.com/go-playground/universal-translator"

var InsuficcientBalanceWithdrawalErr = NewInvalidRequestError(
	"ERR_WITHDRAW_10000",
	"errors::wallet::withdraw::insufficient_balance",
	nil,
)

var IdenticalSourceAndDestinationTransferErr = NewInvalidRequestError(
	"ERR_TRANSFER_10001",
	"errors::wallet::transfer::identical_source_and_destination",
	nil,
)

var InvalidTransferRecipientPhoneNumberErr = NewInvalidRequestError(
	"ERR_TRANSFER_10002",
	"errors::wallet::transfer::invalid_recipient_phone_number",
	nil,
)

var InsufficientBalanceTransferErr = NewInvalidRequestError(
	"ERR_TRANSFER_10003",
	"errors::wallet::transfer::insufficient_balance",
	nil,
)

func MinimumTransferAmountRequiredErr(ut ut.Translator, formattedAmount string) *InvalidRequestError {
	message, _ := ut.T("errors::wallet::transfer::minimum_transfer_amount_required", formattedAmount)
	return NewInvalidRequestError(
		"ERR_TRANSFER_10004",
		message,
		nil,
	)
}
func MinimumDepositAmountRequiredErr(ut ut.Translator, formattedAmount string) *InvalidRequestError {

	message, _ := ut.T("errors::wallet::deposit::minimum_deposit_amount_required", formattedAmount)
	return NewInvalidRequestError(
		"ERR_DEPOSIT_10001",
		message,
		nil,
	)
}
func MinimumWithdrawalAmountRequiredErr(ut ut.Translator, formattedAmount string) *InvalidRequestError {

	message, _ := ut.T("errors::wallet::withdraw::minimum_withdrawal_amount_required", formattedAmount)
	return NewInvalidRequestError(
		"ERR_WITHDRAW_10001",
		message,
		nil,
	)
}
