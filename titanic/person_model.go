package titanic

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Person is a person from the Titanic dataset.
type Person struct {
	Survived                bool    `json:"survived"`
	PassengerClass          int     `json:"passengerClass"`
	Name                    string  `json:"name"`
	Sex                     Sex     `json:"sex"`
	Age                     int     `json:"age"`
	SiblingsOrSpousesAboard int     `json:"siblingsOrSpousesAboard"`
	ParentsOrChildrenAboard int     `json:"parentsOrChildrenAboard"`
	Fare                    float64 `json:"fare"`
	UUID                    string  `json:"uuid" gorm:"type:uuid;primary_key"`
}

// TableName overrides the table name used by Person. Without this, the table
// name would be `peoples`.
func (Person) TableName() string {
	return "people"
}

// Sex of the person from the Titanic dataset.
type Sex string

// Sex represents the sex of persons from the Titanic dataset.
const (
	Male   Sex = "male"
	Female Sex = "female"
	Other  Sex = "other"
)

// People represents a list of persons.
type People []Person

// BeforeCreate generates a new uuid for the person before inserting it on
// the database.
func (p *Person) BeforeCreate(tx *gorm.DB) (err error) {
	p.UUID = uuid.New().String()
	return nil
}
