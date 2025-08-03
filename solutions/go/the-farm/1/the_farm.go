package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fodder FodderCalculator, cows int) (float64, error) {
	amount, amountErr := fodder.FodderAmount(cows)
	factor, factorErr := fodder.FatteningFactor()
	if amountErr != nil {
		return 0, amountErr
	}
	if factorErr != nil {
		return 0, factorErr
	}
	return amount / float64(cows) * factor, nil
}

func ValidateInputAndDivideFood(fodder FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fodder, cows)
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%v cows are invalid: %v", e.cows, e.message)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows, "there are no negative cows"}
	}
	if cows == 0 {
		return &InvalidCowsError{cows, "no cows don't need food"}
	}
	return nil
}
