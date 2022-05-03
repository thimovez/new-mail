package models

type Parcel struct {
	ID                      uint64
	Name                    string
	Weight                  float64
	Volume                  float64
	Price                   float64
	SenderID                uint64
	ReceiverID              uint64
	SourceDepartmentID      uint64
	DestinationDepartmentID uint64
}
