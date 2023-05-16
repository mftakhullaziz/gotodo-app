package validator

import (
	"errors"
	"github.com/badoux/checkmail"
	"gotodo/internal/utils/logger"
)

func ValidateEmail(email string) error {
	if err := checkmail.ValidateFormat(email); err != nil {
		return err
	}
	return nil
}

func ValidateIntValue(val ...int) error {
	log := logger.LoggerParent()
	for _, v := range val {
		if v <= 0 {
			log.Errorln(v)
			return errors.New("invalid int value")
		}
	}
	return nil
}
