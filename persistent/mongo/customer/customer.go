package customer

import (
	"fmt"
	"newMail/models"
)

func Create(customer models.Customer) (models.Customer, error) {
	//save to DB
	customer.ID = 1
	return customer, nil
}

func Read(id uint64) (models.Customer, error) {
	//read from DB
	result := models.Customer{
		ID:      id,
		Name:    "mock",
		SurName: "mock",
		Phone:   0,
	}
	return result, nil
}

func Update(customer models.Customer) (models.Customer, error) {
	//update customer
	return customer, nil
}

func Delete(id uint64) error {
	//TODO delete
	return fmt.Errorf("mock error")
}

func ReadAll() ([]models.Customer, error) {
	mockData := []models.Customer{
		{
			ID:   1,
			Name: "mock",
		},
		{
			ID:   2,
			Name: "mock",
		},
	}
	return mockData, nil
}
