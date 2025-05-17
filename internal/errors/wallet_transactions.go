package errors

var InsuficcientBalanceWithdrawalErr = NewInvalidRequestError(
	"ERR_WITHDRAW_10001",
	"errors::wallet::withdraw::insufficient_balance",
	nil,
)

var InsufficientBalanceTransferErr = NewInvalidRequestError(
	"ERR_TRANSFER_10003",
	"errors::wallet::transfer::insufficient_balance",
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
