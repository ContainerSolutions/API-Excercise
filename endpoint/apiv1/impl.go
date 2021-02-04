package apiv1

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"gitlab.com/DGuedes/API-Exercise/storage/postgres"
	"gitlab.com/DGuedes/API-Exercise/titanic"
	"gitlab.com/DGuedes/API-Exercise/util"
)

// APIImpl holds the v1 implementation of the endpoints.
type APIImpl struct {
	PersonDAO      *postgres.PersonDAO
	TemplateFolder string
}

// GetPeople handle GET /people endpoint and return the result of the query
// that select all persons from the titanic dataset.
func (a *APIImpl) GetPeople(w http.ResponseWriter, r *http.Request) {
	log.Info("GET /people")
	result, _ := a.PersonDAO.GetAll()

	switch r.Header.Get("Content-Type") {
	case "application/json":
		a.handleGetPeopleJSON(w, r, result)

	default:
		a.handleGetPeopleHTML(w, r, result)
	}
}

func (a *APIImpl) handleGetPeopleJSON(w http.ResponseWriter, r *http.Request, result titanic.People) {
	jsonReply(w, http.StatusOK, result)
}

func (a *APIImpl) handleGetPeopleHTML(w http.ResponseWriter, r *http.Request, result titanic.People) {
	templatePath := fmt.Sprintf("%shtml_templates/people/index.go.tmpl", a.TemplateFolder)
	tmpl, err := template.ParseFiles(templatePath)
	util.Check(err)
	tmpl.Execute(w, map[string]titanic.People{"Data": result})
}

// CreatePerson handle POST /people endpoint, try to create a new person with
// the given parameters and return the result of this operation.
func (a *APIImpl) CreatePerson(w http.ResponseWriter, r *http.Request) {
	log.Info("POST /people")
	var p titanic.Person
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		jsonReply(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	defer r.Body.Close()

	a.PersonDAO.Insert(&p)
	jsonReply(w, http.StatusOK, p)
}

// GetPerson handle GET /people/:uuid endpoint and return the result of the query
// that selects one person based on the given :uuid.
func (a *APIImpl) GetPerson(w http.ResponseWriter, r *http.Request) {
	log.Info("GET /people/:uuid")
	vars := mux.Vars(r)
	u, _ := vars["uuid"]
	result, err := a.PersonDAO.Find(u)

	if err != nil {
		jsonReply(w, http.StatusNotFound, map[string]string{"error": "Person not found."})
		return
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		a.handleGetPersonJSON(w, r, result)

	default:
		a.handleGetPersonHTML(w, r, result)
	}
}

func (a *APIImpl) handleGetPersonJSON(w http.ResponseWriter, r *http.Request, result *titanic.Person) {
	jsonReply(w, http.StatusOK, result)
}

func (a *APIImpl) handleGetPersonHTML(w http.ResponseWriter, r *http.Request, result *titanic.Person) {
	templatePath := fmt.Sprintf("%shtml_templates/people/show.go.tmpl", a.TemplateFolder)
	tmpl, err := template.ParseFiles(templatePath)
	util.Check(err)
	tmpl.Execute(w, map[string]titanic.Person{"Data": *result})
}

// EditPerson handle PUT /people/:uuid endpoint and return the result of the
// operation that modifies one person based on the given parameters.
func (a *APIImpl) EditPerson(w http.ResponseWriter, r *http.Request) {
	log.Info("PUT /people/:uuid")
	vars := mux.Vars(r)
	u, _ := vars["uuid"]
	decoder := json.NewDecoder(r.Body)
	p := titanic.Person{}
	if err := decoder.Decode(&p); err != nil {
		jsonReply(w, http.StatusNotFound, map[string]string{"error": "Invalid request payload"})
		return
	}
	defer r.Body.Close()
	p.UUID = u
	a.PersonDAO.Update(&p)
	emptyPeople := []titanic.People{}
	jsonReply(w, http.StatusOK, emptyPeople)
}

// DeletePerson handle DELETE /people/:uuid endpoint and return the result of
// the deletion.
func (a *APIImpl) DeletePerson(w http.ResponseWriter, r *http.Request) {
	log.Info("DELETE /people/:uuid")
	vars := mux.Vars(r)
	err := a.PersonDAO.Destroy(vars["uuid"])
	if err != nil {
		jsonReply(w, http.StatusNotFound, map[string]string{"error": "Couldn't delete person"})
		return
	}
	jsonReply(w, http.StatusOK, nil)
}

func jsonReply(w http.ResponseWriter, httpCode int, body interface{}) {
	response, _ := json.Marshal(body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(response)
}
