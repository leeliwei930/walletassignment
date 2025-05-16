package formatter

import (
	"fmt"
	"math"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func FormatCurrencyAmount(cents int, currencyCode string, decimalPlaces int) string {

	// Convert cents to the amount in the currency's smallest unit.
	amount := float64(cents) / math.Pow(10, float64(decimalPlaces))

	cur, err := currency.ParseISO(currencyCode)
	if err != nil {
		return fmt.Sprintf("$ %.2f", amount)
	}

	printer := message.NewPrinter(language.English)

	return printer.Sprint(cur.Amount(amount))
}
