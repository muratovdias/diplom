package service

import (
	"net/mail"

	"github.com/muratovdias/diplom/internal/models"
)

func isValidEmail(email string) error {
	_, err = mail.ParseAddress(email)
	if err != nil {
		return ErrInvalidEmail
	}
	return nil
}

// func isValidPassword(pwd string) error {
// 	passwordRegex := "^(?:.*[a-z])(?:.*[A-Z])(?:.*[0-9])(?:.*[!@#\\$%^&\\*])(?:.{8,})"
// 	re := regexp.MustCompile(passwordRegex)
// 	if !re.MatchString(pwd) {
// 		return ErrInvalidPwd
// 	}
// 	return nil
// }

func validate(user models.User) error {
	// err1 := isValidPassword(user.Pwd)
	err2 := isValidEmail(user.Email)
	// if err1 != nil && err2 != nil {
	// 	return ErrEmailAndPwd
	// } else if err1 != nil {
	// 	return ErrInvalidPwd
	// } else if err2 != nil {
	if err2 != nil {
		return ErrInvalidEmail
	}
	return nil
}
