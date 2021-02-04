package parser

import (
	"encoding/csv"
	"os"
	"strconv"

	"gitlab.com/DGuedes/API-Exercise/titanic"
	"gitlab.com/DGuedes/API-Exercise/util"
)

// TitanicParser parses the Titanic dataset.
type TitanicParser struct {
}

// Parse will try parsing the file in the given path.
func (p *TitanicParser) Parse(path string) titanic.People {
	r, err := os.Open(path)
	util.Check(err)

	reader := csv.NewReader(r)

	// Read to ignore the header line
	reader.Read()

	// Read remaining lines of the file
	csvLines, err := reader.ReadAll()
	util.Check(err)

	var result []titanic.Person
	for _, entry := range csvLines {
		person, err := p.ParseLine(entry)
		util.Check(err)

		result = append(result, *person)
	}
	return result
}

// ParseLine parses a csv line as a Person instance.
func (p *TitanicParser) ParseLine(line []string) (*titanic.Person, error) {
	survived, err := strconv.ParseBool(line[0])
	util.Check(err)

	passengerClass, err := strconv.Atoi(line[1])
	util.Check(err)

	// Although age is listed as an int, in some fields it is a float.
	age, err := strconv.ParseFloat(line[4], 32)
	util.Check(err)

	siblings, err := strconv.Atoi(line[5])
	util.Check(err)

	parents, err := strconv.Atoi(line[6])
	util.Check(err)

	fare, err := strconv.ParseFloat(line[7], 64)
	util.Check(err)

	return &titanic.Person{
		Survived:                survived,
		PassengerClass:          passengerClass,
		Name:                    line[2],
		Sex:                     titanic.Sex(line[3]),
		Age:                     int(age),
		SiblingsOrSpousesAboard: siblings,
		ParentsOrChildrenAboard: parents,
		Fare:                    fare,
	}, nil
}
