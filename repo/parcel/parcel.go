package parcel

import (
	"fmt"
	"newMail/models"
	persistent "newMail/persistent/mapStorage/parcel"
)

const ERRORNOTFOUND = "not found"

func Create(parcel models.Parcel) (newParcel models.Parcel, err error) {
	newParcel, err = persistent.Create(parcel)
	if err != nil {
		if err.Error() == ERRORNOTFOUND {
			err = fmt.Errorf("not found")
			return
		}
		return
	}
	return
}

func Read(id uint64) (models.Parcel, error) {
	return persistent.Read(id)
}

func Update(parcel models.Parcel) (updatedParcel models.Parcel, err error) {
	return persistent.Update(parcel)
}

func Delete(id uint64) error {
	return persistent.Delete(id)
}

func ReadAll() ([]models.Parcel, error) {
	return persistent.ReadAll()
}
