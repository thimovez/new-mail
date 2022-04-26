package customer

import (
	"fmt"
	"newMail/models"
	persistent "newMail/persistent/mapStorage/customer"
)

const ERRORNOTFOUND = "not found"

func Create(customer models.Customer) (newCustomer models.Customer, err error) {
	newCustomer, err = persistent.Create(customer)
	if err != nil {
		if err.Error() == ERRORNOTFOUND {
			err = fmt.Errorf("not found")
			return
		}
		return
	}
	return
}

func Read(id uint64) (customer models.Customer, err error) {
	return persistent.Read(id)
}

func Update(customer models.Customer) (updatedCustomer models.Customer, err error) {
	updatedCustomer, err = persistent.Update(customer)
	if err != nil {
		if err.Error() == ERRORNOTFOUND {
			err = fmt.Errorf("not found")
			return
		}
		return
	}
	return
}

func Delete(id uint64) error {
	return persistent.Delete(id)
}

func ReadAll() ([]models.Customer, error) {
	return persistent.ReadAll()
}
