package validator

import "regexp"

func EmailValidator(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}

func PasswordValidator(password string) bool {
	regex := regexp.MustCompile(`^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$`)
	return regex.MatchString(password)
}

func PhoneValidator(contact string) bool {
	return len(contact) == 12
}
