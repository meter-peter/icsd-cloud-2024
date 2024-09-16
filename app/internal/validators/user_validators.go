package validators

import (
	"errors"
	"regexp"
	"time"

	"github.com/meter-peter/icsd-cloud-2024/internal/models"
)

func ValidateUser(user *models.User) error {
	if err := validateIdentityNumber(user.IdentityNumber); err != nil {
		return err
	}
	if err := validateName(user.FirstName); err != nil {
		return err
	}
	if err := validateName(user.LastName); err != nil {
		return err
	}
	if err := validateGender(user.Gender); err != nil {
		return err
	}
	if err := validateDateOfBirth(user.DateOfBirth); err != nil {
		return err
	}
	if err := validateAddresses(user.Addresses); err != nil {
		return err
	}
	if err := validatePhoneNumbers(user.PhoneNumbers); err != nil {
		return err
	}
	return nil
}

func validateIdentityNumber(id string) error {
	matched, _ := regexp.MatchString(`^[A-Z]{2}\d{6}$`, id)
	if !matched {
		return errors.New("invalid identity number format")
	}
	return nil
}

func validateName(name string) error {
	matched, _ := regexp.MatchString(`^[A-Z][a-z]{2,}$`, name)
	if !matched {
		return errors.New("invalid name format")
	}
	return nil
}

func validateGender(gender string) error {
	if gender != "male" && gender != "female" {
		return errors.New("invalid gender")
	}
	return nil
}

func validateDateOfBirth(dob time.Time) error {
	age := time.Since(dob).Hours() / 24 / 365
	if age < 16 {
		return errors.New("user must be at least 16 years old")
	}
	return nil
}

func validateAddresses(addresses []string) error {
	for _, address := range addresses {
		matched, _ := regexp.MatchString(`^.+\s\d{1,3}$`, address)
		if !matched {
			return errors.New("invalid address format")
		}
	}
	return nil
}

func validatePhoneNumbers(numbers []string) error {
	for _, number := range numbers {
		matched, _ := regexp.MatchString(`^\d{10}$`, number)
		if !matched {
			return errors.New("invalid phone number format")
		}
	}
	return nil
}
