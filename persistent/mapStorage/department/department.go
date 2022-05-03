package department

import (
	"fmt"
	"newMail/models"
	"sync"
)

type storage struct {
	sync.Mutex
	data map[uint64]models.PostOffice
}

type currID struct {
	sync.Mutex
	id uint64
}

var counter currID
var departments = storage{
	data: make(map[uint64]models.PostOffice),
}

func newID() uint64 {
	departments.Lock()
	counter.id += 1
	newId := counter.id
	departments.Unlock()
	return newId
}

func Create(department models.PostOffice) (models.PostOffice, error) {
	department.ID = newID()
	departments.Lock()
	departments.data[department.ID] = department
	departments.Unlock()
	return department, nil
}

func Read(id uint64) (models.PostOffice, error) {
	departments.Lock()
	value, ok := departments.data[id]
	if !ok {
		return value, fmt.Errorf("id does not exist")
	}
	result := departments.data[id]
	departments.Unlock()
	return result, nil
}

func Update(department models.PostOffice) (models.PostOffice, error) {
	counter.Lock()
	departments.data[department.ID] = department
	counter.Unlock()
	return department, nil
}

func Delete(id uint64) error {
	counter.Lock()
	delete(departments.data, id)
	counter.Unlock()
	return nil
}

func ReadAll() ([]models.PostOffice, error) {
	departments.Lock()
	result := make([]models.PostOffice, 0, len(departments.data))
	for _, value := range departments.data {
		result = append(result, value)
	}
	departments.Unlock()
	return result, nil
}
