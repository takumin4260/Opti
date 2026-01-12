package model

type UserContext struct {
	ID            string
	UserID        string
	ResidenceInfo ResidenceInfo
}

type ResidenceInfo struct {
	Type      ResidenceType
	Age       int
	Layout    string
	Ownership Ownership
}

type ResidenceType string

const (
	ResidenceTypeApartment ResidenceType = "apartment"
	ResidenceTypeHouse     ResidenceType = "house"
	ResidenceTypeTownhouse ResidenceType = "townhouse"
	ResidenceTypeOther     ResidenceType = "other"
)

type Ownership string

const (
	OwnershipOwned  Ownership = "owned"
	OwnershipRented Ownership = "rented"
	OwnershipOther  Ownership = "other"
)
