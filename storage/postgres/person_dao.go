package postgres

import (
	"github.com/google/uuid"
	"gitlab.com/DGuedes/API-Exercise/titanic"
	"gorm.io/gorm"
)

// PersonDAO exposes an API to interact with Person data in a Postgres database.
// It relies heavily on "gorm" package to interact with Postgres. You might find
// more about it at: https://gorm.io/docs/index.html
type PersonDAO struct {
	DB *gorm.DB
}

// Insert a given person in the Postgres database.
func (p *PersonDAO) Insert(person *titanic.Person) error {
	p.DB.Create(person)
	return p.DB.Error
}

// BulkInsert inserts a list of persons in bulk.
func (p *PersonDAO) BulkInsert(people titanic.People) error {
	p.DB.Create(&people)
	return p.DB.Error
}

// GetAll queries all persons from titanic dataset.
func (p *PersonDAO) GetAll() (titanic.People, error) {
	var people titanic.People
	p.DB.Find(&people)

	if p.DB.Error != nil {
		return nil, p.DB.Error
	}

	return people, nil
}

// Find query a person with same uuid from titanic dataset.
func (p *PersonDAO) Find(stringfiedUUID string) (*titanic.Person, error) {
	var result titanic.Person
	u, _ := uuid.Parse(stringfiedUUID)
	p.DB.Find(&result, u)

	if p.DB.Error != nil {
		return nil, p.DB.Error
	}

	return &result, nil
}

// Update updates a person in the database with the new attrs.
func (p *PersonDAO) Update(newPerson *titanic.Person) (*titanic.Person, error) {
	p.DB.Save(newPerson)

	if p.DB.Error != nil {
		return nil, p.DB.Error
	}

	return newPerson, nil
}

// Destroy deletes the person that have the given id as its uuid.
func (p *PersonDAO) Destroy(id string) error {
	u, _ := uuid.Parse(id)
	p.DB.Delete(&titanic.Person{}, u)
	return p.DB.Error
}

// AutoMigrate applies to the database migrations related to Person.
func (p *PersonDAO) AutoMigrate() {
	p.DB.AutoMigrate(&titanic.Person{})
}
