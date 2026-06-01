package type2

import (
	"errors"
	"regexp"
)

type BaseProfile struct {
	Name         string
	Email        string
	EnrollmentID string
}

var reEmail = regexp.MustCompile(`^[a-zA-Z]+@[a-zA-Z]+\.(com|org)$`)
var reID = regexp.MustCompile(`^EDU-[0-9]{4}$`)

func (bp BaseProfile) Validate() error {
	isValidEmail, isValidID := reEmail.MatchString(bp.Email), reID.MatchString(bp.EnrollmentID)

	if !isValidEmail {
		return errors.New("invalid email")
	}

	if !isValidID {
		return errors.New("invalid enrollment-ID")
	}
	return nil
}
