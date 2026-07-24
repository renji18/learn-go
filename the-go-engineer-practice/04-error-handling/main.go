package main

import (
	"errors"
	"fmt"
)

var ErrInvalidAmount = errors.New("Amount cannot be <= 0")

type OverLimitError struct {
	Amount  float64
	Message string
}

func (o *OverLimitError) Error() string {
	return fmt.Sprintf("Overlimit Error for amount %.2f: %s", o.Amount, o.Message)
}

func ProcessPayment(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	} else if amount > 10000 {
		return &OverLimitError{Amount: amount, Message: "Amount cannot be more than 10000"}
	}

	return nil
}

func Checkout(amount float64) error {
	err := ProcessPayment(amount)
	if err != nil {
		return fmt.Errorf("Error during checkout: %w", err)
	}

	return nil
}

func HandleCheckout(amount float64) (int, string) {
	err := Checkout(amount)
	var overLimitErr *OverLimitError

	if err != nil {
		if errors.Is(err, ErrInvalidAmount) {
			return 400, "invalid amount"
		} else if errors.As(err, &overLimitErr) {
			return 403, fmt.Sprintf("Forbidden: amount %.2f exceeds limit\n", overLimitErr.Amount)
		} else {
			return 500, "internal error"
		}
	}

	return 200, "success"
}

func main() {
	status, message := HandleCheckout(0)
	fmt.Printf("Status: %d. Message: %s\n", status, message)

	status, message = HandleCheckout(10)
	fmt.Printf("Status: %d. Message: %s\n", status, message)

	status, message = HandleCheckout(-120)
	fmt.Printf("Status: %d. Message: %s\n", status, message)

	status, message = HandleCheckout(1000)
	fmt.Printf("Status: %d. Message: %s\n", status, message)

	status, message = HandleCheckout(10000)
	fmt.Printf("Status: %d. Message: %s\n", status, message)

	status, message = HandleCheckout(10001)
	fmt.Printf("Status: %d. Message: %s\n", status, message)

}
