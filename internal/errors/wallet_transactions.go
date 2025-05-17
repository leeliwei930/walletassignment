package errors

var InsuficcientBalanceWithdrawalErr = NewInvalidRequestError(
	"ERR_WITHDRAW_10001",
	"errors::wallet::withdraw::insufficient_balance",
	nil,
)
