package postgres

import (
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DAOFactory manages the creation of new data access objects to Postgres.
type DAOFactory struct {
	DB     *gorm.DB
	Config *Config
}

// NewPersonDAO creates a new person DAO with this same postgres instance,
// applies its migrations and returns the fresh new instance.
func (f *DAOFactory) NewPersonDAO() *PersonDAO {
	newDAO := &PersonDAO{DB: f.DB}
	newDAO.AutoMigrate()
	return newDAO
}

// Connect connects the DAO to the configured Postgres instance.
func (f *DAOFactory) Connect() error {
	dsn := f.Config.BuildDSN()
	db, err := gorm.Open(driver.Open(dsn), &gorm.Config{})
	f.DB = db
	return err
}

// BuildDAOFactory builds a data access object used to interact with
// Postgres.
func BuildDAOFactory(cfg *Config) *DAOFactory {
	return &DAOFactory{Config: cfg}
}
