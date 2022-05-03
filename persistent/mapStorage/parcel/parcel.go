package parcel

import (
	"fmt"
	"newMail/models"
	"sync"
)

type storage struct {
	sync.Mutex
	data map[uint64]models.Parcel
}

type currID struct {
	sync.Mutex
	id uint64
}

var counter currID
var parcels = storage{
	data: make(map[uint64]models.Parcel),
}

func newID() uint64 {
	parcels.Lock()
	counter.id += 1
	newId := counter.id
	parcels.Unlock()
	return newId
}

func Create(parcel models.Parcel) (models.Parcel, error) {
	parcel.ID = newID()
	parcels.Lock()
	parcels.data[parcel.ID] = parcel
	parcels.Unlock()
	return parcel, nil
}

func Read(id uint64) (models.Parcel, error) {
	parcels.Lock()
	value, ok := parcels.data[id]
	if !ok {
		return value, fmt.Errorf("id does not exist")
	}
	result := parcels.data[id]
	parcels.Unlock()
	return result, nil
}

func Update(parcel models.Parcel) (models.Parcel, error) {
	counter.Lock()
	parcels.data[parcel.ID] = parcel
	counter.Unlock()
	return parcel, nil
}

func Delete(id uint64) error {
	counter.Lock()
	delete(parcels.data, id)
	counter.Unlock()
	return nil
}

func ReadAll() ([]models.Parcel, error) {
	parcels.Lock()
	result := make([]models.Parcel, 0, len(parcels.data))
	for _, value := range parcels.data {
		result = append(result, value)
	}
	parcels.Unlock()
	return result, nil
}
