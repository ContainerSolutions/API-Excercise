package parser_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"gitlab.com/DGuedes/API-Exercise/parser"
	"gitlab.com/DGuedes/API-Exercise/titanic"
	"gitlab.com/DGuedes/API-Exercise/util"

	"testing"
)

var tempFilePath string

func TestExamples(t *testing.T) {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	logrus.SetOutput(os.Stdout)
	RegisterFailHandler(Fail)
	tempFilePath = "./testtitanicinput.csv"
	writeTempCsvTestFile(tempFilePath)
	RunSpecs(t, "Postgres Titanic Paser Suite")
	removeTempCsvTestFile(tempFilePath)
}

// Remove the temporary csv file created before.
func removeTempCsvTestFile(path string) {
	err := os.Remove(path)
	if err != nil {
		panic(err)
	}
}

// Write a temporary csv file used to test the csv file parsing.
func writeTempCsvTestFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	testInput := `
Survived,Pclass,Name,Sex,Age,Siblings/Spouses Aboard,Parents/Children Aboard,Fare
0,3,Mr. Owen Harris Braund,male,22,1,0,7.25
1,1,Mrs. John Bradley (Florence Briggs Thayer) Cumings,female,38,1,0,71.2833`

	f.WriteString(testInput)
}

var _ = Describe("Titanic Parser", func() {
	It("Parser given file successfully", func() {
		parser := parser.TitanicParser{}
		result := parser.Parse(tempFilePath)
		Expect(result).To(HaveLen(2))
	})

	It("Parses a csv line correctly", func() {
		parser := parser.TitanicParser{}
		line := []string{"0", "3", "Mr. Owen Harris Braund", "male", "22", "1", "0", "7.25"}
		person, err := parser.ParseLine(line)
		util.Check(err)
		Expect(person.Survived).To(Equal(false))
		Expect(person.PassengerClass).To(Equal(3))
		Expect(person.Name).To(Equal("Mr. Owen Harris Braund"))
		Expect(person.Sex).To(Equal(titanic.Male))
		Expect(person.Age).To(Equal(22))
		Expect(person.SiblingsOrSpousesAboard).To(Equal(1))
		Expect(person.ParentsOrChildrenAboard).To(Equal(0))
		Expect(person.Fare).To(Equal(7.25))
	})

	It("Parses wrongly given float as int correctly", func() {
		// Line 58 of the dataset has an entry with a float age.
		parser := parser.TitanicParser{}
		line := []string{"0", "3", "Mr. Mansouer Novel", "male", "28.5", "0", "0", "7.2292"}
		person, err := parser.ParseLine(line)
		util.Check(err)
		Expect(person.Survived).To(Equal(false))
		Expect(person.PassengerClass).To(Equal(3))
		Expect(person.Name).To(Equal("Mr. Mansouer Novel"))
		Expect(person.Sex).To(Equal(titanic.Male))
		Expect(person.Age).To(Equal(28))
		Expect(person.SiblingsOrSpousesAboard).To(Equal(0))
		Expect(person.ParentsOrChildrenAboard).To(Equal(0))
		Expect(person.Fare).To(Equal(7.2292))
	})
})
