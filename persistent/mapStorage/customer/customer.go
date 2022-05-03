package customer

import (
	"fmt"
	"newMail/models"
	"sync"
)

type storage struct {
	sync.Mutex
	data map[uint64]models.Customer
}

type currID struct {
	sync.Mutex
	id uint64
}

var counter currID
var customers = storage{
	data: make(map[uint64]models.Customer),
}

//func newID() uint64 {
//	counter.Lock()
//	counter.id += 1
//	newId := counter.id
//	counter.Unlock()
//	return newId
//}

func Create(customer models.Customer) (models.Customer, error) {
	//customer.ID = newID()
	id := customer.ID
	counter.Lock()
	customers.data[id] = customer
	counter.Unlock()
	return customer, nil
}

func Read(id uint64) (models.Customer, error) {
	counter.Lock()
	value, ok := customers.data[id]
	if !ok {
		return value, fmt.Errorf("id does not exist")
	}
	result := customers.data[id]
	counter.Unlock()
	return result, nil
}

func Update(customer models.Customer) (models.Customer, error) {
	customers.Lock()
	customers.data[customer.ID] = customer
	customers.Unlock()
	return customer, nil
}

func Delete(id uint64) error {
	counter.Lock()
	delete(customers.data, id)
	counter.Unlock()
	return nil
}

func ReadAll() ([]models.Customer, error) {
	customers.Lock()
	result := make([]models.Customer, 0, len(customers.data))
	for _, value := range customers.data {
		result = append(result, value)
	}
	customers.Unlock()
	return result, nil
}
