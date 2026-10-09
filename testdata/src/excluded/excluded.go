package excluded

import (
	"dao"
	"domain"
)

func insertRaised() error {
	_, err := dao.Invoice.Insert(1)
	if err != nil {
		return domain.WrapError(err)
	}

	return nil
}
