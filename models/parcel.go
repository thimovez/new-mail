package models

type Parcel struct {
	ID       uint64
	Name     string
	Weight   float64
	Volume   float64
	Price    float64
	Sender   uint64
	Receiver uint64
}
