package controllers

func MaskCreditCardNumber(cardNumber string) string {
	if len(cardNumber) > 4 {
		masked := "***" + cardNumber[len(cardNumber)-4:]
		return masked
	}
	return cardNumber
}