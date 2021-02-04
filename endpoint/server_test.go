package endpoint_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"gitlab.com/DGuedes/API-Exercise/endpoint"
	"gitlab.com/DGuedes/API-Exercise/endpoint/apiv1"
	"gitlab.com/DGuedes/API-Exercise/titanic"
	"gitlab.com/DGuedes/API-Exercise/util"
)

var apiImpl *apiv1.APIImpl
var svr *endpoint.Server

func TestExamples(t *testing.T) {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	logrus.SetOutput(os.Stdout)
	RegisterFailHandler(Fail)
	runServer()
	RunSpecs(t, "Server Integration Test Suite")
}

var _ = Describe("resource people", func() {
	BeforeEach(func() {
		// Since I'm dropping the table after each test, I have to rerun the
		// migrations.
		svr.APIImpl.PersonDAO.AutoMigrate()
	})

	AfterEach(func() {
		svr.DAOFactory.DB.Exec("DROP TABLE if exists people cascade")
	})

	It("DELETE /people/:uuid deletes person successfully", func() {
		person := insertPerson()
		q1, _ := svr.APIImpl.PersonDAO.GetAll()
		// Guarantee that we have one person
		Expect(q1).To(HaveLen(1))

		path := fmt.Sprintf("/people/%s", person.UUID)
		req, _ := http.NewRequest("DELETE", path, nil)
		req.Header.Set("Content-Type", "application/json")
		util.ExecuteRequest(req, svr.Router)

		// Check that we have 0 persons again
		q2, _ := svr.APIImpl.PersonDAO.GetAll()
		Expect(q2).To(HaveLen(0))
	})

	It("PUT /people/:uuid succesfully edits person", func() {
		person := insertPerson()
		reqBody := `{ "survived": false, "passengerClass": 3, "name": "newname", "sex": "male", "age": 22, "siblingsOrSpousesAboard": 1, "parentsOrChildrenAboard":0, "fare":7.25}`
		jsonStr := []byte(reqBody)
		path := fmt.Sprintf("/people/%s", person.UUID)
		req, _ := http.NewRequest("PUT", path, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		util.ExecuteRequest(req, svr.Router)

		// Check that the person was inserted successfully
		query, _ := svr.APIImpl.PersonDAO.Find(person.UUID)
		Expect(query.Survived).To(Equal(false))
		Expect(query.Name).To(Equal("newname"))
	})

	It("POST /people successfully inserts a new person", func() {
		reqBody := `{ "survived": true, "passengerClass": 3, "name": "Mr. New Person", "sex": "male", "age": 22, "siblingsOrSpousesAboard": 1, "parentsOrChildrenAboard":0, "fare":7.25}`
		jsonStr := []byte(reqBody)
		req, _ := http.NewRequest("POST", "/people", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		util.ExecuteRequest(req, svr.Router)

		// Check that the person was inserted successfully
		query, _ := svr.APIImpl.PersonDAO.GetAll()
		Expect(query).To(HaveLen(1))
		firstPerson := query[0]
		Expect(firstPerson.Survived).To(Equal(true))
		Expect(firstPerson.Name).To(Equal("Mr. New Person"))
	})

	It("GET /people/:uuid returns the correct person successfully", func() {
		person := insertPerson()
		path := fmt.Sprintf("/people/%s", person.UUID)
		req, err := http.NewRequest("GET", path, nil)
		req.Header.Set("Content-Type", "application/json")
		util.Check(err)
		response := util.ExecuteRequest(req, svr.Router)
		Expect(response.Code).To(Equal(http.StatusOK))

		// Parse response as json
		var gotPerson map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &gotPerson)
		util.Check(err)

		Expect(gotPerson).To(HaveKeyWithValue("passengerClass", float64(3)))
		Expect(gotPerson).To(HaveKeyWithValue("name", "newperson"))
		Expect(gotPerson).To(HaveKeyWithValue("sex", "male"))
		Expect(gotPerson).To(HaveKeyWithValue("age", float64(22)))
		Expect(gotPerson).To(HaveKeyWithValue("siblingsOrSpousesAboard", float64(1)))
		Expect(gotPerson).To(HaveKeyWithValue("parentsOrChildrenAboard", float64(0)))
		Expect(gotPerson).To(HaveKeyWithValue("fare", 7.25))
	})

	It("GET /people returns the list of people successfully", func() {
		insertPerson()
		req, err := http.NewRequest("GET", "/people", nil)
		req.Header.Set("Content-Type", "application/json")
		util.Check(err)
		response := util.ExecuteRequest(req, svr.Router)
		Expect(response.Code).To(Equal(http.StatusOK))

		// Parse response as json
		var m []map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &m)
		util.Check(err)
		Expect(m).To(HaveLen(1))
		person := m[0]

		Expect(person).To(HaveKeyWithValue("survived", false))
		Expect(person).To(HaveKey("uuid"))
		// Here I cast int fields to float64 because Unmarshal treat all numbers
		// as float64.
		// Source: https://golang.org/pkg/encoding/json/#Unmarshal
		Expect(person).To(HaveKeyWithValue("passengerClass", float64(3)))
		Expect(person).To(HaveKeyWithValue("name", "newperson"))
		Expect(person).To(HaveKeyWithValue("sex", "male"))
		Expect(person).To(HaveKeyWithValue("age", float64(22)))
		Expect(person).To(HaveKeyWithValue("siblingsOrSpousesAboard", float64(1)))
		Expect(person).To(HaveKeyWithValue("parentsOrChildrenAboard", float64(0)))
		Expect(person).To(HaveKeyWithValue("fare", 7.25))
	})
})

func runServer() {
	logrus.Info("Running test server")
	svrCfg := endpoint.BuildSvrConfig("test")
	svr = endpoint.BuildSvr(svrCfg)

	err := svr.DAOFactory.Connect() // Connect to Postgres database.
	util.Check(err)
	svr.InitializeRoutes()
	server := httptest.NewServer(svr.Router)
	defer server.Close()
}

func insertPerson() *titanic.Person {
	person := titanic.Person{
		Survived:                false,
		PassengerClass:          3,
		Name:                    "newperson",
		Sex:                     titanic.Male,
		Age:                     22,
		SiblingsOrSpousesAboard: 1,
		ParentsOrChildrenAboard: 0,
		Fare:                    7.25,
	}
	err := svr.APIImpl.PersonDAO.Insert(&person)
	util.Check(err)
	return &person
}
