package parcel

import (
	"fmt"
	"newMail/models"
	"newMail/persistent/mapStorage/customer"
	persistent "newMail/persistent/mapStorage/parcel"
	"newMail/repo/department"
)

const ERRORNOTFOUND = "not found"

func Create(parcel models.Parcel) (newParcel models.Parcel, err error) {
	v, err := validateParcel(parcel)
	if err != nil {
		return v, err
	}

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

func validateParcel(parcel models.Parcel) (models.Parcel, error) {
	//проверка номер телефона на существование
	sender, err := customer.Read(parcel.SenderID)
	if err != nil {
		return parcel, fmt.Errorf("sender not exist")
	}
	switch {
	case parcel.SenderID != sender.ID:
		return parcel, fmt.Errorf("sender is not correct")
	}

	//проверка на вес и обьем и отделение на которое посылка едет
	dep, err := department.Read(parcel.DestinationDepartmentID)
	if err != nil {
		return parcel, fmt.Errorf("departament error")
	}
	switch {
	case parcel.Volume > dep.MaxVolume:
		return parcel, fmt.Errorf("volume is not correct")
	case parcel.Weight > dep.MaxWeight:
		return parcel, fmt.Errorf("weight is not correct")
	}
	return parcel, nil
}
