package department

import (
	"fmt"
	"newMail/models"
	persistent "newMail/persistent/mapStorage/department"
)

const ERRORNOTFOUND = "not found"

func Create(department models.PostOffice) (newDepartment models.PostOffice, err error) {
	newDepartment, err = persistent.Create(department)
	if err != nil {
		if err.Error() == ERRORNOTFOUND {
			err = fmt.Errorf("not found")
			return
		}
		return
	}
	return
}
func Read(id uint64) (department models.PostOffice, err error) {
	return persistent.Read(id)
}
func Update(department models.PostOffice) (updateDepartment models.PostOffice, err error) {
	return persistent.Update(department)
}
func Delete(id uint64) error {
	return persistent.Delete(id)
}
func ReadAll() ([]models.PostOffice, error) {
	return persistent.ReadAll()
}
