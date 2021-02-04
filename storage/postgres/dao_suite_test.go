package postgres_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"gitlab.com/DGuedes/API-Exercise/storage/postgres"

	"testing"
)

var p *postgres.PersonDAO

func TestExamples(t *testing.T) {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	logrus.SetOutput(os.Stdout)
	RegisterFailHandler(Fail)
	cfg := postgres.BuildConfig("test")
	factory := postgres.BuildDAOFactory(&cfg)
	factory.Connect()
	p = factory.NewPersonDAO()
	RunSpecs(t, "Postgres Titanic DAO Suite")
}
